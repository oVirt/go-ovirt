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

func addFenceAgent() {
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

	// Find the host
	hostsService := conn.SystemService().HostsService()
	listResp, err := hostsService.List().Search("name=myhost").Send()
	if err != nil {
		fmt.Printf("Failed to search host list, reason: %v\n", err)
		return
	}
	hostSlice, _ := listResp.Hosts()
	host := hostSlice.Slice()[0]

	// Find the service that manages the hosts
	hostService := hostsService.HostService(host.MustId())

	fenceAgentsService := hostService.FenceAgentsService()
	// Use the "add" method to create a fence agents:
	_, err = fenceAgentsService.Add().
		Agent(
			ovirtsdk4.NewAgentBuilder().
				Address("1.2.3.4").
				Type("ipmilan").
				Username("myusername").
				Password("mypassword").
				OptionsOfAny(
					ovirtsdk4.NewOptionBuilder().
						Name("myname").
						Value("myvalue").
						MustBuild()).
				Order(0).
				MustBuild()).
		Send()
	if err != nil {
		fmt.Printf("Failed to add fenceAgents, reason: %v\n", err)
		return
	}

	// Prepare the update host object:
	hostUpdate := &ovirtsdk4.Host{}

	powerManagement, bool := host.PowerManagement()
	kdumpEnabled := false
	if bool && powerManagement != nil {
		newPow := &ovirtsdk4.PowerManagement{}

		// If power management isn't enabled, enable it. Note that
		// power management can be enabled only if at least one fence
		// agent exists for host:
		enable, _ := powerManagement.Enabled()
		if !enable {
			newPow.SetEnabled(true)
		}

		//	If kdump isn't enabled, enable it. Note that kdump
		//	can be enabled only if power management on the host
		//	is enabled:
		kdumpDetection, _ := powerManagement.KdumpDetection()
		if !kdumpDetection {
			kdumpEnabled = true
			newPow.SetKdumpDetection(true)
		}
		hostUpdate.SetPowerManagement(newPow)
	}

	// Send the update request:
	hostService.Update().Host(hostUpdate).Send()

	// Note that after kdump integration is enabled the host must be
	// reinstalled so the kdump will take effect. First we need to
	// switch host to maintenance, because host must be in maintenance
	// mode, for reinstall process:
	if kdumpEnabled {
		hostService.Deactivate().MustSend()
		// Wait a bit so the host status is updated in API
		time.Sleep(5 * time.Second)
		for {
			getHostResp, err := hostService.Get().Send()
			if err != nil {
				continue
			}
			if getHost, ok := getHostResp.Host(); ok {
				if getHost.MustStatus() == ovirtsdk4.HOSTSTATUS_MAINTENANCE {
					break
				}
			}
			time.Sleep(2 * time.Second)
		}

		// Then we can execute the reinstall process
		ssh := &ovirtsdk4.Ssh{}
		ssh.SetAuthenticationMethod(ovirtsdk4.SSHAUTHENTICATIONMETHOD_PUBLICKEY)
		hostService.Install().Ssh(ssh).Send()

		// Wait a bit so the host status is updated in API
		time.Sleep(5 * time.Second)
		for {
			getHostResp, err := hostService.Get().Send()
			if err != nil {
				continue
			}
			if getHost, ok := getHostResp.Host(); ok {
				if getHost.MustStatus() == ovirtsdk4.HOSTSTATUS_MAINTENANCE {
					break
				}
			}
			time.Sleep(10 * time.Second)
		}

		// Activate the host after reinstall
		hostService.Activate().MustSend()
		time.Sleep(5 * time.Second)
		for {
			getHostResp, err := hostService.Get().Send()
			if err != nil {
				continue
			}
			if getHost, ok := getHostResp.Host(); ok {
				if getHost.MustStatus() == ovirtsdk4.HOSTSTATUS_UP {
					break
				}
			}
			time.Sleep(2 * time.Second)
		}

	}
}
