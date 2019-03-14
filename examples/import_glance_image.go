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

func importGlanceImage() {
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

	// Get the root of the services tree
	systemService := conn.SystemService()

	// Find the Glance storage domain that is available for default in any oVirt installation
	sdsService := systemService.StorageDomainsService()
	sd := sdsService.List().
		Search("name=ovirt-image-repository").
		MustSend().
		MustStorageDomains().
		Slice()[0]

	// Find the service that manages the Glance storage domain
	sdService := sdsService.StorageDomainService(sd.MustId())

	// Find the service that manages the images available in that storage domain
	imagesService := sdService.ImagesService()

	// The images service doesn't support search, so in order to find an image
	// we need to retrieve all of them and do the filtering explicitly
	imageSlice := imagesService.List().MustSend().MustImages()
	var image *ovirtsdk4.Image
	for _, img := range imageSlice.Slice() {
		if img.MustName() == "CirrOS 0.3.4 for x86_64" {
			image = img
			break
		}
	}

	// Find the service that manages the image that we found in previous step
	imageService := imagesService.ImageService(image.MustId())

	// Import the image
	imageService.Import().
		ImportAsTemplate(true).
		Template(
			ovirtsdk4.NewTemplateBuilder().
				Name("mytemplate").
				MustBuild()).
		Cluster(
			ovirtsdk4.NewClusterBuilder().
				Name("mycluster").
				MustBuild()).
		StorageDomain(
			ovirtsdk4.NewStorageDomainBuilder().
				Name("mydata").
				MustBuild()).
		MustSend()
}
