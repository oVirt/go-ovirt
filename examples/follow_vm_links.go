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

func followVMLinks() {
	inputRawURL := "https://10.1.111.229/ovirt-engine/api"
	// Create the connection to the server:
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

	// Locate the service that manages the virtual machines
	vmsService := conn.SystemService().VmsService()

	// Find the virtual machine
	vm := vmsService.List().
		Search("name=test4joey").
		MustSend().
		MustVms().
		Slice()[0]

	// When the server returns a virtual machine it will return links to related objects, like the cluster,
	// template and permissions something like this:
	//
	// <link href="/api/vms/123/permissions" rel="permissions"/>
	// ...
	// <cluster id="123" href="/api/clusters/123"/>
	// <template id="456" href="/api/templates/456"/>
	//
	// The SDK provides a "FollowLink" method that can be used to retrieve the complete content of these related
	// objects.
	cluster, _ := conn.FollowLink(vm.MustCluster())
	if cluster, ok := cluster.(*ovirtsdk4.Cluster); ok {
		fmt.Printf("cluster: %v\n", cluster.MustName())
	}
	template, _ := conn.FollowLink(vm.MustTemplate())
	if template, ok := template.(*ovirtsdk4.Template); ok {
		fmt.Printf("template: %v\n", template.MustName())
	}
	permissionSlice, _ := conn.FollowLink(vm.MustPermissions())
	if permissionSlice, ok := permissionSlice.(*ovirtsdk4.PermissionSlice); ok {
		for _, per := range permissionSlice.Slice() {
			fmt.Printf("role: %v\n", per.MustRole().MustId())
		}
	}
}
