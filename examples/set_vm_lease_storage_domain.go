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

func setVMLeaseStorageDomain() {
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

	// Get the reference to the root of the tree of services
	systemService := conn.SystemService()

	// Find the virtual machine
	vmsService := systemService.VmsService()
	vm := vmsService.List().
		Search("name=myvm").
		MustSend().
		MustVms().
		Slice()[0]

	// Find the storage domain
	sdsService := systemService.StorageDomainsService()
	sd := sdsService.List().
		Search("name=mydata").
		MustSend().
		MustStorageDomains().
		Slice()[0]

	// Update the virtual machine so that high availability is enabled and the lease is created in the selected
	// storage domain
	vmService := vmsService.VmService(vm.MustId())
	vmService.Update().
		Vm(
			ovirtsdk4.NewVmBuilder().
				HighAvailability(
					ovirtsdk4.NewHighAvailabilityBuilder().
						Enabled(true).
						MustBuild()).
				Lease(
					ovirtsdk4.NewStorageDomainLeaseBuilder().
						StorageDomain(
							ovirtsdk4.NewStorageDomainBuilder().
								Id(sd.MustId()).
								MustBuild()).
						MustBuild()).
				MustBuild()).
		MustSend()
}
