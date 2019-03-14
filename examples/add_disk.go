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
	"bytes"
	"fmt"
	"time"

	ovirtsdk4 "github.com/ovirt/go-ovirt"
)

func addDisk() {
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

	diskBuilder := ovirtsdk4.NewDiskBuilder().
		Name("joey_disk").
		Format(ovirtsdk4.DISKFORMAT_COW).
		ProvisionedSize(53687091200).
		StorageDomainsOfAny(
			ovirtsdk4.NewStorageDomainBuilder().
				Id("cadbe661-0e35-4fcb-a70d-2b17e2559d9c").
				MustBuild())
	disk, err := diskBuilder.Build()
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	var body bytes.Buffer
	writer := ovirtsdk4.NewXMLWriter(&body)
	err = ovirtsdk4.XMLDiskWriteOne(writer, disk, "")
	writer.Flush()

	fmt.Println("*********************" + body.String())

	_, err = conn.SystemService().DisksService().Add().Disk(disk).Send()

	if err != nil {
		fmt.Printf("error is %v\n", err)
	}

}
