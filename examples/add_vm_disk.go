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

func addVMDisk() {
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

	// Locate the service that manages the disk attachments of the virtual machine
	dasService := vmsService.VmService(vm.MustId()).DiskAttachmentsService()

	// Use the `Add` method of the disk attachments service to add the disk. Note that the size of the disk,
	// the `provionedSize` attribute, is specified in bytes, so to create a disk of 10 GiB the value should
	// be 10 * 2^30.
	resp, err := dasService.Add().
		Attachment(
			ovirtsdk4.NewDiskAttachmentBuilder().
				Disk(
					ovirtsdk4.NewDiskBuilder().
						Name("mydisk").
						Description("My disk description").
						Format(ovirtsdk4.DISKFORMAT_COW).
						ProvisionedSize(int64(10) * int64(1<<30)).
						StorageDomainsOfAny(
							ovirtsdk4.NewStorageDomainBuilder().
								Name("mydata").
								MustBuild()).
						MustBuild()).
				Interface(ovirtsdk4.DISKINTERFACE_VIRTIO).
				Bootable(false).
				Active(true).
				MustBuild()).
		Send()
	if err != nil {
		fmt.Printf("Faild to add disk attachment to vm, reason: %v\n", err)
		return
	}

	if da, ok := resp.Attachment(); ok {
		// Wait till the disk is OK
		disksService := conn.SystemService().DisksService()
		diskService := disksService.DiskService(da.MustDisk().MustId())
		for {
			time.Sleep(5 * time.Second)
			diskResp, _ := diskService.Get().Send()
			if disk, ok := diskResp.Disk(); ok {
				if disk.MustStatus() == ovirtsdk4.DISKSTATUS_OK {
					break
				}
			}
		}
	}

}
