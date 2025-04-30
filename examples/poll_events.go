//
// The oVirt Project - oVirt Engine Go SDK
//
// Copyright (c) oVirt Authors
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
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	ovirtsdk4 "github.com/ovirt/go-ovirt/v4"
)

func pollEvents() {
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

	// Find the service that manages the collection of events:
	eventsService := conn.SystemService().EventsService()

	/* If there is no index stored yet, then retrieve the last event and
	start with it. Events are ordered by index, ascending, and 'max=1'
	requests just one event, so we will get the last event.*/
	if readIndex() == 0 {
		events := eventsService.List().Max(1).MustSend()
		if len(events.MustEvents().Slice()) > 0 {
			event := events.MustEvents().Slice()[0]
			fmt.Printf("%v - %v", event.MustId(), event.MustDescription())

			eventId, err := strconv.ParseInt(event.MustId(), 10, 64)
			fmt.Printf("Failed to parse the string, reason: %v\n", err)
			writeIndex(eventId)
		}
	}

	/* Do a loop to retrieve events, starting always with the last index, and
	waiting a bit before repeating. Note the use of the 'from' parameter
	to specify that we want to get events newer than the last index that
	we already processed. Note also that we don't use the 'max' parameter
	here because we want to get all the pending events.*/
	for {
		time.Sleep(5 * time.Second)
		eventResp, err := eventsService.List().From(readIndex()).Send()
		if err != nil {
			fmt.Printf("Failed to get event list, reason: %v\n", err)
			return
		}
		for _, event := range eventResp.MustEvents().Slice() {
			fmt.Printf("%v - %v", event.MustId(), event.MustDescription())
			eventId, err := strconv.ParseInt(event.MustId(), 10, 64)
			fmt.Printf("Failed to parse the string, reason: %v\n", err)
			writeIndex(eventId)
		}
	}

}

func readIndex() (b int64) {
	_, err := os.Stat("index.txt")
	if err == nil {
		// File exist
		file, err := os.Open("index.txt")
		defer file.Close()
		if err != nil {
			fmt.Printf("Failed to open the file, reason:%v\n", err)
			return 0
		}
		readBuf := make([]byte, 1024)
		for {
			n, err := file.Read(readBuf)
			if n == 0 {
				break
			}
			fmt.Println(int64(n), err)
			return int64(n)
		}

	}
	return 0
}

func writeIndex(data int64) {
	bytebuf := bytes.NewBuffer([]byte{})
	binary.Write(bytebuf, binary.BigEndian, data)
	//将指定内容写入到文件中
	err := ioutil.WriteFile("index.txt", bytebuf.Bytes(), 0666)
	if err != nil {
		fmt.Printf("ioutil WriteFile error: %v\n ", err)
	}
}
