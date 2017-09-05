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

package ovirtsdk4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXMLClusterReadOne(t *testing.T) {
	assert := assert.New(t)
	xmlstring := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<cluster href="/ovirt-engine/api/clusters/00000002-0002-0002-0002-000000000310" id="00000002-0002-0002-0002-000000000310">
    <actions>
        <link href="/ovirt-engine/api/clusters/00000002-0002-0002-0002-000000000310/resetemulatedmachine" rel="resetemulatedmachine"/>
    </actions>
    <name>Default</name>
    <description>The default server cluster</description>
    <link href="/ovirt-engine/api/clusters/00000002-0002-0002-0002-000000000310/networks" rel="networks"/>
    <link href="/ovirt-engine/api/clusters/00000002-0002-0002-0002-000000000310/permissions" rel="permissions"/>
    <link href="/ovirt-engine/api/clusters/00000002-0002-0002-0002-000000000310/glustervolumes" rel="glustervolumes"/>
    <link href="/ovirt-engine/api/clusters/00000002-0002-0002-0002-000000000310/glusterhooks" rel="glusterhooks"/>
    <link href="/ovirt-engine/api/clusters/00000002-0002-0002-0002-000000000310/affinitygroups" rel="affinitygroups"/>
    <link href="/ovirt-engine/api/clusters/00000002-0002-0002-0002-000000000310/cpuprofiles" rel="cpuprofiles"/>
    <ballooning_enabled>false</ballooning_enabled>
    <cpu>
        <architecture>x86_64</architecture>
        <type>Intel SandyBridge Family</type>
    </cpu>
    <error_handling>
        <on_error>migrate</on_error>
    </error_handling>
    <fencing_policy>
        <enabled>true</enabled>
        <skip_if_connectivity_broken>
            <enabled>false</enabled>
            <threshold>50</threshold>
        </skip_if_connectivity_broken>
        <skip_if_sd_active>
            <enabled>false</enabled>
        </skip_if_sd_active>
    </fencing_policy>
    <gluster_service>false</gluster_service>
    <ha_reservation>false</ha_reservation>
    <ksm>
        <enabled>true</enabled>
        <merge_across_nodes>true</merge_across_nodes>
    </ksm>
    <maintenance_reason_required>false</maintenance_reason_required>
    <memory_policy>
        <over_commit>
            <percent>100</percent>
        </over_commit>
        <transparent_hugepages>
            <enabled>true</enabled>
        </transparent_hugepages>
    </memory_policy>
    <migration>
        <auto_converge>inherit</auto_converge>
        <bandwidth>
            <assignment_method>auto</assignment_method>
        </bandwidth>
        <compressed>inherit</compressed>
        <policy id="80554327-0569-496b-bdeb-fcbbf52b827b"/>
    </migration>
    <optional_reason>false</optional_reason>
    <required_rng_sources>
		<required_rng_source>random</required_rng_source>
		<required_rng_source>get</required_rng_source>
		<required_rng_source>post</required_rng_source>
    </required_rng_sources>
    <switch_type>legacy</switch_type>
    <threads_as_cores>false</threads_as_cores>
    <trusted_service>false</trusted_service>
    <tunnel_migration>false</tunnel_migration>
    <version>
        <major>4</major>
        <minor>0</minor>
    </version>
    <virt_service>true</virt_service>
    <data_center href="/ovirt-engine/api/datacenters/00000001-0001-0001-0001-0000000002ed" id="00000001-0001-0001-0001-0000000002ed"/>
    <scheduling_policy href="/ovirt-engine/api/schedulingpolicies/b4ed2332-a7ac-4d5f-9596-99a439cb2812" id="b4ed2332-a7ac-4d5f-9596-99a439cb2812"/>
</cluster>
`

	reader := NewXMLReader([]byte(xmlstring))

	cluster, err := XMLClusterReadOne(reader, nil, "")
	assert.Nil(err)
	assert.NotNil(cluster)
	// Cluster>Id
	// assert.Equal("00000002-0002-0002-0002-000000000310", *cluster.Id)
	// Cluster>Name
	name, ok := cluster.Name()
	assert.True(ok)
	assert.Equal("Default", name, "Name should be `Default`")

	// Cluster>CPU>Architecture
	cpu, ok := cluster.Cpu()
	assert.True(ok)
	arch, ok := cpu.Architecture()
	assert.True(ok)
	assert.Equal(Architecture("x86_64"), arch, "CPU Arch should be `x86_64`")

	// Cluster>BallooningEnabled
	ballooningEnabled, ok := cluster.BallooningEnabled()
	assert.True(ok)
	assert.False(ballooningEnabled, "Cluster>BallooningEnabled should be false")

	// Cluster>Description
	desc, ok := cluster.Description()
	assert.True(ok)
	assert.Equal("The default server cluster", desc)

	// Cluster>ErrorHandling>OnError>MigrateOnError
	errorHandling, ok := cluster.ErrorHandling()
	assert.True(ok)
	errorHandlingOnError, ok := errorHandling.OnError()
	assert.Equal(MigrateOnError("migrate"), errorHandlingOnError)

	// Cluster>FencingPolicy
	fencingPolicy, ok := cluster.FencingPolicy()
	assert.True(ok)
	// 		>Enable
	fencingPolicyEnabled, ok := fencingPolicy.Enabled()
	assert.True(ok)
	assert.True(fencingPolicyEnabled, "Cluster>FencingPolicy>Enable should be true")
	// 		>SkipIfConnectivityBroken>Threshold
	skipIfConnectiveBroken, ok := fencingPolicy.SkipIfConnectivityBroken()
	assert.True(ok)
	threshold, ok := skipIfConnectiveBroken.Threshold()
	assert.True(ok)
	assert.Equal(int64(50), threshold,
		"Cluster>FencingPolicy>SkipIfConnectivityBroken>Threshold should be 50")
	skipIfSdActive, ok := fencingPolicy.SkipIfSdActive()
	assert.True(ok)
	sdActiveEnable, ok := skipIfSdActive.Enabled()
	assert.True(ok)
	assert.False(sdActiveEnable, "Cluster>FencingPolicy>SkipIfSdActive>Enabled should be false")

	// Cluster>SchedulingPolicy
	schedulingPolicy, ok := cluster.SchedulingPolicy()
	assert.True(ok)
	schedulingPolicyID, ok := schedulingPolicy.Id()
	assert.True(ok)
	assert.Equal("b4ed2332-a7ac-4d5f-9596-99a439cb2812", schedulingPolicyID)
	schedulingPolicyHref, ok := schedulingPolicy.Href()
	assert.True(ok)
	assert.Equal("/ovirt-engine/api/schedulingpolicies/b4ed2332-a7ac-4d5f-9596-99a439cb2812", schedulingPolicyHref)

	// Cluster Links
	glusterhooks, ok := cluster.GlusterHooks()
	assert.True(ok)
	href, ok := glusterhooks.Href()
	assert.True(ok)
	assert.Equal("/ovirt-engine/api/clusters/00000002-0002-0002-0002-000000000310/glusterhooks", href)
}

func TestXMLErrorHandlingReadOne(t *testing.T) {
	assert := assert.New(t)
	xmlstring := `
    <error_handling>
        <on_error>migrate</on_error>
	</error_handling>`
	reader := NewXMLReader([]byte(xmlstring))
	eh, err := XMLErrorHandlingReadOne(reader, nil, "")
	assert.Nil(err)
	assert.NotNil(eh)
	onError, ok := eh.OnError()
	assert.True(ok)
	assert.Equal(MigrateOnError("migrate"), onError)
}

func TestXMLCPUReadOne(t *testing.T) {
	assert := assert.New(t)
	xmlstring := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<cpu>
	<architecture>x86_64</architecture>
	<type>Intel SandyBridge Family</type>
</cpu>
`

	reader := NewXMLReader([]byte(xmlstring))

	cpu, err := XMLCpuReadOne(reader, nil, "")
	assert.Nil(err)
	cpuArch, ok := cpu.Architecture()
	assert.True(ok)
	assert.Equal(Architecture("x86_64"), cpuArch)
	cpuType, ok := cpu.Type()
	assert.True(ok)
	assert.Equal("Intel SandyBridge Family", cpuType)
}

func TestArchitectureReadMany(t *testing.T) {
	assert := assert.New(t)
	xmlstring := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<architectures>
	<architecture>x86_64</architecture>
	<architecture>x86</architecture>
	<architecture>mips</architecture>
	<architecture>powerpc</architecture>
</architectures>
`

	reader := NewXMLReader([]byte(xmlstring))

	archs, err := XMLArchitectureReadMany(reader, nil)
	assert.Nil(err)
	assert.Equal([]Architecture{
		Architecture("x86_64"), Architecture("x86"),
		Architecture("mips"), Architecture("powerpc")}, archs)
}

func TestFaultReadOne(t *testing.T) {
	assert := assert.New(t)
	xmlstring := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<architectures>
	<architecture>x86_64</architecture>
	<architecture>x86</architecture>
	<architecture>mips</architecture>
	<architecture>powerpc</architecture>
</architectures>`
	reader := NewXMLReader([]byte(xmlstring))
	fault, err := XMLFaultReadOne(reader, nil, "")
	assert.Nil(fault)
	err2, ok := err.(XMLTagNotMatchError)
	assert.True(ok)
	assert.Equal("architectures", err2.ActualTag)
	assert.Equal("fault", err2.ExpectedTag)
}
