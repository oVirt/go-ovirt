//
// Copyright (c) 2020 huihui <huihui.fu@cs2c.com.cn>.
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

func addOpenstackImageProvider() {
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

	// Get the list of OpenStack image providers (a.k.a. Glance providers)
	// that match the name that we want to use:
	openstackImageProvidersService := conn.SystemService().OpenstackImageProvidersService()

	listResp, err := openstackImageProvidersService.List().Search("name=myglance").Send()
	if err != nil {
		fmt.Printf("Failed to search dataCenter list, reason: %v\n", err)
		return
	}
	providersSlice, _ := listResp.Providers()

	// There is no such provider, then add it
	if len(providersSlice.Slice()) == 0 {
		_, err = openstackImageProvidersService.Add().
			Provider(
				ovirtsdk4.NewOpenStackImageProviderBuilder().
					Name("myglance").
					Description("My Glance").
					Url("http://glance.ovirt.org:9292").
					RequiresAuthentication(false).
					MustBuild()).
			Send()
	}

	// Note that the provider that we are using in this example is public
	// and doesn't require any authentication. If your provider requires
	// authentication then you will need to specify additional security
	// related attributes:
	//_, err = openstackImageProvidersService.Add().
	//	Provider(
	//		ovirtsdk4.NewOpenStackImageProviderBuilder().
	//			Name("myglance").
	//			Description("My Glance").
	//			Url("http://glance.ovirt.org:9292").
	//			RequiresAuthentication(true).
	//			AuthenticationUrl("http://mykeystone").
	//			Username("myuser").
	//			Password("mypassword").
	//			TenantName("mytenant").
	//			MustBuild()).
	//	Send()

}
