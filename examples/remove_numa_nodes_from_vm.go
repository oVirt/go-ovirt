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
	"sort"
	"time"

	ovirtsdk4 "github.com/ovirt/go-ovirt/v4"
)

func removeNumaNodesFromVm() {
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

	// Get the reference to the vm service:
	vmsService := conn.SystemService().VmsService()

	// Retrieve the description of the virtual machine:
	vmsResp, err := vmsService.List().Search("name=myvm").Send()
	if err != nil {
		fmt.Printf("Failed to get vm list, reason: %v\n", err)
		return
	}
	vm := vmsResp.MustVms().Slice()[0]

	//In order to update the virtual machine we need a reference to the service
	// the manages it:
	vmService := vmsService.VmService(vm.MustId())

	// Get the reference to the numa nodes service of the virtual machine:
	vmNodesService := vmService.NumaNodesService()

	// For each numa node of the virtual machine we call the remove method
	// to remove the numa node. Note that we must remove the numa nodes from
	// highest numa node index to lowest numa node index:
	vmNodeResp, err := vmNodesService.List().Send()
	vmNodes := vmNodeResp.MustNodes().Slice()
	sort.Slice(vmNodes, func(i, j int) bool {
		return vmNodes[i].MustIndex() > vmNodes[j].MustIndex() // reverse=True
	})

	for _, vmNode := range vmNodes {
		fmt.Printf("Removing node with id %v\n", vmNode.MustId())
		vmNodeService := vmNodesService.NodeService(vmNode.MustId())
		vmNodeService.Remove().Send()
	}

}
