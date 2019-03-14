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

func addVMNic() {
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

	// Locate the virtual machines service and use it to find the virtual machine
	vmsService := conn.SystemService().VmsService()
	vmsResp, _ := vmsService.List().Search("name=myvm").Send()
	vmsSlice, _ := vmsResp.Vms()
	vm := vmsSlice.Slice()[0]

	// In order to specify the network that the new interface will be connected to, we need to specify the
	// identifier of the virtual network interface profile, so we need to find it
	profilesService := conn.SystemService().VnicProfilesService()
	var profileID string
	pfsResp, _ := profilesService.List().Send()
	pfSlice, _ := pfsResp.Profiles()
	for _, pf := range pfSlice.Slice() {
		if pf.MustName() == "mynetwork" {
			profileID = pf.MustId()
			break
		}
	}

	// Locate the service that manages the NICs of the virtual machine
	nicsService := vmsService.VmService(vm.MustId()).NicsService()
	nicsService.Add().
		Nic(
			ovirtsdk4.NewNicBuilder().
				Name("mynic").
				Description("My network interface card").
				VnicProfile(
					ovirtsdk4.NewVnicProfileBuilder().
						Id(profileID).
						MustBuild()).
				MustBuild()).
		Send()
}
