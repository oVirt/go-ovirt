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

func listAffinityLabels() {
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

	affinityLableService := conn.SystemService().AffinityLabelsService()

	resp, err := affinityLableService.List().Send()

	if err != nil {
		fmt.Printf("Failed to get affinity label list, reason: %v\n", err)
		return
	}
	labels, _ := resp.Labels()
	for _, label := range labels.Slice() {
		fmt.Printf("Affinity labels (")
		if name, ok := label.Name(); ok {
			fmt.Printf(" name: %v", name)
		}
		if desc, ok := label.Description(); ok {
			fmt.Printf(" desc: %v", desc)
		}
		fmt.Println(")")

		// Print all affinity labels names and virtual machines
		// which has assigned that affinity label
		if vmSlice, ok := label.Vms(); ok {
			vms, err := conn.FollowLink(vmSlice)
			if err != nil {
				if href, ok := vmSlice.Href(); ok {
					fmt.Printf("Failed to follow vms link: %v, reason: %v\n", href, err)
					return
				}
			}
			if vms, ok := vms.(*ovirtsdk4.VmSlice); ok {
				for _, vmLink := range vms.Slice() {
					vm, err := conn.FollowLink(vmLink)
					if err != nil {
						fmt.Printf("Failed to follow vm link: %v, reason: %v\n", vmLink.MustHref(), err)
						return
					}
					if vm, ok := vm.(*ovirtsdk4.Vm); ok {
						fmt.Printf("**** attached VM (%+v)\n", vm.MustName())
					}
				}
			}
		}
	}

}
