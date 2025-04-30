//
// Copyright (c) 2020 huihui0311 <huihui.fu@cs2c.com.cn>.
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

func addCluster() {
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

	// Get the reference to the clusters service:
	clustersService := conn.SystemService().ClustersService()

	// Use the "add" method to create a cluster:
	_, err = clustersService.Add().
		Cluster(
			ovirtsdk4.NewClusterBuilder().
				Name("mycluster").
				Description("My cluster").
				Cpu(
					ovirtsdk4.NewCpuBuilder().
						Architecture(ovirtsdk4.ARCHITECTURE_X86_64).
						Type("Intel Conroe Family").
						MustBuild()).
				DataCenter(
					ovirtsdk4.NewDataCenterBuilder().
						Name("mydc").
						MustBuild()).
				MustBuild()).
		Send()
	if err != nil {
		fmt.Printf("Failed to add cluster, reason: %v\n", err)
		return
	}
}
