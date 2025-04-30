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

// To cover the issue: https://github.com/oVirt/ovirt-engine-sdk-go/issues/121
func TestGlusterBrickReadMany(t *testing.T) {
	assert := assert.New(t)
	xmlstring := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<bricks>
	<actions>
		<link href="/ovirt-engine/api/clusters/5b6d20ce-00b5-019a-02d2-00000000037c/glustervolumes/c667523c-7b0d-4602-82ff-3a721e760835/glusterbricks/migrate" rel="migrate"/>
		<link href="/ovirt-engine/api/clusters/5b6d20ce-00b5-019a-02d2-00000000037c/glustervolumes/c667523c-7b0d-4602-82ff-3a721e760835/glusterbricks/stopmigrate" rel="stopmigrate"/>
		<link href="/ovirt-engine/api/clusters/5b6d20ce-00b5-019a-02d2-00000000037c/glustervolumes/c667523c-7b0d-4602-82ff-3a721e760835/glusterbricks/activate" rel="activate"/>
	</actions>
	<brick href="/ovirt-engine/api/clusters/5b6d20ce-00b5-019a-02d2-00000000037c/glustervolumes/c667523c-7b0d-4602-82ff-3a721e760835/glusterbricks/06de18a8-8534-4ce4-9328-ff487ac1355f" id="06de18a8-8534-4ce4-9328-ff487ac1355f">
		<actions>
			<link href="/ovirt-engine/api/clusters/5b6d20ce-00b5-019a-02d2-00000000037c/glustervolumes/c667523c-7b0d-4602-82ff-3a721e760835/glusterbricks/06de18a8-8534-4ce4-9328-ff487ac1355f/replace" rel="replace"/>
		</actions>
		<name>10.1.87.223:/tt112</name>
		<link href="/ovirt-engine/api/clusters/5b6d20ce-00b5-019a-02d2-00000000037c/glustervolumes/c667523c-7b0d-4602-82ff-3a721e760835/glusterbricks/06de18a8-8534-4ce4-9328-ff487ac1355f/statistics" rel="statistics"/>
		<brick_dir>/tt112</brick_dir>
		<server_id>75b39f00-fa03-410c-b29c-db80cf162a87</server_id>
		<status>up</status>
		<gluster_volume href="/ovirt-engine/api/clusters/5b6d20ce-00b5-019a-02d2-00000000037c/glustervolumes/c667523c-7b0d-4602-82ff-3a721e760835" id="c667523c-7b0d-4602-82ff-3a721e760835"/>
	</brick>
</bricks>
`
	reader := NewXMLReader([]byte(xmlstring))
	bricks, err := XMLGlusterBrickReadMany(reader, nil)
	assert.Nil(err)
	assert.NotNil(bricks)
	assert.Equal(1, len(bricks.Slice()))

}

func TestNetworkReadOne(t *testing.T) {
	assert := assert.New(t)
	xmlstring := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<network href="/ovirt-engine/api/clusters/ee0a136a-941d-11ea-bdb1-525400d65c05/networks/bf9915f3-c4e8-400b-9d38-17e6772bcd5d" id="bf9915f3-c4e8-400b-9d38-17e6772bcd5d">
	<name>ocp</name>
	<description/>
	<comment/>
	<display>false</display>
	<dns_resolver_configuration>
		<name_servers>
			<name_server>8.8.8.8</name_server>
		</name_servers>
	</dns_resolver_configuration>
	<mtu>0</mtu>
	<required>true</required>
	<status>non_operational</status>
	<stp>false</stp>
	<usages>
	<usage>vm</usage>
	</usages>
	<vdsm_name>ocp</vdsm_name>
	<cluster href="/ovirt-engine/api/clusters/ee0a136a-941d-11ea-bdb1-525400d65c05" id="ee0a136a-941d-11ea-bdb1-525400d65c05"/>
	<data_center href="/ovirt-engine/api/datacenters/ee081c4a-941d-11ea-9cf8-525400d65c05" id="ee081c4a-941d-11ea-9cf8-525400d65c05"/>
</network>
`

	reader := NewXMLReader([]byte(xmlstring))
	network, err := XMLNetworkReadOne(reader, nil, "")
	assert.Nil(err)
	mtu, ok := network.Mtu()
	assert.True(ok)
	assert.Equal(int64(0), mtu)
}

