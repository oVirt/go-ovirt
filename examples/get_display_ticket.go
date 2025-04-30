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

func getDisplayTicket() {
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

	// Get a reference to the vm service:
	vmsService := conn.SystemService().VmsService()

	//# Look up fot the vm by name:
	vmResp, err := vmsService.List().Search("name=myvm").Send()
	if err != nil {
		fmt.Printf("Failed to search vm list, reason: %v\n", err)
		return
	}
	vmSlice, _ := vmResp.Vms()
	vm := vmSlice.Slice()[0]

	consolesService := vmsService.VmService(vm.MustId()).GraphicsConsolesService()

	// The method that lists the graphics consoles doesn't support search, so in
	// order to find the console corresponding to the access protocol that we are
	// interested on (SPICE in this example) we need to get all of them and filter
	// explicitly. In addition the `current` parameter must be `True`, as otherwise
	// you will *not* get important values like the `address` and `port` where the
	// console is available.
	consolesResp, err := consolesService.List().Send()
	if err != nil {
		fmt.Printf("Failed to search console list, reason: %v\n", err)
		return
	}
	consoleSlice, _ := consolesResp.Consoles()
	newConsole := &ovirtsdk4.GraphicsConsole{}
	for _, console := range consoleSlice.Slice() {
		if console.MustProtocol() == ovirtsdk4.GRAPHICSTYPE_SPICE {
			newConsole = console
		}
	}

	// Find the service that manages the graphics console that was selected in the
	// previous step:
	consoleService := consolesService.ConsoleService(newConsole.MustId())

	// Request the ticket. The virtual machine must be up and running, as it
	// doesn't make sense to get a console ticket for a virtual machine that
	// is down. If you try that, the request will fail.
	ticket := consoleService.Ticket().MustSend().MustTicket()

	// Print the details needed to connect to the console (the ticket value is the
	// password):
	fmt.Printf("address: %v\n", newConsole.MustAddress())
	fmt.Printf("port: %v\n", newConsole.Port())
	fmt.Printf("tls_port: %v\n", newConsole.TlsPort())
	fmt.Printf("password: %v\n", ticket.MustValue())

}
