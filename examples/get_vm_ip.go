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

func getVmIp() {
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

	// Get the reference to the vm service for a particular vm:
	vmsService := conn.SystemService().VmsService()

	//# Look up fot the vm by name:
	vmResp, err := vmsService.List().Search("name=myvm").Send()
	if err != nil {
		fmt.Printf("Failed to search vm list, reason: %v\n", err)
		return
	}
	vmSlice, _ := vmResp.Vms()
	vm := vmSlice.Slice()[0]

	// Get the reported-devices service for this vm:
	reportedDevicesService := vmsService.VmService(vm.MustId()).ReportedDevicesService()

	// Get the guest reported devices
	reportedDeviceResp, err := reportedDevicesService.List().Send()
	if err != nil {
		fmt.Printf("Failed to get reported devices list, reason: %v\n", err)
		return
	}
	reportedDeviceSlice, _ := reportedDeviceResp.ReportedDevice()
	for _, reportedDevice := range reportedDeviceSlice.Slice() {
		ips := reportedDevice.MustIps()
		if ips != nil {
			for _, ip := range ips.Slice() {
				fmt.Printf(" - %v", ip.MustAddress())
			}
		}
	}

}
