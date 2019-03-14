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

func changeVMCd() {
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

	// Locate the service that manages the CDROM devices of the virtual machine
	cdromsService := vmService.CdromsService()

	// Get the first CDROM
	getcdsResp, _ := cdromsService.List().Send()
	cdSlice, _ := getcdsResp.Cdroms()
	cdrom := cdSlice.Slice()[0]

	// Locate the service that manages the CDROM device found in the previous step
	cdromService := cdromsService.CdromService(cdrom.MustId())

	// Change the CDROM disk of the virtual machine to 'my_iso_file.iso'. By default the below operation changes
	// permanently the disk that will be visible to the virtual machine after the next boot, but it doesn't have
	// any effect on the currently running virtual machine. If you want to change the disk that is visible to the
	// current running virtual machine, change the value of the 'current' parameter to 'true'.
	cdromService.Update().
		Cdrom(
			ovirtsdk4.NewCdromBuilder().
				File(
					ovirtsdk4.NewFileBuilder().
						Id("my_iso_file.iso").
						MustBuild()).
				MustBuild()).
		Current(false).
		Send()
}
