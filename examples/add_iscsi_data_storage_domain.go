//
// Copyright (c) 2020 huihui <huihui.fu@cs2c.com.cn>.
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

func addStorageDomain() {
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

	// use add method to create storageDomain
	// Note that 'myhost' should be in a dc with version >= 4.1 for enabling
	// the 'Discard After Delete' (DAD) parameter, since only from that
	// version, the host returns the luns' discard related information.
	storageDomainService := conn.SystemService().StorageDomainsService()

	storageDomain, err := storageDomainService.Add().
		StorageDomain(
			ovirtsdk4.NewStorageDomainBuilder().
				Name("1.2.3.4").
				Description("My iSCSI With Discard After Delete").
				Type(ovirtsdk4.STORAGEDOMAINTYPE_DATA).
				DiscardAfterDelete(true).
				DataCenter(
					ovirtsdk4.NewDataCenterBuilder().
						Name("mydc").
						MustBuild()).
				Host(
					ovirtsdk4.NewHostBuilder().
						Name("myhost").
						MustBuild()).
				StorageFormat(ovirtsdk4.STORAGEFORMAT_V4).
				Storage(
					ovirtsdk4.NewHostStorageBuilder().
						Type(ovirtsdk4.STORAGETYPE_ISCSI).
						OverrideLuns(true).
						VolumeGroup(
							ovirtsdk4.NewVolumeGroupBuilder().
								LogicalUnitsOfAny(
									ovirtsdk4.NewLogicalUnitBuilder().
										Id("36001405e396b3d9c9a54284a51685f9a").
										Address("192.168.201.3").
										Port(int64(3260)).
										Target("iqn.2014-07.org.ovirt:storage").
										Username("username").
										Password("password").
										MustBuild()).
								MustBuild()).
						MustBuild()).
				MustBuild()).
		Send()
	if err != nil {
		fmt.Printf("Failed to add storageDomain, reason: %v\n", err)
		return
	}

	// Locate the service that manages the data centers and use it to
	// search for the data center:
	dataCentersService := conn.SystemService().DataCentersService()
	listResp, err := dataCentersService.List().Search("name=mydc").Send()
	if err != nil {
		fmt.Printf("Failed to search dataCenter list, reason: %v\n", err)
		return
	}
	dataCentersSlice, _ := listResp.DataCenters()
	dataCenter := dataCentersSlice.Slice()[0]

	// Locate the service that manages the data center where we want to
	// attach the storage domain:
	dataCenterService := dataCentersService.DataCenterService(dataCenter.MustId())

	// Locate the service that manages the storage domains that are attached
	// to the data centers:
	attachedSdsService := dataCenterService.StorageDomainsService()

	// Use the "add" method of service that manages the attached storage
	// domains to attach it.
	// Note that attaching this SD to the dc will succeed only if the
	// dc version is >= 4.1.
	_, err = attachedSdsService.Add().
		StorageDomain(
			ovirtsdk4.NewStorageDomainBuilder().
				Id(storageDomain.MustStorageDomain().MustId()).
				MustBuild()).
		Send()

	// Wait till the storage domain is active:
	attachedSdService := attachedSdsService.StorageDomainService(storageDomain.MustStorageDomain().MustId())
	for {
		getSdsResp, err := attachedSdService.Get().Send()
		if err != nil {
			continue
		}
		if getSd, ok := getSdsResp.StorageDomain(); ok {
			if getSd.MustStatus() == ovirtsdk4.STORAGEDOMAINSTATUS_ACTIVE {
				break
			}
		}
		time.Sleep(5 * time.Second)
	}
}
