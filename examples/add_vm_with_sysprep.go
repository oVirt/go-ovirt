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

package main

import (
	"fmt"
	"time"

	"github.com/imjoey/go-ovirt"
)

func main() {
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

	// Find the service that manages the collection of virtual machines
	vmsService := conn.SystemService().VmsService()

	// Create the virtual machine. Note that no Sysprep stuff is needed here, when creating it, it will be used
	// later, when starting it
	addVMResp, _ := vmsService.Add().
		Vm(
			ovirtsdk4.NewVmBuilder().
				Name("myvm").
				Cluster(
					ovirtsdk4.NewClusterBuilder().
						Name("mycluster").
						MustBuild()).
				Template(
					ovirtsdk4.NewTemplateBuilder().
						Name("mytemplate").
						MustBuild()).
				MustBuild()).
		Send()

	vm, _ := addVMResp.Vm()

	// Find the service that manages the virtual machine
	vmService := vmsService.VmService(vm.MustId())

	// Wait till the virtual machine is down, which indicates that all the disks have been created
	for {
		if vm.MustStatus() == ovirtsdk4.VMSTATUS_DOWN {
			break
		}
		getResp, _ := vmService.Get().Send()
		vm, _ := getResp.Vm()
	}

	// The content of the Unattend.xml file. Note that this is an incomplete file, make sure to use a complete one,
	// maybe reading it from an external file:
	unattendXML :=
		"<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n" +
			"<unattend xmlns=\"urn:schemas-microsoft-com:unattend\">\n" +
			"  ...\n" +
			"</unattend>\n"

	// Start the virtual machine enabling Sysprep. Make sure to use a Windows operating system, either in the
	// template, or overriding it explicitly here. Without that the Sysprep logic won't be triggered.
	vmService.Start().
		UseSysprep(true).
		Vm(
			ovirtsdk4.NewVmBuilder().
				Os(
					ovirtsdk4.NewOperatingSystemBuilder().
						Type("windows_7x64").
						MustBuild()).
				MustBuild()).
		Send()
}
