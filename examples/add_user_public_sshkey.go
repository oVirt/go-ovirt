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

func addUserPublicSSHKey() {
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

	// Get the reference to the service that manages the users
	usersService := conn.SystemService().UsersService()

	// Find the user
	getUserResp, _ := usersService.List().Search("name=myuser").Send()
	userSlice, _ := getUserResp.Users()
	user := userSlice.Slice()[0]

	// Get the reference to the service that manages the user that we found in the previous step
	userService := usersService.UserService(user.MustId())

	// Get a reference to the service that manages the public SSH keys
	keysService := userService.SshPublicKeysService()

	// Add a new SSH public key
	keysService.Add().
		Key(
			ovirtsdk4.NewSshPublicKeyBuilder().
				Content("ssh-rsa AAA...mu9 myuser@example.com").
				MustBuild()).
		Send()

	// Note that the above operation will fail because the example SSH public key is not valid, make sure to use a
	// valid key.
}