func TestNetworkReadMany(t *testing.T) {
	assert := assert.New(t)
	xmlstring := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<networks>
	<network href="/ovirt-engine/api/clusters/ee0a136a-941d-11ea-bdb1-525400d65c05/networks/acfcf2fa-dd36-4047-9da9-5e349516a679" id="acfcf2fa-dd36-4047-9da9-5e349516a679">
		<name>lab</name>
		<description>Lab network</description>
		<comment/>
		<display>false</display>
		<mtu>0</mtu>
		<required>true</required>
		<status>non_operational</status>
		<stp>false</stp>
		<usages>
		<usage>vm</usage>
		</usages>
		<vdsm_name>lab</vdsm_name>
		<cluster href="/ovirt-engine/api/clusters/ee0a136a-941d-11ea-bdb1-525400d65c05" id="ee0a136a-941d-11ea-bdb1-525400d65c05"/>
		<data_center href="/ovirt-engine/api/datacenters/ee081c4a-941d-11ea-9cf8-525400d65c05" id="ee081c4a-941d-11ea-9cf8-525400d65c05"/>
	</network>
	<network href="/ovirt-engine/api/clusters/ee0a136a-941d-11ea-bdb1-525400d65c05/networks/bf9915f3-c4e8-400b-9d38-17e6772bcd5d" id="bf9915f3-c4e8-400b-9d38-17e6772bcd5d">
		<name>ocp</name>
		<description/>
		<comment/>
		<display>false</display>
		<dns_resolver_configuration>
			<name_servers>
				<name_server>8.8.8.8</name_server>
			</name_servers>
		</dns_resolver_configuration>
		<mtu>0</mtu>
		<required>true</required>
		<status>non_operational</status>
		<stp>false</stp>
		<usages>
		<usage>vm</usage>
		</usages>
		<vdsm_name>ocp</vdsm_name>
		<cluster href="/ovirt-engine/api/clusters/ee0a136a-941d-11ea-bdb1-525400d65c05" id="ee0a136a-941d-11ea-bdb1-525400d65c05"/>
		<data_center href="/ovirt-engine/api/datacenters/ee081c4a-941d-11ea-9cf8-525400d65c05" id="ee081c4a-941d-11ea-9cf8-525400d65c05"/>
	</network>
	<network href="/ovirt-engine/api/clusters/ee0a136a-941d-11ea-bdb1-525400d65c05/networks/d73e5b4c-f527-4057-b071-65b60302286c" id="d73e5b4c-f527-4057-b071-65b60302286c">
		<name>ocpnr</name>
		<description/>
		<comment/>
		<display>false</display>
		<mtu>0</mtu>
		<required>true</required>
		<status>non_operational</status>
		<stp>false</stp>
		<usages>
		<usage>vm</usage>
		</usages>
		<vdsm_name>ocpnr</vdsm_name>
		<cluster href="/ovirt-engine/api/clusters/ee0a136a-941d-11ea-bdb1-525400d65c05" id="ee0a136a-941d-11ea-bdb1-525400d65c05"/>
		<data_center href="/ovirt-engine/api/datacenters/ee081c4a-941d-11ea-9cf8-525400d65c05" id="ee081c4a-941d-11ea-9cf8-525400d65c05"/>
	</network>
</networks>
`
	reader := NewXMLReader([]byte(xmlstring))
	networks, err := XMLNetworkReadMany(reader, nil)
	assert.Nil(err)
	assert.NotNil(networks)
	assert.Equal(3, len(networks.Slice()))
}
