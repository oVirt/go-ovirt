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

func assignAffinityLabelToVM() {
	inputRawURL := "https://10.1.111.229/ovirt-engine/api"
	// Create the connection to the server
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

	// Get the reference to the "vms" service
	vmsService := conn.SystemService().VmsService()

	// Use the "List" method of the "vms" service to search the virtual machines
	vmsResponse, err := vmsService.List().Search("name=test4joey").CaseSensitive(false).Send()

	if err != nil {
		fmt.Printf("Failed to get vm list, reason: %v\n", err)
		return
	}
	// Find the vm:
	vm := vmsResponse.MustVms().Slice()[0]

	// Get the reference to the affinity labels service
	affinityLabelsService := conn.SystemService().AffinityLabelsService()

	// Find the id of the affinity label
	var affinityLabelID string
	resp, err := affinityLabelsService.List().Send()
	if err != nil {
		fmt.Printf("Failed to get affinity label list, reason: %v\n", err)
		return
	}
	labels, _ := resp.Labels()
	for _, label := range labels.Slice() {
		if label.MustName() == "myaffinitylabel" {
			affinityLabelID = label.MustId()
			break
		}
	}
	if affinityLabelID == "" {
		fmt.Println("Affinity label of name (myaffinitylabel) not exists")
		return
	}
	// Locate the service that manages the affinity label named `myaffinitylabel`
	labelService := affinityLabelsService.LabelService(affinityLabelID)

	// Get the reference to the service that manages the set of virtual machines that have the affinity label
	// named `myaffinitylabel` assigned
	labelVmsService := labelService.VmsService()

	// Assign affinity label to virtual machine:
	_, err = labelVmsService.Add().
		Vm(ovirtsdk4.NewVmBuilder().
			Id(vm.MustId()).
			MustBuild()).
		Send()

	if err != nil {
		fmt.Printf("Failed to assign affinity label to vm, reason: %v\n", err)
		return
	}
	fmt.Printf("Assign affinity label (myaffinitylabel) to vm (%v) successfully\n", vm.MustId())
}
