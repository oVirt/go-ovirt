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

func assignNetworkToCluster() {
	inputRawURL := "https://10.1.111.229/ovirt-engine/api"
	// Create the connection to the server
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

	// Locate the networks service and use it to find the network
	networksService := conn.SystemService().NetworksService()
	getNSResp, _ := networksService.List().
		Search("name=mynetwork and datacenter=mydatacenter").
		Send()
	nsSlice, _ := getNSResp.Networks()
	network := nsSlice.Slice()[0]

	// Locate the clusters service and use it to find the cluster
	clustersService := conn.SystemService().ClustersService()
	getClustersResp, _ := clustersService.List().Search("name=mycluster").Send()
	clustersSlice, _ := getClustersResp.Clusters()
	cluster := clustersSlice.Slice()[0]

	// Locate the service that manages the networks of the cluster
	cluService := clustersService.ClusterService(cluster.MustId())
	assignedNSService := cluService.NetworksService()

	// Use the "add" method to assign network to cluster
	assignedNSService.Add().
		Network(
			ovirtsdk4.NewNetworkBuilder().
				Id(network.MustId()).
				Required(true).
				MustBuild()).
		Send()
}
