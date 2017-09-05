package ovirtsdk4

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
