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

func addTemplate() {
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

	// Get the reference to the root of the services tree
	systemService := conn.SystemService()

	// Find the original virtual machine
	vmsService := systemService.VmsService()
	vmlistResp, _ := vmsService.List().Search("name=myvm").Send()
	vmSlice, _ := vmlistResp.Vms()
	vm := vmSlice.Slice()[0]

	// Get the identifiers of the disks attached to the virtual machine. We # need this because we want to tell the
	// server to create the disks of the template using a format different to the format used by the original disks	.
	var diskIDs []string
	dasInterface, _ := conn.FollowLink(vm.MustDiskAttachments())
	if das, ok := dasInterface.(*ovirtsdk4.DiskAttachmentSlice); ok {
		for _, da := range das.Slice() {
			diskIDs = append(diskIDs, da.MustDisk().MustId())
		}
	}

	// Create a customized list of disk attachments, explicitly indicating that we want COW disks, regardless of
	// format the original disks had
	attachments := new(ovirtsdk4.DiskAttachmentSlice)
	for _, diskID := range diskIDs {
		one := ovirtsdk4.NewDiskAttachmentBuilder().
			Disk(
				ovirtsdk4.NewDiskBuilder().
					Id(diskID).
					Format(ovirtsdk4.DISKFORMAT_COW).
					MustBuild()).
			MustBuild()
		attachments.SetSlice(append(attachments.Slice(), one))
	}

	// Send the request to create the template. Note that the way to specify the original virtual machine, and the
	// customizations, is to use the 'vm' attribute of the 'Template' type.
	templatesService := systemService.TemplatesService()
	addTpResp, _ := templatesService.Add().
		Template(
			ovirtsdk4.NewTemplateBuilder().
				Name("mytemplate").
				Vm(
					ovirtsdk4.NewVmBuilder().
						Id(vm.MustId()).
						DiskAttachments(attachments).
						MustBuild()).
				MustBuild()).
		Send()
	template, _ := addTpResp.Template()

	// Wait till the status of the template is OK, as that means that it is completely created and ready to use
	tpService := templatesService.TemplateService(template.MustId())
	for {
		time.Sleep(5 * time.Second)
		getTpResp, _ := tpService.Get().Send()
		getTp, _ := getTpResp.Template()
		if getTp.MustStatus() == ovirtsdk4.TEMPLATESTATUS_OK {
			break
		}
	}

}
