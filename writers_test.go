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

package ovirtsdk

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCpuWriteOne(t *testing.T) {
	assert := assert.New(t)
	var b bytes.Buffer
	writer := NewXMLWriter(&b)
	cpu, err := NewCpuBuilder().Type("Intel SandyBridge Family").Architecture(ARCHITECTURE_X86_64).Build()
	assert.Nil(err)
	err = XMLCpuWriteOne(writer, cpu, "")
	assert.Nil(err)
	writer.Flush()
	assert.Equal("<cpu><architecture>x86_64</architecture><type>Intel SandyBridge Family</type></cpu>",
		string(b.Bytes()))

}

func TestBootDevice(t *testing.T) {
	assert := assert.New(t)
	var b bytes.Buffer
	writer := NewXMLWriter(&b)
	boot, err := NewBootBuilder().DevicesOfAny("network").Build()
	assert.Nil(err)
	err = XMLBootWriteOne(writer, boot, "")
	assert.Nil(err)
	writer.Flush()
	assert.Equal("<boot><devices><device>network</device></devices></boot>",
		string(b.Bytes()))
}

func TestBootDevices(t *testing.T) {
	assert := assert.New(t)
	var b bytes.Buffer
	writer := NewXMLWriter(&b)
	boot, err := NewBootBuilder().DevicesOfAny("network", "hd").Build()
	assert.Nil(err)
	err = XMLBootWriteOne(writer, boot, "")
	assert.Nil(err)
	writer.Flush()
	assert.Equal("<boot><devices><device>network</device><device>hd</device></devices></boot>",
		string(b.Bytes()))
}
