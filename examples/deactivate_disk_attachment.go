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
	"time"

	ovirtsdk4 "github.com/ovirt/go-ovirt/v4"
)

func deactivateDiskAttachment() {
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

	// Locate the virtual machines service and use it to find the virtual machine
	vmsService := conn.SystemService().VmsService()
	vmsResp, err := vmsService.List().Search("name=myvm").Send()
	if err != nil {
		fmt.Printf("Failed to search vm list, reason: %v\n", err)
		return
	}
	vmSlice, _ := vmsResp.Vms()
	vm := vmSlice.Slice()[0]

	// Get disks service and find the disk we want:
	disksService := conn.SystemService().DisksService()
	disksResp, err := disksService.List().Search("name=mydisk").Send()
	if err != nil {
		fmt.Printf("Failed to search disk list, reason: %v\n", err)
		return
	}
	diskSlice, _ := disksResp.Disks()
	disk := diskSlice.Slice()[0]

	// Locate the service that manages the disk attachments of the virtual machine
	dasService := vmsService.VmService(vm.MustId()).DiskAttachmentsService()
	diskAttachments, err := dasService.List().Send()

	diskAttachmentInVm := &ovirtsdk4.DiskAttachment{}

	for _, diskAttachment := range diskAttachments.MustAttachments().Slice() {
		if diskAttachment.MustDisk().MustId() == disk.MustId() {
			diskAttachmentInVm = diskAttachment
			break
		}
	}

	// Deactivate the disk we found
	// or print an error if there is no such disk attached:
	if diskAttachmentInVm != nil {
		// Locate the service that manages the disk attachment that we found
		// in the previous step:
		diskAttachmentService := dasService.AttachmentService(diskAttachmentInVm.MustId())

		// Deactivate the disk attachment
		diskAttachmentService.Update().
			DiskAttachment(
				ovirtsdk4.NewDiskAttachmentBuilder().
					Active(false).
					MustBuild()).
			Send()

		// Wait till the disk attachment not active:
		for {
			time.Sleep(5 * time.Second)
			diskAttachmentResp, _ := diskAttachmentService.Get().Send()
			if diskAttachment, ok := diskAttachmentResp.Attachment(); ok {
				if !diskAttachment.MustActive() {
					break
				}
			}
		}

	} else {
		fmt.Printf("There's no disk attachment for %v.", disk.MustName())
	}

}
