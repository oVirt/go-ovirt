//
// Copyright (c) 2017 Joey <majunjiev@gmail.com>.
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

	ovirtsdk4 "github.com/ovirt/go-ovirt"
)

func addLunDiskToVM() {
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
	vmResp, _ := vmsService.List().Search("name=myvm").Send()
	vmSlice, _ := vmResp.Vms()
	vm := vmSlice.Slice()[0]

	// Locate the service that manages the disk attachments of the virtual machine
	diskAttachmentsService := vmsService.VmService(vm.MustId()).DiskAttachmentsService()

	// Use the "add" method of the disk attachments service to add the LUN disk.
	diskAttachmentsService.Add().
		Attachment(
			ovirtsdk4.NewDiskAttachmentBuilder().
				Disk(
					ovirtsdk4.NewDiskBuilder().
						Name("myiscsidisk").
						LunStorage(
							ovirtsdk4.NewHostStorageBuilder().
								Type(ovirtsdk4.STORAGETYPE_ISCSI).
								LogicalUnitsOfAny(
									ovirtsdk4.NewLogicalUnitBuilder().
										Address("192.168.200.3").
										Port(3260).
										Target("iqn.2014-07.org.ovirt:storage").
										Id("36001405e793bf9c57a840f58c9a8a220").
										Username("username").
										Password("password").
										MustBuild()).
								MustBuild()).
						MustBuild()).
				Interface(ovirtsdk4.DISKINTERFACE_VIRTIO).
				Bootable(false).
				Active(true).
				MustBuild()).
		Send()

}
