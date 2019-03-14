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

func addISODataStorageDomain() {
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

	// Get the reference to the storage domains service
	sdsService := conn.SystemService().StorageDomainsService()

	// Create a new ISO storage domain
	addSdResp, err := sdsService.Add().
		StorageDomain(
			ovirtsdk4.NewStorageDomainBuilder().
				Name("myiso").
				Description("My ISO").
				Type(ovirtsdk4.STORAGEDOMAINTYPE_ISO).
				Host(
					ovirtsdk4.NewHostBuilder().
						Name("myhost").
						MustBuild()).
				Storage(
					ovirtsdk4.NewHostStorageBuilder().
						Type(ovirtsdk4.STORAGETYPE_NFS).
						Address("server0.example.com").
						Path("/nfs/ovirt/40/myiso").
						MustBuild()).
				MustBuild()).
		Send()
	if err != nil {
		fmt.Printf("Faild to add sd, reason: %v\n", err)
		return
	}

	// Wait till the storage domain is unattached
	if sd, ok := addSdResp.StorageDomain(); ok {
		sdService := sdsService.StorageDomainService(sd.MustId())
		for {
			time.Sleep(5 * time.Second)
			getSdResp, _ := sdService.Get().Send()
			getSd, _ := getSdResp.StorageDomain()
			if getSd.MustStatus() == ovirtsdk4.STORAGEDOMAINSTATUS_UNATTACHED {
				break
			}
		}
	}
}
