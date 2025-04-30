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

func asynchronousInventory() {
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

	// Send requests for all the collections, but don't wait for the results. This
	// way the requests will be sent simultaneously, using the multipl connections
	systemService := conn.SystemService()
	fmt.Printf("Requesting data...")
	dcsListFuture, _ := systemService.DataCentersService().List().Send()
	clustersListFuture, _ := systemService.ClustersService().List().Send()
	sdsListFuture, _ := systemService.StorageDomainsService().List().Send()
	networksListFuture, _ := systemService.NetworksService().List().Send()
	hostsListFuture, _ := systemService.HostsService().List().Send()
	vmsListFuture, _ := systemService.VmsService().List().Send()
	disksListFuture, _ := systemService.DisksService().List().Send()

	//# Wait for the results of the requests that we sent. The calls to the `wait`
	//# method will perform all the work, for all the pending requests, and will
	//# eventually return the requested data.
	fmt.Printf("Waiting for data centers ...")
	dcs := dcsListFuture.MustDataCenters()
	fmt.Printf("Loaded %v data centers.", len(dcs.Slice()))

	fmt.Printf("Waiting for clusters...")
	clusters := clustersListFuture.MustClusters()
	fmt.Printf("Loaded %v clusters.", len(clusters.Slice()))

	fmt.Printf("Waiting for storage domains...")
	sds := sdsListFuture.MustStorageDomains()
	fmt.Printf("Loaded %v storage domains.", len(sds.Slice()))

	fmt.Printf("Waiting for networks ...")
	nets := networksListFuture.MustNetworks()
	fmt.Printf("Loaded %v networks.", len(nets.Slice()))

	fmt.Printf("Waiting for hosts ...")
	hosts := hostsListFuture.MustHosts()
	fmt.Printf("Loaded %v hosts.", len(hosts.Slice()))

	fmt.Printf("Waiting for VMs ...")
	vms := vmsListFuture.MustVms()
	fmt.Printf("Loaded %v VMs.", len(vms.Slice()))

	fmt.Printf("Waiting for disks ...")
	disks := disksListFuture.MustDisks()
	fmt.Printf("Loaded %v disks.", len(disks.Slice()))
}
