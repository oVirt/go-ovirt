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

func addMacPool() {
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

	// Get the reference to the service that manages the MAC address pools
	poolsService := conn.SystemService().MacPoolsService()

	// Add a new MAC address pool
	addPoolResp, err := poolsService.Add().
		Pool(
			ovirtsdk4.NewMacPoolBuilder().
				Name("mymacpool").
				RangesOfAny(
					ovirtsdk4.NewRangeBuilder().
						From("02:00:00:00:00:00").
						To("02:00:00:01:00:00").
						MustBuild()).
				MustBuild()).
		Send()
	if err != nil {
		fmt.Printf("Faild to add mac pool, reason: %v\n", err)
		return
	}
	if pool, ok := addPoolResp.Pool(); ok {
		// Find the service that manages clusters, as we need it in order to find the cluster where we wnt to set the
		// MAC address pool
		clustersService := conn.SystemService().ClustersService()

		// Find the cluster
		searchClusterResp, _ := clustersService.List().Search("name=mycluster").Send()
		clustersSlice, _ := searchClusterResp.Clusters()
		cluster := clustersSlice.Slice()[0]

		// Find the service that manages the cluster, as we need it in order to do the update
		clusterService := clustersService.ClusterService(cluster.MustId())

		// Update the service so that it uses the new MAC pool
		clusterService.Update().
			Cluster(
				ovirtsdk4.NewClusterBuilder().
					MacPool(
						ovirtsdk4.NewMacPoolBuilder().
							Id(pool.MustId()).
							MustBuild()).
					MustBuild()).
			Send()
	}
}
