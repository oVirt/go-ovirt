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

func assignPermissionToVm() {
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

	// This example will connect to the server and assign UserVmManager
	// role to myuser1 on virtual machine name myvm1, myvm2 nad myvm3:
	// List of virtual machine where user should have assigned permission:
	myVms := []string{"myvm1", "myvm2", "myvm3"}

	//Role which we want to assign to the user on the virtual machines:
	roleName := "UserVmManager"

	// Locate the users service and use it to find the user
	usersService := conn.SystemService().UsersService()
	userListResp, err := usersService.List().Search("userName=user2@internal-authz").Send()
	if err != nil {
		fmt.Printf("Failed to search user list, reason: %v\n", err)
		return
	}
	usersSlice, _ := userListResp.Users()
	user := usersSlice.Slice()[0]

	// Locate the roles service and use it to find the role
	role := &ovirtsdk4.Role{}
	rolesService := conn.SystemService().RolesService()
	roleListResp, err := rolesService.List().Send()
	if err != nil {
		fmt.Printf("Failed to search role list, reason: %v\n", err)
		return
	}
	rolesSlice, _ := roleListResp.Roles()
	for _, da := range rolesSlice.Slice() {
		if da.MustName() == roleName {
			role = da
			break
		}
	}
	if role == nil {
		permitSlice := &ovirtsdk4.PermitSlice{}
		permitSlice.SetSlice(genPermits([]string{"1", "1300"}))
		role = ovirtsdk4.NewRoleBuilder().
			Name(roleName).
			Description("My role").
			Administrative(false).
			Permits(permitSlice).
			MustBuild()
	}

	for vmName := range myVms {
		vmNamestr := fmt.Sprintf("name=%v", vmName)

		// Locate the virtual machine service and use it to find the specific
		// virtual machines:
		vmsService := conn.SystemService().VmsService()
		vmListResp, err := vmsService.List().Search(vmNamestr).Send()
		if err != nil {
			fmt.Printf("Failed to search vm list, reason: %v\n", err)
			return
		}
		vmsSlice, _ := vmListResp.Vms()
		vm := vmsSlice.Slice()[0]

		// Locate the service that manages the permissions of the virtual machine:
		permissionsService := vmsService.VmService(vm.MustId()).PermissionsService()

		// Use the "add" method to assign UserVmManager role to user on virtual machine:
		_, err = permissionsService.Add().
			Permission(
				ovirtsdk4.NewPermissionBuilder().
					User(user).
					Role(role).
					MustBuild()).
			Send()
	}

}
