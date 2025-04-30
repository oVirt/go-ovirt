//
// Copyright (c) 2020 huihui <huihui.fu@cs2c.com.cn>.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package examples

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"os"
	"time"

	ovirtsdk4 "github.com/ovirt/go-ovirt/v4"
)

func vmBackup() {
	inputRawURL := "https://10.1.111.229/ovirt-engine/api"

	conn, err := ovirtsdk4.NewConnectionBuilder().
		URL(inputRawURL).
		Username("admin@internal").
		Password("qwer1234").
		Insecure(true).
		Compress(true).
		Timeout(time.Second * 10).
		Build()
	if err != nil {
		fmt.Printf("Make connection failed, reason: %v\n", err)
		return
	}
	defer conn.Close()

	// To use `Must` methods, you should recover it if panics
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Panics occurs, try the non-Must methods to find the reason")
		}
	}()

	// Get the reference to the service that we will use to send events to
	// the audit log:
	eventsService := conn.SystemService().EventsService()

	// Get the reference to the "vms" service:
	vmsService := conn.SystemService().VmsService()

	// Find the virtual machine:
	vmsResp, err := vmsService.List().Search("name=myvm").
		AllContent(true).Send()
	if err != nil {
		fmt.Printf("Failed to get data vm list, reason: %v\n", err)
		return
	}

	// Locate the service that manages the virtual machine, as that is where
	// the action methods are defined:
	dataVm := vmsResp.MustVms().Slice()[0]
	fmt.Printf("Find the data vm %v, the id is %v\n",
		dataVm.MustName(), dataVm.MustId())

	//Find the virtual machine were we will attach the disks in order to do
	// the backup:
	agentVmsResp, err := vmsService.List().Search("name=myagent").Send()
	if err != nil {
		fmt.Printf("Failed to get agent vm list, reason: %v\n", err)
		return
	}

	agentVm := agentVmsResp.MustVms().Slice()[0]
	fmt.Printf("Find the agent vm %v, the id is %v\n",
		agentVm.MustName(), agentVm.MustId())

	// Find the services that manage the data and agent virtual machines:
	dataVmService := vmsService.VmService(dataVm.MustId())
	agentVmService := vmsService.VmService(agentVm.MustId())

	//Create an unique description for the snapshot, so that it is easier
	// for the administrator to identify this snapshot as a temporary one
	// created just for backup purposes:
	snapDesc := fmt.Sprintf("%v-backup-%v", dataVm.MustName(), uuid.NewV4())

	// Send an external event to indicate to the administrator that the
	// backup of the virtual machine is starting. Note that the description
	// of the event contains the name of the virtual machine and the name of
	// the temporary snapshot, this way, if something fails, the administrator
	// will know what snapshot was used and remove it manually.
	eventId := time.Now().Unix()
	desc := fmt.Sprintf("Backup of virtual machine %v using snapshot %v is starting",
		dataVm.MustName(), snapDesc)

	eventsService.Add().
		Event(
			ovirtsdk4.NewEventBuilder().
				Vm(
					ovirtsdk4.NewVmBuilder().
						Id(dataVm.MustId()).MustBuild()).
				Origin("mybackup").
				Severity(ovirtsdk4.LOGSEVERITY_NORMAL).
				CustomId(eventId).
				Description(desc).
				MustBuild()).
		Send()

	eventId += 1
	// Save the OVF to a file, so that we can use to restore the virtual
	// machine later. The name of the file is the name of the virtual
	// machine, followed by a dash and the identifier of the virtual machine,
	// to make it unique:
	ovfData := dataVm.MustInitialization().MustConfiguration().MustData()
	ovfFile := fmt.Sprintf("%v-%v.ovf", dataVm.MustName(), dataVm.MustId())

	file, err := os.OpenFile(ovfFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, reason:", err)
		return
	}
	defer file.Close()

	file.WriteString(ovfData)

	// Send the request to create the snapshot. Note that this will return
	// before the snapshot is completely created, so we will later need to
	// wait till the snapshot is completely created.
	// The snapshot will not include memory. Change to True the parameter
	// persist_memorystate to get it (in that case the VM will be paused for a while).
	snapsService := dataVmService.SnapshotsService()
	snap, err := snapsService.Add().
		Snapshot(
			ovirtsdk4.NewSnapshotBuilder().
				Description(snapDesc).
				PersistMemorystate(false).
				MustBuild()).
		Send()

	fmt.Printf("Sent request to create snapshot %v, the id is %v",
		snap.MustSnapshot().MustDescription(), snap.MustSnapshot().MustId())

	// Poll and wait till the status of the snapshot is 'ok', which means
	// that it is completely created:
	snapService := snapsService.SnapshotService(snap.MustSnapshot().MustId())
	for {
		time.Sleep(1 * time.Second)
		snap := snapService.Get().MustSend().MustSnapshot()
		if snap.MustSnapshotStatus() == ovirtsdk4.SNAPSHOTSTATUS_OK {
			break
		}
		fmt.Printf("Waiting till the snapshot is created, the satus is %v nes",
			snap.MustSnapshotStatus())
	}

	// Retrieve the descriptions of the disks of the snapshot:
	snapDisksService := snapService.DisksService()
	snapDisks := snapDisksService.List().MustSend()

	// Attach all the disks of the snapshot to the agent virtual machine, and
	// save the resulting disk attachments in a list so that we can later
	// detach them easily:
	attachmentsService := agentVmService.DiskAttachmentsService()
	attachments := make([]*ovirtsdk4.DiskAttachment, 0)
	for _, snapDisk := range snapDisks.MustDisks().Slice() {
		attachment := attachmentsService.Add().
			Attachment(
				ovirtsdk4.NewDiskAttachmentBuilder().
					Disk(
						ovirtsdk4.NewDiskBuilder().
							Id(snapDisk.MustId()).
							Snapshot(
								ovirtsdk4.NewSnapshotBuilder().
									Id(snap.MustSnapshot().MustId()).
									MustBuild()).
							MustBuild()).
					Active(true).
					Bootable(false).
					Interface(ovirtsdk4.DISKINTERFACE_VIRTIO).
					MustBuild()).
			MustSend()
		attachments = append(attachments, attachment.MustAttachment())
		fmt.Println("Attached disk %v to the agent vm.",
			attachment.MustAttachment().MustId())
	}

	// Now the disks are attached to the virtual agent virtual machine, we
	// can then ask that virtual machine to perform the backup. Doing that
	// requires a mechanism to talk to the backup software that runs inside the
	// agent virtual machine. That is outside of the scope of the SDK. But if
	// the guest agent is installed in the virtual machine then we can
	// provide useful information, like the identifiers of the disks that have
	// just been attached.
	for _, attachment := range attachments {
		if attachment.MustLogicalName() != "" {
			fmt.Printf("The logical name for disk %v is %v",
				attachment.MustDisk().MustId(), attachment.MustLogicalName())
		} else {
			fmt.Printf("The logical name for disk %v is not available. Is the guest agent installed?",
				attachment.MustDisk().MustId())
		}
	}
	//Insert here the code to contact the backup agent and do the actual
	// backup ...
	fmt.Printf("Doing the actual backup ...")

	// Detach the disks from the agent virtual machine:
	for _, attachment := range attachments {
		attachmentService := attachmentsService.AttachmentService(attachment.MustId())
		attachmentService.Remove().Send()
		fmt.Printf("Detached disk %v to from the agent virtual machine.",
			attachment.MustDisk().MustId())
	}

	snapService.Remove().Send()
	fmt.Printf("Remove the snapshot %v.", snap.MustSnapshot().MustDescription())

	// Send an external event to indicate to the administrator that the
	// backup of the virtual machine is completed:
	eventsService.Add().
		Event(
			ovirtsdk4.NewEventBuilder().
				Vm(
					ovirtsdk4.NewVmBuilder().
						Id(dataVm.MustId()).MustBuild()).
				Origin("mybackup").
				Severity(ovirtsdk4.LOGSEVERITY_NORMAL).
				CustomId(eventId).
				Description(desc).
				MustBuild()).
		Send()

	eventId += 1
}
