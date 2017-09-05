package ovirtsdk4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadStrings(t *testing.T) {
	assert := assert.New(t)
	xmlstring := `
    <cpu>
        <type>x86_64</type>
        <type>Intel SandyBridge Family</type>
    </cpu>
	`
	reader := NewXMLReader([]byte(xmlstring))

	strs, err := reader.ReadStrings(nil)
	assert.Nil(err)

	assert.Equal([]string{"x86_64", "Intel SandyBridge Family"}, strs)
}
