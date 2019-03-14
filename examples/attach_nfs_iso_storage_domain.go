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

func attachISODataStorageDomain() {
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

	// Locate the service that manages the storage domains, and use it to search for the storage domain
	sdsService := conn.SystemService().StorageDomainsService()
	sd := sdsService.List().
		Search("name=myiso").
		MustSend().
		MustStorageDomains().
		Slice()[0]

	// Locate the service that manages the data centers and use it to search for the data center
	dcsService := conn.SystemService().DataCentersService()
	dc := dcsService.List().
		Search("name=mydc").
		MustSend().
		MustDataCenters().
		Slice()[0]

	// Locate the service that manages the data center where we want to attach the storage domain
	dcService := dcsService.DataCenterService(dc.MustId())

	// Locate the service that manages the storage domains that are attached to the data center
	attachedSdsService := dcService.StorageDomainsService()

	// Use the "add" method of the service that manages the attached storage domains to attach it
	attachedSdsService.Add().
		StorageDomain(
			ovirtsdk4.NewStorageDomainBuilder().
				Id(sd.MustId()).
				MustBuild()).
		MustSend()

	// Wait till the storage domain is active
	attachedSdService := attachedSdsService.StorageDomainService(sd.MustId())
	for {
		time.Sleep(5 * time.Second)
		sd = attachedSdService.Get().MustSend().MustStorageDomain()
		if sd.MustStatus() == ovirtsdk4.STORAGEDOMAINSTATUS_ACTIVE {
			break
		}
	}
}
