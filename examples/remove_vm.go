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

func removeVM() {
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

	// Find the service that manages VMs
	vmsService := conn.SystemService().VmsService()

	// Find the VM
	vm := vmsService.List().
		Search("name=myvm").
		MustSend().
		MustVms().
		Slice()[0]

	// Note that the "vm" variable that we assigned above contains only the data of the VM, it doesn't have any
	// method like "remove". Methods are defined in the services. So now that we have the description of the VM
	// we can find the service that manages it, calling the locator method "vmService" defined in the "vms"
	// service. This locator method receives as parameter the identifier of the VM and returns a reference to the
	// service that manages that VM.
	vmService := vmsService.VmService(vm.MustId())

	// Now that we have the reference to the service that manages the VM we can use it to remove the VM. Note that
	// this method doesn't need any parameter, as the identifier of the VM is already known by the service that we
	// located in the previous step.
	vmService.Remove().MustSend()

}
