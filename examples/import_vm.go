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

func importVM() {
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

	// Get storage domains service
	sdsService := conn.SystemService().StorageDomainsService()

	// Get export storage domain
	exportDomain := sdsService.List().
		Search("name=myexport").
		MustSend().
		MustStorageDomains().
		Slice()[0]

	// Get target storage domain
	targetDomain := sdsService.List().
		Search("name=mydata").
		MustSend().
		MustStorageDomains().
		Slice()[0]

	// Get cluster service
	clustersService := conn.SystemService().ClustersService()

	// Get the cluster we import the VM to
	cluster := clustersService.List().
		Search("name=mycluster").
		MustSend().
		MustClusters().
		Slice()[0]

	// Get VM service for export storage domain
	vmsService := sdsService.StorageDomainService(exportDomain.MustId()).VmsService()

	// Get the first exported VM, assuming we have one
	exportedVM := vmsService.List().
		MustSend().
		MustVm().
		Slice()[0]

	// Import the exported VM into target storage domain, 'mydata'
	vmsService.VmService(exportedVM.MustId()).
		Import().
		StorageDomain(
			ovirtsdk4.NewStorageDomainBuilder().
				Id(targetDomain.MustId()).
				MustBuild()).
		Cluster(
			ovirtsdk4.NewClusterBuilder().
				Id(cluster.MustId()).
				MustBuild()).
		Vm(
			ovirtsdk4.NewVmBuilder().
				Id(exportedVM.MustId()).
				MustBuild()).
		MustSend()
}
