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

func pageVms() {
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

	// Get the reference to the vm service:
	vmsService := conn.SystemService().VmsService()

	// List the vm page by page:
	size := int64(1)
	index := int64(1)
	for {
		pageIndex := fmt.Sprintf("page %v", index)
		vmsResp, err := vmsService.List().Search(pageIndex).Max(size).Send()
		if err != nil {
			fmt.Printf("Failed to get vm list, reason: %v\n", err)
			return
		}
		if len(vmsResp.MustVms().Slice()) == 0 {
			break
		}

		for _, vm := range vmsResp.MustVms().Slice() {
			fmt.Printf("Vm name: %v\n", vm.MustName())
		}

		index += int64(1)
		fmt.Printf("Index: %v\n", index)
	}

}
