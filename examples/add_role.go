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

func addRole() {
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

	// Get the reference to the roles service:
	rolesService := conn.SystemService().RolesService()

	// Use the "add" method to create a role:
	permitSlice := &ovirtsdk4.PermitSlice{}
	permitSlice.SetSlice(genPermits([]string{"1", "1300"}))
	_, err = rolesService.Add().
		Role(
			ovirtsdk4.NewRoleBuilder().
				Name("myrole").
				Description("My role").
				Administrative(false).
				Permits(permitSlice).
				MustBuild()).
		Send()
	if err != nil {
		fmt.Printf("Failed to add role, reason: %v\n", err)
		return
	}
}

func genPermits(ids []string) []*ovirtsdk4.Permit {
	permits := make([]*ovirtsdk4.Permit, 0)
	for _, id := range ids {
		if id == "" {
			continue
		}
		permit := &ovirtsdk4.Permit{}
		permit.SetId(id)
		permits = append(permits, permit)
	}
	return permits
}
