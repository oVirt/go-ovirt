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

func registerVM() {
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

	// Get the reference to the "StorageDomains" service
	sdsService := conn.SystemService().StorageDomainsService()

	// Find the storage domain with unregistered VM
	sd := sdsService.List().Search("name=mysd").MustSend().MustStorageDomains().Slice()[0]

	// Locate the service that manages the storage domain, as that is where the action methods are defined
	sdService := sdsService.StorageDomainService(sd.MustId())

	// Locate the service that manages the VMs in storage domain
	sdVmsService := sdService.VmsService()

	// Find the unregistered VM we want to register
	unregVMSlice := sdVmsService.List().Unregistered(true).MustSend().MustVm()
	var vm *ovirtsdk4.Vm
	for _, v := range unregVMSlice.Slice() {
		if v.MustName() == "myvm" {
			vm = v
			break
		}
	}

	// Locate the service that manages virtual machine in the storage domain, as that is where the action methods
	// are defined
	sdVMService := sdVmsService.VmService(vm.MustId())

	// Register the VM into the system
	sdVMService.Register().
		Vm(
			ovirtsdk4.NewVmBuilder().
				Name("exported_myvm").
				MustBuild()).
		Cluster(
			ovirtsdk4.NewClusterBuilder().
				Name("mycluster").
				MustBuild()).
		VnicProfileMappingsOfAny(
			ovirtsdk4.NewVnicProfileMappingBuilder().
				SourceNetworkName("mynetwork").
				SourceNetworkProfileName("mynetwork").
				TargetVnicProfile(
					ovirtsdk4.NewVnicProfileBuilder().
						Name("mynetwork").
						MustBuild()).
				MustBuild()).
		ReassignBadMacs(true).
		MustSend()
}
