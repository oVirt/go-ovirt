//
// Copyright (c) 2019 Red Hat, Inc.
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
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	ovirtsdk4 "github.com/ovirt/go-ovirt"
)

func downloadDisk() {
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

	// Get a disk identified by uuid
	disksService := conn.SystemService().DisksService()
	diskService := disksService.DiskService("bec4ad05-2178-4e44-840b-c8f3b2906974")
	diskRequest := diskService.Get()
	disk := diskRequest.MustSend().MustDisk()

	// Prepare image transfer request
	transfersService := conn.SystemService().ImageTransfersService()
	transfer := transfersService.Add()
	imageTransfer := ovirtsdk4.NewImageTransferBuilder().Image(
		ovirtsdk4.NewImageBuilder().Id(disk.MustId()).MustBuild(),
	).Direction(
		ovirtsdk4.IMAGETRANSFERDIRECTION_DOWNLOAD,
	).MustBuild()

	// Initialize image transfer and lock the disk
	transfer.ImageTransfer(imageTransfer)
	it := transfer.MustSend().MustImageTransfer()
	for {
		if it.MustPhase() == ovirtsdk4.IMAGETRANSFERPHASE_INITIALIZING {
			time.Sleep(1 * time.Second)
			it = conn.SystemService().ImageTransfersService().ImageTransferService(it.MustId()).Get().MustSend().MustImageTransfer()
		} else {
			break
		}
	}

	// Create target disk
	out, err := os.Create(disk.MustName())
	if err != nil {
		fmt.Printf("Disk file creation failed, reason: %v\n", err)
		return
	}
	defer out.Close()

	// Load api certificate. To obtain run: "wget https://{URL}/ovirt-engine/services/pki-resource?resource=ca-certificate&format=X509-PEM-CA"
	caCert, err := ioutil.ReadFile("ca.pem")
	if err != nil {
		fmt.Printf("Reading ca file failed, reason: %v\n", err)
		return
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Create http client
	tlsConfig := &tls.Config{
		RootCAs: caCertPool,
	}
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: transport}

	// Prepare disk GET request
	req, err := http.NewRequest("GET", it.MustTransferUrl(), nil)
	req.Header.Set("Authorization", it.MustSignedTicket())

	// Run the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Sending request failed, reason: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("bad status: %s", resp.Status)
		return
	}

	// Copy body to local disk file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Printf("Writing disk file failed, reason: %v\n", err)
		return
	}
}
