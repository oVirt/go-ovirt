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

func startVM() {
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

	// Get the reference to the "vms" service
	vmsService := conn.SystemService().VmsService()

	// Find the virtual machine
	vm := vmsService.List().Search("name=myvm").MustSend().MustVms().Slice()[0]

	// Locate the service that manages the virtual machine, as that is where the action methods are defined
	vmService := vmsService.VmService(vm.MustId())

	// Call the "start" method of the service to start it
	vmService.Start().MustSend()

	// Wait till the virtual machine is up
	for {
		time.Sleep(5 * time.Second)
		vm = vmService.Get().MustSend().MustVm()
		if vm.MustStatus() == ovirtsdk4.VMSTATUS_UP {
			break
		}
	}
}
