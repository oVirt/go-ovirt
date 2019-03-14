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

func addVMFromTemplate() {
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

	// Get the reference to the service that manages the storage domains
	sdsService := conn.SystemService().StorageDomainsService()

	// Find the storage domain we want to be used for virtual machine disks
	sdsResp, err := sdsService.List().Search("name=mydata").Send()
	if err != nil {
		fmt.Printf("Failed to search storage domains, reason: %v\n", err)
		return
	}

	if sdSlice, ok := sdsResp.StorageDomains(); ok {
		sd := sdSlice.Slice()[0]

		// Get the reference to the service that manages the templates
		templatesService := conn.SystemService().TemplatesService()
		// When a template has multiple versions they all have the same name, so we need to explicitly find the one that
		// has the version name or version number that we want to use. In this case we want to use version 3 of the
		// template.
		tpsResp, err := templatesService.List().Search("name=mytemplate").Send()
		if err != nil {
			fmt.Printf("Failed to search templates, reason: %v\n", err)
			return
		}
		tpSlice, _ := tpsResp.Templates()
		var templateID string
		for _, tp := range tpSlice.Slice() {
			if tp.MustVersion().MustVersionNumber() == 3 {
				templateID = tp.MustId()
				break
			}
		}

		// Find the template disk we want be created on specific storage domain
		// for our virtual machine
		if templateID != "" {
			tpService := templatesService.TemplateService(templateID)
			tpGetResp, _ := tpService.Get().Send()
			tp, _ := tpGetResp.Template()
			disks, _ := conn.FollowLink(tp.MustDiskAttachments())
			if disks, ok := disks.(*ovirtsdk4.DiskAttachmentSlice); ok {
				disk := disks.Slice()[0].MustDisk()
				// Get the reference to the service that manages the virtual machines
				vmsService := conn.SystemService().VmsService()
				vmsService.Add().
					Vm(
						ovirtsdk4.NewVmBuilder().
							Name("myvm").
							Cluster(
								ovirtsdk4.NewClusterBuilder().
									Name("mycluster").
									MustBuild()).
							Template(
								ovirtsdk4.NewTemplateBuilder().
									Id(templateID).
									MustBuild()).
							// Use the vararg syntax to add just one or some
							DiskAttachmentsOfAny(
								ovirtsdk4.NewDiskAttachmentBuilder().
									Disk(
										ovirtsdk4.NewDiskBuilder().
											Id(disk.MustId()).
											Format(ovirtsdk4.DISKFORMAT_COW).
											StorageDomainsOfAny(
												ovirtsdk4.NewStorageDomainBuilder().
													Id(sd.MustId()).
													MustBuild()).
											MustBuild()).
									MustBuild()).
							MustBuild()).
					Send()
			}
		}
	}
}
