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

func addHost() {
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

	// Get the reference to the hosts service
	hostsService := conn.SystemService().HostsService()

	// Add the host
	addHostResp, err := hostsService.Add().
		Host(
			ovirtsdk4.NewHostBuilder().
				Name("myhost").
				Description("My host").
				Address("node40.example.com").
				RootPassword("123456").
				Cluster(
					ovirtsdk4.NewClusterBuilder().
						Name("mycluster").
						MustBuild()).
				MustBuild()).
		Send()
	if err != nil {
		fmt.Printf("Failed to add host, reason: %v\n", err)
		return
	}
	if host, ok := addHostResp.Host(); ok {
		hostService := hostsService.HostService(host.MustId())
		// Wait till the host is up
		for {
			time.Sleep(5 * time.Second)
			getHostResp, err := hostService.Get().Send()
			if err != nil {
				continue
			}
			if getHost, ok := getHostResp.Host(); ok {
				if getHost.MustStatus() == ovirtsdk4.HOSTSTATUS_UP {
					break
				}
			}
		}
	}
}
