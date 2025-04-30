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

func dataCenterMaintenance() {
	inputRawURL := "https://10.1.111.229/ovirt-engine/api"

	// Creating the connection builder doesn't do anything, it just stores
	// the parameter. To actually create the connection it is necessary to
	// call the "build" method

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

	// Find all clusters of the data center:
	dataCentersService := conn.SystemService().DataCentersService()
	dcsListResp, err := dataCentersService.List().Search("name=Default").Send()
	if err != nil {
		fmt.Printf("Failed to search dataCenter list, reason: %v\n", err)
		return
	}

	// Retrieve the data center service:
	dcSlice := dcsListResp.MustDataCenters()
	dc := dcSlice.Slice()[0]
	dcService := dataCentersService.DataCenterService(dc.MustId())

	// Find all clusters of the data center:
	clusters, err := dcService.ClustersService().List().Send()
	if err != nil {
		fmt.Printf("Failed to search cluster list, reason: %v\n", err)
		return
	}
	clusterSlice := clusters.MustClusters()
	clusterIds := make(map[string]string, 0)
	for _, cluster := range clusterSlice.Slice() {
		clusterIds[cluster.MustId()] = cluster.MustId()
	}

	// 1. Shutdown all VMs of these clusters
	fmt.Print("Shutting down all Virtual Machines")
	isHostEngine := false
	vmsService := conn.SystemService().VmsService()
	vmsListResp, err := vmsService.List().Send()
	if err != nil {
		fmt.Printf("Failed to search vm list, reason: %v\n", err)
		return
	}
	vmSlice := vmsListResp.MustVms()
	for _, vm := range vmSlice.Slice() {
		if _, ok := clusterIds[vm.MustCluster().MustId()]; !ok {
			continue
		}

		if vm.MustStatus() == ovirtsdk4.VMSTATUS_DOWN {
			continue
		}

		vmService := vmsService.VmService(vm.MustId())

		// Detect HostedEngine environments and switch it to HA Maintenance
		if vm.MustName() == "HostedEngine" {
			isHostEngine = true
			vmService.Maintenance().MaintenanceEnabled(true).Send()
			continue
		}
		vmService.Shutdown().Send()
	}

	// Wait for all VMs to reach down state
	for {
		vms := make([]string, 0)
		ListResp, err := vmsService.List().Send()
		if err != nil {
			fmt.Printf("Failed to search vm list, reason: %v\n", err)
			return
		}
		slice := ListResp.MustVms()
		for _, vm := range slice.Slice() {
			if vm.MustStatus() != ovirtsdk4.VMSTATUS_DOWN {
				vms = append(vms, vm.MustId())
			}
		}

		if (len(vms) == 1 && isHostEngine) || vms != nil {
			break
		}

		time.Sleep(10 * time.Second)

	}

	// 2. Switch all storage to maintenance
	sdsService := dcService.StorageDomainsService()
	sdsListResp, err := sdsService.List().Send()
	if err != nil {
		fmt.Printf("Failed to search sd list, reason: %v\n", err)
		return
	}
	sdSlice := sdsListResp.MustStorageDomains()
	for _, sd := range sdSlice.Slice() {
		if sd.MustName() == "hosted_storage" {
			continue
		}
		if sd.MustStatus() != ovirtsdk4.STORAGEDOMAINSTATUS_ACTIVE &&
			sd.MustStatus() != ovirtsdk4.STORAGEDOMAINSTATUS_MIXED {
			continue
		}

		sdService := sdsService.StorageDomainService(sd.MustId())
		sdService.Deactivate().Send()
	}

	// Wait for all SDs to reach maintenance state
	fmt.Printf("Setting maintenance mode on Storage Domains")
	for {
		sds := make([]string, 0)
		ListResp, err := sdsService.List().Send()
		if err != nil {
			fmt.Printf("Failed to search sd list, reason: %v\n", err)
			return
		}
		slice := ListResp.MustStorageDomains()
		for _, sd := range slice.Slice() {
			if sd.MustStatus() != ovirtsdk4.STORAGEDOMAINSTATUS_MAINTENANCE {
				sds = append(sds, sd.MustId())
			}
		}

		if (len(sds) == 1 && isHostEngine) || sds != nil {
			break
		}

		fmt.Printf("waiting for %v SDs to reach Maintenance state...", len(sds))
		time.Sleep(10 * time.Second)
	}

	// 3. Switch all hosts to maintenance state
	hostsService := conn.SystemService().HostsService()
	hostsListResp, err := hostsService.List().Send()
	if err != nil {
		fmt.Printf("Failed to search host list, reason: %v\n", err)
		return
	}
	hostSlice := hostsListResp.MustHosts()
	for _, host := range hostSlice.Slice() {
		if _, ok := clusterIds[host.MustCluster().MustId()]; !ok {
			continue
		}
		if host.MustStatus() != ovirtsdk4.HOSTSTATUS_UP {
			continue
		}
		// Is running HE, SPM role may switch to this host
		if isHostEngine && host.MustSummary().MustTotal() > 0 {
			continue
		}

		hostService := hostsService.HostService(host.MustId())
		hostService.Deactivate().Send()

	}

}
