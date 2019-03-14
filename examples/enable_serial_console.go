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

func enableSerialConsole() {
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

	// Find the virtual machine. Note the use of the  `all_content` parameter, it is required in order to obtain
	// additional information that isn't retrieved by default, like the configuration of the serial console.
	vmsResp, _ := vmsService.List().
		Search("name=myvm").
		AllContent(true).
		Send()
	vmsSlice, _ := vmsResp.Vms()
	vm := vmsSlice.Slice()[0]

	// Check if the serial console is enabled, and if it isn't then update the virtual machine to enable it
	if !vm.MustConsole().MustEnabled() {
		vmService := vmsService.VmService(vm.MustId())
		vmService.Update().
			Vm(
				ovirtsdk4.NewVmBuilder().
					Console(
						ovirtsdk4.NewConsoleBuilder().
							Enabled(true).
							MustBuild()).
					MustBuild()).
			Send()
	}

}
