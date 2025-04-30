//
// The oVirt Project - oVirt Engine Go SDK
//
// Copyright (c) oVirt Authors
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

func sparsifyDisk() {
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

	// Get the reference to the "disks" service:
	disksService := conn.SystemService().DisksService()

	// Find the virtual machine:
	disksResp, err := disksService.List().Search("name=mydisk").Send()
	if err != nil {
		fmt.Printf("Failed to get vm list, reason: %v\n", err)
		return
	}

	//# Find the disk we want to sparsify:
	//# Locate the service that manage the disk we want to sparsify:
	disk := disksResp.MustDisks().Slice()[0]
	diskService := disksService.DiskService(disk.MustId())

	// Sparsify the disk. Note that the virtual machine where the disk is attached
	// must not be running so the sparsification is executed. If the virtual
	// machine will be running the sparsify operation will fail:
	diskService.Sparsify().Send()

	// Wait till the disk is ok:
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
