//
// Copyright (c) 2017 Joey <majunjiev@gmail.com>.
// Copyright (c) 2020 Douglas Schilling Landgraf <dougsland@redhat.com>
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
	ovirtsdk4 "github.com/ovirt/go-ovirt/v4"
	"time"
)

func listClusterFromDatacenter() {
	inputRawURL := "https://foobar.mydomain.home/ovirt-engine/api"

	conn, err := ovirtsdk4.NewConnectionBuilder().
		URL(inputRawURL).
		Username("admin@internal").
		Password("mysuperpass").
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
	dcsListResp, err := dataCentersService.List().Search("status=up").Send()
	if err != nil {
		fmt.Printf("Failed to search dataCenter list, reason: %v\n", err)
		return
	}
	dcs := dcsListResp.MustDataCenters()
	for _, dc := range dcs.Slice() {
		dcService := dataCentersService.DataCenterService(dc.MustId())
		fmt.Println("Datacenter:", dc.MustName())
		clusters, err := dcService.ClustersService().List().Send()
		if err != nil {
			fmt.Printf("Failed to search dataCenter list, reason: %v\n", err)
			return
		}
		clusterSlice := clusters.MustClusters()
		clusterIds := make(map[string]string, 0)
		for _, cluster := range clusterSlice.Slice() {
			clusterIds[cluster.MustId()] = cluster.MustId()
			fmt.Println("\tCluster:", cluster.MustName())
		}
	}
}
