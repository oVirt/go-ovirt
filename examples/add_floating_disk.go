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

func addFloatingDisk() {
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

	// Get the reference to the disks service
	disksService := conn.SystemService().DisksService()

	// Add the disk. Note that the size of the disk, the `provisionedSize` attribute, is specified in bytes,
	// so to create a disk of 10 GiB the value should be 10 * 2^30.
	addDiskResp, err := disksService.Add().
		Disk(
			ovirtsdk4.NewDiskBuilder().
				Name("mydisk").
				Description("My disk").
				Format(ovirtsdk4.DISKFORMAT_COW).
				ProvisionedSize(int64(10) * int64(1<<30)).
				StorageDomainsOfAny(
					ovirtsdk4.NewStorageDomainBuilder().
						Name("mydata").
						MustBuild()).
				MustBuild()).
		Send()
	if err != nil {
		fmt.Printf("Faild to add disk, reason: %v\n", err)
		return
	}

	// Get disk from response
	if disk, ok := addDiskResp.Disk(); ok {
		// Wait till the disk is completely created
		diskService := disksService.DiskService(disk.MustId())
		for {
			time.Sleep(5 * time.Second)
			getResp, err := diskService.Get().Send()
			if err != nil {
				fmt.Printf("Faild to get disk status, reason: %v\n", err)
				continue
			}
			// Check if the disk status is OK
			if d, ok := getResp.Disk(); ok {
				if d.MustStatus() == ovirtsdk4.DISKSTATUS_OK {
					break
				}
			}
		}
	}

}
