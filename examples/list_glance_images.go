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

func listGlanceImages() {
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

	// Find the Glance storage domain that is available for default in any oVirt installation
	sdsService := conn.SystemService().StorageDomainsService()
	resp, err := sdsService.List().
		Search("name=ovirt-image-repository").
		Send()
	if err != nil {
		fmt.Printf("Failed to get storage domains list, reason: %v\n", err)
		return
	}
	if sds, ok := resp.StorageDomains(); ok {
		sd := sds.Slice()[0]
		sdService := sdsService.StorageDomainService(sd.MustId())
		imagesService := sdService.ImagesService()
		// Not recommended, response error should be always be checked
		imgResp, _ := imagesService.List().Send()
		images, _ := imgResp.Images()
		for _, img := range images.Slice() {
			fmt.Printf("Image name is %v\n", img.MustName())
		}
	}
}
