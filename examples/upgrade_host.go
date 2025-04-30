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
	"strconv"
	"time"

	ovirtsdk4 "github.com/ovirt/go-ovirt/v4"
)

func upgradeHosts() {
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

	// Get the reference to the service that manages the hosts
	hostsService := conn.SystemService().HostsService()

	// Get the reference to the service that manages the host
	eventsService := conn.SystemService().EventsService()

	// Find the host
	host := hostsService.List().
		Search("name=myhost").
		MustSend().
		MustHosts().
		Slice()[0]

	// Get the reference to the service that manages the fencing agents used by the host that we found in the
	// previous step
	hostService := hostsService.HostService(host.MustId())

	// Check if the host has available update, if not run check for
	// upgrade action:
	if !host.MustUpdateAvailable() {
		// Run the check for upgrade action:
		hostService.UpgradeCheck().Send()

		// Wait until event with id 839 or 887 occur, which means
		// that check for upgrade action failed.
		// Or wait until event with id 885 occured, which means
		// that check for upgrade action succeed.
		lastEvent := eventsService.List().Max(int64(1)).MustSend().
			MustEvents().Slice()[0]
		timeOut := int64(180)
		lastEventId, _ := strconv.ParseInt(lastEvent.MustId(), 10, 64)
		hostName := fmt.Sprintf("host.name=%v", host.MustName())

		for {
			if timeOut > 0 {
				eventResp := eventsService.List().
					From(lastEventId).
					Search(hostName).
					MustSend()

				for _, event := range eventResp.MustEvents().Slice() {
					if (event.MustCode() == 839) || (event.MustCode() == 887) {
						fmt.Printf("Check for upgrade failed.")
						break
					}
					if (event.MustCode() == 839) || (event.MustCode() == 887) {
						fmt.Printf("Check for upgrade done.")
						break
					}
				}
				time.Sleep(2 * time.Second)
				timeOut -= 2
			} else {
				break
			}
		}
	}

	// Refresh the host object and run the upgrade action
	// if host has available updates:
	getResp, err := hostService.Get().Send()
	if err != nil {
		fmt.Printf("Faild to get host, reason: %v\n", err)
	}
	// Check if the host is OK
	if host, ok := getResp.Host(); ok {
		if host.MustUpdateAvailable() {
			hostService.Upgrade().Send()
		}
	}

}
