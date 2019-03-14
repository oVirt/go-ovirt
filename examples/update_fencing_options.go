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

func updateFencingOptions() {
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

	// The name and value of the option that we want to add or update:
	name := "lanplus"
	value := "1"

	// Get the reference to the service that manages the hosts
	hostsService := conn.SystemService().HostsService()

	// Find the host
	host := hostsService.List().
		Search("name=myhost").
		MustSend().
		MustHosts().
		Slice()[0]

	// Get the reference to the service that manages the fencing agents used by the host that we found in the
	// previous step
	hostService := hostsService.HostService(host.MustId())
	agentsService := hostService.FenceAgentsService()

	// The host may have multiple fencing agents, so we need to locate the first of type 'ipmilan'
	agentSlice := agentsService.List().MustSend().MustAgents()
	var agent *ovirtsdk4.Agent
	for _, ag := range agentSlice.Slice() {
		if ag.MustType() == "ipmlan" {
			agent = ag
			break
		}
	}

	// Get the options of the fencing agent. There may be no options, in that case we need to use an empty list.
	optionSlice, ok := agent.Options()
	if !ok {
		optionSlice = new(ovirtsdk4.OptionSlice)
	}

	// Create a list of modified options, containing all the original options except the one with the name we want
	// to modify, as we will add that with the right value later
	var modifiedList []*ovirtsdk4.Option
	for _, ori := range optionSlice.Slice() {
		if name != ori.MustName() {
			modifiedList = append(modifiedList, ori)
		}
	}

	// Add the modified option to the list of modified options
	option := ovirtsdk4.NewOptionBuilder().
		Name(name).
		Value(value).
		MustBuild()
	modifiedList = append(modifiedList, option)

	// Find the service that manages the fence agent
	agentService := agentsService.AgentService(agent.MustId())

	// Send the update request containing the modified list of options
	agentService.Update().
		Agent(
			ovirtsdk4.NewAgentBuilder().
				OptionsOfAny(modifiedList...).
				MustBuild()).
		MustSend()
}
