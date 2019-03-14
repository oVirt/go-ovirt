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

func addTag() {
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

	tagService := conn.SystemService().TagsService()

	newTag, err := ovirtsdk4.NewTagBuilder().
		Name("mytag").
		Description("mytag desc").
		Build()
	if err != nil {
		fmt.Printf("Construct a new tag failed, reason: %v\n", err)
	}

	resp, err := tagService.Add().Tag(newTag).Send()
	if err != nil {
		fmt.Printf("Failed to create a tag, reason: %v\n", err)
		return
	}
	tagAdded, _ := resp.Tag()
	fmt.Printf("Tag with name-(%v) and desc-(%v) added successfully\n",
		tagAdded.MustName(), tagAdded.MustDescription())

	// If tag added successfully, print out:
	// 		`Tag with name-(mytag) and desc-(mytag desc) added successfully`
	// If failed due to duplicate name, print out:
	// 		`Failed to create a tag, reason: Fault reason is "Operation Failed".
	// 		Fault detail is "[The specified Tag name already exists.]".
	//		HTTP response code is "409". HTTP response message is "409 Conflict".`

}
