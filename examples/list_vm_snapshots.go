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

func listVMSnapshots() {
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

	// Find all the virtual machines and store the id and name in a
	// map, so that looking them up later will be faster
	vmsService := conn.SystemService().VmsService()
	vmSlice := vmsService.List().MustSend().MustVms()
	vmsMap := make(map[string]string)
	for _, vm := range vmSlice.Slice() {
		vmsMap[vm.MustId()] = vm.MustName()
	}

	// Same for storage domains
	sdsService := conn.SystemService().StorageDomainsService()
	sdSlice := sdsService.List().MustSend().MustStorageDomains()
	sdsMap := make(map[string]string)
	for _, sd := range sdSlice.Slice() {
		sdsMap[sd.MustId()] = sd.MustName()
	}

	// For each virtual machine find its snapshots, then for each snapshot
	// find its disks
	for vmID, vmName := range vmsMap {
		vmService := vmsService.VmService(vmID)
		snapshotsService := vmService.SnapshotsService()
		snapshotsMap := make(map[string]string)
		for _, snapshot := range snapshotsService.List().MustSend().MustSnapshots().Slice() {
			snapshotsMap[snapshot.MustId()] = snapshot.MustDescription()
		}

		for snapshotID, snapshotDesc := range snapshotsMap {
			snapshotService := snapshotsService.SnapshotService(snapshotID)
			snapDisksService := snapshotService.DisksService()
			for _, disk := range snapDisksService.List().MustSend().MustDisks().Slice() {
				sdID := disk.MustStorageDomains().Slice()[0].MustId()
				sdName := sdsMap[sdID]
				fmt.Printf("%v:%v:%v:%v\n", vmName, snapshotDesc, disk.MustAlias(), sdName)
			}
		}
	}
}
