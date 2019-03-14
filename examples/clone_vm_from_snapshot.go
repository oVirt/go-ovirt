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

func cloneVMFromSnapshot() {
	inputRawURL := "https://10.1.111.229/ovirt-engine/api"
	// Create the connection to the server:
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

	// Locate the service that manages the virtual machines
	vmsService := conn.SystemService().VmsService()

	// Find the virtual machine
	vmsResp, _ := vmsService.List().Search("name=myvm").Send()
	vmsSlice, _ := vmsResp.Vms()
	vm := vmsSlice.Slice()[0]

	// Locate the service that manages the virtual machine
	vmService := vmsService.VmService(vm.MustId())

	// Find the snapshot. Note that the snapshots collection doesn't support search, so we need to retrieve
	// the complete list and then look for the snapshot that has the description that we are looking for
	snapsService := vmService.SnapshotsService()
	listSnapsResp, _ := snapsService.List().Send()
	snapSlice, _ := listSnapsResp.Snapshots()
	var snap *ovirtsdk4.Snapshot
	for _, sn := range snapSlice.Slice() {
		if sn.MustDescription() == "mysnap" {
			snap = sn
		}
	}

	// Create a new virtual machine, cloning it from the snapshot
	addVMResp, _ := vmsService.Add().
		Vm(
			ovirtsdk4.NewVmBuilder().
				Name("myclonedvm").
				SnapshotsOfAny(
					ovirtsdk4.NewSnapshotBuilder().
						Id(snap.MustId()).
						MustBuild()).
				Cluster(
					ovirtsdk4.NewClusterBuilder().
						Name("mycluster").
						MustBuild()).
				MustBuild()).
		Send()
	if clonedVM, ok := addVMResp.Vm(); ok {
		// Find the service that manages the cloned virtual machine
		clonedVMService := vmsService.VmService(clonedVM.MustId())

		// Wait till the virtual machine is down, as that means that the creation of the disks has been completed
		for {
			time.Sleep(5 * time.Second)
			getClonedResp, _ := clonedVMService.Get().Send()
			clonedVM, _ := getClonedResp.Vm()
			if clonedVM.MustStatus() == ovirtsdk4.VMSTATUS_DOWN {
				break
			}
		}
	}
}
