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

func assignPermission() {
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

	// Locate the networks service and use it to find the network:
	networksService := conn.SystemService().NetworksService()
	networkListResp, err := networksService.List().Search("name=mynetwork").Send()
	if err != nil {
		fmt.Printf("Failed to search network list, reason: %v\n", err)
		return
	}
	networksSlice, _ := networkListResp.Networks()
	network := networksSlice.Slice()[0]

	//Locate the users service and use it to find the user
	usersService := conn.SystemService().UsersService()
	userListResp, err := usersService.List().Search("userName=myuser@mydomain-authz").Send()
	if err != nil {
		fmt.Printf("Failed to search user list, reason: %v\n", err)
		return
	}
	usersSlice, _ := userListResp.Users()
	user := usersSlice.Slice()[0]

	//Locate the role service and use it to find the role
	rolesService := conn.SystemService().RolesService()
	roleListResp, err := rolesService.List().Send()
	if err != nil {
		fmt.Printf("Failed to search role list, reason: %v\n", err)
		return
	}
	rolesSlice, _ := roleListResp.Roles()
	role := &ovirtsdk4.Role{}
	for _, da := range rolesSlice.Slice() {
		if da.MustName() == "GlusterAdmin" {
			role = da
			break
		}
	}
	if role == nil {
		permitSlice := &ovirtsdk4.PermitSlice{}
		permitSlice.SetSlice(genPermits([]string{"1", "1300"}))
		role = ovirtsdk4.NewRoleBuilder().
			Name("GlusterAdmin").
			Description("My role").
			Administrative(false).
			Permits(permitSlice).
			MustBuild()
	}

	// Locate the service that manages the permissions of the network
	permissionsService := networksService.NetworkService(network.MustId()).PermissionsService()

	// Use the "add" method to assign GlusterAdmin role to user on network
	_, err = permissionsService.Add().
		Permission(
			ovirtsdk4.NewPermissionBuilder().
				User(user).
				Role(role).
				MustBuild()).
		Send()

}
