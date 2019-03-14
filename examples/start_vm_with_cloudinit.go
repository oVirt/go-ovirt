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

func startVMWithCloudinit() {
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

	// Get the reference to the "vms" service
	vmsService := conn.SystemService().VmsService()

	// Find the virtual machine
	vm := vmsService.List().Search("name=myvm").MustSend().MustVms().Slice()[0]

	// Locate the service that manages the virtual machine, as that is where the action methods are defined
	vmService := vmsService.VmService(vm.MustId())

	// Create cloud-init script, which you want to execute in deployed virtual machine, please note that
	// the script must be properly formatted and indented as it's using YAML.
	myScript := "write_files:\n" +
		"  - content: |\n" +
		"      Hello, world!\n" +
		"    path: /tmp/greeting.txt\n" +
		"    permissions: '0644'\n"

	// Start the virtual machine enabling cloud-init and providing the password for the `root` user and the network
	// configuration
	vmService.Start().
		UseCloudInit(true).
		Vm(
			ovirtsdk4.NewVmBuilder().
				Initialization(
					ovirtsdk4.NewInitializationBuilder().
						UserName("root").
						RootPassword("redhat123").
						HostName("myvm.example.com").
						NicConfigurationsOfAny(
							ovirtsdk4.NewNicConfigurationBuilder().
								Name("eth0").
								OnBoot(true).
								BootProtocol(ovirtsdk4.BOOTPROTOCOL_STATIC).
								Ip(
									ovirtsdk4.NewIpBuilder().
										Version(ovirtsdk4.IPVERSION_V4).
										Address("192.168.0.100").
										Netmask("255.255.255.0").
										Gateway("192.168.0.1").
										MustBuild()).
								MustBuild()).
						DnsServers("192.168.0.1 192.168.0.2").
						DnsSearch("example.com").
						CustomScript(myScript).
						MustBuild()).
				MustBuild()).
		MustSend()

}
