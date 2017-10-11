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
package ovirtsdk

import (
	"fmt"
)

func XMLSeLinuxWriteOne(writer *XMLWriter, object *SeLinux, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "se_linux"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Mode(); ok {
		XMLSeLinuxModeWriteOne(writer, r, "mode")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLSeLinuxWriteMany(writer *XMLWriter, structSlice *SeLinuxSlice, plural, singular string) error {
	if plural == "" {
		plural = "se_linuxs"
	}
	if singular == "" {
		singular = "se_linux"
	}
	writer.WriteStart("", "se_linuxs", nil)
	for _, o := range structSlice.Slice() {
		XMLSeLinuxWriteOne(writer, &o, "se_linux")
	}
	writer.WriteEnd("se_linuxs")
	return nil
}

func XMLGraphicsConsoleWriteOne(writer *XMLWriter, object *GraphicsConsole, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "graphics_console"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Address(); ok {
		writer.WriteCharacter("address", r)
	}
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.InstanceType(); ok {
		XMLInstanceTypeWriteOne(writer, r, "instance_type")
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Port(); ok {
		writer.WriteInt64("port", r)
	}
	if r, ok := object.Protocol(); ok {
		XMLGraphicsTypeWriteOne(writer, r, "protocol")
	}
	if r, ok := object.Template(); ok {
		XMLTemplateWriteOne(writer, r, "template")
	}
	if r, ok := object.TlsPort(); ok {
		writer.WriteInt64("tls_port", r)
	}
	if r, ok := object.Vm(); ok {
		XMLVmWriteOne(writer, r, "vm")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLGraphicsConsoleWriteMany(writer *XMLWriter, structSlice *GraphicsConsoleSlice, plural, singular string) error {
	if plural == "" {
		plural = "graphics_consoles"
	}
	if singular == "" {
		singular = "graphics_console"
	}
	writer.WriteStart("", "graphics_consoles", nil)
	for _, o := range structSlice.Slice() {
		XMLGraphicsConsoleWriteOne(writer, &o, "graphics_console")
	}
	writer.WriteEnd("graphics_consoles")
	return nil
}

func XMLPowerManagementWriteOne(writer *XMLWriter, object *PowerManagement, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "power_management"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Address(); ok {
		writer.WriteCharacter("address", r)
	}
	if r, ok := object.Agents(); ok {
		XMLAgentWriteMany(writer, r, "agents", "agent")
	}
	if r, ok := object.AutomaticPmEnabled(); ok {
		writer.WriteBool("automatic_pm_enabled", r)
	}
	if r, ok := object.Enabled(); ok {
		writer.WriteBool("enabled", r)
	}
	if r, ok := object.KdumpDetection(); ok {
		writer.WriteBool("kdump_detection", r)
	}
	if r, ok := object.Options(); ok {
		XMLOptionWriteMany(writer, r, "options", "option")
	}
	if r, ok := object.Password(); ok {
		writer.WriteCharacter("password", r)
	}
	if r, ok := object.PmProxies(); ok {
		XMLPmProxyWriteMany(writer, r, "pm_proxies", "pm_proxy")
	}
	if r, ok := object.Status(); ok {
		XMLPowerManagementStatusWriteOne(writer, r, "status")
	}
	if r, ok := object.Type(); ok {
		writer.WriteCharacter("type", r)
	}
	if r, ok := object.Username(); ok {
		writer.WriteCharacter("username", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLPowerManagementWriteMany(writer *XMLWriter, structSlice *PowerManagementSlice, plural, singular string) error {
	if plural == "" {
		plural = "power_managements"
	}
	if singular == "" {
		singular = "power_management"
	}
	writer.WriteStart("", "power_managements", nil)
	for _, o := range structSlice.Slice() {
		XMLPowerManagementWriteOne(writer, &o, "power_management")
	}
	writer.WriteEnd("power_managements")
	return nil
}

func XMLBondingWriteOne(writer *XMLWriter, object *Bonding, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "bonding"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.ActiveSlave(); ok {
		XMLHostNicWriteOne(writer, r, "active_slave")
	}
	if r, ok := object.AdPartnerMac(); ok {
		XMLMacWriteOne(writer, r, "ad_partner_mac")
	}
	if r, ok := object.Options(); ok {
		XMLOptionWriteMany(writer, r, "options", "option")
	}
	if r, ok := object.Slaves(); ok {
		XMLHostNicWriteMany(writer, r, "slaves", "host_nic")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLBondingWriteMany(writer *XMLWriter, structSlice *BondingSlice, plural, singular string) error {
	if plural == "" {
		plural = "bondings"
	}
	if singular == "" {
		singular = "bonding"
	}
	writer.WriteStart("", "bondings", nil)
	for _, o := range structSlice.Slice() {
		XMLBondingWriteOne(writer, &o, "bonding")
	}
	writer.WriteEnd("bondings")
	return nil
}

func XMLVnicPassThroughWriteOne(writer *XMLWriter, object *VnicPassThrough, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "vnic_pass_through"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Mode(); ok {
		XMLVnicPassThroughModeWriteOne(writer, r, "mode")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLVnicPassThroughWriteMany(writer *XMLWriter, structSlice *VnicPassThroughSlice, plural, singular string) error {
	if plural == "" {
		plural = "vnic_pass_throughs"
	}
	if singular == "" {
		singular = "vnic_pass_through"
	}
	writer.WriteStart("", "vnic_pass_throughs", nil)
	for _, o := range structSlice.Slice() {
		XMLVnicPassThroughWriteOne(writer, &o, "vnic_pass_through")
	}
	writer.WriteEnd("vnic_pass_throughs")
	return nil
}

func XMLPmProxyWriteOne(writer *XMLWriter, object *PmProxy, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "pm_proxy"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Type(); ok {
		XMLPmProxyTypeWriteOne(writer, r, "type")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLPmProxyWriteMany(writer *XMLWriter, structSlice *PmProxySlice, plural, singular string) error {
	if plural == "" {
		plural = "pm_proxies"
	}
	if singular == "" {
		singular = "pm_proxy"
	}
	writer.WriteStart("", "pm_proxies", nil)
	for _, o := range structSlice.Slice() {
		XMLPmProxyWriteOne(writer, &o, "pm_proxy")
	}
	writer.WriteEnd("pm_proxies")
	return nil
}

func XMLConsoleWriteOne(writer *XMLWriter, object *Console, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "console"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Enabled(); ok {
		writer.WriteBool("enabled", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLConsoleWriteMany(writer *XMLWriter, structSlice *ConsoleSlice, plural, singular string) error {
	if plural == "" {
		plural = "consoles"
	}
	if singular == "" {
		singular = "console"
	}
	writer.WriteStart("", "consoles", nil)
	for _, o := range structSlice.Slice() {
		XMLConsoleWriteOne(writer, &o, "console")
	}
	writer.WriteEnd("consoles")
	return nil
}

func XMLDiskAttachmentWriteOne(writer *XMLWriter, object *DiskAttachment, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "disk_attachment"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Active(); ok {
		writer.WriteBool("active", r)
	}
	if r, ok := object.Bootable(); ok {
		writer.WriteBool("bootable", r)
	}
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Disk(); ok {
		XMLDiskWriteOne(writer, r, "disk")
	}
	if r, ok := object.Interface(); ok {
		XMLDiskInterfaceWriteOne(writer, r, "interface")
	}
	if r, ok := object.LogicalName(); ok {
		writer.WriteCharacter("logical_name", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.PassDiscard(); ok {
		writer.WriteBool("pass_discard", r)
	}
	if r, ok := object.Template(); ok {
		XMLTemplateWriteOne(writer, r, "template")
	}
	if r, ok := object.UsesScsiReservation(); ok {
		writer.WriteBool("uses_scsi_reservation", r)
	}
	if r, ok := object.Vm(); ok {
		XMLVmWriteOne(writer, r, "vm")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLDiskAttachmentWriteMany(writer *XMLWriter, structSlice *DiskAttachmentSlice, plural, singular string) error {
	if plural == "" {
		plural = "disk_attachments"
	}
	if singular == "" {
		singular = "disk_attachment"
	}
	writer.WriteStart("", "disk_attachments", nil)
	for _, o := range structSlice.Slice() {
		XMLDiskAttachmentWriteOne(writer, &o, "disk_attachment")
	}
	writer.WriteEnd("disk_attachments")
	return nil
}

func XMLStorageConnectionWriteOne(writer *XMLWriter, object *StorageConnection, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "storage_connection"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Address(); ok {
		writer.WriteCharacter("address", r)
	}
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Host(); ok {
		XMLHostWriteOne(writer, r, "host")
	}
	if r, ok := object.MountOptions(); ok {
		writer.WriteCharacter("mount_options", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.NfsRetrans(); ok {
		writer.WriteInt64("nfs_retrans", r)
	}
	if r, ok := object.NfsTimeo(); ok {
		writer.WriteInt64("nfs_timeo", r)
	}
	if r, ok := object.NfsVersion(); ok {
		XMLNfsVersionWriteOne(writer, r, "nfs_version")
	}
	if r, ok := object.Password(); ok {
		writer.WriteCharacter("password", r)
	}
	if r, ok := object.Path(); ok {
		writer.WriteCharacter("path", r)
	}
	if r, ok := object.Port(); ok {
		writer.WriteInt64("port", r)
	}
	if r, ok := object.Portal(); ok {
		writer.WriteCharacter("portal", r)
	}
	if r, ok := object.Target(); ok {
		writer.WriteCharacter("target", r)
	}
	if r, ok := object.Type(); ok {
		XMLStorageTypeWriteOne(writer, r, "type")
	}
	if r, ok := object.Username(); ok {
		writer.WriteCharacter("username", r)
	}
	if r, ok := object.VfsType(); ok {
		writer.WriteCharacter("vfs_type", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLStorageConnectionWriteMany(writer *XMLWriter, structSlice *StorageConnectionSlice, plural, singular string) error {
	if plural == "" {
		plural = "storage_connections"
	}
	if singular == "" {
		singular = "storage_connection"
	}
	writer.WriteStart("", "storage_connections", nil)
	for _, o := range structSlice.Slice() {
		XMLStorageConnectionWriteOne(writer, &o, "storage_connection")
	}
	writer.WriteEnd("storage_connections")
	return nil
}

func XMLInstanceTypeWriteOne(writer *XMLWriter, object *InstanceType, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "instance_type"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Bios(); ok {
		XMLBiosWriteOne(writer, r, "bios")
	}
	if r, ok := object.Cdroms(); ok {
		XMLCdromWriteMany(writer, r, "cdroms", "cdrom")
	}
	if r, ok := object.Cluster(); ok {
		XMLClusterWriteOne(writer, r, "cluster")
	}
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Console(); ok {
		XMLConsoleWriteOne(writer, r, "console")
	}
	if r, ok := object.Cpu(); ok {
		XMLCpuWriteOne(writer, r, "cpu")
	}
	if r, ok := object.CpuProfile(); ok {
		XMLCpuProfileWriteOne(writer, r, "cpu_profile")
	}
	if r, ok := object.CpuShares(); ok {
		writer.WriteInt64("cpu_shares", r)
	}
	if r, ok := object.CreationTime(); ok {
		writer.WriteDate("creation_time", r)
	}
	if r, ok := object.CustomCompatibilityVersion(); ok {
		XMLVersionWriteOne(writer, r, "custom_compatibility_version")
	}
	if r, ok := object.CustomCpuModel(); ok {
		writer.WriteCharacter("custom_cpu_model", r)
	}
	if r, ok := object.CustomEmulatedMachine(); ok {
		writer.WriteCharacter("custom_emulated_machine", r)
	}
	if r, ok := object.CustomProperties(); ok {
		XMLCustomPropertyWriteMany(writer, r, "custom_properties", "custom_property")
	}
	if r, ok := object.DeleteProtected(); ok {
		writer.WriteBool("delete_protected", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.DiskAttachments(); ok {
		XMLDiskAttachmentWriteMany(writer, r, "disk_attachments", "disk_attachment")
	}
	if r, ok := object.Display(); ok {
		XMLDisplayWriteOne(writer, r, "display")
	}
	if r, ok := object.Domain(); ok {
		XMLDomainWriteOne(writer, r, "domain")
	}
	if r, ok := object.GraphicsConsoles(); ok {
		XMLGraphicsConsoleWriteMany(writer, r, "graphics_consoles", "graphics_console")
	}
	if r, ok := object.HighAvailability(); ok {
		XMLHighAvailabilityWriteOne(writer, r, "high_availability")
	}
	if r, ok := object.Initialization(); ok {
		XMLInitializationWriteOne(writer, r, "initialization")
	}
	if r, ok := object.Io(); ok {
		XMLIoWriteOne(writer, r, "io")
	}
	if r, ok := object.LargeIcon(); ok {
		XMLIconWriteOne(writer, r, "large_icon")
	}
	if r, ok := object.Lease(); ok {
		XMLStorageDomainLeaseWriteOne(writer, r, "lease")
	}
	if r, ok := object.Memory(); ok {
		writer.WriteInt64("memory", r)
	}
	if r, ok := object.MemoryPolicy(); ok {
		XMLMemoryPolicyWriteOne(writer, r, "memory_policy")
	}
	if r, ok := object.Migration(); ok {
		XMLMigrationOptionsWriteOne(writer, r, "migration")
	}
	if r, ok := object.MigrationDowntime(); ok {
		writer.WriteInt64("migration_downtime", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Nics(); ok {
		XMLNicWriteMany(writer, r, "nics", "nic")
	}
	if r, ok := object.Origin(); ok {
		writer.WriteCharacter("origin", r)
	}
	if r, ok := object.Os(); ok {
		XMLOperatingSystemWriteOne(writer, r, "os")
	}
	if r, ok := object.Permissions(); ok {
		XMLPermissionWriteMany(writer, r, "permissions", "permission")
	}
	if r, ok := object.Quota(); ok {
		XMLQuotaWriteOne(writer, r, "quota")
	}
	if r, ok := object.RngDevice(); ok {
		XMLRngDeviceWriteOne(writer, r, "rng_device")
	}
	if r, ok := object.SerialNumber(); ok {
		XMLSerialNumberWriteOne(writer, r, "serial_number")
	}
	if r, ok := object.SmallIcon(); ok {
		XMLIconWriteOne(writer, r, "small_icon")
	}
	if r, ok := object.SoundcardEnabled(); ok {
		writer.WriteBool("soundcard_enabled", r)
	}
	if r, ok := object.Sso(); ok {
		XMLSsoWriteOne(writer, r, "sso")
	}
	if r, ok := object.StartPaused(); ok {
		writer.WriteBool("start_paused", r)
	}
	if r, ok := object.Stateless(); ok {
		writer.WriteBool("stateless", r)
	}
	if r, ok := object.Status(); ok {
		XMLTemplateStatusWriteOne(writer, r, "status")
	}
	if r, ok := object.StorageDomain(); ok {
		XMLStorageDomainWriteOne(writer, r, "storage_domain")
	}
	if r, ok := object.Tags(); ok {
		XMLTagWriteMany(writer, r, "tags", "tag")
	}
	if r, ok := object.TimeZone(); ok {
		XMLTimeZoneWriteOne(writer, r, "time_zone")
	}
	if r, ok := object.TunnelMigration(); ok {
		writer.WriteBool("tunnel_migration", r)
	}
	if r, ok := object.Type(); ok {
		XMLVmTypeWriteOne(writer, r, "type")
	}
	if r, ok := object.Usb(); ok {
		XMLUsbWriteOne(writer, r, "usb")
	}
	if r, ok := object.Version(); ok {
		XMLTemplateVersionWriteOne(writer, r, "version")
	}
	if r, ok := object.VirtioScsi(); ok {
		XMLVirtioScsiWriteOne(writer, r, "virtio_scsi")
	}
	if r, ok := object.Vm(); ok {
		XMLVmWriteOne(writer, r, "vm")
	}
	if r, ok := object.Watchdogs(); ok {
		XMLWatchdogWriteMany(writer, r, "watchdogs", "watchdog")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLInstanceTypeWriteMany(writer *XMLWriter, structSlice *InstanceTypeSlice, plural, singular string) error {
	if plural == "" {
		plural = "instance_types"
	}
	if singular == "" {
		singular = "instance_type"
	}
	writer.WriteStart("", "instance_types", nil)
	for _, o := range structSlice.Slice() {
		XMLInstanceTypeWriteOne(writer, &o, "instance_type")
	}
	writer.WriteEnd("instance_types")
	return nil
}

func XMLPackageWriteOne(writer *XMLWriter, object *Package, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "package"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLPackageWriteMany(writer *XMLWriter, structSlice *PackageSlice, plural, singular string) error {
	if plural == "" {
		plural = "packages"
	}
	if singular == "" {
		singular = "package"
	}
	writer.WriteStart("", "packages", nil)
	for _, o := range structSlice.Slice() {
		XMLPackageWriteOne(writer, &o, "package")
	}
	writer.WriteEnd("packages")
	return nil
}

func XMLSshPublicKeyWriteOne(writer *XMLWriter, object *SshPublicKey, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "ssh_public_key"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Content(); ok {
		writer.WriteCharacter("content", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.User(); ok {
		XMLUserWriteOne(writer, r, "user")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLSshPublicKeyWriteMany(writer *XMLWriter, structSlice *SshPublicKeySlice, plural, singular string) error {
	if plural == "" {
		plural = "ssh_public_keys"
	}
	if singular == "" {
		singular = "ssh_public_key"
	}
	writer.WriteStart("", "ssh_public_keys", nil)
	for _, o := range structSlice.Slice() {
		XMLSshPublicKeyWriteOne(writer, &o, "ssh_public_key")
	}
	writer.WriteEnd("ssh_public_keys")
	return nil
}

func XMLOpenStackImageWriteOne(writer *XMLWriter, object *OpenStackImage, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "openstack_image"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.OpenstackImageProvider(); ok {
		XMLOpenStackImageProviderWriteOne(writer, r, "openstack_image_provider")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLOpenStackImageWriteMany(writer *XMLWriter, structSlice *OpenStackImageSlice, plural, singular string) error {
	if plural == "" {
		plural = "openstack_images"
	}
	if singular == "" {
		singular = "openstack_image"
	}
	writer.WriteStart("", "openstack_images", nil)
	for _, o := range structSlice.Slice() {
		XMLOpenStackImageWriteOne(writer, &o, "openstack_image")
	}
	writer.WriteEnd("openstack_images")
	return nil
}

func XMLProductWriteOne(writer *XMLWriter, object *Product, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "product"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLProductWriteMany(writer *XMLWriter, structSlice *ProductSlice, plural, singular string) error {
	if plural == "" {
		plural = "products"
	}
	if singular == "" {
		singular = "product"
	}
	writer.WriteStart("", "products", nil)
	for _, o := range structSlice.Slice() {
		XMLProductWriteOne(writer, &o, "product")
	}
	writer.WriteEnd("products")
	return nil
}

func XMLVcpuPinWriteOne(writer *XMLWriter, object *VcpuPin, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "vcpu_pin"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.CpuSet(); ok {
		writer.WriteCharacter("cpu_set", r)
	}
	if r, ok := object.Vcpu(); ok {
		writer.WriteInt64("vcpu", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLVcpuPinWriteMany(writer *XMLWriter, structSlice *VcpuPinSlice, plural, singular string) error {
	if plural == "" {
		plural = "vcpu_pins"
	}
	if singular == "" {
		singular = "vcpu_pin"
	}
	writer.WriteStart("", "vcpu_pins", nil)
	for _, o := range structSlice.Slice() {
		XMLVcpuPinWriteOne(writer, &o, "vcpu_pin")
	}
	writer.WriteEnd("vcpu_pins")
	return nil
}

func XMLVmBaseWriteOne(writer *XMLWriter, object *VmBase, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "vm_base"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Bios(); ok {
		XMLBiosWriteOne(writer, r, "bios")
	}
	if r, ok := object.Cluster(); ok {
		XMLClusterWriteOne(writer, r, "cluster")
	}
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Console(); ok {
		XMLConsoleWriteOne(writer, r, "console")
	}
	if r, ok := object.Cpu(); ok {
		XMLCpuWriteOne(writer, r, "cpu")
	}
	if r, ok := object.CpuProfile(); ok {
		XMLCpuProfileWriteOne(writer, r, "cpu_profile")
	}
	if r, ok := object.CpuShares(); ok {
		writer.WriteInt64("cpu_shares", r)
	}
	if r, ok := object.CreationTime(); ok {
		writer.WriteDate("creation_time", r)
	}
	if r, ok := object.CustomCompatibilityVersion(); ok {
		XMLVersionWriteOne(writer, r, "custom_compatibility_version")
	}
	if r, ok := object.CustomCpuModel(); ok {
		writer.WriteCharacter("custom_cpu_model", r)
	}
	if r, ok := object.CustomEmulatedMachine(); ok {
		writer.WriteCharacter("custom_emulated_machine", r)
	}
	if r, ok := object.CustomProperties(); ok {
		XMLCustomPropertyWriteMany(writer, r, "custom_properties", "custom_property")
	}
	if r, ok := object.DeleteProtected(); ok {
		writer.WriteBool("delete_protected", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Display(); ok {
		XMLDisplayWriteOne(writer, r, "display")
	}
	if r, ok := object.Domain(); ok {
		XMLDomainWriteOne(writer, r, "domain")
	}
	if r, ok := object.HighAvailability(); ok {
		XMLHighAvailabilityWriteOne(writer, r, "high_availability")
	}
	if r, ok := object.Initialization(); ok {
		XMLInitializationWriteOne(writer, r, "initialization")
	}
	if r, ok := object.Io(); ok {
		XMLIoWriteOne(writer, r, "io")
	}
	if r, ok := object.LargeIcon(); ok {
		XMLIconWriteOne(writer, r, "large_icon")
	}
	if r, ok := object.Lease(); ok {
		XMLStorageDomainLeaseWriteOne(writer, r, "lease")
	}
	if r, ok := object.Memory(); ok {
		writer.WriteInt64("memory", r)
	}
	if r, ok := object.MemoryPolicy(); ok {
		XMLMemoryPolicyWriteOne(writer, r, "memory_policy")
	}
	if r, ok := object.Migration(); ok {
		XMLMigrationOptionsWriteOne(writer, r, "migration")
	}
	if r, ok := object.MigrationDowntime(); ok {
		writer.WriteInt64("migration_downtime", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Origin(); ok {
		writer.WriteCharacter("origin", r)
	}
	if r, ok := object.Os(); ok {
		XMLOperatingSystemWriteOne(writer, r, "os")
	}
	if r, ok := object.Quota(); ok {
		XMLQuotaWriteOne(writer, r, "quota")
	}
	if r, ok := object.RngDevice(); ok {
		XMLRngDeviceWriteOne(writer, r, "rng_device")
	}
	if r, ok := object.SerialNumber(); ok {
		XMLSerialNumberWriteOne(writer, r, "serial_number")
	}
	if r, ok := object.SmallIcon(); ok {
		XMLIconWriteOne(writer, r, "small_icon")
	}
	if r, ok := object.SoundcardEnabled(); ok {
		writer.WriteBool("soundcard_enabled", r)
	}
	if r, ok := object.Sso(); ok {
		XMLSsoWriteOne(writer, r, "sso")
	}
	if r, ok := object.StartPaused(); ok {
		writer.WriteBool("start_paused", r)
	}
	if r, ok := object.Stateless(); ok {
		writer.WriteBool("stateless", r)
	}
	if r, ok := object.StorageDomain(); ok {
		XMLStorageDomainWriteOne(writer, r, "storage_domain")
	}
	if r, ok := object.TimeZone(); ok {
		XMLTimeZoneWriteOne(writer, r, "time_zone")
	}
	if r, ok := object.TunnelMigration(); ok {
		writer.WriteBool("tunnel_migration", r)
	}
	if r, ok := object.Type(); ok {
		XMLVmTypeWriteOne(writer, r, "type")
	}
	if r, ok := object.Usb(); ok {
		XMLUsbWriteOne(writer, r, "usb")
	}
	if r, ok := object.VirtioScsi(); ok {
		XMLVirtioScsiWriteOne(writer, r, "virtio_scsi")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLVmBaseWriteMany(writer *XMLWriter, structSlice *VmBaseSlice, plural, singular string) error {
	if plural == "" {
		plural = "vm_bases"
	}
	if singular == "" {
		singular = "vm_base"
	}
	writer.WriteStart("", "vm_bases", nil)
	for _, o := range structSlice.Slice() {
		XMLVmBaseWriteOne(writer, &o, "vm_base")
	}
	writer.WriteEnd("vm_bases")
	return nil
}

func XMLHookWriteOne(writer *XMLWriter, object *Hook, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "hook"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.EventName(); ok {
		writer.WriteCharacter("event_name", r)
	}
	if r, ok := object.Host(); ok {
		XMLHostWriteOne(writer, r, "host")
	}
	if r, ok := object.Md5(); ok {
		writer.WriteCharacter("md5", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLHookWriteMany(writer *XMLWriter, structSlice *HookSlice, plural, singular string) error {
	if plural == "" {
		plural = "hooks"
	}
	if singular == "" {
		singular = "hook"
	}
	writer.WriteStart("", "hooks", nil)
	for _, o := range structSlice.Slice() {
		XMLHookWriteOne(writer, &o, "hook")
	}
	writer.WriteEnd("hooks")
	return nil
}

func XMLVlanWriteOne(writer *XMLWriter, object *Vlan, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "vlan"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = writer.FormatInt64(r)
	}
	writer.WriteStart("", tag, attrs)
	writer.WriteEnd(tag)
	return nil
}

func XMLVlanWriteMany(writer *XMLWriter, structSlice *VlanSlice, plural, singular string) error {
	if plural == "" {
		plural = "vlans"
	}
	if singular == "" {
		singular = "vlan"
	}
	writer.WriteStart("", "vlans", nil)
	for _, o := range structSlice.Slice() {
		XMLVlanWriteOne(writer, &o, "vlan")
	}
	writer.WriteEnd("vlans")
	return nil
}

func XMLUsbWriteOne(writer *XMLWriter, object *Usb, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "usb"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Enabled(); ok {
		writer.WriteBool("enabled", r)
	}
	if r, ok := object.Type(); ok {
		XMLUsbTypeWriteOne(writer, r, "type")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLUsbWriteMany(writer *XMLWriter, structSlice *UsbSlice, plural, singular string) error {
	if plural == "" {
		plural = "usbs"
	}
	if singular == "" {
		singular = "usb"
	}
	writer.WriteStart("", "usbs", nil)
	for _, o := range structSlice.Slice() {
		XMLUsbWriteOne(writer, &o, "usb")
	}
	writer.WriteEnd("usbs")
	return nil
}

func XMLExternalComputeResourceWriteOne(writer *XMLWriter, object *ExternalComputeResource, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "external_compute_resource"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.ExternalHostProvider(); ok {
		XMLExternalHostProviderWriteOne(writer, r, "external_host_provider")
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Provider(); ok {
		writer.WriteCharacter("provider", r)
	}
	if r, ok := object.Url(); ok {
		writer.WriteCharacter("url", r)
	}
	if r, ok := object.User(); ok {
		writer.WriteCharacter("user", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLExternalComputeResourceWriteMany(writer *XMLWriter, structSlice *ExternalComputeResourceSlice, plural, singular string) error {
	if plural == "" {
		plural = "external_compute_resources"
	}
	if singular == "" {
		singular = "external_compute_resource"
	}
	writer.WriteStart("", "external_compute_resources", nil)
	for _, o := range structSlice.Slice() {
		XMLExternalComputeResourceWriteOne(writer, &o, "external_compute_resource")
	}
	writer.WriteEnd("external_compute_resources")
	return nil
}

func XMLAgentWriteOne(writer *XMLWriter, object *Agent, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "agent"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Address(); ok {
		writer.WriteCharacter("address", r)
	}
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Concurrent(); ok {
		writer.WriteBool("concurrent", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.EncryptOptions(); ok {
		writer.WriteBool("encrypt_options", r)
	}
	if r, ok := object.Host(); ok {
		XMLHostWriteOne(writer, r, "host")
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Options(); ok {
		XMLOptionWriteMany(writer, r, "options", "option")
	}
	if r, ok := object.Order(); ok {
		writer.WriteInt64("order", r)
	}
	if r, ok := object.Password(); ok {
		writer.WriteCharacter("password", r)
	}
	if r, ok := object.Port(); ok {
		writer.WriteInt64("port", r)
	}
	if r, ok := object.Type(); ok {
		writer.WriteCharacter("type", r)
	}
	if r, ok := object.Username(); ok {
		writer.WriteCharacter("username", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLAgentWriteMany(writer *XMLWriter, structSlice *AgentSlice, plural, singular string) error {
	if plural == "" {
		plural = "agents"
	}
	if singular == "" {
		singular = "agent"
	}
	writer.WriteStart("", "agents", nil)
	for _, o := range structSlice.Slice() {
		XMLAgentWriteOne(writer, &o, "agent")
	}
	writer.WriteEnd("agents")
	return nil
}

func XMLEntityProfileDetailWriteOne(writer *XMLWriter, object *EntityProfileDetail, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "entity_profile_detail"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.ProfileDetails(); ok {
		XMLProfileDetailWriteMany(writer, r, "profile_details", "profile_detail")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLEntityProfileDetailWriteMany(writer *XMLWriter, structSlice *EntityProfileDetailSlice, plural, singular string) error {
	if plural == "" {
		plural = "entity_profile_details"
	}
	if singular == "" {
		singular = "entity_profile_detail"
	}
	writer.WriteStart("", "entity_profile_details", nil)
	for _, o := range structSlice.Slice() {
		XMLEntityProfileDetailWriteOne(writer, &o, "entity_profile_detail")
	}
	writer.WriteEnd("entity_profile_details")
	return nil
}

func XMLHostDeviceWriteOne(writer *XMLWriter, object *HostDevice, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "host_device"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Capability(); ok {
		writer.WriteCharacter("capability", r)
	}
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Host(); ok {
		XMLHostWriteOne(writer, r, "host")
	}
	if r, ok := object.IommuGroup(); ok {
		writer.WriteInt64("iommu_group", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.ParentDevice(); ok {
		XMLHostDeviceWriteOne(writer, r, "parent_device")
	}
	if r, ok := object.PhysicalFunction(); ok {
		XMLHostDeviceWriteOne(writer, r, "physical_function")
	}
	if r, ok := object.Placeholder(); ok {
		writer.WriteBool("placeholder", r)
	}
	if r, ok := object.Product(); ok {
		XMLProductWriteOne(writer, r, "product")
	}
	if r, ok := object.Vendor(); ok {
		XMLVendorWriteOne(writer, r, "vendor")
	}
	if r, ok := object.VirtualFunctions(); ok {
		writer.WriteInt64("virtual_functions", r)
	}
	if r, ok := object.Vm(); ok {
		XMLVmWriteOne(writer, r, "vm")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLHostDeviceWriteMany(writer *XMLWriter, structSlice *HostDeviceSlice, plural, singular string) error {
	if plural == "" {
		plural = "host_devices"
	}
	if singular == "" {
		singular = "host_device"
	}
	writer.WriteStart("", "host_devices", nil)
	for _, o := range structSlice.Slice() {
		XMLHostDeviceWriteOne(writer, &o, "host_device")
	}
	writer.WriteEnd("host_devices")
	return nil
}

func XMLCpuTuneWriteOne(writer *XMLWriter, object *CpuTune, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "cpu_tune"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.VcpuPins(); ok {
		XMLVcpuPinWriteMany(writer, r, "vcpu_pins", "vcpu_pin")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLCpuTuneWriteMany(writer *XMLWriter, structSlice *CpuTuneSlice, plural, singular string) error {
	if plural == "" {
		plural = "cpu_tunes"
	}
	if singular == "" {
		singular = "cpu_tune"
	}
	writer.WriteStart("", "cpu_tunes", nil)
	for _, o := range structSlice.Slice() {
		XMLCpuTuneWriteOne(writer, &o, "cpu_tune")
	}
	writer.WriteEnd("cpu_tunes")
	return nil
}

func XMLDomainWriteOne(writer *XMLWriter, object *Domain, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "domain"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Groups(); ok {
		XMLGroupWriteMany(writer, r, "groups", "group")
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.User(); ok {
		XMLUserWriteOne(writer, r, "user")
	}
	if r, ok := object.Users(); ok {
		XMLUserWriteMany(writer, r, "users", "user")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLDomainWriteMany(writer *XMLWriter, structSlice *DomainSlice, plural, singular string) error {
	if plural == "" {
		plural = "domains"
	}
	if singular == "" {
		singular = "domain"
	}
	writer.WriteStart("", "domains", nil)
	for _, o := range structSlice.Slice() {
		XMLDomainWriteOne(writer, &o, "domain")
	}
	writer.WriteEnd("domains")
	return nil
}

func XMLProductInfoWriteOne(writer *XMLWriter, object *ProductInfo, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "product_info"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Vendor(); ok {
		writer.WriteCharacter("vendor", r)
	}
	if r, ok := object.Version(); ok {
		XMLVersionWriteOne(writer, r, "version")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLProductInfoWriteMany(writer *XMLWriter, structSlice *ProductInfoSlice, plural, singular string) error {
	if plural == "" {
		plural = "product_infos"
	}
	if singular == "" {
		singular = "product_info"
	}
	writer.WriteStart("", "product_infos", nil)
	for _, o := range structSlice.Slice() {
		XMLProductInfoWriteOne(writer, &o, "product_info")
	}
	writer.WriteEnd("product_infos")
	return nil
}

func XMLVmPoolWriteOne(writer *XMLWriter, object *VmPool, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "vm_pool"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.AutoStorageSelect(); ok {
		writer.WriteBool("auto_storage_select", r)
	}
	if r, ok := object.Cluster(); ok {
		XMLClusterWriteOne(writer, r, "cluster")
	}
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Display(); ok {
		XMLDisplayWriteOne(writer, r, "display")
	}
	if r, ok := object.InstanceType(); ok {
		XMLInstanceTypeWriteOne(writer, r, "instance_type")
	}
	if r, ok := object.MaxUserVms(); ok {
		writer.WriteInt64("max_user_vms", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Permissions(); ok {
		XMLPermissionWriteMany(writer, r, "permissions", "permission")
	}
	if r, ok := object.PrestartedVms(); ok {
		writer.WriteInt64("prestarted_vms", r)
	}
	if r, ok := object.RngDevice(); ok {
		XMLRngDeviceWriteOne(writer, r, "rng_device")
	}
	if r, ok := object.Size(); ok {
		writer.WriteInt64("size", r)
	}
	if r, ok := object.SoundcardEnabled(); ok {
		writer.WriteBool("soundcard_enabled", r)
	}
	if r, ok := object.Stateful(); ok {
		writer.WriteBool("stateful", r)
	}
	if r, ok := object.Template(); ok {
		XMLTemplateWriteOne(writer, r, "template")
	}
	if r, ok := object.Type(); ok {
		XMLVmPoolTypeWriteOne(writer, r, "type")
	}
	if r, ok := object.UseLatestTemplateVersion(); ok {
		writer.WriteBool("use_latest_template_version", r)
	}
	if r, ok := object.Vm(); ok {
		XMLVmWriteOne(writer, r, "vm")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLVmPoolWriteMany(writer *XMLWriter, structSlice *VmPoolSlice, plural, singular string) error {
	if plural == "" {
		plural = "vm_pools"
	}
	if singular == "" {
		singular = "vm_pool"
	}
	writer.WriteStart("", "vm_pools", nil)
	for _, o := range structSlice.Slice() {
		XMLVmPoolWriteOne(writer, &o, "vm_pool")
	}
	writer.WriteEnd("vm_pools")
	return nil
}

func XMLUnmanagedNetworkWriteOne(writer *XMLWriter, object *UnmanagedNetwork, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "unmanaged_network"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Host(); ok {
		XMLHostWriteOne(writer, r, "host")
	}
	if r, ok := object.HostNic(); ok {
		XMLHostNicWriteOne(writer, r, "host_nic")
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLUnmanagedNetworkWriteMany(writer *XMLWriter, structSlice *UnmanagedNetworkSlice, plural, singular string) error {
	if plural == "" {
		plural = "unmanaged_networks"
	}
	if singular == "" {
		singular = "unmanaged_network"
	}
	writer.WriteStart("", "unmanaged_networks", nil)
	for _, o := range structSlice.Slice() {
		XMLUnmanagedNetworkWriteOne(writer, &o, "unmanaged_network")
	}
	writer.WriteEnd("unmanaged_networks")
	return nil
}

func XMLGlusterClientWriteOne(writer *XMLWriter, object *GlusterClient, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "gluster_client"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.BytesRead(); ok {
		writer.WriteInt64("bytes_read", r)
	}
	if r, ok := object.BytesWritten(); ok {
		writer.WriteInt64("bytes_written", r)
	}
	if r, ok := object.ClientPort(); ok {
		writer.WriteInt64("client_port", r)
	}
	if r, ok := object.HostName(); ok {
		writer.WriteCharacter("host_name", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLGlusterClientWriteMany(writer *XMLWriter, structSlice *GlusterClientSlice, plural, singular string) error {
	if plural == "" {
		plural = "gluster_clients"
	}
	if singular == "" {
		singular = "gluster_client"
	}
	writer.WriteStart("", "gluster_clients", nil)
	for _, o := range structSlice.Slice() {
		XMLGlusterClientWriteOne(writer, &o, "gluster_client")
	}
	writer.WriteEnd("gluster_clients")
	return nil
}

func XMLRateWriteOne(writer *XMLWriter, object *Rate, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "rate"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Bytes(); ok {
		writer.WriteInt64("bytes", r)
	}
	if r, ok := object.Period(); ok {
		writer.WriteInt64("period", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLRateWriteMany(writer *XMLWriter, structSlice *RateSlice, plural, singular string) error {
	if plural == "" {
		plural = "rates"
	}
	if singular == "" {
		singular = "rate"
	}
	writer.WriteStart("", "rates", nil)
	for _, o := range structSlice.Slice() {
		XMLRateWriteOne(writer, &o, "rate")
	}
	writer.WriteEnd("rates")
	return nil
}

func XMLAffinityLabelWriteOne(writer *XMLWriter, object *AffinityLabel, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "affinity_label"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Hosts(); ok {
		XMLHostWriteMany(writer, r, "hosts", "host")
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.ReadOnly(); ok {
		writer.WriteBool("read_only", r)
	}
	if r, ok := object.Vms(); ok {
		XMLVmWriteMany(writer, r, "vms", "vm")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLAffinityLabelWriteMany(writer *XMLWriter, structSlice *AffinityLabelSlice, plural, singular string) error {
	if plural == "" {
		plural = "affinity_labels"
	}
	if singular == "" {
		singular = "affinity_label"
	}
	writer.WriteStart("", "affinity_labels", nil)
	for _, o := range structSlice.Slice() {
		XMLAffinityLabelWriteOne(writer, &o, "affinity_label")
	}
	writer.WriteEnd("affinity_labels")
	return nil
}

func XMLAffinityRuleWriteOne(writer *XMLWriter, object *AffinityRule, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "affinity_rule"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Enabled(); ok {
		writer.WriteBool("enabled", r)
	}
	if r, ok := object.Enforcing(); ok {
		writer.WriteBool("enforcing", r)
	}
	if r, ok := object.Positive(); ok {
		writer.WriteBool("positive", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLAffinityRuleWriteMany(writer *XMLWriter, structSlice *AffinityRuleSlice, plural, singular string) error {
	if plural == "" {
		plural = "affinity_rules"
	}
	if singular == "" {
		singular = "affinity_rule"
	}
	writer.WriteStart("", "affinity_rules", nil)
	for _, o := range structSlice.Slice() {
		XMLAffinityRuleWriteOne(writer, &o, "affinity_rule")
	}
	writer.WriteEnd("affinity_rules")
	return nil
}

func XMLKatelloErratumWriteOne(writer *XMLWriter, object *KatelloErratum, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "katello_erratum"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Host(); ok {
		XMLHostWriteOne(writer, r, "host")
	}
	if r, ok := object.Issued(); ok {
		writer.WriteDate("issued", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Packages(); ok {
		XMLPackageWriteMany(writer, r, "packages", "package")
	}
	if r, ok := object.Severity(); ok {
		writer.WriteCharacter("severity", r)
	}
	if r, ok := object.Solution(); ok {
		writer.WriteCharacter("solution", r)
	}
	if r, ok := object.Summary(); ok {
		writer.WriteCharacter("summary", r)
	}
	if r, ok := object.Title(); ok {
		writer.WriteCharacter("title", r)
	}
	if r, ok := object.Type(); ok {
		writer.WriteCharacter("type", r)
	}
	if r, ok := object.Vm(); ok {
		XMLVmWriteOne(writer, r, "vm")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLKatelloErratumWriteMany(writer *XMLWriter, structSlice *KatelloErratumSlice, plural, singular string) error {
	if plural == "" {
		plural = "katello_errata"
	}
	if singular == "" {
		singular = "katello_erratum"
	}
	writer.WriteStart("", "katello_errata", nil)
	for _, o := range structSlice.Slice() {
		XMLKatelloErratumWriteOne(writer, &o, "katello_erratum")
	}
	writer.WriteEnd("katello_errata")
	return nil
}

func XMLClusterWriteOne(writer *XMLWriter, object *Cluster, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "cluster"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.AffinityGroups(); ok {
		XMLAffinityGroupWriteMany(writer, r, "affinity_groups", "affinity_group")
	}
	if r, ok := object.BallooningEnabled(); ok {
		writer.WriteBool("ballooning_enabled", r)
	}
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Cpu(); ok {
		XMLCpuWriteOne(writer, r, "cpu")
	}
	if r, ok := object.CpuProfiles(); ok {
		XMLCpuProfileWriteMany(writer, r, "cpu_profiles", "cpu_profile")
	}
	if r, ok := object.CustomSchedulingPolicyProperties(); ok {
		XMLPropertyWriteMany(writer, r, "custom_scheduling_policy_properties", "property")
	}
	if r, ok := object.DataCenter(); ok {
		XMLDataCenterWriteOne(writer, r, "data_center")
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Display(); ok {
		XMLDisplayWriteOne(writer, r, "display")
	}
	if r, ok := object.ErrorHandling(); ok {
		XMLErrorHandlingWriteOne(writer, r, "error_handling")
	}
	if r, ok := object.FencingPolicy(); ok {
		XMLFencingPolicyWriteOne(writer, r, "fencing_policy")
	}
	if r, ok := object.GlusterHooks(); ok {
		XMLGlusterHookWriteMany(writer, r, "gluster_hooks", "gluster_hook")
	}
	if r, ok := object.GlusterService(); ok {
		writer.WriteBool("gluster_service", r)
	}
	if r, ok := object.GlusterTunedProfile(); ok {
		writer.WriteCharacter("gluster_tuned_profile", r)
	}
	if r, ok := object.GlusterVolumes(); ok {
		XMLGlusterVolumeWriteMany(writer, r, "gluster_volumes", "gluster_volume")
	}
	if r, ok := object.HaReservation(); ok {
		writer.WriteBool("ha_reservation", r)
	}
	if r, ok := object.Ksm(); ok {
		XMLKsmWriteOne(writer, r, "ksm")
	}
	if r, ok := object.MacPool(); ok {
		XMLMacPoolWriteOne(writer, r, "mac_pool")
	}
	if r, ok := object.MaintenanceReasonRequired(); ok {
		writer.WriteBool("maintenance_reason_required", r)
	}
	if r, ok := object.ManagementNetwork(); ok {
		XMLNetworkWriteOne(writer, r, "management_network")
	}
	if r, ok := object.MemoryPolicy(); ok {
		XMLMemoryPolicyWriteOne(writer, r, "memory_policy")
	}
	if r, ok := object.Migration(); ok {
		XMLMigrationOptionsWriteOne(writer, r, "migration")
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.NetworkFilters(); ok {
		XMLNetworkFilterWriteMany(writer, r, "network_filters", "network_filter")
	}
	if r, ok := object.Networks(); ok {
		XMLNetworkWriteMany(writer, r, "networks", "network")
	}
	if r, ok := object.OptionalReason(); ok {
		writer.WriteBool("optional_reason", r)
	}
	if r, ok := object.Permissions(); ok {
		XMLPermissionWriteMany(writer, r, "permissions", "permission")
	}
	if r, ok := object.RequiredRngSources(); ok {
		XMLRngSourceWriteMany(writer, r, "required_rng_sources", "rng_source")
	}
	if r, ok := object.SchedulingPolicy(); ok {
		XMLSchedulingPolicyWriteOne(writer, r, "scheduling_policy")
	}
	if r, ok := object.SerialNumber(); ok {
		XMLSerialNumberWriteOne(writer, r, "serial_number")
	}
	if r, ok := object.SupportedVersions(); ok {
		XMLVersionWriteMany(writer, r, "supported_versions", "version")
	}
	if r, ok := object.SwitchType(); ok {
		XMLSwitchTypeWriteOne(writer, r, "switch_type")
	}
	if r, ok := object.ThreadsAsCores(); ok {
		writer.WriteBool("threads_as_cores", r)
	}
	if r, ok := object.TrustedService(); ok {
		writer.WriteBool("trusted_service", r)
	}
	if r, ok := object.TunnelMigration(); ok {
		writer.WriteBool("tunnel_migration", r)
	}
	if r, ok := object.Version(); ok {
		XMLVersionWriteOne(writer, r, "version")
	}
	if r, ok := object.VirtService(); ok {
		writer.WriteBool("virt_service", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLClusterWriteMany(writer *XMLWriter, structSlice *ClusterSlice, plural, singular string) error {
	if plural == "" {
		plural = "clusters"
	}
	if singular == "" {
		singular = "cluster"
	}
	writer.WriteStart("", "clusters", nil)
	for _, o := range structSlice.Slice() {
		XMLClusterWriteOne(writer, &o, "cluster")
	}
	writer.WriteEnd("clusters")
	return nil
}

func XMLDnsResolverConfigurationWriteOne(writer *XMLWriter, object *DnsResolverConfiguration, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "dns_resolver_configuration"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.NameServers(); ok {
		writer.WriteCharacters("name_servers", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLDnsResolverConfigurationWriteMany(writer *XMLWriter, structSlice *DnsResolverConfigurationSlice, plural, singular string) error {
	if plural == "" {
		plural = "dns_resolver_configurations"
	}
	if singular == "" {
		singular = "dns_resolver_configuration"
	}
	writer.WriteStart("", "dns_resolver_configurations", nil)
	for _, o := range structSlice.Slice() {
		XMLDnsResolverConfigurationWriteOne(writer, &o, "dns_resolver_configuration")
	}
	writer.WriteEnd("dns_resolver_configurations")
	return nil
}

func XMLExternalProviderWriteOne(writer *XMLWriter, object *ExternalProvider, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "external_provider"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.AuthenticationUrl(); ok {
		writer.WriteCharacter("authentication_url", r)
	}
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Password(); ok {
		writer.WriteCharacter("password", r)
	}
	if r, ok := object.Properties(); ok {
		XMLPropertyWriteMany(writer, r, "properties", "property")
	}
	if r, ok := object.RequiresAuthentication(); ok {
		writer.WriteBool("requires_authentication", r)
	}
	if r, ok := object.Url(); ok {
		writer.WriteCharacter("url", r)
	}
	if r, ok := object.Username(); ok {
		writer.WriteCharacter("username", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLExternalProviderWriteMany(writer *XMLWriter, structSlice *ExternalProviderSlice, plural, singular string) error {
	if plural == "" {
		plural = "external_providers"
	}
	if singular == "" {
		singular = "external_provider"
	}
	writer.WriteStart("", "external_providers", nil)
	for _, o := range structSlice.Slice() {
		XMLExternalProviderWriteOne(writer, &o, "external_provider")
	}
	writer.WriteEnd("external_providers")
	return nil
}

func XMLExternalHostProviderWriteOne(writer *XMLWriter, object *ExternalHostProvider, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "external_host_provider"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.AuthenticationUrl(); ok {
		writer.WriteCharacter("authentication_url", r)
	}
	if r, ok := object.Certificates(); ok {
		XMLCertificateWriteMany(writer, r, "certificates", "certificate")
	}
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.ComputeResources(); ok {
		XMLExternalComputeResourceWriteMany(writer, r, "compute_resources", "external_compute_resource")
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.DiscoveredHosts(); ok {
		XMLExternalDiscoveredHostWriteMany(writer, r, "discovered_hosts", "external_discovered_host")
	}
	if r, ok := object.HostGroups(); ok {
		XMLExternalHostGroupWriteMany(writer, r, "host_groups", "external_host_group")
	}
	if r, ok := object.Hosts(); ok {
		XMLHostWriteMany(writer, r, "hosts", "host")
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Password(); ok {
		writer.WriteCharacter("password", r)
	}
	if r, ok := object.Properties(); ok {
		XMLPropertyWriteMany(writer, r, "properties", "property")
	}
	if r, ok := object.RequiresAuthentication(); ok {
		writer.WriteBool("requires_authentication", r)
	}
	if r, ok := object.Url(); ok {
		writer.WriteCharacter("url", r)
	}
	if r, ok := object.Username(); ok {
		writer.WriteCharacter("username", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLExternalHostProviderWriteMany(writer *XMLWriter, structSlice *ExternalHostProviderSlice, plural, singular string) error {
	if plural == "" {
		plural = "external_host_providers"
	}
	if singular == "" {
		singular = "external_host_provider"
	}
	writer.WriteStart("", "external_host_providers", nil)
	for _, o := range structSlice.Slice() {
		XMLExternalHostProviderWriteOne(writer, &o, "external_host_provider")
	}
	writer.WriteEnd("external_host_providers")
	return nil
}

func XMLGlusterVolumeProfileDetailsWriteOne(writer *XMLWriter, object *GlusterVolumeProfileDetails, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "gluster_volume_profile_details"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.BrickProfileDetails(); ok {
		XMLBrickProfileDetailWriteMany(writer, r, "brick_profile_details", "brick_profile_detail")
	}
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.NfsProfileDetails(); ok {
		XMLNfsProfileDetailWriteMany(writer, r, "nfs_profile_details", "nfs_profile_detail")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLGlusterVolumeProfileDetailsWriteMany(writer *XMLWriter, structSlice *GlusterVolumeProfileDetailsSlice, plural, singular string) error {
	if plural == "" {
		plural = "gluster_volume_profile_detailss"
	}
	if singular == "" {
		singular = "gluster_volume_profile_details"
	}
	writer.WriteStart("", "gluster_volume_profile_detailss", nil)
	for _, o := range structSlice.Slice() {
		XMLGlusterVolumeProfileDetailsWriteOne(writer, &o, "gluster_volume_profile_details")
	}
	writer.WriteEnd("gluster_volume_profile_detailss")
	return nil
}

func XMLWatchdogWriteOne(writer *XMLWriter, object *Watchdog, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "watchdog"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Action(); ok {
		XMLWatchdogActionWriteOne(writer, r, "action")
	}
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.InstanceType(); ok {
		XMLInstanceTypeWriteOne(writer, r, "instance_type")
	}
	if r, ok := object.Model(); ok {
		XMLWatchdogModelWriteOne(writer, r, "model")
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Template(); ok {
		XMLTemplateWriteOne(writer, r, "template")
	}
	if r, ok := object.Vm(); ok {
		XMLVmWriteOne(writer, r, "vm")
	}
	if r, ok := object.Vms(); ok {
		XMLVmWriteMany(writer, r, "vms", "vm")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLWatchdogWriteMany(writer *XMLWriter, structSlice *WatchdogSlice, plural, singular string) error {
	if plural == "" {
		plural = "watchdogs"
	}
	if singular == "" {
		singular = "watchdog"
	}
	writer.WriteStart("", "watchdogs", nil)
	for _, o := range structSlice.Slice() {
		XMLWatchdogWriteOne(writer, &o, "watchdog")
	}
	writer.WriteEnd("watchdogs")
	return nil
}

func XMLGlusterBrickWriteOne(writer *XMLWriter, object *GlusterBrick, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "brick"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.BrickDir(); ok {
		writer.WriteCharacter("brick_dir", r)
	}
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Device(); ok {
		writer.WriteCharacter("device", r)
	}
	if r, ok := object.FsName(); ok {
		writer.WriteCharacter("fs_name", r)
	}
	if r, ok := object.GlusterClients(); ok {
		XMLGlusterClientWriteMany(writer, r, "gluster_clients", "gluster_client")
	}
	if r, ok := object.GlusterVolume(); ok {
		XMLGlusterVolumeWriteOne(writer, r, "gluster_volume")
	}
	if r, ok := object.InstanceType(); ok {
		XMLInstanceTypeWriteOne(writer, r, "instance_type")
	}
	if r, ok := object.MemoryPools(); ok {
		XMLGlusterMemoryPoolWriteMany(writer, r, "memory_pools", "memory_pool")
	}
	if r, ok := object.MntOptions(); ok {
		writer.WriteCharacter("mnt_options", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Pid(); ok {
		writer.WriteInt64("pid", r)
	}
	if r, ok := object.Port(); ok {
		writer.WriteInt64("port", r)
	}
	if r, ok := object.ServerId(); ok {
		writer.WriteCharacter("server_id", r)
	}
	if r, ok := object.Statistics(); ok {
		XMLStatisticWriteMany(writer, r, "statistics", "statistic")
	}
	if r, ok := object.Status(); ok {
		XMLGlusterBrickStatusWriteOne(writer, r, "status")
	}
	if r, ok := object.Template(); ok {
		XMLTemplateWriteOne(writer, r, "template")
	}
	if r, ok := object.Vm(); ok {
		XMLVmWriteOne(writer, r, "vm")
	}
	if r, ok := object.Vms(); ok {
		XMLVmWriteMany(writer, r, "vms", "vm")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLGlusterBrickWriteMany(writer *XMLWriter, structSlice *GlusterBrickSlice, plural, singular string) error {
	if plural == "" {
		plural = "bricks"
	}
	if singular == "" {
		singular = "brick"
	}
	writer.WriteStart("", "bricks", nil)
	for _, o := range structSlice.Slice() {
		XMLGlusterBrickWriteOne(writer, &o, "brick")
	}
	writer.WriteEnd("bricks")
	return nil
}

func XMLSkipIfConnectivityBrokenWriteOne(writer *XMLWriter, object *SkipIfConnectivityBroken, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "skip_if_connectivity_broken"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Enabled(); ok {
		writer.WriteBool("enabled", r)
	}
	if r, ok := object.Threshold(); ok {
		writer.WriteInt64("threshold", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLSkipIfConnectivityBrokenWriteMany(writer *XMLWriter, structSlice *SkipIfConnectivityBrokenSlice, plural, singular string) error {
	if plural == "" {
		plural = "skip_if_connectivity_brokens"
	}
	if singular == "" {
		singular = "skip_if_connectivity_broken"
	}
	writer.WriteStart("", "skip_if_connectivity_brokens", nil)
	for _, o := range structSlice.Slice() {
		XMLSkipIfConnectivityBrokenWriteOne(writer, &o, "skip_if_connectivity_broken")
	}
	writer.WriteEnd("skip_if_connectivity_brokens")
	return nil
}

func XMLOpenStackProviderWriteOne(writer *XMLWriter, object *OpenStackProvider, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "open_stack_provider"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.AuthenticationUrl(); ok {
		writer.WriteCharacter("authentication_url", r)
	}
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Password(); ok {
		writer.WriteCharacter("password", r)
	}
	if r, ok := object.Properties(); ok {
		XMLPropertyWriteMany(writer, r, "properties", "property")
	}
	if r, ok := object.RequiresAuthentication(); ok {
		writer.WriteBool("requires_authentication", r)
	}
	if r, ok := object.TenantName(); ok {
		writer.WriteCharacter("tenant_name", r)
	}
	if r, ok := object.Url(); ok {
		writer.WriteCharacter("url", r)
	}
	if r, ok := object.Username(); ok {
		writer.WriteCharacter("username", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLOpenStackProviderWriteMany(writer *XMLWriter, structSlice *OpenStackProviderSlice, plural, singular string) error {
	if plural == "" {
		plural = "open_stack_providers"
	}
	if singular == "" {
		singular = "open_stack_provider"
	}
	writer.WriteStart("", "open_stack_providers", nil)
	for _, o := range structSlice.Slice() {
		XMLOpenStackProviderWriteOne(writer, &o, "open_stack_provider")
	}
	writer.WriteEnd("open_stack_providers")
	return nil
}

func XMLCpuProfileWriteOne(writer *XMLWriter, object *CpuProfile, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "cpu_profile"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Cluster(); ok {
		XMLClusterWriteOne(writer, r, "cluster")
	}
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Permissions(); ok {
		XMLPermissionWriteMany(writer, r, "permissions", "permission")
	}
	if r, ok := object.Qos(); ok {
		XMLQosWriteOne(writer, r, "qos")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLCpuProfileWriteMany(writer *XMLWriter, structSlice *CpuProfileSlice, plural, singular string) error {
	if plural == "" {
		plural = "cpu_profiles"
	}
	if singular == "" {
		singular = "cpu_profile"
	}
	writer.WriteStart("", "cpu_profiles", nil)
	for _, o := range structSlice.Slice() {
		XMLCpuProfileWriteOne(writer, &o, "cpu_profile")
	}
	writer.WriteEnd("cpu_profiles")
	return nil
}

func XMLNumaNodePinWriteOne(writer *XMLWriter, object *NumaNodePin, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "numa_node_pin"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.HostNumaNode(); ok {
		XMLNumaNodeWriteOne(writer, r, "host_numa_node")
	}
	if r, ok := object.Index(); ok {
		writer.WriteInt64("index", r)
	}
	if r, ok := object.Pinned(); ok {
		writer.WriteBool("pinned", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLNumaNodePinWriteMany(writer *XMLWriter, structSlice *NumaNodePinSlice, plural, singular string) error {
	if plural == "" {
		plural = "numa_node_pins"
	}
	if singular == "" {
		singular = "numa_node_pin"
	}
	writer.WriteStart("", "numa_node_pins", nil)
	for _, o := range structSlice.Slice() {
		XMLNumaNodePinWriteOne(writer, &o, "numa_node_pin")
	}
	writer.WriteEnd("numa_node_pins")
	return nil
}

func XMLIscsiBondWriteOne(writer *XMLWriter, object *IscsiBond, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "iscsi_bond"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.DataCenter(); ok {
		XMLDataCenterWriteOne(writer, r, "data_center")
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Networks(); ok {
		XMLNetworkWriteMany(writer, r, "networks", "network")
	}
	if r, ok := object.StorageConnections(); ok {
		XMLStorageConnectionWriteMany(writer, r, "storage_connections", "storage_connection")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLIscsiBondWriteMany(writer *XMLWriter, structSlice *IscsiBondSlice, plural, singular string) error {
	if plural == "" {
		plural = "iscsi_bonds"
	}
	if singular == "" {
		singular = "iscsi_bond"
	}
	writer.WriteStart("", "iscsi_bonds", nil)
	for _, o := range structSlice.Slice() {
		XMLIscsiBondWriteOne(writer, &o, "iscsi_bond")
	}
	writer.WriteEnd("iscsi_bonds")
	return nil
}

func XMLBiosWriteOne(writer *XMLWriter, object *Bios, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "bios"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.BootMenu(); ok {
		XMLBootMenuWriteOne(writer, r, "boot_menu")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLBiosWriteMany(writer *XMLWriter, structSlice *BiosSlice, plural, singular string) error {
	if plural == "" {
		plural = "bioss"
	}
	if singular == "" {
		singular = "bios"
	}
	writer.WriteStart("", "bioss", nil)
	for _, o := range structSlice.Slice() {
		XMLBiosWriteOne(writer, &o, "bios")
	}
	writer.WriteEnd("bioss")
	return nil
}

func XMLCloudInitWriteOne(writer *XMLWriter, object *CloudInit, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "cloud_init"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.AuthorizedKeys(); ok {
		XMLAuthorizedKeyWriteMany(writer, r, "authorized_keys", "authorized_key")
	}
	if r, ok := object.Files(); ok {
		XMLFileWriteMany(writer, r, "files", "file")
	}
	if r, ok := object.Host(); ok {
		XMLHostWriteOne(writer, r, "host")
	}
	if r, ok := object.NetworkConfiguration(); ok {
		XMLNetworkConfigurationWriteOne(writer, r, "network_configuration")
	}
	if r, ok := object.RegenerateSshKeys(); ok {
		writer.WriteBool("regenerate_ssh_keys", r)
	}
	if r, ok := object.Timezone(); ok {
		writer.WriteCharacter("timezone", r)
	}
	if r, ok := object.Users(); ok {
		XMLUserWriteMany(writer, r, "users", "user")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLCloudInitWriteMany(writer *XMLWriter, structSlice *CloudInitSlice, plural, singular string) error {
	if plural == "" {
		plural = "cloud_inits"
	}
	if singular == "" {
		singular = "cloud_init"
	}
	writer.WriteStart("", "cloud_inits", nil)
	for _, o := range structSlice.Slice() {
		XMLCloudInitWriteOne(writer, &o, "cloud_init")
	}
	writer.WriteEnd("cloud_inits")
	return nil
}

func XMLMemoryOverCommitWriteOne(writer *XMLWriter, object *MemoryOverCommit, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "memory_over_commit"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Percent(); ok {
		writer.WriteInt64("percent", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLMemoryOverCommitWriteMany(writer *XMLWriter, structSlice *MemoryOverCommitSlice, plural, singular string) error {
	if plural == "" {
		plural = "memory_over_commits"
	}
	if singular == "" {
		singular = "memory_over_commit"
	}
	writer.WriteStart("", "memory_over_commits", nil)
	for _, o := range structSlice.Slice() {
		XMLMemoryOverCommitWriteOne(writer, &o, "memory_over_commit")
	}
	writer.WriteEnd("memory_over_commits")
	return nil
}

func XMLRoleWriteOne(writer *XMLWriter, object *Role, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "role"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Administrative(); ok {
		writer.WriteBool("administrative", r)
	}
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Mutable(); ok {
		writer.WriteBool("mutable", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Permits(); ok {
		XMLPermitWriteMany(writer, r, "permits", "permit")
	}
	if r, ok := object.User(); ok {
		XMLUserWriteOne(writer, r, "user")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLRoleWriteMany(writer *XMLWriter, structSlice *RoleSlice, plural, singular string) error {
	if plural == "" {
		plural = "roles"
	}
	if singular == "" {
		singular = "role"
	}
	writer.WriteStart("", "roles", nil)
	for _, o := range structSlice.Slice() {
		XMLRoleWriteOne(writer, &o, "role")
	}
	writer.WriteEnd("roles")
	return nil
}

func XMLErrorHandlingWriteOne(writer *XMLWriter, object *ErrorHandling, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "error_handling"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.OnError(); ok {
		XMLMigrateOnErrorWriteOne(writer, r, "on_error")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLErrorHandlingWriteMany(writer *XMLWriter, structSlice *ErrorHandlingSlice, plural, singular string) error {
	if plural == "" {
		plural = "error_handlings"
	}
	if singular == "" {
		singular = "error_handling"
	}
	writer.WriteStart("", "error_handlings", nil)
	for _, o := range structSlice.Slice() {
		XMLErrorHandlingWriteOne(writer, &o, "error_handling")
	}
	writer.WriteEnd("error_handlings")
	return nil
}

func XMLSchedulingPolicyWriteOne(writer *XMLWriter, object *SchedulingPolicy, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "scheduling_policy"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Balances(); ok {
		XMLBalanceWriteMany(writer, r, "balances", "balance")
	}
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.DefaultPolicy(); ok {
		writer.WriteBool("default_policy", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Filters(); ok {
		XMLFilterWriteMany(writer, r, "filters", "filter")
	}
	if r, ok := object.Locked(); ok {
		writer.WriteBool("locked", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Properties(); ok {
		XMLPropertyWriteMany(writer, r, "properties", "property")
	}
	if r, ok := object.Weight(); ok {
		XMLWeightWriteMany(writer, r, "weight", "weight")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLSchedulingPolicyWriteMany(writer *XMLWriter, structSlice *SchedulingPolicySlice, plural, singular string) error {
	if plural == "" {
		plural = "scheduling_policies"
	}
	if singular == "" {
		singular = "scheduling_policy"
	}
	writer.WriteStart("", "scheduling_policies", nil)
	for _, o := range structSlice.Slice() {
		XMLSchedulingPolicyWriteOne(writer, &o, "scheduling_policy")
	}
	writer.WriteEnd("scheduling_policies")
	return nil
}

func XMLGuestOperatingSystemWriteOne(writer *XMLWriter, object *GuestOperatingSystem, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "guest_operating_system"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Architecture(); ok {
		writer.WriteCharacter("architecture", r)
	}
	if r, ok := object.Codename(); ok {
		writer.WriteCharacter("codename", r)
	}
	if r, ok := object.Distribution(); ok {
		writer.WriteCharacter("distribution", r)
	}
	if r, ok := object.Family(); ok {
		writer.WriteCharacter("family", r)
	}
	if r, ok := object.Kernel(); ok {
		XMLKernelWriteOne(writer, r, "kernel")
	}
	if r, ok := object.Version(); ok {
		XMLVersionWriteOne(writer, r, "version")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLGuestOperatingSystemWriteMany(writer *XMLWriter, structSlice *GuestOperatingSystemSlice, plural, singular string) error {
	if plural == "" {
		plural = "guest_operating_systems"
	}
	if singular == "" {
		singular = "guest_operating_system"
	}
	writer.WriteStart("", "guest_operating_systems", nil)
	for _, o := range structSlice.Slice() {
		XMLGuestOperatingSystemWriteOne(writer, &o, "guest_operating_system")
	}
	writer.WriteEnd("guest_operating_systems")
	return nil
}

func XMLWeightWriteOne(writer *XMLWriter, object *Weight, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "weight"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Factor(); ok {
		writer.WriteInt64("factor", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.SchedulingPolicy(); ok {
		XMLSchedulingPolicyWriteOne(writer, r, "scheduling_policy")
	}
	if r, ok := object.SchedulingPolicyUnit(); ok {
		XMLSchedulingPolicyUnitWriteOne(writer, r, "scheduling_policy_unit")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLWeightWriteMany(writer *XMLWriter, structSlice *WeightSlice, plural, singular string) error {
	if plural == "" {
		plural = "weights"
	}
	if singular == "" {
		singular = "weight"
	}
	writer.WriteStart("", "weights", nil)
	for _, o := range structSlice.Slice() {
		XMLWeightWriteOne(writer, &o, "weight")
	}
	writer.WriteEnd("weights")
	return nil
}

func XMLVmPlacementPolicyWriteOne(writer *XMLWriter, object *VmPlacementPolicy, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "vm_placement_policy"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Affinity(); ok {
		XMLVmAffinityWriteOne(writer, r, "affinity")
	}
	if r, ok := object.Hosts(); ok {
		XMLHostWriteMany(writer, r, "hosts", "host")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLVmPlacementPolicyWriteMany(writer *XMLWriter, structSlice *VmPlacementPolicySlice, plural, singular string) error {
	if plural == "" {
		plural = "vm_placement_policies"
	}
	if singular == "" {
		singular = "vm_placement_policy"
	}
	writer.WriteStart("", "vm_placement_policies", nil)
	for _, o := range structSlice.Slice() {
		XMLVmPlacementPolicyWriteOne(writer, &o, "vm_placement_policy")
	}
	writer.WriteEnd("vm_placement_policies")
	return nil
}

func XMLSshWriteOne(writer *XMLWriter, object *Ssh, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "ssh"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.AuthenticationMethod(); ok {
		XMLSshAuthenticationMethodWriteOne(writer, r, "authentication_method")
	}
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Fingerprint(); ok {
		writer.WriteCharacter("fingerprint", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Port(); ok {
		writer.WriteInt64("port", r)
	}
	if r, ok := object.User(); ok {
		XMLUserWriteOne(writer, r, "user")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLSshWriteMany(writer *XMLWriter, structSlice *SshSlice, plural, singular string) error {
	if plural == "" {
		plural = "sshs"
	}
	if singular == "" {
		singular = "ssh"
	}
	writer.WriteStart("", "sshs", nil)
	for _, o := range structSlice.Slice() {
		XMLSshWriteOne(writer, &o, "ssh")
	}
	writer.WriteEnd("sshs")
	return nil
}

func XMLApplicationWriteOne(writer *XMLWriter, object *Application, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "application"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Vm(); ok {
		XMLVmWriteOne(writer, r, "vm")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLApplicationWriteMany(writer *XMLWriter, structSlice *ApplicationSlice, plural, singular string) error {
	if plural == "" {
		plural = "applications"
	}
	if singular == "" {
		singular = "application"
	}
	writer.WriteStart("", "applications", nil)
	for _, o := range structSlice.Slice() {
		XMLApplicationWriteOne(writer, &o, "application")
	}
	writer.WriteEnd("applications")
	return nil
}

func XMLNetworkConfigurationWriteOne(writer *XMLWriter, object *NetworkConfiguration, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "network_configuration"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Dns(); ok {
		XMLDnsWriteOne(writer, r, "dns")
	}
	if r, ok := object.Nics(); ok {
		XMLNicWriteMany(writer, r, "nics", "nic")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLNetworkConfigurationWriteMany(writer *XMLWriter, structSlice *NetworkConfigurationSlice, plural, singular string) error {
	if plural == "" {
		plural = "network_configurations"
	}
	if singular == "" {
		singular = "network_configuration"
	}
	writer.WriteStart("", "network_configurations", nil)
	for _, o := range structSlice.Slice() {
		XMLNetworkConfigurationWriteOne(writer, &o, "network_configuration")
	}
	writer.WriteEnd("network_configurations")
	return nil
}

func XMLSpmWriteOne(writer *XMLWriter, object *Spm, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "spm"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Priority(); ok {
		writer.WriteInt64("priority", r)
	}
	if r, ok := object.Status(); ok {
		XMLSpmStatusWriteOne(writer, r, "status")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLSpmWriteMany(writer *XMLWriter, structSlice *SpmSlice, plural, singular string) error {
	if plural == "" {
		plural = "spms"
	}
	if singular == "" {
		singular = "spm"
	}
	writer.WriteStart("", "spms", nil)
	for _, o := range structSlice.Slice() {
		XMLSpmWriteOne(writer, &o, "spm")
	}
	writer.WriteEnd("spms")
	return nil
}

func XMLNetworkAttachmentWriteOne(writer *XMLWriter, object *NetworkAttachment, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "network_attachment"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.DnsResolverConfiguration(); ok {
		XMLDnsResolverConfigurationWriteOne(writer, r, "dns_resolver_configuration")
	}
	if r, ok := object.Host(); ok {
		XMLHostWriteOne(writer, r, "host")
	}
	if r, ok := object.HostNic(); ok {
		XMLHostNicWriteOne(writer, r, "host_nic")
	}
	if r, ok := object.InSync(); ok {
		writer.WriteBool("in_sync", r)
	}
	if r, ok := object.IpAddressAssignments(); ok {
		XMLIpAddressAssignmentWriteMany(writer, r, "ip_address_assignments", "ip_address_assignment")
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Network(); ok {
		XMLNetworkWriteOne(writer, r, "network")
	}
	if r, ok := object.Properties(); ok {
		XMLPropertyWriteMany(writer, r, "properties", "property")
	}
	if r, ok := object.Qos(); ok {
		XMLQosWriteOne(writer, r, "qos")
	}
	if r, ok := object.ReportedConfigurations(); ok {
		XMLReportedConfigurationWriteMany(writer, r, "reported_configurations", "reported_configuration")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLNetworkAttachmentWriteMany(writer *XMLWriter, structSlice *NetworkAttachmentSlice, plural, singular string) error {
	if plural == "" {
		plural = "network_attachments"
	}
	if singular == "" {
		singular = "network_attachment"
	}
	writer.WriteStart("", "network_attachments", nil)
	for _, o := range structSlice.Slice() {
		XMLNetworkAttachmentWriteOne(writer, &o, "network_attachment")
	}
	writer.WriteEnd("network_attachments")
	return nil
}

func XMLStorageConnectionExtensionWriteOne(writer *XMLWriter, object *StorageConnectionExtension, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "storage_connection_extension"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Host(); ok {
		XMLHostWriteOne(writer, r, "host")
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Password(); ok {
		writer.WriteCharacter("password", r)
	}
	if r, ok := object.Target(); ok {
		writer.WriteCharacter("target", r)
	}
	if r, ok := object.Username(); ok {
		writer.WriteCharacter("username", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLStorageConnectionExtensionWriteMany(writer *XMLWriter, structSlice *StorageConnectionExtensionSlice, plural, singular string) error {
	if plural == "" {
		plural = "storage_connection_extensions"
	}
	if singular == "" {
		singular = "storage_connection_extension"
	}
	writer.WriteStart("", "storage_connection_extensions", nil)
	for _, o := range structSlice.Slice() {
		XMLStorageConnectionExtensionWriteOne(writer, &o, "storage_connection_extension")
	}
	writer.WriteEnd("storage_connection_extensions")
	return nil
}

func XMLPayloadWriteOne(writer *XMLWriter, object *Payload, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "payload"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Files(); ok {
		XMLFileWriteMany(writer, r, "files", "file")
	}
	if r, ok := object.Type(); ok {
		XMLVmDeviceTypeWriteOne(writer, r, "type")
	}
	if r, ok := object.VolumeId(); ok {
		writer.WriteCharacter("volume_id", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLPayloadWriteMany(writer *XMLWriter, structSlice *PayloadSlice, plural, singular string) error {
	if plural == "" {
		plural = "payloads"
	}
	if singular == "" {
		singular = "payload"
	}
	writer.WriteStart("", "payloads", nil)
	for _, o := range structSlice.Slice() {
		XMLPayloadWriteOne(writer, &o, "payload")
	}
	writer.WriteEnd("payloads")
	return nil
}

func XMLRngDeviceWriteOne(writer *XMLWriter, object *RngDevice, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "rng_device"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Rate(); ok {
		XMLRateWriteOne(writer, r, "rate")
	}
	if r, ok := object.Source(); ok {
		XMLRngSourceWriteOne(writer, r, "source")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLRngDeviceWriteMany(writer *XMLWriter, structSlice *RngDeviceSlice, plural, singular string) error {
	if plural == "" {
		plural = "rng_devices"
	}
	if singular == "" {
		singular = "rng_device"
	}
	writer.WriteStart("", "rng_devices", nil)
	for _, o := range structSlice.Slice() {
		XMLRngDeviceWriteOne(writer, &o, "rng_device")
	}
	writer.WriteEnd("rng_devices")
	return nil
}

func XMLGlusterServerHookWriteOne(writer *XMLWriter, object *GlusterServerHook, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "server_hook"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Checksum(); ok {
		writer.WriteCharacter("checksum", r)
	}
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.ContentType(); ok {
		XMLHookContentTypeWriteOne(writer, r, "content_type")
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Host(); ok {
		XMLHostWriteOne(writer, r, "host")
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Status(); ok {
		XMLGlusterHookStatusWriteOne(writer, r, "status")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLGlusterServerHookWriteMany(writer *XMLWriter, structSlice *GlusterServerHookSlice, plural, singular string) error {
	if plural == "" {
		plural = "server_hooks"
	}
	if singular == "" {
		singular = "server_hook"
	}
	writer.WriteStart("", "server_hooks", nil)
	for _, o := range structSlice.Slice() {
		XMLGlusterServerHookWriteOne(writer, &o, "server_hook")
	}
	writer.WriteEnd("server_hooks")
	return nil
}

func XMLTemplateVersionWriteOne(writer *XMLWriter, object *TemplateVersion, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "template_version"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.BaseTemplate(); ok {
		XMLTemplateWriteOne(writer, r, "base_template")
	}
	if r, ok := object.VersionName(); ok {
		writer.WriteCharacter("version_name", r)
	}
	if r, ok := object.VersionNumber(); ok {
		writer.WriteInt64("version_number", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLTemplateVersionWriteMany(writer *XMLWriter, structSlice *TemplateVersionSlice, plural, singular string) error {
	if plural == "" {
		plural = "template_versions"
	}
	if singular == "" {
		singular = "template_version"
	}
	writer.WriteStart("", "template_versions", nil)
	for _, o := range structSlice.Slice() {
		XMLTemplateVersionWriteOne(writer, &o, "template_version")
	}
	writer.WriteEnd("template_versions")
	return nil
}

func XMLGlusterBrickAdvancedDetailsWriteOne(writer *XMLWriter, object *GlusterBrickAdvancedDetails, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "gluster_brick_advanced_details"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Device(); ok {
		writer.WriteCharacter("device", r)
	}
	if r, ok := object.FsName(); ok {
		writer.WriteCharacter("fs_name", r)
	}
	if r, ok := object.GlusterClients(); ok {
		XMLGlusterClientWriteMany(writer, r, "gluster_clients", "gluster_client")
	}
	if r, ok := object.InstanceType(); ok {
		XMLInstanceTypeWriteOne(writer, r, "instance_type")
	}
	if r, ok := object.MemoryPools(); ok {
		XMLGlusterMemoryPoolWriteMany(writer, r, "memory_pools", "memory_pool")
	}
	if r, ok := object.MntOptions(); ok {
		writer.WriteCharacter("mnt_options", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Pid(); ok {
		writer.WriteInt64("pid", r)
	}
	if r, ok := object.Port(); ok {
		writer.WriteInt64("port", r)
	}
	if r, ok := object.Template(); ok {
		XMLTemplateWriteOne(writer, r, "template")
	}
	if r, ok := object.Vm(); ok {
		XMLVmWriteOne(writer, r, "vm")
	}
	if r, ok := object.Vms(); ok {
		XMLVmWriteMany(writer, r, "vms", "vm")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLGlusterBrickAdvancedDetailsWriteMany(writer *XMLWriter, structSlice *GlusterBrickAdvancedDetailsSlice, plural, singular string) error {
	if plural == "" {
		plural = "gluster_brick_advanced_detailss"
	}
	if singular == "" {
		singular = "gluster_brick_advanced_details"
	}
	writer.WriteStart("", "gluster_brick_advanced_detailss", nil)
	for _, o := range structSlice.Slice() {
		XMLGlusterBrickAdvancedDetailsWriteOne(writer, &o, "gluster_brick_advanced_details")
	}
	writer.WriteEnd("gluster_brick_advanced_detailss")
	return nil
}

func XMLVnicProfileWriteOne(writer *XMLWriter, object *VnicProfile, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "vnic_profile"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.CustomProperties(); ok {
		XMLCustomPropertyWriteMany(writer, r, "custom_properties", "custom_property")
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Migratable(); ok {
		writer.WriteBool("migratable", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Network(); ok {
		XMLNetworkWriteOne(writer, r, "network")
	}
	if r, ok := object.NetworkFilter(); ok {
		XMLNetworkFilterWriteOne(writer, r, "network_filter")
	}
	if r, ok := object.PassThrough(); ok {
		XMLVnicPassThroughWriteOne(writer, r, "pass_through")
	}
	if r, ok := object.Permissions(); ok {
		XMLPermissionWriteMany(writer, r, "permissions", "permission")
	}
	if r, ok := object.PortMirroring(); ok {
		writer.WriteBool("port_mirroring", r)
	}
	if r, ok := object.Qos(); ok {
		XMLQosWriteOne(writer, r, "qos")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLVnicProfileWriteMany(writer *XMLWriter, structSlice *VnicProfileSlice, plural, singular string) error {
	if plural == "" {
		plural = "vnic_profiles"
	}
	if singular == "" {
		singular = "vnic_profile"
	}
	writer.WriteStart("", "vnic_profiles", nil)
	for _, o := range structSlice.Slice() {
		XMLVnicProfileWriteOne(writer, &o, "vnic_profile")
	}
	writer.WriteEnd("vnic_profiles")
	return nil
}

func XMLFopStatisticWriteOne(writer *XMLWriter, object *FopStatistic, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "fop_statistic"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Statistics(); ok {
		XMLStatisticWriteMany(writer, r, "statistics", "statistic")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLFopStatisticWriteMany(writer *XMLWriter, structSlice *FopStatisticSlice, plural, singular string) error {
	if plural == "" {
		plural = "fop_statistics"
	}
	if singular == "" {
		singular = "fop_statistic"
	}
	writer.WriteStart("", "fop_statistics", nil)
	for _, o := range structSlice.Slice() {
		XMLFopStatisticWriteOne(writer, &o, "fop_statistic")
	}
	writer.WriteEnd("fop_statistics")
	return nil
}

func XMLCpuTopologyWriteOne(writer *XMLWriter, object *CpuTopology, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "cpu_topology"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Cores(); ok {
		writer.WriteInt64("cores", r)
	}
	if r, ok := object.Sockets(); ok {
		writer.WriteInt64("sockets", r)
	}
	if r, ok := object.Threads(); ok {
		writer.WriteInt64("threads", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLCpuTopologyWriteMany(writer *XMLWriter, structSlice *CpuTopologySlice, plural, singular string) error {
	if plural == "" {
		plural = "cpu_topologies"
	}
	if singular == "" {
		singular = "cpu_topology"
	}
	writer.WriteStart("", "cpu_topologies", nil)
	for _, o := range structSlice.Slice() {
		XMLCpuTopologyWriteOne(writer, &o, "cpu_topology")
	}
	writer.WriteEnd("cpu_topologies")
	return nil
}

func XMLStorageDomainWriteOne(writer *XMLWriter, object *StorageDomain, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "storage_domain"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Available(); ok {
		writer.WriteInt64("available", r)
	}
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Committed(); ok {
		writer.WriteInt64("committed", r)
	}
	if r, ok := object.CriticalSpaceActionBlocker(); ok {
		writer.WriteInt64("critical_space_action_blocker", r)
	}
	if r, ok := object.DataCenter(); ok {
		XMLDataCenterWriteOne(writer, r, "data_center")
	}
	if r, ok := object.DataCenters(); ok {
		XMLDataCenterWriteMany(writer, r, "data_centers", "data_center")
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.DiscardAfterDelete(); ok {
		writer.WriteBool("discard_after_delete", r)
	}
	if r, ok := object.DiskProfiles(); ok {
		XMLDiskProfileWriteMany(writer, r, "disk_profiles", "disk_profile")
	}
	if r, ok := object.DiskSnapshots(); ok {
		XMLDiskSnapshotWriteMany(writer, r, "disk_snapshots", "disk_snapshot")
	}
	if r, ok := object.Disks(); ok {
		XMLDiskWriteMany(writer, r, "disks", "disk")
	}
	if r, ok := object.ExternalStatus(); ok {
		XMLExternalStatusWriteOne(writer, r, "external_status")
	}
	if r, ok := object.Files(); ok {
		XMLFileWriteMany(writer, r, "files", "file")
	}
	if r, ok := object.Host(); ok {
		XMLHostWriteOne(writer, r, "host")
	}
	if r, ok := object.Images(); ok {
		XMLImageWriteMany(writer, r, "images", "image")
	}
	if r, ok := object.Import(); ok {
		writer.WriteBool("import", r)
	}
	if r, ok := object.Master(); ok {
		writer.WriteBool("master", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Permissions(); ok {
		XMLPermissionWriteMany(writer, r, "permissions", "permission")
	}
	if r, ok := object.Status(); ok {
		XMLStorageDomainStatusWriteOne(writer, r, "status")
	}
	if r, ok := object.Storage(); ok {
		XMLHostStorageWriteOne(writer, r, "storage")
	}
	if r, ok := object.StorageConnections(); ok {
		XMLStorageConnectionWriteMany(writer, r, "storage_connections", "storage_connection")
	}
	if r, ok := object.StorageFormat(); ok {
		XMLStorageFormatWriteOne(writer, r, "storage_format")
	}
	if r, ok := object.SupportsDiscard(); ok {
		writer.WriteBool("supports_discard", r)
	}
	if r, ok := object.SupportsDiscardZeroesData(); ok {
		writer.WriteBool("supports_discard_zeroes_data", r)
	}
	if r, ok := object.Templates(); ok {
		XMLTemplateWriteMany(writer, r, "templates", "template")
	}
	if r, ok := object.Type(); ok {
		XMLStorageDomainTypeWriteOne(writer, r, "type")
	}
	if r, ok := object.Used(); ok {
		writer.WriteInt64("used", r)
	}
	if r, ok := object.Vms(); ok {
		XMLVmWriteMany(writer, r, "vms", "vm")
	}
	if r, ok := object.WarningLowSpaceIndicator(); ok {
		writer.WriteInt64("warning_low_space_indicator", r)
	}
	if r, ok := object.WipeAfterDelete(); ok {
		writer.WriteBool("wipe_after_delete", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLStorageDomainWriteMany(writer *XMLWriter, structSlice *StorageDomainSlice, plural, singular string) error {
	if plural == "" {
		plural = "storage_domains"
	}
	if singular == "" {
		singular = "storage_domain"
	}
	writer.WriteStart("", "storage_domains", nil)
	for _, o := range structSlice.Slice() {
		XMLStorageDomainWriteOne(writer, &o, "storage_domain")
	}
	writer.WriteEnd("storage_domains")
	return nil
}

func XMLCustomPropertyWriteOne(writer *XMLWriter, object *CustomProperty, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "custom_property"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Regexp(); ok {
		writer.WriteCharacter("regexp", r)
	}
	if r, ok := object.Value(); ok {
		writer.WriteCharacter("value", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLCustomPropertyWriteMany(writer *XMLWriter, structSlice *CustomPropertySlice, plural, singular string) error {
	if plural == "" {
		plural = "custom_properties"
	}
	if singular == "" {
		singular = "custom_property"
	}
	writer.WriteStart("", "custom_properties", nil)
	for _, o := range structSlice.Slice() {
		XMLCustomPropertyWriteOne(writer, &o, "custom_property")
	}
	writer.WriteEnd("custom_properties")
	return nil
}

func XMLMemoryPolicyWriteOne(writer *XMLWriter, object *MemoryPolicy, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "memory_policy"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Ballooning(); ok {
		writer.WriteBool("ballooning", r)
	}
	if r, ok := object.Guaranteed(); ok {
		writer.WriteInt64("guaranteed", r)
	}
	if r, ok := object.Max(); ok {
		writer.WriteInt64("max", r)
	}
	if r, ok := object.OverCommit(); ok {
		XMLMemoryOverCommitWriteOne(writer, r, "over_commit")
	}
	if r, ok := object.TransparentHugePages(); ok {
		XMLTransparentHugePagesWriteOne(writer, r, "transparent_hugepages")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLMemoryPolicyWriteMany(writer *XMLWriter, structSlice *MemoryPolicySlice, plural, singular string) error {
	if plural == "" {
		plural = "memory_policies"
	}
	if singular == "" {
		singular = "memory_policy"
	}
	writer.WriteStart("", "memory_policies", nil)
	for _, o := range structSlice.Slice() {
		XMLMemoryPolicyWriteOne(writer, &o, "memory_policy")
	}
	writer.WriteEnd("memory_policies")
	return nil
}

func XMLNetworkWriteOne(writer *XMLWriter, object *Network, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "network"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Cluster(); ok {
		XMLClusterWriteOne(writer, r, "cluster")
	}
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.DataCenter(); ok {
		XMLDataCenterWriteOne(writer, r, "data_center")
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Display(); ok {
		writer.WriteBool("display", r)
	}
	if r, ok := object.DnsResolverConfiguration(); ok {
		XMLDnsResolverConfigurationWriteOne(writer, r, "dns_resolver_configuration")
	}
	if r, ok := object.Ip(); ok {
		XMLIpWriteOne(writer, r, "ip")
	}
	if r, ok := object.Mtu(); ok {
		writer.WriteInt64("mtu", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.NetworkLabels(); ok {
		XMLNetworkLabelWriteMany(writer, r, "network_labels", "network_label")
	}
	if r, ok := object.Permissions(); ok {
		XMLPermissionWriteMany(writer, r, "permissions", "permission")
	}
	if r, ok := object.ProfileRequired(); ok {
		writer.WriteBool("profile_required", r)
	}
	if r, ok := object.Qos(); ok {
		XMLQosWriteOne(writer, r, "qos")
	}
	if r, ok := object.Required(); ok {
		writer.WriteBool("required", r)
	}
	if r, ok := object.Status(); ok {
		XMLNetworkStatusWriteOne(writer, r, "status")
	}
	if r, ok := object.Stp(); ok {
		writer.WriteBool("stp", r)
	}
	if r, ok := object.Usages(); ok {
		XMLNetworkUsageWriteMany(writer, r, "usages", "network_usage")
	}
	if r, ok := object.Vlan(); ok {
		XMLVlanWriteOne(writer, r, "vlan")
	}
	if r, ok := object.VnicProfiles(); ok {
		XMLVnicProfileWriteMany(writer, r, "vnic_profiles", "vnic_profile")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLNetworkWriteMany(writer *XMLWriter, structSlice *NetworkSlice, plural, singular string) error {
	if plural == "" {
		plural = "networks"
	}
	if singular == "" {
		singular = "network"
	}
	writer.WriteStart("", "networks", nil)
	for _, o := range structSlice.Slice() {
		XMLNetworkWriteOne(writer, &o, "network")
	}
	writer.WriteEnd("networks")
	return nil
}

func XMLAgentConfigurationWriteOne(writer *XMLWriter, object *AgentConfiguration, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "agent_configuration"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Address(); ok {
		writer.WriteCharacter("address", r)
	}
	if r, ok := object.BrokerType(); ok {
		XMLMessageBrokerTypeWriteOne(writer, r, "broker_type")
	}
	if r, ok := object.NetworkMappings(); ok {
		writer.WriteCharacter("network_mappings", r)
	}
	if r, ok := object.Password(); ok {
		writer.WriteCharacter("password", r)
	}
	if r, ok := object.Port(); ok {
		writer.WriteInt64("port", r)
	}
	if r, ok := object.Username(); ok {
		writer.WriteCharacter("username", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLAgentConfigurationWriteMany(writer *XMLWriter, structSlice *AgentConfigurationSlice, plural, singular string) error {
	if plural == "" {
		plural = "agent_configurations"
	}
	if singular == "" {
		singular = "agent_configuration"
	}
	writer.WriteStart("", "agent_configurations", nil)
	for _, o := range structSlice.Slice() {
		XMLAgentConfigurationWriteOne(writer, &o, "agent_configuration")
	}
	writer.WriteEnd("agent_configurations")
	return nil
}

func XMLAffinityGroupWriteOne(writer *XMLWriter, object *AffinityGroup, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "affinity_group"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Cluster(); ok {
		XMLClusterWriteOne(writer, r, "cluster")
	}
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Enforcing(); ok {
		writer.WriteBool("enforcing", r)
	}
	if r, ok := object.Hosts(); ok {
		XMLHostWriteMany(writer, r, "hosts", "host")
	}
	if r, ok := object.HostsRule(); ok {
		XMLAffinityRuleWriteOne(writer, r, "hosts_rule")
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Positive(); ok {
		writer.WriteBool("positive", r)
	}
	if r, ok := object.Vms(); ok {
		XMLVmWriteMany(writer, r, "vms", "vm")
	}
	if r, ok := object.VmsRule(); ok {
		XMLAffinityRuleWriteOne(writer, r, "vms_rule")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLAffinityGroupWriteMany(writer *XMLWriter, structSlice *AffinityGroupSlice, plural, singular string) error {
	if plural == "" {
		plural = "affinity_groups"
	}
	if singular == "" {
		singular = "affinity_group"
	}
	writer.WriteStart("", "affinity_groups", nil)
	for _, o := range structSlice.Slice() {
		XMLAffinityGroupWriteOne(writer, &o, "affinity_group")
	}
	writer.WriteEnd("affinity_groups")
	return nil
}

func XMLHostedEngineWriteOne(writer *XMLWriter, object *HostedEngine, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "hosted_engine"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Active(); ok {
		writer.WriteBool("active", r)
	}
	if r, ok := object.Configured(); ok {
		writer.WriteBool("configured", r)
	}
	if r, ok := object.GlobalMaintenance(); ok {
		writer.WriteBool("global_maintenance", r)
	}
	if r, ok := object.LocalMaintenance(); ok {
		writer.WriteBool("local_maintenance", r)
	}
	if r, ok := object.Score(); ok {
		writer.WriteInt64("score", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLHostedEngineWriteMany(writer *XMLWriter, structSlice *HostedEngineSlice, plural, singular string) error {
	if plural == "" {
		plural = "hosted_engines"
	}
	if singular == "" {
		singular = "hosted_engine"
	}
	writer.WriteStart("", "hosted_engines", nil)
	for _, o := range structSlice.Slice() {
		XMLHostedEngineWriteOne(writer, &o, "hosted_engine")
	}
	writer.WriteEnd("hosted_engines")
	return nil
}

func XMLKernelWriteOne(writer *XMLWriter, object *Kernel, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "kernel"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Version(); ok {
		XMLVersionWriteOne(writer, r, "version")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLKernelWriteMany(writer *XMLWriter, structSlice *KernelSlice, plural, singular string) error {
	if plural == "" {
		plural = "kernels"
	}
	if singular == "" {
		singular = "kernel"
	}
	writer.WriteStart("", "kernels", nil)
	for _, o := range structSlice.Slice() {
		XMLKernelWriteOne(writer, &o, "kernel")
	}
	writer.WriteEnd("kernels")
	return nil
}

func XMLTicketWriteOne(writer *XMLWriter, object *Ticket, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "ticket"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Expiry(); ok {
		writer.WriteInt64("expiry", r)
	}
	if r, ok := object.Value(); ok {
		writer.WriteCharacter("value", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLTicketWriteMany(writer *XMLWriter, structSlice *TicketSlice, plural, singular string) error {
	if plural == "" {
		plural = "tickets"
	}
	if singular == "" {
		singular = "ticket"
	}
	writer.WriteStart("", "tickets", nil)
	for _, o := range structSlice.Slice() {
		XMLTicketWriteOne(writer, &o, "ticket")
	}
	writer.WriteEnd("tickets")
	return nil
}

func XMLImageWriteOne(writer *XMLWriter, object *Image, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "image"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.StorageDomain(); ok {
		XMLStorageDomainWriteOne(writer, r, "storage_domain")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLImageWriteMany(writer *XMLWriter, structSlice *ImageSlice, plural, singular string) error {
	if plural == "" {
		plural = "images"
	}
	if singular == "" {
		singular = "image"
	}
	writer.WriteStart("", "images", nil)
	for _, o := range structSlice.Slice() {
		XMLImageWriteOne(writer, &o, "image")
	}
	writer.WriteEnd("images")
	return nil
}

func XMLBlockStatisticWriteOne(writer *XMLWriter, object *BlockStatistic, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "block_statistic"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Statistics(); ok {
		XMLStatisticWriteMany(writer, r, "statistics", "statistic")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLBlockStatisticWriteMany(writer *XMLWriter, structSlice *BlockStatisticSlice, plural, singular string) error {
	if plural == "" {
		plural = "block_statistics"
	}
	if singular == "" {
		singular = "block_statistic"
	}
	writer.WriteStart("", "block_statistics", nil)
	for _, o := range structSlice.Slice() {
		XMLBlockStatisticWriteOne(writer, &o, "block_statistic")
	}
	writer.WriteEnd("block_statistics")
	return nil
}

func XMLMacPoolWriteOne(writer *XMLWriter, object *MacPool, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "mac_pool"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.AllowDuplicates(); ok {
		writer.WriteBool("allow_duplicates", r)
	}
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.DefaultPool(); ok {
		writer.WriteBool("default_pool", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Ranges(); ok {
		XMLRangeWriteMany(writer, r, "ranges", "range")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLMacPoolWriteMany(writer *XMLWriter, structSlice *MacPoolSlice, plural, singular string) error {
	if plural == "" {
		plural = "mac_pools"
	}
	if singular == "" {
		singular = "mac_pool"
	}
	writer.WriteStart("", "mac_pools", nil)
	for _, o := range structSlice.Slice() {
		XMLMacPoolWriteOne(writer, &o, "mac_pool")
	}
	writer.WriteEnd("mac_pools")
	return nil
}

func XMLGlusterMemoryPoolWriteOne(writer *XMLWriter, object *GlusterMemoryPool, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "memory_pool"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.AllocCount(); ok {
		writer.WriteInt64("alloc_count", r)
	}
	if r, ok := object.ColdCount(); ok {
		writer.WriteInt64("cold_count", r)
	}
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.HotCount(); ok {
		writer.WriteInt64("hot_count", r)
	}
	if r, ok := object.MaxAlloc(); ok {
		writer.WriteInt64("max_alloc", r)
	}
	if r, ok := object.MaxStdalloc(); ok {
		writer.WriteInt64("max_stdalloc", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.PaddedSize(); ok {
		writer.WriteInt64("padded_size", r)
	}
	if r, ok := object.PoolMisses(); ok {
		writer.WriteInt64("pool_misses", r)
	}
	if r, ok := object.Type(); ok {
		writer.WriteCharacter("type", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLGlusterMemoryPoolWriteMany(writer *XMLWriter, structSlice *GlusterMemoryPoolSlice, plural, singular string) error {
	if plural == "" {
		plural = "memory_pools"
	}
	if singular == "" {
		singular = "memory_pool"
	}
	writer.WriteStart("", "memory_pools", nil)
	for _, o := range structSlice.Slice() {
		XMLGlusterMemoryPoolWriteOne(writer, &o, "memory_pool")
	}
	writer.WriteEnd("memory_pools")
	return nil
}

func XMLPropertyWriteOne(writer *XMLWriter, object *Property, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "property"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Value(); ok {
		writer.WriteCharacter("value", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLPropertyWriteMany(writer *XMLWriter, structSlice *PropertySlice, plural, singular string) error {
	if plural == "" {
		plural = "properties"
	}
	if singular == "" {
		singular = "property"
	}
	writer.WriteStart("", "properties", nil)
	for _, o := range structSlice.Slice() {
		XMLPropertyWriteOne(writer, &o, "property")
	}
	writer.WriteEnd("properties")
	return nil
}

func XMLStepWriteOne(writer *XMLWriter, object *Step, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "step"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.EndTime(); ok {
		writer.WriteDate("end_time", r)
	}
	if r, ok := object.ExecutionHost(); ok {
		XMLHostWriteOne(writer, r, "execution_host")
	}
	if r, ok := object.External(); ok {
		writer.WriteBool("external", r)
	}
	if r, ok := object.ExternalType(); ok {
		XMLExternalSystemTypeWriteOne(writer, r, "external_type")
	}
	if r, ok := object.Job(); ok {
		XMLJobWriteOne(writer, r, "job")
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Number(); ok {
		writer.WriteInt64("number", r)
	}
	if r, ok := object.ParentStep(); ok {
		XMLStepWriteOne(writer, r, "parent_step")
	}
	if r, ok := object.Progress(); ok {
		writer.WriteInt64("progress", r)
	}
	if r, ok := object.StartTime(); ok {
		writer.WriteDate("start_time", r)
	}
	if r, ok := object.Statistics(); ok {
		XMLStatisticWriteMany(writer, r, "statistics", "statistic")
	}
	if r, ok := object.Status(); ok {
		XMLStepStatusWriteOne(writer, r, "status")
	}
	if r, ok := object.Type(); ok {
		XMLStepEnumWriteOne(writer, r, "type")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLStepWriteMany(writer *XMLWriter, structSlice *StepSlice, plural, singular string) error {
	if plural == "" {
		plural = "steps"
	}
	if singular == "" {
		singular = "step"
	}
	writer.WriteStart("", "steps", nil)
	for _, o := range structSlice.Slice() {
		XMLStepWriteOne(writer, &o, "step")
	}
	writer.WriteEnd("steps")
	return nil
}

func XMLOperatingSystemWriteOne(writer *XMLWriter, object *OperatingSystem, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "os"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Boot(); ok {
		XMLBootWriteOne(writer, r, "boot")
	}
	if r, ok := object.Cmdline(); ok {
		writer.WriteCharacter("cmdline", r)
	}
	if r, ok := object.CustomKernelCmdline(); ok {
		writer.WriteCharacter("custom_kernel_cmdline", r)
	}
	if r, ok := object.Initrd(); ok {
		writer.WriteCharacter("initrd", r)
	}
	if r, ok := object.Kernel(); ok {
		writer.WriteCharacter("kernel", r)
	}
	if r, ok := object.ReportedKernelCmdline(); ok {
		writer.WriteCharacter("reported_kernel_cmdline", r)
	}
	if r, ok := object.Type(); ok {
		writer.WriteCharacter("type", r)
	}
	if r, ok := object.Version(); ok {
		XMLVersionWriteOne(writer, r, "version")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLOperatingSystemWriteMany(writer *XMLWriter, structSlice *OperatingSystemSlice, plural, singular string) error {
	if plural == "" {
		plural = "oss"
	}
	if singular == "" {
		singular = "os"
	}
	writer.WriteStart("", "oss", nil)
	for _, o := range structSlice.Slice() {
		XMLOperatingSystemWriteOne(writer, &o, "os")
	}
	writer.WriteEnd("oss")
	return nil
}

func XMLHostDevicePassthroughWriteOne(writer *XMLWriter, object *HostDevicePassthrough, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "host_device_passthrough"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Enabled(); ok {
		writer.WriteBool("enabled", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLHostDevicePassthroughWriteMany(writer *XMLWriter, structSlice *HostDevicePassthroughSlice, plural, singular string) error {
	if plural == "" {
		plural = "host_device_passthroughs"
	}
	if singular == "" {
		singular = "host_device_passthrough"
	}
	writer.WriteStart("", "host_device_passthroughs", nil)
	for _, o := range structSlice.Slice() {
		XMLHostDevicePassthroughWriteOne(writer, &o, "host_device_passthrough")
	}
	writer.WriteEnd("host_device_passthroughs")
	return nil
}

func XMLKsmWriteOne(writer *XMLWriter, object *Ksm, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "ksm"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Enabled(); ok {
		writer.WriteBool("enabled", r)
	}
	if r, ok := object.MergeAcrossNodes(); ok {
		writer.WriteBool("merge_across_nodes", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLKsmWriteMany(writer *XMLWriter, structSlice *KsmSlice, plural, singular string) error {
	if plural == "" {
		plural = "ksms"
	}
	if singular == "" {
		singular = "ksm"
	}
	writer.WriteStart("", "ksms", nil)
	for _, o := range structSlice.Slice() {
		XMLKsmWriteOne(writer, &o, "ksm")
	}
	writer.WriteEnd("ksms")
	return nil
}

func XMLIconWriteOne(writer *XMLWriter, object *Icon, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "icon"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Data(); ok {
		writer.WriteCharacter("data", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.MediaType(); ok {
		writer.WriteCharacter("media_type", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLIconWriteMany(writer *XMLWriter, structSlice *IconSlice, plural, singular string) error {
	if plural == "" {
		plural = "icons"
	}
	if singular == "" {
		singular = "icon"
	}
	writer.WriteStart("", "icons", nil)
	for _, o := range structSlice.Slice() {
		XMLIconWriteOne(writer, &o, "icon")
	}
	writer.WriteEnd("icons")
	return nil
}

func XMLConfigurationWriteOne(writer *XMLWriter, object *Configuration, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "configuration"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Data(); ok {
		writer.WriteCharacter("data", r)
	}
	if r, ok := object.Type(); ok {
		XMLConfigurationTypeWriteOne(writer, r, "type")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLConfigurationWriteMany(writer *XMLWriter, structSlice *ConfigurationSlice, plural, singular string) error {
	if plural == "" {
		plural = "configurations"
	}
	if singular == "" {
		singular = "configuration"
	}
	writer.WriteStart("", "configurations", nil)
	for _, o := range structSlice.Slice() {
		XMLConfigurationWriteOne(writer, &o, "configuration")
	}
	writer.WriteEnd("configurations")
	return nil
}

func XMLGlusterVolumeWriteOne(writer *XMLWriter, object *GlusterVolume, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "gluster_volume"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Bricks(); ok {
		XMLGlusterBrickWriteMany(writer, r, "bricks", "brick")
	}
	if r, ok := object.Cluster(); ok {
		XMLClusterWriteOne(writer, r, "cluster")
	}
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.DisperseCount(); ok {
		writer.WriteInt64("disperse_count", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Options(); ok {
		XMLOptionWriteMany(writer, r, "options", "option")
	}
	if r, ok := object.RedundancyCount(); ok {
		writer.WriteInt64("redundancy_count", r)
	}
	if r, ok := object.ReplicaCount(); ok {
		writer.WriteInt64("replica_count", r)
	}
	if r, ok := object.Statistics(); ok {
		XMLStatisticWriteMany(writer, r, "statistics", "statistic")
	}
	if r, ok := object.Status(); ok {
		XMLGlusterVolumeStatusWriteOne(writer, r, "status")
	}
	if r, ok := object.StripeCount(); ok {
		writer.WriteInt64("stripe_count", r)
	}
	if r, ok := object.TransportTypes(); ok {
		XMLTransportTypeWriteMany(writer, r, "transport_types", "transport_type")
	}
	if r, ok := object.VolumeType(); ok {
		XMLGlusterVolumeTypeWriteOne(writer, r, "volume_type")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLGlusterVolumeWriteMany(writer *XMLWriter, structSlice *GlusterVolumeSlice, plural, singular string) error {
	if plural == "" {
		plural = "gluster_volumes"
	}
	if singular == "" {
		singular = "gluster_volume"
	}
	writer.WriteStart("", "gluster_volumes", nil)
	for _, o := range structSlice.Slice() {
		XMLGlusterVolumeWriteOne(writer, &o, "gluster_volume")
	}
	writer.WriteEnd("gluster_volumes")
	return nil
}

func XMLSpecialObjectsWriteOne(writer *XMLWriter, object *SpecialObjects, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "special_objects"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.BlankTemplate(); ok {
		XMLTemplateWriteOne(writer, r, "blank_template")
	}
	if r, ok := object.RootTag(); ok {
		XMLTagWriteOne(writer, r, "root_tag")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLSpecialObjectsWriteMany(writer *XMLWriter, structSlice *SpecialObjectsSlice, plural, singular string) error {
	if plural == "" {
		plural = "special_objectss"
	}
	if singular == "" {
		singular = "special_objects"
	}
	writer.WriteStart("", "special_objectss", nil)
	for _, o := range structSlice.Slice() {
		XMLSpecialObjectsWriteOne(writer, &o, "special_objects")
	}
	writer.WriteEnd("special_objectss")
	return nil
}

func XMLBootWriteOne(writer *XMLWriter, object *Boot, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "boot"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Devices(); ok {
		XMLBootDeviceWriteMany(writer, r, "devices", "boot_device")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLBootWriteMany(writer *XMLWriter, structSlice *BootSlice, plural, singular string) error {
	if plural == "" {
		plural = "boots"
	}
	if singular == "" {
		singular = "boot"
	}
	writer.WriteStart("", "boots", nil)
	for _, o := range structSlice.Slice() {
		XMLBootWriteOne(writer, &o, "boot")
	}
	writer.WriteEnd("boots")
	return nil
}

func XMLHostWriteOne(writer *XMLWriter, object *Host, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "host"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Address(); ok {
		writer.WriteCharacter("address", r)
	}
	if r, ok := object.AffinityLabels(); ok {
		XMLAffinityLabelWriteMany(writer, r, "affinity_labels", "affinity_label")
	}
	if r, ok := object.Agents(); ok {
		XMLAgentWriteMany(writer, r, "agents", "agent")
	}
	if r, ok := object.AutoNumaStatus(); ok {
		XMLAutoNumaStatusWriteOne(writer, r, "auto_numa_status")
	}
	if r, ok := object.Certificate(); ok {
		XMLCertificateWriteOne(writer, r, "certificate")
	}
	if r, ok := object.Cluster(); ok {
		XMLClusterWriteOne(writer, r, "cluster")
	}
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Cpu(); ok {
		XMLCpuWriteOne(writer, r, "cpu")
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.DevicePassthrough(); ok {
		XMLHostDevicePassthroughWriteOne(writer, r, "device_passthrough")
	}
	if r, ok := object.Devices(); ok {
		XMLDeviceWriteMany(writer, r, "devices", "device")
	}
	if r, ok := object.Display(); ok {
		XMLDisplayWriteOne(writer, r, "display")
	}
	if r, ok := object.ExternalHostProvider(); ok {
		XMLExternalHostProviderWriteOne(writer, r, "external_host_provider")
	}
	if r, ok := object.ExternalStatus(); ok {
		XMLExternalStatusWriteOne(writer, r, "external_status")
	}
	if r, ok := object.HardwareInformation(); ok {
		XMLHardwareInformationWriteOne(writer, r, "hardware_information")
	}
	if r, ok := object.Hooks(); ok {
		XMLHookWriteMany(writer, r, "hooks", "hook")
	}
	if r, ok := object.HostedEngine(); ok {
		XMLHostedEngineWriteOne(writer, r, "hosted_engine")
	}
	if r, ok := object.Iscsi(); ok {
		XMLIscsiDetailsWriteOne(writer, r, "iscsi")
	}
	if r, ok := object.KatelloErrata(); ok {
		XMLKatelloErratumWriteMany(writer, r, "katello_errata", "katello_erratum")
	}
	if r, ok := object.KdumpStatus(); ok {
		XMLKdumpStatusWriteOne(writer, r, "kdump_status")
	}
	if r, ok := object.Ksm(); ok {
		XMLKsmWriteOne(writer, r, "ksm")
	}
	if r, ok := object.LibvirtVersion(); ok {
		XMLVersionWriteOne(writer, r, "libvirt_version")
	}
	if r, ok := object.MaxSchedulingMemory(); ok {
		writer.WriteInt64("max_scheduling_memory", r)
	}
	if r, ok := object.Memory(); ok {
		writer.WriteInt64("memory", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.NetworkAttachments(); ok {
		XMLNetworkAttachmentWriteMany(writer, r, "network_attachments", "network_attachment")
	}
	if r, ok := object.Nics(); ok {
		XMLNicWriteMany(writer, r, "nics", "nic")
	}
	if r, ok := object.NumaNodes(); ok {
		XMLNumaNodeWriteMany(writer, r, "host_numa_nodes", "host_numa_node")
	}
	if r, ok := object.NumaSupported(); ok {
		writer.WriteBool("numa_supported", r)
	}
	if r, ok := object.Os(); ok {
		XMLOperatingSystemWriteOne(writer, r, "os")
	}
	if r, ok := object.OverrideIptables(); ok {
		writer.WriteBool("override_iptables", r)
	}
	if r, ok := object.Permissions(); ok {
		XMLPermissionWriteMany(writer, r, "permissions", "permission")
	}
	if r, ok := object.Port(); ok {
		writer.WriteInt64("port", r)
	}
	if r, ok := object.PowerManagement(); ok {
		XMLPowerManagementWriteOne(writer, r, "power_management")
	}
	if r, ok := object.Protocol(); ok {
		XMLHostProtocolWriteOne(writer, r, "protocol")
	}
	if r, ok := object.RootPassword(); ok {
		writer.WriteCharacter("root_password", r)
	}
	if r, ok := object.SeLinux(); ok {
		XMLSeLinuxWriteOne(writer, r, "se_linux")
	}
	if r, ok := object.Spm(); ok {
		XMLSpmWriteOne(writer, r, "spm")
	}
	if r, ok := object.Ssh(); ok {
		XMLSshWriteOne(writer, r, "ssh")
	}
	if r, ok := object.Statistics(); ok {
		XMLStatisticWriteMany(writer, r, "statistics", "statistic")
	}
	if r, ok := object.Status(); ok {
		XMLHostStatusWriteOne(writer, r, "status")
	}
	if r, ok := object.StatusDetail(); ok {
		writer.WriteCharacter("status_detail", r)
	}
	if r, ok := object.StorageConnectionExtensions(); ok {
		XMLStorageConnectionExtensionWriteMany(writer, r, "storage_connection_extensions", "storage_connection_extension")
	}
	if r, ok := object.Storages(); ok {
		XMLHostStorageWriteMany(writer, r, "storages", "host_storage")
	}
	if r, ok := object.Summary(); ok {
		XMLVmSummaryWriteOne(writer, r, "summary")
	}
	if r, ok := object.Tags(); ok {
		XMLTagWriteMany(writer, r, "tags", "tag")
	}
	if r, ok := object.TransparentHugePages(); ok {
		XMLTransparentHugePagesWriteOne(writer, r, "transparent_hugepages")
	}
	if r, ok := object.Type(); ok {
		XMLHostTypeWriteOne(writer, r, "type")
	}
	if r, ok := object.UnmanagedNetworks(); ok {
		XMLUnmanagedNetworkWriteMany(writer, r, "unmanaged_networks", "unmanaged_network")
	}
	if r, ok := object.UpdateAvailable(); ok {
		writer.WriteBool("update_available", r)
	}
	if r, ok := object.Version(); ok {
		XMLVersionWriteOne(writer, r, "version")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLHostWriteMany(writer *XMLWriter, structSlice *HostSlice, plural, singular string) error {
	if plural == "" {
		plural = "hosts"
	}
	if singular == "" {
		singular = "host"
	}
	writer.WriteStart("", "hosts", nil)
	for _, o := range structSlice.Slice() {
		XMLHostWriteOne(writer, &o, "host")
	}
	writer.WriteEnd("hosts")
	return nil
}

func XMLCpuWriteOne(writer *XMLWriter, object *Cpu, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "cpu"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Architecture(); ok {
		XMLArchitectureWriteOne(writer, r, "architecture")
	}
	if r, ok := object.Cores(); ok {
		XMLCoreWriteMany(writer, r, "cores", "core")
	}
	if r, ok := object.CpuTune(); ok {
		XMLCpuTuneWriteOne(writer, r, "cpu_tune")
	}
	if r, ok := object.Level(); ok {
		writer.WriteInt64("level", r)
	}
	if r, ok := object.Mode(); ok {
		XMLCpuModeWriteOne(writer, r, "mode")
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Speed(); ok {
		writer.WriteFloat64("speed", r)
	}
	if r, ok := object.Topology(); ok {
		XMLCpuTopologyWriteOne(writer, r, "topology")
	}
	if r, ok := object.Type(); ok {
		writer.WriteCharacter("type", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLCpuWriteMany(writer *XMLWriter, structSlice *CpuSlice, plural, singular string) error {
	if plural == "" {
		plural = "cpus"
	}
	if singular == "" {
		singular = "cpu"
	}
	writer.WriteStart("", "cpus", nil)
	for _, o := range structSlice.Slice() {
		XMLCpuWriteOne(writer, &o, "cpu")
	}
	writer.WriteEnd("cpus")
	return nil
}

func XMLHighAvailabilityWriteOne(writer *XMLWriter, object *HighAvailability, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "high_availability"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Enabled(); ok {
		writer.WriteBool("enabled", r)
	}
	if r, ok := object.Priority(); ok {
		writer.WriteInt64("priority", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLHighAvailabilityWriteMany(writer *XMLWriter, structSlice *HighAvailabilitySlice, plural, singular string) error {
	if plural == "" {
		plural = "high_availabilities"
	}
	if singular == "" {
		singular = "high_availability"
	}
	writer.WriteStart("", "high_availabilities", nil)
	for _, o := range structSlice.Slice() {
		XMLHighAvailabilityWriteOne(writer, &o, "high_availability")
	}
	writer.WriteEnd("high_availabilities")
	return nil
}

func XMLSsoWriteOne(writer *XMLWriter, object *Sso, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "sso"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Methods(); ok {
		XMLMethodWriteMany(writer, r, "methods", "method")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLSsoWriteMany(writer *XMLWriter, structSlice *SsoSlice, plural, singular string) error {
	if plural == "" {
		plural = "ssos"
	}
	if singular == "" {
		singular = "sso"
	}
	writer.WriteStart("", "ssos", nil)
	for _, o := range structSlice.Slice() {
		XMLSsoWriteOne(writer, &o, "sso")
	}
	writer.WriteEnd("ssos")
	return nil
}

func XMLStatisticWriteOne(writer *XMLWriter, object *Statistic, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "statistic"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Brick(); ok {
		XMLGlusterBrickWriteOne(writer, r, "brick")
	}
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Disk(); ok {
		XMLDiskWriteOne(writer, r, "disk")
	}
	if r, ok := object.GlusterVolume(); ok {
		XMLGlusterVolumeWriteOne(writer, r, "gluster_volume")
	}
	if r, ok := object.Host(); ok {
		XMLHostWriteOne(writer, r, "host")
	}
	if r, ok := object.HostNic(); ok {
		XMLHostNicWriteOne(writer, r, "host_nic")
	}
	if r, ok := object.HostNumaNode(); ok {
		XMLNumaNodeWriteOne(writer, r, "host_numa_node")
	}
	if r, ok := object.Kind(); ok {
		XMLStatisticKindWriteOne(writer, r, "kind")
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Nic(); ok {
		XMLNicWriteOne(writer, r, "nic")
	}
	if r, ok := object.Step(); ok {
		XMLStepWriteOne(writer, r, "step")
	}
	if r, ok := object.Type(); ok {
		XMLValueTypeWriteOne(writer, r, "type")
	}
	if r, ok := object.Unit(); ok {
		XMLStatisticUnitWriteOne(writer, r, "unit")
	}
	if r, ok := object.Values(); ok {
		XMLValueWriteMany(writer, r, "values", "value")
	}
	if r, ok := object.Vm(); ok {
		XMLVmWriteOne(writer, r, "vm")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLStatisticWriteMany(writer *XMLWriter, structSlice *StatisticSlice, plural, singular string) error {
	if plural == "" {
		plural = "statistics"
	}
	if singular == "" {
		singular = "statistic"
	}
	writer.WriteStart("", "statistics", nil)
	for _, o := range structSlice.Slice() {
		XMLStatisticWriteOne(writer, &o, "statistic")
	}
	writer.WriteEnd("statistics")
	return nil
}

func XMLNetworkFilterWriteOne(writer *XMLWriter, object *NetworkFilter, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "network_filter"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Version(); ok {
		XMLVersionWriteOne(writer, r, "version")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLNetworkFilterWriteMany(writer *XMLWriter, structSlice *NetworkFilterSlice, plural, singular string) error {
	if plural == "" {
		plural = "network_filters"
	}
	if singular == "" {
		singular = "network_filter"
	}
	writer.WriteStart("", "network_filters", nil)
	for _, o := range structSlice.Slice() {
		XMLNetworkFilterWriteOne(writer, &o, "network_filter")
	}
	writer.WriteEnd("network_filters")
	return nil
}

func XMLVirtioScsiWriteOne(writer *XMLWriter, object *VirtioScsi, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "virtio_scsi"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Enabled(); ok {
		writer.WriteBool("enabled", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLVirtioScsiWriteMany(writer *XMLWriter, structSlice *VirtioScsiSlice, plural, singular string) error {
	if plural == "" {
		plural = "virtio_scsis"
	}
	if singular == "" {
		singular = "virtio_scsi"
	}
	writer.WriteStart("", "virtio_scsis", nil)
	for _, o := range structSlice.Slice() {
		XMLVirtioScsiWriteOne(writer, &o, "virtio_scsi")
	}
	writer.WriteEnd("virtio_scsis")
	return nil
}

func XMLPermissionWriteOne(writer *XMLWriter, object *Permission, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "permission"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Cluster(); ok {
		XMLClusterWriteOne(writer, r, "cluster")
	}
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.DataCenter(); ok {
		XMLDataCenterWriteOne(writer, r, "data_center")
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Disk(); ok {
		XMLDiskWriteOne(writer, r, "disk")
	}
	if r, ok := object.Group(); ok {
		XMLGroupWriteOne(writer, r, "group")
	}
	if r, ok := object.Host(); ok {
		XMLHostWriteOne(writer, r, "host")
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Role(); ok {
		XMLRoleWriteOne(writer, r, "role")
	}
	if r, ok := object.StorageDomain(); ok {
		XMLStorageDomainWriteOne(writer, r, "storage_domain")
	}
	if r, ok := object.Template(); ok {
		XMLTemplateWriteOne(writer, r, "template")
	}
	if r, ok := object.User(); ok {
		XMLUserWriteOne(writer, r, "user")
	}
	if r, ok := object.Vm(); ok {
		XMLVmWriteOne(writer, r, "vm")
	}
	if r, ok := object.VmPool(); ok {
		XMLVmPoolWriteOne(writer, r, "vm_pool")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLPermissionWriteMany(writer *XMLWriter, structSlice *PermissionSlice, plural, singular string) error {
	if plural == "" {
		plural = "permissions"
	}
	if singular == "" {
		singular = "permission"
	}
	writer.WriteStart("", "permissions", nil)
	for _, o := range structSlice.Slice() {
		XMLPermissionWriteOne(writer, &o, "permission")
	}
	writer.WriteEnd("permissions")
	return nil
}

func XMLDnsWriteOne(writer *XMLWriter, object *Dns, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "dns"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.SearchDomains(); ok {
		XMLHostWriteMany(writer, r, "search_domains", "host")
	}
	if r, ok := object.Servers(); ok {
		XMLHostWriteMany(writer, r, "servers", "host")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLDnsWriteMany(writer *XMLWriter, structSlice *DnsSlice, plural, singular string) error {
	if plural == "" {
		plural = "dnss"
	}
	if singular == "" {
		singular = "dns"
	}
	writer.WriteStart("", "dnss", nil)
	for _, o := range structSlice.Slice() {
		XMLDnsWriteOne(writer, &o, "dns")
	}
	writer.WriteEnd("dnss")
	return nil
}

func XMLNetworkLabelWriteOne(writer *XMLWriter, object *NetworkLabel, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "network_label"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.HostNic(); ok {
		XMLHostNicWriteOne(writer, r, "host_nic")
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Network(); ok {
		XMLNetworkWriteOne(writer, r, "network")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLNetworkLabelWriteMany(writer *XMLWriter, structSlice *NetworkLabelSlice, plural, singular string) error {
	if plural == "" {
		plural = "network_labels"
	}
	if singular == "" {
		singular = "network_label"
	}
	writer.WriteStart("", "network_labels", nil)
	for _, o := range structSlice.Slice() {
		XMLNetworkLabelWriteOne(writer, &o, "network_label")
	}
	writer.WriteEnd("network_labels")
	return nil
}

func XMLOpenstackVolumeAuthenticationKeyWriteOne(writer *XMLWriter, object *OpenstackVolumeAuthenticationKey, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "openstack_volume_authentication_key"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.CreationDate(); ok {
		writer.WriteDate("creation_date", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.OpenstackVolumeProvider(); ok {
		XMLOpenStackVolumeProviderWriteOne(writer, r, "openstack_volume_provider")
	}
	if r, ok := object.UsageType(); ok {
		XMLOpenstackVolumeAuthenticationKeyUsageTypeWriteOne(writer, r, "usage_type")
	}
	if r, ok := object.Uuid(); ok {
		writer.WriteCharacter("uuid", r)
	}
	if r, ok := object.Value(); ok {
		writer.WriteCharacter("value", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLOpenstackVolumeAuthenticationKeyWriteMany(writer *XMLWriter, structSlice *OpenstackVolumeAuthenticationKeySlice, plural, singular string) error {
	if plural == "" {
		plural = "openstack_volume_authentication_keys"
	}
	if singular == "" {
		singular = "openstack_volume_authentication_key"
	}
	writer.WriteStart("", "openstack_volume_authentication_keys", nil)
	for _, o := range structSlice.Slice() {
		XMLOpenstackVolumeAuthenticationKeyWriteOne(writer, &o, "openstack_volume_authentication_key")
	}
	writer.WriteEnd("openstack_volume_authentication_keys")
	return nil
}

func XMLOperatingSystemInfoWriteOne(writer *XMLWriter, object *OperatingSystemInfo, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "operating_system"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.LargeIcon(); ok {
		XMLIconWriteOne(writer, r, "large_icon")
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.SmallIcon(); ok {
		XMLIconWriteOne(writer, r, "small_icon")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLOperatingSystemInfoWriteMany(writer *XMLWriter, structSlice *OperatingSystemInfoSlice, plural, singular string) error {
	if plural == "" {
		plural = "operation_systems"
	}
	if singular == "" {
		singular = "operating_system"
	}
	writer.WriteStart("", "operation_systems", nil)
	for _, o := range structSlice.Slice() {
		XMLOperatingSystemInfoWriteOne(writer, &o, "operating_system")
	}
	writer.WriteEnd("operation_systems")
	return nil
}

func XMLSnapshotWriteOne(writer *XMLWriter, object *Snapshot, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "snapshot"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.AffinityLabels(); ok {
		XMLAffinityLabelWriteMany(writer, r, "affinity_labels", "affinity_label")
	}
	if r, ok := object.Applications(); ok {
		XMLApplicationWriteMany(writer, r, "applications", "application")
	}
	if r, ok := object.Bios(); ok {
		XMLBiosWriteOne(writer, r, "bios")
	}
	if r, ok := object.Cdroms(); ok {
		XMLCdromWriteMany(writer, r, "cdroms", "cdrom")
	}
	if r, ok := object.Cluster(); ok {
		XMLClusterWriteOne(writer, r, "cluster")
	}
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Console(); ok {
		XMLConsoleWriteOne(writer, r, "console")
	}
	if r, ok := object.Cpu(); ok {
		XMLCpuWriteOne(writer, r, "cpu")
	}
	if r, ok := object.CpuProfile(); ok {
		XMLCpuProfileWriteOne(writer, r, "cpu_profile")
	}
	if r, ok := object.CpuShares(); ok {
		writer.WriteInt64("cpu_shares", r)
	}
	if r, ok := object.CreationTime(); ok {
		writer.WriteDate("creation_time", r)
	}
	if r, ok := object.CustomCompatibilityVersion(); ok {
		XMLVersionWriteOne(writer, r, "custom_compatibility_version")
	}
	if r, ok := object.CustomCpuModel(); ok {
		writer.WriteCharacter("custom_cpu_model", r)
	}
	if r, ok := object.CustomEmulatedMachine(); ok {
		writer.WriteCharacter("custom_emulated_machine", r)
	}
	if r, ok := object.CustomProperties(); ok {
		XMLCustomPropertyWriteMany(writer, r, "custom_properties", "custom_property")
	}
	if r, ok := object.Date(); ok {
		writer.WriteDate("date", r)
	}
	if r, ok := object.DeleteProtected(); ok {
		writer.WriteBool("delete_protected", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.DiskAttachments(); ok {
		XMLDiskAttachmentWriteMany(writer, r, "disk_attachments", "disk_attachment")
	}
	if r, ok := object.Display(); ok {
		XMLDisplayWriteOne(writer, r, "display")
	}
	if r, ok := object.Domain(); ok {
		XMLDomainWriteOne(writer, r, "domain")
	}
	if r, ok := object.ExternalHostProvider(); ok {
		XMLExternalHostProviderWriteOne(writer, r, "external_host_provider")
	}
	if r, ok := object.Floppies(); ok {
		XMLFloppyWriteMany(writer, r, "floppies", "floppy")
	}
	if r, ok := object.Fqdn(); ok {
		writer.WriteCharacter("fqdn", r)
	}
	if r, ok := object.GraphicsConsoles(); ok {
		XMLGraphicsConsoleWriteMany(writer, r, "graphics_consoles", "graphics_console")
	}
	if r, ok := object.GuestOperatingSystem(); ok {
		XMLGuestOperatingSystemWriteOne(writer, r, "guest_operating_system")
	}
	if r, ok := object.GuestTimeZone(); ok {
		XMLTimeZoneWriteOne(writer, r, "guest_time_zone")
	}
	if r, ok := object.HighAvailability(); ok {
		XMLHighAvailabilityWriteOne(writer, r, "high_availability")
	}
	if r, ok := object.Host(); ok {
		XMLHostWriteOne(writer, r, "host")
	}
	if r, ok := object.HostDevices(); ok {
		XMLHostDeviceWriteMany(writer, r, "host_devices", "host_device")
	}
	if r, ok := object.Initialization(); ok {
		XMLInitializationWriteOne(writer, r, "initialization")
	}
	if r, ok := object.InstanceType(); ok {
		XMLInstanceTypeWriteOne(writer, r, "instance_type")
	}
	if r, ok := object.Io(); ok {
		XMLIoWriteOne(writer, r, "io")
	}
	if r, ok := object.KatelloErrata(); ok {
		XMLKatelloErratumWriteMany(writer, r, "katello_errata", "katello_erratum")
	}
	if r, ok := object.LargeIcon(); ok {
		XMLIconWriteOne(writer, r, "large_icon")
	}
	if r, ok := object.Lease(); ok {
		XMLStorageDomainLeaseWriteOne(writer, r, "lease")
	}
	if r, ok := object.Memory(); ok {
		writer.WriteInt64("memory", r)
	}
	if r, ok := object.MemoryPolicy(); ok {
		XMLMemoryPolicyWriteOne(writer, r, "memory_policy")
	}
	if r, ok := object.Migration(); ok {
		XMLMigrationOptionsWriteOne(writer, r, "migration")
	}
	if r, ok := object.MigrationDowntime(); ok {
		writer.WriteInt64("migration_downtime", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.NextRunConfigurationExists(); ok {
		writer.WriteBool("next_run_configuration_exists", r)
	}
	if r, ok := object.Nics(); ok {
		XMLNicWriteMany(writer, r, "nics", "nic")
	}
	if r, ok := object.NumaNodes(); ok {
		XMLNumaNodeWriteMany(writer, r, "host_numa_nodes", "host_numa_node")
	}
	if r, ok := object.NumaTuneMode(); ok {
		XMLNumaTuneModeWriteOne(writer, r, "numa_tune_mode")
	}
	if r, ok := object.Origin(); ok {
		writer.WriteCharacter("origin", r)
	}
	if r, ok := object.OriginalTemplate(); ok {
		XMLTemplateWriteOne(writer, r, "original_template")
	}
	if r, ok := object.Os(); ok {
		XMLOperatingSystemWriteOne(writer, r, "os")
	}
	if r, ok := object.Payloads(); ok {
		XMLPayloadWriteMany(writer, r, "payloads", "payload")
	}
	if r, ok := object.Permissions(); ok {
		XMLPermissionWriteMany(writer, r, "permissions", "permission")
	}
	if r, ok := object.PersistMemorystate(); ok {
		writer.WriteBool("persist_memorystate", r)
	}
	if r, ok := object.PlacementPolicy(); ok {
		XMLVmPlacementPolicyWriteOne(writer, r, "placement_policy")
	}
	if r, ok := object.Quota(); ok {
		XMLQuotaWriteOne(writer, r, "quota")
	}
	if r, ok := object.ReportedDevices(); ok {
		XMLReportedDeviceWriteMany(writer, r, "reported_devices", "reported_device")
	}
	if r, ok := object.RngDevice(); ok {
		XMLRngDeviceWriteOne(writer, r, "rng_device")
	}
	if r, ok := object.RunOnce(); ok {
		writer.WriteBool("run_once", r)
	}
	if r, ok := object.SerialNumber(); ok {
		XMLSerialNumberWriteOne(writer, r, "serial_number")
	}
	if r, ok := object.Sessions(); ok {
		XMLSessionWriteMany(writer, r, "sessions", "session")
	}
	if r, ok := object.SmallIcon(); ok {
		XMLIconWriteOne(writer, r, "small_icon")
	}
	if r, ok := object.SnapshotStatus(); ok {
		XMLSnapshotStatusWriteOne(writer, r, "snapshot_status")
	}
	if r, ok := object.SnapshotType(); ok {
		XMLSnapshotTypeWriteOne(writer, r, "snapshot_type")
	}
	if r, ok := object.Snapshots(); ok {
		XMLSnapshotWriteMany(writer, r, "snapshots", "snapshot")
	}
	if r, ok := object.SoundcardEnabled(); ok {
		writer.WriteBool("soundcard_enabled", r)
	}
	if r, ok := object.Sso(); ok {
		XMLSsoWriteOne(writer, r, "sso")
	}
	if r, ok := object.StartPaused(); ok {
		writer.WriteBool("start_paused", r)
	}
	if r, ok := object.StartTime(); ok {
		writer.WriteDate("start_time", r)
	}
	if r, ok := object.Stateless(); ok {
		writer.WriteBool("stateless", r)
	}
	if r, ok := object.Statistics(); ok {
		XMLStatisticWriteMany(writer, r, "statistics", "statistic")
	}
	if r, ok := object.Status(); ok {
		XMLVmStatusWriteOne(writer, r, "status")
	}
	if r, ok := object.StatusDetail(); ok {
		writer.WriteCharacter("status_detail", r)
	}
	if r, ok := object.StopReason(); ok {
		writer.WriteCharacter("stop_reason", r)
	}
	if r, ok := object.StopTime(); ok {
		writer.WriteDate("stop_time", r)
	}
	if r, ok := object.StorageDomain(); ok {
		XMLStorageDomainWriteOne(writer, r, "storage_domain")
	}
	if r, ok := object.Tags(); ok {
		XMLTagWriteMany(writer, r, "tags", "tag")
	}
	if r, ok := object.Template(); ok {
		XMLTemplateWriteOne(writer, r, "template")
	}
	if r, ok := object.TimeZone(); ok {
		XMLTimeZoneWriteOne(writer, r, "time_zone")
	}
	if r, ok := object.TunnelMigration(); ok {
		writer.WriteBool("tunnel_migration", r)
	}
	if r, ok := object.Type(); ok {
		XMLVmTypeWriteOne(writer, r, "type")
	}
	if r, ok := object.Usb(); ok {
		XMLUsbWriteOne(writer, r, "usb")
	}
	if r, ok := object.UseLatestTemplateVersion(); ok {
		writer.WriteBool("use_latest_template_version", r)
	}
	if r, ok := object.VirtioScsi(); ok {
		XMLVirtioScsiWriteOne(writer, r, "virtio_scsi")
	}
	if r, ok := object.Vm(); ok {
		XMLVmWriteOne(writer, r, "vm")
	}
	if r, ok := object.VmPool(); ok {
		XMLVmPoolWriteOne(writer, r, "vm_pool")
	}
	if r, ok := object.Watchdogs(); ok {
		XMLWatchdogWriteMany(writer, r, "watchdogs", "watchdog")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLSnapshotWriteMany(writer *XMLWriter, structSlice *SnapshotSlice, plural, singular string) error {
	if plural == "" {
		plural = "snapshots"
	}
	if singular == "" {
		singular = "snapshot"
	}
	writer.WriteStart("", "snapshots", nil)
	for _, o := range structSlice.Slice() {
		XMLSnapshotWriteOne(writer, &o, "snapshot")
	}
	writer.WriteEnd("snapshots")
	return nil
}

func XMLExternalHostGroupWriteOne(writer *XMLWriter, object *ExternalHostGroup, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "external_host_group"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.ArchitectureName(); ok {
		writer.WriteCharacter("architecture_name", r)
	}
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.DomainName(); ok {
		writer.WriteCharacter("domain_name", r)
	}
	if r, ok := object.ExternalHostProvider(); ok {
		XMLExternalHostProviderWriteOne(writer, r, "external_host_provider")
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.OperatingSystemName(); ok {
		writer.WriteCharacter("operating_system_name", r)
	}
	if r, ok := object.SubnetName(); ok {
		writer.WriteCharacter("subnet_name", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLExternalHostGroupWriteMany(writer *XMLWriter, structSlice *ExternalHostGroupSlice, plural, singular string) error {
	if plural == "" {
		plural = "external_host_groups"
	}
	if singular == "" {
		singular = "external_host_group"
	}
	writer.WriteStart("", "external_host_groups", nil)
	for _, o := range structSlice.Slice() {
		XMLExternalHostGroupWriteOne(writer, &o, "external_host_group")
	}
	writer.WriteEnd("external_host_groups")
	return nil
}

func XMLVersionWriteOne(writer *XMLWriter, object *Version, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "version"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Build_(); ok {
		writer.WriteInt64("build", r)
	}
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.FullVersion(); ok {
		writer.WriteCharacter("full_version", r)
	}
	if r, ok := object.Major(); ok {
		writer.WriteInt64("major", r)
	}
	if r, ok := object.Minor(); ok {
		writer.WriteInt64("minor", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Revision(); ok {
		writer.WriteInt64("revision", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLVersionWriteMany(writer *XMLWriter, structSlice *VersionSlice, plural, singular string) error {
	if plural == "" {
		plural = "versions"
	}
	if singular == "" {
		singular = "version"
	}
	writer.WriteStart("", "versions", nil)
	for _, o := range structSlice.Slice() {
		XMLVersionWriteOne(writer, &o, "version")
	}
	writer.WriteEnd("versions")
	return nil
}

func XMLVendorWriteOne(writer *XMLWriter, object *Vendor, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "vendor"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLVendorWriteMany(writer *XMLWriter, structSlice *VendorSlice, plural, singular string) error {
	if plural == "" {
		plural = "vendors"
	}
	if singular == "" {
		singular = "vendor"
	}
	writer.WriteStart("", "vendors", nil)
	for _, o := range structSlice.Slice() {
		XMLVendorWriteOne(writer, &o, "vendor")
	}
	writer.WriteEnd("vendors")
	return nil
}

func XMLPortMirroringWriteOne(writer *XMLWriter, object *PortMirroring, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "port_mirroring"
	}
	writer.WriteStart("", tag, nil)
	writer.WriteEnd(tag)
	return nil
}

func XMLPortMirroringWriteMany(writer *XMLWriter, structSlice *PortMirroringSlice, plural, singular string) error {
	if plural == "" {
		plural = "port_mirrorings"
	}
	if singular == "" {
		singular = "port_mirroring"
	}
	writer.WriteStart("", "port_mirrorings", nil)
	for _, o := range structSlice.Slice() {
		XMLPortMirroringWriteOne(writer, &o, "port_mirroring")
	}
	writer.WriteEnd("port_mirrorings")
	return nil
}

func XMLSerialNumberWriteOne(writer *XMLWriter, object *SerialNumber, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "serial_number"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Policy(); ok {
		XMLSerialNumberPolicyWriteOne(writer, r, "policy")
	}
	if r, ok := object.Value(); ok {
		writer.WriteCharacter("value", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLSerialNumberWriteMany(writer *XMLWriter, structSlice *SerialNumberSlice, plural, singular string) error {
	if plural == "" {
		plural = "serial_numbers"
	}
	if singular == "" {
		singular = "serial_number"
	}
	writer.WriteStart("", "serial_numbers", nil)
	for _, o := range structSlice.Slice() {
		XMLSerialNumberWriteOne(writer, &o, "serial_number")
	}
	writer.WriteEnd("serial_numbers")
	return nil
}

func XMLJobWriteOne(writer *XMLWriter, object *Job, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "job"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.AutoCleared(); ok {
		writer.WriteBool("auto_cleared", r)
	}
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.EndTime(); ok {
		writer.WriteDate("end_time", r)
	}
	if r, ok := object.External(); ok {
		writer.WriteBool("external", r)
	}
	if r, ok := object.LastUpdated(); ok {
		writer.WriteDate("last_updated", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Owner(); ok {
		XMLUserWriteOne(writer, r, "owner")
	}
	if r, ok := object.StartTime(); ok {
		writer.WriteDate("start_time", r)
	}
	if r, ok := object.Status(); ok {
		XMLJobStatusWriteOne(writer, r, "status")
	}
	if r, ok := object.Steps(); ok {
		XMLStepWriteMany(writer, r, "steps", "step")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLJobWriteMany(writer *XMLWriter, structSlice *JobSlice, plural, singular string) error {
	if plural == "" {
		plural = "jobs"
	}
	if singular == "" {
		singular = "job"
	}
	writer.WriteStart("", "jobs", nil)
	for _, o := range structSlice.Slice() {
		XMLJobWriteOne(writer, &o, "job")
	}
	writer.WriteEnd("jobs")
	return nil
}

func XMLTimeZoneWriteOne(writer *XMLWriter, object *TimeZone, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "time_zone"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.UtcOffset(); ok {
		writer.WriteCharacter("utc_offset", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLTimeZoneWriteMany(writer *XMLWriter, structSlice *TimeZoneSlice, plural, singular string) error {
	if plural == "" {
		plural = "time_zones"
	}
	if singular == "" {
		singular = "time_zone"
	}
	writer.WriteStart("", "time_zones", nil)
	for _, o := range structSlice.Slice() {
		XMLTimeZoneWriteOne(writer, &o, "time_zone")
	}
	writer.WriteEnd("time_zones")
	return nil
}

func XMLQuotaWriteOne(writer *XMLWriter, object *Quota, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "quota"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.ClusterHardLimitPct(); ok {
		writer.WriteInt64("cluster_hard_limit_pct", r)
	}
	if r, ok := object.ClusterSoftLimitPct(); ok {
		writer.WriteInt64("cluster_soft_limit_pct", r)
	}
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.DataCenter(); ok {
		XMLDataCenterWriteOne(writer, r, "data_center")
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Disks(); ok {
		XMLDiskWriteMany(writer, r, "disks", "disk")
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Permissions(); ok {
		XMLPermissionWriteMany(writer, r, "permissions", "permission")
	}
	if r, ok := object.QuotaClusterLimits(); ok {
		XMLQuotaClusterLimitWriteMany(writer, r, "quota_cluster_limits", "quota_cluster_limit")
	}
	if r, ok := object.QuotaStorageLimits(); ok {
		XMLQuotaStorageLimitWriteMany(writer, r, "quota_storage_limits", "quota_storage_limit")
	}
	if r, ok := object.StorageHardLimitPct(); ok {
		writer.WriteInt64("storage_hard_limit_pct", r)
	}
	if r, ok := object.StorageSoftLimitPct(); ok {
		writer.WriteInt64("storage_soft_limit_pct", r)
	}
	if r, ok := object.Users(); ok {
		XMLUserWriteMany(writer, r, "users", "user")
	}
	if r, ok := object.Vms(); ok {
		XMLVmWriteMany(writer, r, "vms", "vm")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLQuotaWriteMany(writer *XMLWriter, structSlice *QuotaSlice, plural, singular string) error {
	if plural == "" {
		plural = "quotas"
	}
	if singular == "" {
		singular = "quota"
	}
	writer.WriteStart("", "quotas", nil)
	for _, o := range structSlice.Slice() {
		XMLQuotaWriteOne(writer, &o, "quota")
	}
	writer.WriteEnd("quotas")
	return nil
}

func XMLVmWriteOne(writer *XMLWriter, object *Vm, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "vm"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.AffinityLabels(); ok {
		XMLAffinityLabelWriteMany(writer, r, "affinity_labels", "affinity_label")
	}
	if r, ok := object.Applications(); ok {
		XMLApplicationWriteMany(writer, r, "applications", "application")
	}
	if r, ok := object.Bios(); ok {
		XMLBiosWriteOne(writer, r, "bios")
	}
	if r, ok := object.Cdroms(); ok {
		XMLCdromWriteMany(writer, r, "cdroms", "cdrom")
	}
	if r, ok := object.Cluster(); ok {
		XMLClusterWriteOne(writer, r, "cluster")
	}
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Console(); ok {
		XMLConsoleWriteOne(writer, r, "console")
	}
	if r, ok := object.Cpu(); ok {
		XMLCpuWriteOne(writer, r, "cpu")
	}
	if r, ok := object.CpuProfile(); ok {
		XMLCpuProfileWriteOne(writer, r, "cpu_profile")
	}
	if r, ok := object.CpuShares(); ok {
		writer.WriteInt64("cpu_shares", r)
	}
	if r, ok := object.CreationTime(); ok {
		writer.WriteDate("creation_time", r)
	}
	if r, ok := object.CustomCompatibilityVersion(); ok {
		XMLVersionWriteOne(writer, r, "custom_compatibility_version")
	}
	if r, ok := object.CustomCpuModel(); ok {
		writer.WriteCharacter("custom_cpu_model", r)
	}
	if r, ok := object.CustomEmulatedMachine(); ok {
		writer.WriteCharacter("custom_emulated_machine", r)
	}
	if r, ok := object.CustomProperties(); ok {
		XMLCustomPropertyWriteMany(writer, r, "custom_properties", "custom_property")
	}
	if r, ok := object.DeleteProtected(); ok {
		writer.WriteBool("delete_protected", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.DiskAttachments(); ok {
		XMLDiskAttachmentWriteMany(writer, r, "disk_attachments", "disk_attachment")
	}
	if r, ok := object.Display(); ok {
		XMLDisplayWriteOne(writer, r, "display")
	}
	if r, ok := object.Domain(); ok {
		XMLDomainWriteOne(writer, r, "domain")
	}
	if r, ok := object.ExternalHostProvider(); ok {
		XMLExternalHostProviderWriteOne(writer, r, "external_host_provider")
	}
	if r, ok := object.Floppies(); ok {
		XMLFloppyWriteMany(writer, r, "floppies", "floppy")
	}
	if r, ok := object.Fqdn(); ok {
		writer.WriteCharacter("fqdn", r)
	}
	if r, ok := object.GraphicsConsoles(); ok {
		XMLGraphicsConsoleWriteMany(writer, r, "graphics_consoles", "graphics_console")
	}
	if r, ok := object.GuestOperatingSystem(); ok {
		XMLGuestOperatingSystemWriteOne(writer, r, "guest_operating_system")
	}
	if r, ok := object.GuestTimeZone(); ok {
		XMLTimeZoneWriteOne(writer, r, "guest_time_zone")
	}
	if r, ok := object.HighAvailability(); ok {
		XMLHighAvailabilityWriteOne(writer, r, "high_availability")
	}
	if r, ok := object.Host(); ok {
		XMLHostWriteOne(writer, r, "host")
	}
	if r, ok := object.HostDevices(); ok {
		XMLHostDeviceWriteMany(writer, r, "host_devices", "host_device")
	}
	if r, ok := object.Initialization(); ok {
		XMLInitializationWriteOne(writer, r, "initialization")
	}
	if r, ok := object.InstanceType(); ok {
		XMLInstanceTypeWriteOne(writer, r, "instance_type")
	}
	if r, ok := object.Io(); ok {
		XMLIoWriteOne(writer, r, "io")
	}
	if r, ok := object.KatelloErrata(); ok {
		XMLKatelloErratumWriteMany(writer, r, "katello_errata", "katello_erratum")
	}
	if r, ok := object.LargeIcon(); ok {
		XMLIconWriteOne(writer, r, "large_icon")
	}
	if r, ok := object.Lease(); ok {
		XMLStorageDomainLeaseWriteOne(writer, r, "lease")
	}
	if r, ok := object.Memory(); ok {
		writer.WriteInt64("memory", r)
	}
	if r, ok := object.MemoryPolicy(); ok {
		XMLMemoryPolicyWriteOne(writer, r, "memory_policy")
	}
	if r, ok := object.Migration(); ok {
		XMLMigrationOptionsWriteOne(writer, r, "migration")
	}
	if r, ok := object.MigrationDowntime(); ok {
		writer.WriteInt64("migration_downtime", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.NextRunConfigurationExists(); ok {
		writer.WriteBool("next_run_configuration_exists", r)
	}
	if r, ok := object.Nics(); ok {
		XMLNicWriteMany(writer, r, "nics", "nic")
	}
	if r, ok := object.NumaNodes(); ok {
		XMLNumaNodeWriteMany(writer, r, "host_numa_nodes", "host_numa_node")
	}
	if r, ok := object.NumaTuneMode(); ok {
		XMLNumaTuneModeWriteOne(writer, r, "numa_tune_mode")
	}
	if r, ok := object.Origin(); ok {
		writer.WriteCharacter("origin", r)
	}
	if r, ok := object.OriginalTemplate(); ok {
		XMLTemplateWriteOne(writer, r, "original_template")
	}
	if r, ok := object.Os(); ok {
		XMLOperatingSystemWriteOne(writer, r, "os")
	}
	if r, ok := object.Payloads(); ok {
		XMLPayloadWriteMany(writer, r, "payloads", "payload")
	}
	if r, ok := object.Permissions(); ok {
		XMLPermissionWriteMany(writer, r, "permissions", "permission")
	}
	if r, ok := object.PlacementPolicy(); ok {
		XMLVmPlacementPolicyWriteOne(writer, r, "placement_policy")
	}
	if r, ok := object.Quota(); ok {
		XMLQuotaWriteOne(writer, r, "quota")
	}
	if r, ok := object.ReportedDevices(); ok {
		XMLReportedDeviceWriteMany(writer, r, "reported_devices", "reported_device")
	}
	if r, ok := object.RngDevice(); ok {
		XMLRngDeviceWriteOne(writer, r, "rng_device")
	}
	if r, ok := object.RunOnce(); ok {
		writer.WriteBool("run_once", r)
	}
	if r, ok := object.SerialNumber(); ok {
		XMLSerialNumberWriteOne(writer, r, "serial_number")
	}
	if r, ok := object.Sessions(); ok {
		XMLSessionWriteMany(writer, r, "sessions", "session")
	}
	if r, ok := object.SmallIcon(); ok {
		XMLIconWriteOne(writer, r, "small_icon")
	}
	if r, ok := object.Snapshots(); ok {
		XMLSnapshotWriteMany(writer, r, "snapshots", "snapshot")
	}
	if r, ok := object.SoundcardEnabled(); ok {
		writer.WriteBool("soundcard_enabled", r)
	}
	if r, ok := object.Sso(); ok {
		XMLSsoWriteOne(writer, r, "sso")
	}
	if r, ok := object.StartPaused(); ok {
		writer.WriteBool("start_paused", r)
	}
	if r, ok := object.StartTime(); ok {
		writer.WriteDate("start_time", r)
	}
	if r, ok := object.Stateless(); ok {
		writer.WriteBool("stateless", r)
	}
	if r, ok := object.Statistics(); ok {
		XMLStatisticWriteMany(writer, r, "statistics", "statistic")
	}
	if r, ok := object.Status(); ok {
		XMLVmStatusWriteOne(writer, r, "status")
	}
	if r, ok := object.StatusDetail(); ok {
		writer.WriteCharacter("status_detail", r)
	}
	if r, ok := object.StopReason(); ok {
		writer.WriteCharacter("stop_reason", r)
	}
	if r, ok := object.StopTime(); ok {
		writer.WriteDate("stop_time", r)
	}
	if r, ok := object.StorageDomain(); ok {
		XMLStorageDomainWriteOne(writer, r, "storage_domain")
	}
	if r, ok := object.Tags(); ok {
		XMLTagWriteMany(writer, r, "tags", "tag")
	}
	if r, ok := object.Template(); ok {
		XMLTemplateWriteOne(writer, r, "template")
	}
	if r, ok := object.TimeZone(); ok {
		XMLTimeZoneWriteOne(writer, r, "time_zone")
	}
	if r, ok := object.TunnelMigration(); ok {
		writer.WriteBool("tunnel_migration", r)
	}
	if r, ok := object.Type(); ok {
		XMLVmTypeWriteOne(writer, r, "type")
	}
	if r, ok := object.Usb(); ok {
		XMLUsbWriteOne(writer, r, "usb")
	}
	if r, ok := object.UseLatestTemplateVersion(); ok {
		writer.WriteBool("use_latest_template_version", r)
	}
	if r, ok := object.VirtioScsi(); ok {
		XMLVirtioScsiWriteOne(writer, r, "virtio_scsi")
	}
	if r, ok := object.VmPool(); ok {
		XMLVmPoolWriteOne(writer, r, "vm_pool")
	}
	if r, ok := object.Watchdogs(); ok {
		XMLWatchdogWriteMany(writer, r, "watchdogs", "watchdog")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLVmWriteMany(writer *XMLWriter, structSlice *VmSlice, plural, singular string) error {
	if plural == "" {
		plural = "vms"
	}
	if singular == "" {
		singular = "vm"
	}
	writer.WriteStart("", "vms", nil)
	for _, o := range structSlice.Slice() {
		XMLVmWriteOne(writer, &o, "vm")
	}
	writer.WriteEnd("vms")
	return nil
}

func XMLPermitWriteOne(writer *XMLWriter, object *Permit, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "permit"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Administrative(); ok {
		writer.WriteBool("administrative", r)
	}
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Role(); ok {
		XMLRoleWriteOne(writer, r, "role")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLPermitWriteMany(writer *XMLWriter, structSlice *PermitSlice, plural, singular string) error {
	if plural == "" {
		plural = "permits"
	}
	if singular == "" {
		singular = "permit"
	}
	writer.WriteStart("", "permits", nil)
	for _, o := range structSlice.Slice() {
		XMLPermitWriteOne(writer, &o, "permit")
	}
	writer.WriteEnd("permits")
	return nil
}

func XMLVnicProfileMappingWriteOne(writer *XMLWriter, object *VnicProfileMapping, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "vnic_profile_mapping"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.SourceNetworkName(); ok {
		writer.WriteCharacter("source_network_name", r)
	}
	if r, ok := object.SourceNetworkProfileName(); ok {
		writer.WriteCharacter("source_network_profile_name", r)
	}
	if r, ok := object.TargetVnicProfile(); ok {
		XMLVnicProfileWriteOne(writer, r, "target_vnic_profile")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLVnicProfileMappingWriteMany(writer *XMLWriter, structSlice *VnicProfileMappingSlice, plural, singular string) error {
	if plural == "" {
		plural = "vnic_profile_mappings"
	}
	if singular == "" {
		singular = "vnic_profile_mapping"
	}
	writer.WriteStart("", "vnic_profile_mappings", nil)
	for _, o := range structSlice.Slice() {
		XMLVnicProfileMappingWriteOne(writer, &o, "vnic_profile_mapping")
	}
	writer.WriteEnd("vnic_profile_mappings")
	return nil
}

func XMLBrickProfileDetailWriteOne(writer *XMLWriter, object *BrickProfileDetail, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "brick_profile_detail"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Brick(); ok {
		XMLGlusterBrickWriteOne(writer, r, "brick")
	}
	if r, ok := object.ProfileDetails(); ok {
		XMLProfileDetailWriteMany(writer, r, "profile_details", "profile_detail")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLBrickProfileDetailWriteMany(writer *XMLWriter, structSlice *BrickProfileDetailSlice, plural, singular string) error {
	if plural == "" {
		plural = "brick_profile_details"
	}
	if singular == "" {
		singular = "brick_profile_detail"
	}
	writer.WriteStart("", "brick_profile_details", nil)
	for _, o := range structSlice.Slice() {
		XMLBrickProfileDetailWriteOne(writer, &o, "brick_profile_detail")
	}
	writer.WriteEnd("brick_profile_details")
	return nil
}

func XMLQuotaClusterLimitWriteOne(writer *XMLWriter, object *QuotaClusterLimit, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "quota_cluster_limit"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Cluster(); ok {
		XMLClusterWriteOne(writer, r, "cluster")
	}
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.MemoryLimit(); ok {
		writer.WriteFloat64("memory_limit", r)
	}
	if r, ok := object.MemoryUsage(); ok {
		writer.WriteFloat64("memory_usage", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Quota(); ok {
		XMLQuotaWriteOne(writer, r, "quota")
	}
	if r, ok := object.VcpuLimit(); ok {
		writer.WriteInt64("vcpu_limit", r)
	}
	if r, ok := object.VcpuUsage(); ok {
		writer.WriteInt64("vcpu_usage", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLQuotaClusterLimitWriteMany(writer *XMLWriter, structSlice *QuotaClusterLimitSlice, plural, singular string) error {
	if plural == "" {
		plural = "quota_cluster_limits"
	}
	if singular == "" {
		singular = "quota_cluster_limit"
	}
	writer.WriteStart("", "quota_cluster_limits", nil)
	for _, o := range structSlice.Slice() {
		XMLQuotaClusterLimitWriteOne(writer, &o, "quota_cluster_limit")
	}
	writer.WriteEnd("quota_cluster_limits")
	return nil
}

func XMLLogicalUnitWriteOne(writer *XMLWriter, object *LogicalUnit, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "logical_unit"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Address(); ok {
		writer.WriteCharacter("address", r)
	}
	if r, ok := object.DiscardMaxSize(); ok {
		writer.WriteInt64("discard_max_size", r)
	}
	if r, ok := object.DiscardZeroesData(); ok {
		writer.WriteBool("discard_zeroes_data", r)
	}
	if r, ok := object.DiskId(); ok {
		writer.WriteCharacter("disk_id", r)
	}
	if r, ok := object.LunMapping(); ok {
		writer.WriteInt64("lun_mapping", r)
	}
	if r, ok := object.Password(); ok {
		writer.WriteCharacter("password", r)
	}
	if r, ok := object.Paths(); ok {
		writer.WriteInt64("paths", r)
	}
	if r, ok := object.Port(); ok {
		writer.WriteInt64("port", r)
	}
	if r, ok := object.Portal(); ok {
		writer.WriteCharacter("portal", r)
	}
	if r, ok := object.ProductId(); ok {
		writer.WriteCharacter("product_id", r)
	}
	if r, ok := object.Serial(); ok {
		writer.WriteCharacter("serial", r)
	}
	if r, ok := object.Size(); ok {
		writer.WriteInt64("size", r)
	}
	if r, ok := object.Status(); ok {
		XMLLunStatusWriteOne(writer, r, "status")
	}
	if r, ok := object.StorageDomainId(); ok {
		writer.WriteCharacter("storage_domain_id", r)
	}
	if r, ok := object.Target(); ok {
		writer.WriteCharacter("target", r)
	}
	if r, ok := object.Username(); ok {
		writer.WriteCharacter("username", r)
	}
	if r, ok := object.VendorId(); ok {
		writer.WriteCharacter("vendor_id", r)
	}
	if r, ok := object.VolumeGroupId(); ok {
		writer.WriteCharacter("volume_group_id", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLLogicalUnitWriteMany(writer *XMLWriter, structSlice *LogicalUnitSlice, plural, singular string) error {
	if plural == "" {
		plural = "logical_units"
	}
	if singular == "" {
		singular = "logical_unit"
	}
	writer.WriteStart("", "logical_units", nil)
	for _, o := range structSlice.Slice() {
		XMLLogicalUnitWriteOne(writer, &o, "logical_unit")
	}
	writer.WriteEnd("logical_units")
	return nil
}

func XMLDisplayWriteOne(writer *XMLWriter, object *Display, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "display"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Address(); ok {
		writer.WriteCharacter("address", r)
	}
	if r, ok := object.AllowOverride(); ok {
		writer.WriteBool("allow_override", r)
	}
	if r, ok := object.Certificate(); ok {
		XMLCertificateWriteOne(writer, r, "certificate")
	}
	if r, ok := object.CopyPasteEnabled(); ok {
		writer.WriteBool("copy_paste_enabled", r)
	}
	if r, ok := object.DisconnectAction(); ok {
		writer.WriteCharacter("disconnect_action", r)
	}
	if r, ok := object.FileTransferEnabled(); ok {
		writer.WriteBool("file_transfer_enabled", r)
	}
	if r, ok := object.KeyboardLayout(); ok {
		writer.WriteCharacter("keyboard_layout", r)
	}
	if r, ok := object.Monitors(); ok {
		writer.WriteInt64("monitors", r)
	}
	if r, ok := object.Port(); ok {
		writer.WriteInt64("port", r)
	}
	if r, ok := object.Proxy(); ok {
		writer.WriteCharacter("proxy", r)
	}
	if r, ok := object.SecurePort(); ok {
		writer.WriteInt64("secure_port", r)
	}
	if r, ok := object.SingleQxlPci(); ok {
		writer.WriteBool("single_qxl_pci", r)
	}
	if r, ok := object.SmartcardEnabled(); ok {
		writer.WriteBool("smartcard_enabled", r)
	}
	if r, ok := object.Type(); ok {
		XMLDisplayTypeWriteOne(writer, r, "type")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLDisplayWriteMany(writer *XMLWriter, structSlice *DisplaySlice, plural, singular string) error {
	if plural == "" {
		plural = "displays"
	}
	if singular == "" {
		singular = "display"
	}
	writer.WriteStart("", "displays", nil)
	for _, o := range structSlice.Slice() {
		XMLDisplayWriteOne(writer, &o, "display")
	}
	writer.WriteEnd("displays")
	return nil
}

func XMLCertificateWriteOne(writer *XMLWriter, object *Certificate, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "certificate"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Content(); ok {
		writer.WriteCharacter("content", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Organization(); ok {
		writer.WriteCharacter("organization", r)
	}
	if r, ok := object.Subject(); ok {
		writer.WriteCharacter("subject", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLCertificateWriteMany(writer *XMLWriter, structSlice *CertificateSlice, plural, singular string) error {
	if plural == "" {
		plural = "certificates"
	}
	if singular == "" {
		singular = "certificate"
	}
	writer.WriteStart("", "certificates", nil)
	for _, o := range structSlice.Slice() {
		XMLCertificateWriteOne(writer, &o, "certificate")
	}
	writer.WriteEnd("certificates")
	return nil
}

func XMLDiskProfileWriteOne(writer *XMLWriter, object *DiskProfile, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "disk_profile"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Permissions(); ok {
		XMLPermissionWriteMany(writer, r, "permissions", "permission")
	}
	if r, ok := object.Qos(); ok {
		XMLQosWriteOne(writer, r, "qos")
	}
	if r, ok := object.StorageDomain(); ok {
		XMLStorageDomainWriteOne(writer, r, "storage_domain")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLDiskProfileWriteMany(writer *XMLWriter, structSlice *DiskProfileSlice, plural, singular string) error {
	if plural == "" {
		plural = "disk_profiles"
	}
	if singular == "" {
		singular = "disk_profile"
	}
	writer.WriteStart("", "disk_profiles", nil)
	for _, o := range structSlice.Slice() {
		XMLDiskProfileWriteOne(writer, &o, "disk_profile")
	}
	writer.WriteEnd("disk_profiles")
	return nil
}

func XMLMigrationPolicyWriteOne(writer *XMLWriter, object *MigrationPolicy, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "migration_policy"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLMigrationPolicyWriteMany(writer *XMLWriter, structSlice *MigrationPolicySlice, plural, singular string) error {
	if plural == "" {
		plural = "migration_policies"
	}
	if singular == "" {
		singular = "migration_policy"
	}
	writer.WriteStart("", "migration_policies", nil)
	for _, o := range structSlice.Slice() {
		XMLMigrationPolicyWriteOne(writer, &o, "migration_policy")
	}
	writer.WriteEnd("migration_policies")
	return nil
}

func XMLOpenStackSubnetWriteOne(writer *XMLWriter, object *OpenStackSubnet, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "openstack_subnet"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Cidr(); ok {
		writer.WriteCharacter("cidr", r)
	}
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.DnsServers(); ok {
		writer.WriteCharacters("dns_servers", r)
	}
	if r, ok := object.Gateway(); ok {
		writer.WriteCharacter("gateway", r)
	}
	if r, ok := object.IpVersion(); ok {
		writer.WriteCharacter("ip_version", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.OpenstackNetwork(); ok {
		XMLOpenStackNetworkWriteOne(writer, r, "openstack_network")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLOpenStackSubnetWriteMany(writer *XMLWriter, structSlice *OpenStackSubnetSlice, plural, singular string) error {
	if plural == "" {
		plural = "openstack_subnets"
	}
	if singular == "" {
		singular = "openstack_subnet"
	}
	writer.WriteStart("", "openstack_subnets", nil)
	for _, o := range structSlice.Slice() {
		XMLOpenStackSubnetWriteOne(writer, &o, "openstack_subnet")
	}
	writer.WriteEnd("openstack_subnets")
	return nil
}

func XMLFloppyWriteOne(writer *XMLWriter, object *Floppy, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "floppy"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.File(); ok {
		XMLFileWriteOne(writer, r, "file")
	}
	if r, ok := object.InstanceType(); ok {
		XMLInstanceTypeWriteOne(writer, r, "instance_type")
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Template(); ok {
		XMLTemplateWriteOne(writer, r, "template")
	}
	if r, ok := object.Vm(); ok {
		XMLVmWriteOne(writer, r, "vm")
	}
	if r, ok := object.Vms(); ok {
		XMLVmWriteMany(writer, r, "vms", "vm")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLFloppyWriteMany(writer *XMLWriter, structSlice *FloppySlice, plural, singular string) error {
	if plural == "" {
		plural = "floppies"
	}
	if singular == "" {
		singular = "floppy"
	}
	writer.WriteStart("", "floppies", nil)
	for _, o := range structSlice.Slice() {
		XMLFloppyWriteOne(writer, &o, "floppy")
	}
	writer.WriteEnd("floppies")
	return nil
}

func XMLClusterLevelWriteOne(writer *XMLWriter, object *ClusterLevel, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "cluster_level"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.CpuTypes(); ok {
		XMLCpuTypeWriteMany(writer, r, "cpu_types", "cpu_type")
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Permits(); ok {
		XMLPermitWriteMany(writer, r, "permits", "permit")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLClusterLevelWriteMany(writer *XMLWriter, structSlice *ClusterLevelSlice, plural, singular string) error {
	if plural == "" {
		plural = "cluster_levels"
	}
	if singular == "" {
		singular = "cluster_level"
	}
	writer.WriteStart("", "cluster_levels", nil)
	for _, o := range structSlice.Slice() {
		XMLClusterLevelWriteOne(writer, &o, "cluster_level")
	}
	writer.WriteEnd("cluster_levels")
	return nil
}

func XMLGroupWriteOne(writer *XMLWriter, object *Group, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "group"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Domain(); ok {
		XMLDomainWriteOne(writer, r, "domain")
	}
	if r, ok := object.DomainEntryId(); ok {
		writer.WriteCharacter("domain_entry_id", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Namespace(); ok {
		writer.WriteCharacter("namespace", r)
	}
	if r, ok := object.Permissions(); ok {
		XMLPermissionWriteMany(writer, r, "permissions", "permission")
	}
	if r, ok := object.Roles(); ok {
		XMLRoleWriteMany(writer, r, "roles", "role")
	}
	if r, ok := object.Tags(); ok {
		XMLTagWriteMany(writer, r, "tags", "tag")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLGroupWriteMany(writer *XMLWriter, structSlice *GroupSlice, plural, singular string) error {
	if plural == "" {
		plural = "groups"
	}
	if singular == "" {
		singular = "group"
	}
	writer.WriteStart("", "groups", nil)
	for _, o := range structSlice.Slice() {
		XMLGroupWriteOne(writer, &o, "group")
	}
	writer.WriteEnd("groups")
	return nil
}

func XMLReportedConfigurationWriteOne(writer *XMLWriter, object *ReportedConfiguration, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "reported_configuration"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.ActualValue(); ok {
		writer.WriteCharacter("actual_value", r)
	}
	if r, ok := object.ExpectedValue(); ok {
		writer.WriteCharacter("expected_value", r)
	}
	if r, ok := object.InSync(); ok {
		writer.WriteBool("in_sync", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLReportedConfigurationWriteMany(writer *XMLWriter, structSlice *ReportedConfigurationSlice, plural, singular string) error {
	if plural == "" {
		plural = "reported_configurations"
	}
	if singular == "" {
		singular = "reported_configuration"
	}
	writer.WriteStart("", "reported_configurations", nil)
	for _, o := range structSlice.Slice() {
		XMLReportedConfigurationWriteOne(writer, &o, "reported_configuration")
	}
	writer.WriteEnd("reported_configurations")
	return nil
}

func XMLDataCenterWriteOne(writer *XMLWriter, object *DataCenter, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "data_center"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Clusters(); ok {
		XMLClusterWriteMany(writer, r, "clusters", "cluster")
	}
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.IscsiBonds(); ok {
		XMLIscsiBondWriteMany(writer, r, "iscsi_bonds", "iscsi_bond")
	}
	if r, ok := object.Local(); ok {
		writer.WriteBool("local", r)
	}
	if r, ok := object.MacPool(); ok {
		XMLMacPoolWriteOne(writer, r, "mac_pool")
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Networks(); ok {
		XMLNetworkWriteMany(writer, r, "networks", "network")
	}
	if r, ok := object.Permissions(); ok {
		XMLPermissionWriteMany(writer, r, "permissions", "permission")
	}
	if r, ok := object.Qoss(); ok {
		XMLQosWriteMany(writer, r, "qoss", "qos")
	}
	if r, ok := object.QuotaMode(); ok {
		XMLQuotaModeTypeWriteOne(writer, r, "quota_mode")
	}
	if r, ok := object.Quotas(); ok {
		XMLQuotaWriteMany(writer, r, "quotas", "quota")
	}
	if r, ok := object.Status(); ok {
		XMLDataCenterStatusWriteOne(writer, r, "status")
	}
	if r, ok := object.StorageDomains(); ok {
		XMLStorageDomainWriteMany(writer, r, "storage_domains", "storage_domain")
	}
	if r, ok := object.StorageFormat(); ok {
		XMLStorageFormatWriteOne(writer, r, "storage_format")
	}
	if r, ok := object.SupportedVersions(); ok {
		XMLVersionWriteMany(writer, r, "supported_versions", "version")
	}
	if r, ok := object.Version(); ok {
		XMLVersionWriteOne(writer, r, "version")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLDataCenterWriteMany(writer *XMLWriter, structSlice *DataCenterSlice, plural, singular string) error {
	if plural == "" {
		plural = "data_centers"
	}
	if singular == "" {
		singular = "data_center"
	}
	writer.WriteStart("", "data_centers", nil)
	for _, o := range structSlice.Slice() {
		XMLDataCenterWriteOne(writer, &o, "data_center")
	}
	writer.WriteEnd("data_centers")
	return nil
}

func XMLCoreWriteOne(writer *XMLWriter, object *Core, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "core"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Index(); ok {
		writer.WriteInt64("index", r)
	}
	if r, ok := object.Socket(); ok {
		writer.WriteInt64("socket", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLCoreWriteMany(writer *XMLWriter, structSlice *CoreSlice, plural, singular string) error {
	if plural == "" {
		plural = "cores"
	}
	if singular == "" {
		singular = "core"
	}
	writer.WriteStart("", "cores", nil)
	for _, o := range structSlice.Slice() {
		XMLCoreWriteOne(writer, &o, "core")
	}
	writer.WriteEnd("cores")
	return nil
}

func XMLIpWriteOne(writer *XMLWriter, object *Ip, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "ip"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Address(); ok {
		writer.WriteCharacter("address", r)
	}
	if r, ok := object.Gateway(); ok {
		writer.WriteCharacter("gateway", r)
	}
	if r, ok := object.Netmask(); ok {
		writer.WriteCharacter("netmask", r)
	}
	if r, ok := object.Version(); ok {
		XMLIpVersionWriteOne(writer, r, "version")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLIpWriteMany(writer *XMLWriter, structSlice *IpSlice, plural, singular string) error {
	if plural == "" {
		plural = "ips"
	}
	if singular == "" {
		singular = "ip"
	}
	writer.WriteStart("", "ips", nil)
	for _, o := range structSlice.Slice() {
		XMLIpWriteOne(writer, &o, "ip")
	}
	writer.WriteEnd("ips")
	return nil
}

func XMLNicConfigurationWriteOne(writer *XMLWriter, object *NicConfiguration, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "nic_configuration"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.BootProtocol(); ok {
		XMLBootProtocolWriteOne(writer, r, "boot_protocol")
	}
	if r, ok := object.Ip(); ok {
		XMLIpWriteOne(writer, r, "ip")
	}
	if r, ok := object.Ipv6(); ok {
		XMLIpWriteOne(writer, r, "ipv6")
	}
	if r, ok := object.Ipv6BootProtocol(); ok {
		XMLBootProtocolWriteOne(writer, r, "ipv6_boot_protocol")
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.OnBoot(); ok {
		writer.WriteBool("on_boot", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLNicConfigurationWriteMany(writer *XMLWriter, structSlice *NicConfigurationSlice, plural, singular string) error {
	if plural == "" {
		plural = "nic_configurations"
	}
	if singular == "" {
		singular = "nic_configuration"
	}
	writer.WriteStart("", "nic_configurations", nil)
	for _, o := range structSlice.Slice() {
		XMLNicConfigurationWriteOne(writer, &o, "nic_configuration")
	}
	writer.WriteEnd("nic_configurations")
	return nil
}

func XMLValueWriteOne(writer *XMLWriter, object *Value, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "value"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Datum(); ok {
		writer.WriteFloat64("datum", r)
	}
	if r, ok := object.Detail(); ok {
		writer.WriteCharacter("detail", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLValueWriteMany(writer *XMLWriter, structSlice *ValueSlice, plural, singular string) error {
	if plural == "" {
		plural = "values"
	}
	if singular == "" {
		singular = "value"
	}
	writer.WriteStart("", "values", nil)
	for _, o := range structSlice.Slice() {
		XMLValueWriteOne(writer, &o, "value")
	}
	writer.WriteEnd("values")
	return nil
}

func XMLAuthorizedKeyWriteOne(writer *XMLWriter, object *AuthorizedKey, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "authorized_key"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Key(); ok {
		writer.WriteCharacter("key", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.User(); ok {
		XMLUserWriteOne(writer, r, "user")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLAuthorizedKeyWriteMany(writer *XMLWriter, structSlice *AuthorizedKeySlice, plural, singular string) error {
	if plural == "" {
		plural = "authorized_keys"
	}
	if singular == "" {
		singular = "authorized_key"
	}
	writer.WriteStart("", "authorized_keys", nil)
	for _, o := range structSlice.Slice() {
		XMLAuthorizedKeyWriteOne(writer, &o, "authorized_key")
	}
	writer.WriteEnd("authorized_keys")
	return nil
}

func XMLTemplateWriteOne(writer *XMLWriter, object *Template, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "template"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Bios(); ok {
		XMLBiosWriteOne(writer, r, "bios")
	}
	if r, ok := object.Cdroms(); ok {
		XMLCdromWriteMany(writer, r, "cdroms", "cdrom")
	}
	if r, ok := object.Cluster(); ok {
		XMLClusterWriteOne(writer, r, "cluster")
	}
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Console(); ok {
		XMLConsoleWriteOne(writer, r, "console")
	}
	if r, ok := object.Cpu(); ok {
		XMLCpuWriteOne(writer, r, "cpu")
	}
	if r, ok := object.CpuProfile(); ok {
		XMLCpuProfileWriteOne(writer, r, "cpu_profile")
	}
	if r, ok := object.CpuShares(); ok {
		writer.WriteInt64("cpu_shares", r)
	}
	if r, ok := object.CreationTime(); ok {
		writer.WriteDate("creation_time", r)
	}
	if r, ok := object.CustomCompatibilityVersion(); ok {
		XMLVersionWriteOne(writer, r, "custom_compatibility_version")
	}
	if r, ok := object.CustomCpuModel(); ok {
		writer.WriteCharacter("custom_cpu_model", r)
	}
	if r, ok := object.CustomEmulatedMachine(); ok {
		writer.WriteCharacter("custom_emulated_machine", r)
	}
	if r, ok := object.CustomProperties(); ok {
		XMLCustomPropertyWriteMany(writer, r, "custom_properties", "custom_property")
	}
	if r, ok := object.DeleteProtected(); ok {
		writer.WriteBool("delete_protected", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.DiskAttachments(); ok {
		XMLDiskAttachmentWriteMany(writer, r, "disk_attachments", "disk_attachment")
	}
	if r, ok := object.Display(); ok {
		XMLDisplayWriteOne(writer, r, "display")
	}
	if r, ok := object.Domain(); ok {
		XMLDomainWriteOne(writer, r, "domain")
	}
	if r, ok := object.GraphicsConsoles(); ok {
		XMLGraphicsConsoleWriteMany(writer, r, "graphics_consoles", "graphics_console")
	}
	if r, ok := object.HighAvailability(); ok {
		XMLHighAvailabilityWriteOne(writer, r, "high_availability")
	}
	if r, ok := object.Initialization(); ok {
		XMLInitializationWriteOne(writer, r, "initialization")
	}
	if r, ok := object.Io(); ok {
		XMLIoWriteOne(writer, r, "io")
	}
	if r, ok := object.LargeIcon(); ok {
		XMLIconWriteOne(writer, r, "large_icon")
	}
	if r, ok := object.Lease(); ok {
		XMLStorageDomainLeaseWriteOne(writer, r, "lease")
	}
	if r, ok := object.Memory(); ok {
		writer.WriteInt64("memory", r)
	}
	if r, ok := object.MemoryPolicy(); ok {
		XMLMemoryPolicyWriteOne(writer, r, "memory_policy")
	}
	if r, ok := object.Migration(); ok {
		XMLMigrationOptionsWriteOne(writer, r, "migration")
	}
	if r, ok := object.MigrationDowntime(); ok {
		writer.WriteInt64("migration_downtime", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Nics(); ok {
		XMLNicWriteMany(writer, r, "nics", "nic")
	}
	if r, ok := object.Origin(); ok {
		writer.WriteCharacter("origin", r)
	}
	if r, ok := object.Os(); ok {
		XMLOperatingSystemWriteOne(writer, r, "os")
	}
	if r, ok := object.Permissions(); ok {
		XMLPermissionWriteMany(writer, r, "permissions", "permission")
	}
	if r, ok := object.Quota(); ok {
		XMLQuotaWriteOne(writer, r, "quota")
	}
	if r, ok := object.RngDevice(); ok {
		XMLRngDeviceWriteOne(writer, r, "rng_device")
	}
	if r, ok := object.SerialNumber(); ok {
		XMLSerialNumberWriteOne(writer, r, "serial_number")
	}
	if r, ok := object.SmallIcon(); ok {
		XMLIconWriteOne(writer, r, "small_icon")
	}
	if r, ok := object.SoundcardEnabled(); ok {
		writer.WriteBool("soundcard_enabled", r)
	}
	if r, ok := object.Sso(); ok {
		XMLSsoWriteOne(writer, r, "sso")
	}
	if r, ok := object.StartPaused(); ok {
		writer.WriteBool("start_paused", r)
	}
	if r, ok := object.Stateless(); ok {
		writer.WriteBool("stateless", r)
	}
	if r, ok := object.Status(); ok {
		XMLTemplateStatusWriteOne(writer, r, "status")
	}
	if r, ok := object.StorageDomain(); ok {
		XMLStorageDomainWriteOne(writer, r, "storage_domain")
	}
	if r, ok := object.Tags(); ok {
		XMLTagWriteMany(writer, r, "tags", "tag")
	}
	if r, ok := object.TimeZone(); ok {
		XMLTimeZoneWriteOne(writer, r, "time_zone")
	}
	if r, ok := object.TunnelMigration(); ok {
		writer.WriteBool("tunnel_migration", r)
	}
	if r, ok := object.Type(); ok {
		XMLVmTypeWriteOne(writer, r, "type")
	}
	if r, ok := object.Usb(); ok {
		XMLUsbWriteOne(writer, r, "usb")
	}
	if r, ok := object.Version(); ok {
		XMLTemplateVersionWriteOne(writer, r, "version")
	}
	if r, ok := object.VirtioScsi(); ok {
		XMLVirtioScsiWriteOne(writer, r, "virtio_scsi")
	}
	if r, ok := object.Vm(); ok {
		XMLVmWriteOne(writer, r, "vm")
	}
	if r, ok := object.Watchdogs(); ok {
		XMLWatchdogWriteMany(writer, r, "watchdogs", "watchdog")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLTemplateWriteMany(writer *XMLWriter, structSlice *TemplateSlice, plural, singular string) error {
	if plural == "" {
		plural = "templates"
	}
	if singular == "" {
		singular = "template"
	}
	writer.WriteStart("", "templates", nil)
	for _, o := range structSlice.Slice() {
		XMLTemplateWriteOne(writer, &o, "template")
	}
	writer.WriteEnd("templates")
	return nil
}

func XMLUserWriteOne(writer *XMLWriter, object *User, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "user"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Department(); ok {
		writer.WriteCharacter("department", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Domain(); ok {
		XMLDomainWriteOne(writer, r, "domain")
	}
	if r, ok := object.DomainEntryId(); ok {
		writer.WriteCharacter("domain_entry_id", r)
	}
	if r, ok := object.Email(); ok {
		writer.WriteCharacter("email", r)
	}
	if r, ok := object.Groups(); ok {
		XMLGroupWriteMany(writer, r, "groups", "group")
	}
	if r, ok := object.LastName(); ok {
		writer.WriteCharacter("last_name", r)
	}
	if r, ok := object.LoggedIn(); ok {
		writer.WriteBool("logged_in", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Namespace(); ok {
		writer.WriteCharacter("namespace", r)
	}
	if r, ok := object.Password(); ok {
		writer.WriteCharacter("password", r)
	}
	if r, ok := object.Permissions(); ok {
		XMLPermissionWriteMany(writer, r, "permissions", "permission")
	}
	if r, ok := object.Principal(); ok {
		writer.WriteCharacter("principal", r)
	}
	if r, ok := object.Roles(); ok {
		XMLRoleWriteMany(writer, r, "roles", "role")
	}
	if r, ok := object.SshPublicKeys(); ok {
		XMLSshPublicKeyWriteMany(writer, r, "ssh_public_keys", "ssh_public_key")
	}
	if r, ok := object.Tags(); ok {
		XMLTagWriteMany(writer, r, "tags", "tag")
	}
	if r, ok := object.UserName(); ok {
		writer.WriteCharacter("user_name", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLUserWriteMany(writer *XMLWriter, structSlice *UserSlice, plural, singular string) error {
	if plural == "" {
		plural = "users"
	}
	if singular == "" {
		singular = "user"
	}
	writer.WriteStart("", "users", nil)
	for _, o := range structSlice.Slice() {
		XMLUserWriteOne(writer, &o, "user")
	}
	writer.WriteEnd("users")
	return nil
}

func XMLIoWriteOne(writer *XMLWriter, object *Io, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "io"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Threads(); ok {
		writer.WriteInt64("threads", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLIoWriteMany(writer *XMLWriter, structSlice *IoSlice, plural, singular string) error {
	if plural == "" {
		plural = "ios"
	}
	if singular == "" {
		singular = "io"
	}
	writer.WriteStart("", "ios", nil)
	for _, o := range structSlice.Slice() {
		XMLIoWriteOne(writer, &o, "io")
	}
	writer.WriteEnd("ios")
	return nil
}

func XMLMethodWriteOne(writer *XMLWriter, object *Method, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "method"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = string(r)
	}
	writer.WriteStart("", tag, attrs)
	writer.WriteEnd(tag)
	return nil
}

func XMLMethodWriteMany(writer *XMLWriter, structSlice *MethodSlice, plural, singular string) error {
	if plural == "" {
		plural = "methods"
	}
	if singular == "" {
		singular = "method"
	}
	writer.WriteStart("", "methods", nil)
	for _, o := range structSlice.Slice() {
		XMLMethodWriteOne(writer, &o, "method")
	}
	writer.WriteEnd("methods")
	return nil
}

func XMLEventWriteOne(writer *XMLWriter, object *Event, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "event"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Cluster(); ok {
		XMLClusterWriteOne(writer, r, "cluster")
	}
	if r, ok := object.Code(); ok {
		writer.WriteInt64("code", r)
	}
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.CorrelationId(); ok {
		writer.WriteCharacter("correlation_id", r)
	}
	if r, ok := object.CustomData(); ok {
		writer.WriteCharacter("custom_data", r)
	}
	if r, ok := object.CustomId(); ok {
		writer.WriteInt64("custom_id", r)
	}
	if r, ok := object.DataCenter(); ok {
		XMLDataCenterWriteOne(writer, r, "data_center")
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.FloodRate(); ok {
		writer.WriteInt64("flood_rate", r)
	}
	if r, ok := object.Host(); ok {
		XMLHostWriteOne(writer, r, "host")
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Origin(); ok {
		writer.WriteCharacter("origin", r)
	}
	if r, ok := object.Severity(); ok {
		XMLLogSeverityWriteOne(writer, r, "severity")
	}
	if r, ok := object.StorageDomain(); ok {
		XMLStorageDomainWriteOne(writer, r, "storage_domain")
	}
	if r, ok := object.Template(); ok {
		XMLTemplateWriteOne(writer, r, "template")
	}
	if r, ok := object.Time(); ok {
		writer.WriteDate("time", r)
	}
	if r, ok := object.User(); ok {
		XMLUserWriteOne(writer, r, "user")
	}
	if r, ok := object.Vm(); ok {
		XMLVmWriteOne(writer, r, "vm")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLEventWriteMany(writer *XMLWriter, structSlice *EventSlice, plural, singular string) error {
	if plural == "" {
		plural = "events"
	}
	if singular == "" {
		singular = "event"
	}
	writer.WriteStart("", "events", nil)
	for _, o := range structSlice.Slice() {
		XMLEventWriteOne(writer, &o, "event")
	}
	writer.WriteEnd("events")
	return nil
}

func XMLVmSummaryWriteOne(writer *XMLWriter, object *VmSummary, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "vm_summary"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Active(); ok {
		writer.WriteInt64("active", r)
	}
	if r, ok := object.Migrating(); ok {
		writer.WriteInt64("migrating", r)
	}
	if r, ok := object.Total(); ok {
		writer.WriteInt64("total", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLVmSummaryWriteMany(writer *XMLWriter, structSlice *VmSummarySlice, plural, singular string) error {
	if plural == "" {
		plural = "vm_summaries"
	}
	if singular == "" {
		singular = "vm_summary"
	}
	writer.WriteStart("", "vm_summaries", nil)
	for _, o := range structSlice.Slice() {
		XMLVmSummaryWriteOne(writer, &o, "vm_summary")
	}
	writer.WriteEnd("vm_summaries")
	return nil
}

func XMLIscsiDetailsWriteOne(writer *XMLWriter, object *IscsiDetails, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "iscsi_details"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Address(); ok {
		writer.WriteCharacter("address", r)
	}
	if r, ok := object.DiskId(); ok {
		writer.WriteCharacter("disk_id", r)
	}
	if r, ok := object.Initiator(); ok {
		writer.WriteCharacter("initiator", r)
	}
	if r, ok := object.LunMapping(); ok {
		writer.WriteInt64("lun_mapping", r)
	}
	if r, ok := object.Password(); ok {
		writer.WriteCharacter("password", r)
	}
	if r, ok := object.Paths(); ok {
		writer.WriteInt64("paths", r)
	}
	if r, ok := object.Port(); ok {
		writer.WriteInt64("port", r)
	}
	if r, ok := object.Portal(); ok {
		writer.WriteCharacter("portal", r)
	}
	if r, ok := object.ProductId(); ok {
		writer.WriteCharacter("product_id", r)
	}
	if r, ok := object.Serial(); ok {
		writer.WriteCharacter("serial", r)
	}
	if r, ok := object.Size(); ok {
		writer.WriteInt64("size", r)
	}
	if r, ok := object.Status(); ok {
		writer.WriteCharacter("status", r)
	}
	if r, ok := object.StorageDomainId(); ok {
		writer.WriteCharacter("storage_domain_id", r)
	}
	if r, ok := object.Target(); ok {
		writer.WriteCharacter("target", r)
	}
	if r, ok := object.Username(); ok {
		writer.WriteCharacter("username", r)
	}
	if r, ok := object.VendorId(); ok {
		writer.WriteCharacter("vendor_id", r)
	}
	if r, ok := object.VolumeGroupId(); ok {
		writer.WriteCharacter("volume_group_id", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLIscsiDetailsWriteMany(writer *XMLWriter, structSlice *IscsiDetailsSlice, plural, singular string) error {
	if plural == "" {
		plural = "iscsi_detailss"
	}
	if singular == "" {
		singular = "iscsi_details"
	}
	writer.WriteStart("", "iscsi_detailss", nil)
	for _, o := range structSlice.Slice() {
		XMLIscsiDetailsWriteOne(writer, &o, "iscsi_details")
	}
	writer.WriteEnd("iscsi_detailss")
	return nil
}

func XMLDeviceWriteOne(writer *XMLWriter, object *Device, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "device"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.InstanceType(); ok {
		XMLInstanceTypeWriteOne(writer, r, "instance_type")
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Template(); ok {
		XMLTemplateWriteOne(writer, r, "template")
	}
	if r, ok := object.Vm(); ok {
		XMLVmWriteOne(writer, r, "vm")
	}
	if r, ok := object.Vms(); ok {
		XMLVmWriteMany(writer, r, "vms", "vm")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLDeviceWriteMany(writer *XMLWriter, structSlice *DeviceSlice, plural, singular string) error {
	if plural == "" {
		plural = "devices"
	}
	if singular == "" {
		singular = "device"
	}
	writer.WriteStart("", "devices", nil)
	for _, o := range structSlice.Slice() {
		XMLDeviceWriteOne(writer, &o, "device")
	}
	writer.WriteEnd("devices")
	return nil
}

func XMLApiSummaryItemWriteOne(writer *XMLWriter, object *ApiSummaryItem, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "api_summary_item"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Active(); ok {
		writer.WriteInt64("active", r)
	}
	if r, ok := object.Total(); ok {
		writer.WriteInt64("total", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLApiSummaryItemWriteMany(writer *XMLWriter, structSlice *ApiSummaryItemSlice, plural, singular string) error {
	if plural == "" {
		plural = "api_summary_items"
	}
	if singular == "" {
		singular = "api_summary_item"
	}
	writer.WriteStart("", "api_summary_items", nil)
	for _, o := range structSlice.Slice() {
		XMLApiSummaryItemWriteOne(writer, &o, "api_summary_item")
	}
	writer.WriteEnd("api_summary_items")
	return nil
}

func XMLProfileDetailWriteOne(writer *XMLWriter, object *ProfileDetail, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "profile_detail"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.BlockStatistics(); ok {
		XMLBlockStatisticWriteMany(writer, r, "block_statistics", "block_statistic")
	}
	if r, ok := object.Duration(); ok {
		writer.WriteInt64("duration", r)
	}
	if r, ok := object.FopStatistics(); ok {
		XMLFopStatisticWriteMany(writer, r, "fop_statistics", "fop_statistic")
	}
	if r, ok := object.ProfileType(); ok {
		writer.WriteCharacter("profile_type", r)
	}
	if r, ok := object.Statistics(); ok {
		XMLStatisticWriteMany(writer, r, "statistics", "statistic")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLProfileDetailWriteMany(writer *XMLWriter, structSlice *ProfileDetailSlice, plural, singular string) error {
	if plural == "" {
		plural = "profile_details"
	}
	if singular == "" {
		singular = "profile_detail"
	}
	writer.WriteStart("", "profile_details", nil)
	for _, o := range structSlice.Slice() {
		XMLProfileDetailWriteOne(writer, &o, "profile_detail")
	}
	writer.WriteEnd("profile_details")
	return nil
}

func XMLHostNicVirtualFunctionsConfigurationWriteOne(writer *XMLWriter, object *HostNicVirtualFunctionsConfiguration, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "host_nic_virtual_functions_configuration"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.AllNetworksAllowed(); ok {
		writer.WriteBool("all_networks_allowed", r)
	}
	if r, ok := object.MaxNumberOfVirtualFunctions(); ok {
		writer.WriteInt64("max_number_of_virtual_functions", r)
	}
	if r, ok := object.NumberOfVirtualFunctions(); ok {
		writer.WriteInt64("number_of_virtual_functions", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLHostNicVirtualFunctionsConfigurationWriteMany(writer *XMLWriter, structSlice *HostNicVirtualFunctionsConfigurationSlice, plural, singular string) error {
	if plural == "" {
		plural = "host_nic_virtual_functions_configurations"
	}
	if singular == "" {
		singular = "host_nic_virtual_functions_configuration"
	}
	writer.WriteStart("", "host_nic_virtual_functions_configurations", nil)
	for _, o := range structSlice.Slice() {
		XMLHostNicVirtualFunctionsConfigurationWriteOne(writer, &o, "host_nic_virtual_functions_configuration")
	}
	writer.WriteEnd("host_nic_virtual_functions_configurations")
	return nil
}

func XMLHostNicWriteOne(writer *XMLWriter, object *HostNic, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "host_nic"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.AdAggregatorId(); ok {
		writer.WriteInt64("ad_aggregator_id", r)
	}
	if r, ok := object.BaseInterface(); ok {
		writer.WriteCharacter("base_interface", r)
	}
	if r, ok := object.Bonding(); ok {
		XMLBondingWriteOne(writer, r, "bonding")
	}
	if r, ok := object.BootProtocol(); ok {
		XMLBootProtocolWriteOne(writer, r, "boot_protocol")
	}
	if r, ok := object.Bridged(); ok {
		writer.WriteBool("bridged", r)
	}
	if r, ok := object.CheckConnectivity(); ok {
		writer.WriteBool("check_connectivity", r)
	}
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.CustomConfiguration(); ok {
		writer.WriteBool("custom_configuration", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Host(); ok {
		XMLHostWriteOne(writer, r, "host")
	}
	if r, ok := object.Ip(); ok {
		XMLIpWriteOne(writer, r, "ip")
	}
	if r, ok := object.Ipv6(); ok {
		XMLIpWriteOne(writer, r, "ipv6")
	}
	if r, ok := object.Ipv6BootProtocol(); ok {
		XMLBootProtocolWriteOne(writer, r, "ipv6_boot_protocol")
	}
	if r, ok := object.Mac(); ok {
		XMLMacWriteOne(writer, r, "mac")
	}
	if r, ok := object.Mtu(); ok {
		writer.WriteInt64("mtu", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Network(); ok {
		XMLNetworkWriteOne(writer, r, "network")
	}
	if r, ok := object.NetworkLabels(); ok {
		XMLNetworkLabelWriteMany(writer, r, "network_labels", "network_label")
	}
	if r, ok := object.OverrideConfiguration(); ok {
		writer.WriteBool("override_configuration", r)
	}
	if r, ok := object.PhysicalFunction(); ok {
		XMLHostNicWriteOne(writer, r, "physical_function")
	}
	if r, ok := object.Properties(); ok {
		XMLPropertyWriteMany(writer, r, "properties", "property")
	}
	if r, ok := object.Qos(); ok {
		XMLQosWriteOne(writer, r, "qos")
	}
	if r, ok := object.Speed(); ok {
		writer.WriteInt64("speed", r)
	}
	if r, ok := object.Statistics(); ok {
		XMLStatisticWriteMany(writer, r, "statistics", "statistic")
	}
	if r, ok := object.Status(); ok {
		XMLNicStatusWriteOne(writer, r, "status")
	}
	if r, ok := object.VirtualFunctionsConfiguration(); ok {
		XMLHostNicVirtualFunctionsConfigurationWriteOne(writer, r, "virtual_functions_configuration")
	}
	if r, ok := object.Vlan(); ok {
		XMLVlanWriteOne(writer, r, "vlan")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLHostNicWriteMany(writer *XMLWriter, structSlice *HostNicSlice, plural, singular string) error {
	if plural == "" {
		plural = "host_nics"
	}
	if singular == "" {
		singular = "host_nic"
	}
	writer.WriteStart("", "host_nics", nil)
	for _, o := range structSlice.Slice() {
		XMLHostNicWriteOne(writer, &o, "host_nic")
	}
	writer.WriteEnd("host_nics")
	return nil
}

func XMLOpenStackNetworkProviderWriteOne(writer *XMLWriter, object *OpenStackNetworkProvider, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "openstack_network_provider"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.AgentConfiguration(); ok {
		XMLAgentConfigurationWriteOne(writer, r, "agent_configuration")
	}
	if r, ok := object.AuthenticationUrl(); ok {
		writer.WriteCharacter("authentication_url", r)
	}
	if r, ok := object.Certificates(); ok {
		XMLCertificateWriteMany(writer, r, "certificates", "certificate")
	}
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Networks(); ok {
		XMLOpenStackNetworkWriteMany(writer, r, "networks", "openstack_network")
	}
	if r, ok := object.Password(); ok {
		writer.WriteCharacter("password", r)
	}
	if r, ok := object.PluginType(); ok {
		XMLNetworkPluginTypeWriteOne(writer, r, "plugin_type")
	}
	if r, ok := object.Properties(); ok {
		XMLPropertyWriteMany(writer, r, "properties", "property")
	}
	if r, ok := object.ReadOnly(); ok {
		writer.WriteBool("read_only", r)
	}
	if r, ok := object.RequiresAuthentication(); ok {
		writer.WriteBool("requires_authentication", r)
	}
	if r, ok := object.Subnets(); ok {
		XMLOpenStackSubnetWriteMany(writer, r, "subnets", "openstack_subnet")
	}
	if r, ok := object.TenantName(); ok {
		writer.WriteCharacter("tenant_name", r)
	}
	if r, ok := object.Type(); ok {
		XMLOpenStackNetworkProviderTypeWriteOne(writer, r, "type")
	}
	if r, ok := object.Url(); ok {
		writer.WriteCharacter("url", r)
	}
	if r, ok := object.Username(); ok {
		writer.WriteCharacter("username", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLOpenStackNetworkProviderWriteMany(writer *XMLWriter, structSlice *OpenStackNetworkProviderSlice, plural, singular string) error {
	if plural == "" {
		plural = "openstack_network_providers"
	}
	if singular == "" {
		singular = "openstack_network_provider"
	}
	writer.WriteStart("", "openstack_network_providers", nil)
	for _, o := range structSlice.Slice() {
		XMLOpenStackNetworkProviderWriteOne(writer, &o, "openstack_network_provider")
	}
	writer.WriteEnd("openstack_network_providers")
	return nil
}

func XMLCpuTypeWriteOne(writer *XMLWriter, object *CpuType, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "cpu_type"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Architecture(); ok {
		XMLArchitectureWriteOne(writer, r, "architecture")
	}
	if r, ok := object.Level(); ok {
		writer.WriteInt64("level", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLCpuTypeWriteMany(writer *XMLWriter, structSlice *CpuTypeSlice, plural, singular string) error {
	if plural == "" {
		plural = "cpu_types"
	}
	if singular == "" {
		singular = "cpu_type"
	}
	writer.WriteStart("", "cpu_types", nil)
	for _, o := range structSlice.Slice() {
		XMLCpuTypeWriteOne(writer, &o, "cpu_type")
	}
	writer.WriteEnd("cpu_types")
	return nil
}

func XMLBookmarkWriteOne(writer *XMLWriter, object *Bookmark, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "bookmark"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Value(); ok {
		writer.WriteCharacter("value", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLBookmarkWriteMany(writer *XMLWriter, structSlice *BookmarkSlice, plural, singular string) error {
	if plural == "" {
		plural = "bookmarks"
	}
	if singular == "" {
		singular = "bookmark"
	}
	writer.WriteStart("", "bookmarks", nil)
	for _, o := range structSlice.Slice() {
		XMLBookmarkWriteOne(writer, &o, "bookmark")
	}
	writer.WriteEnd("bookmarks")
	return nil
}

func XMLExternalHostWriteOne(writer *XMLWriter, object *ExternalHost, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "external_host"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Address(); ok {
		writer.WriteCharacter("address", r)
	}
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.ExternalHostProvider(); ok {
		XMLExternalHostProviderWriteOne(writer, r, "external_host_provider")
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLExternalHostWriteMany(writer *XMLWriter, structSlice *ExternalHostSlice, plural, singular string) error {
	if plural == "" {
		plural = "external_hosts"
	}
	if singular == "" {
		singular = "external_host"
	}
	writer.WriteStart("", "external_hosts", nil)
	for _, o := range structSlice.Slice() {
		XMLExternalHostWriteOne(writer, &o, "external_host")
	}
	writer.WriteEnd("external_hosts")
	return nil
}

func XMLOptionWriteOne(writer *XMLWriter, object *Option, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "option"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Type(); ok {
		writer.WriteCharacter("type", r)
	}
	if r, ok := object.Value(); ok {
		writer.WriteCharacter("value", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLOptionWriteMany(writer *XMLWriter, structSlice *OptionSlice, plural, singular string) error {
	if plural == "" {
		plural = "options"
	}
	if singular == "" {
		singular = "option"
	}
	writer.WriteStart("", "options", nil)
	for _, o := range structSlice.Slice() {
		XMLOptionWriteOne(writer, &o, "option")
	}
	writer.WriteEnd("options")
	return nil
}

func XMLIdentifiedWriteOne(writer *XMLWriter, object *Identified, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "identified"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLIdentifiedWriteMany(writer *XMLWriter, structSlice *IdentifiedSlice, plural, singular string) error {
	if plural == "" {
		plural = "identifieds"
	}
	if singular == "" {
		singular = "identified"
	}
	writer.WriteStart("", "identifieds", nil)
	for _, o := range structSlice.Slice() {
		XMLIdentifiedWriteOne(writer, &o, "identified")
	}
	writer.WriteEnd("identifieds")
	return nil
}

func XMLRangeWriteOne(writer *XMLWriter, object *Range, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "range"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.From(); ok {
		writer.WriteCharacter("from", r)
	}
	if r, ok := object.To(); ok {
		writer.WriteCharacter("to", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLRangeWriteMany(writer *XMLWriter, structSlice *RangeSlice, plural, singular string) error {
	if plural == "" {
		plural = "ranges"
	}
	if singular == "" {
		singular = "range"
	}
	writer.WriteStart("", "ranges", nil)
	for _, o := range structSlice.Slice() {
		XMLRangeWriteOne(writer, &o, "range")
	}
	writer.WriteEnd("ranges")
	return nil
}

func XMLDiskSnapshotWriteOne(writer *XMLWriter, object *DiskSnapshot, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "disk_snapshot"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Active(); ok {
		writer.WriteBool("active", r)
	}
	if r, ok := object.ActualSize(); ok {
		writer.WriteInt64("actual_size", r)
	}
	if r, ok := object.Alias(); ok {
		writer.WriteCharacter("alias", r)
	}
	if r, ok := object.Bootable(); ok {
		writer.WriteBool("bootable", r)
	}
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Disk(); ok {
		XMLDiskWriteOne(writer, r, "disk")
	}
	if r, ok := object.DiskProfile(); ok {
		XMLDiskProfileWriteOne(writer, r, "disk_profile")
	}
	if r, ok := object.Format(); ok {
		XMLDiskFormatWriteOne(writer, r, "format")
	}
	if r, ok := object.ImageId(); ok {
		writer.WriteCharacter("image_id", r)
	}
	if r, ok := object.InitialSize(); ok {
		writer.WriteInt64("initial_size", r)
	}
	if r, ok := object.InstanceType(); ok {
		XMLInstanceTypeWriteOne(writer, r, "instance_type")
	}
	if r, ok := object.Interface(); ok {
		XMLDiskInterfaceWriteOne(writer, r, "interface")
	}
	if r, ok := object.LogicalName(); ok {
		writer.WriteCharacter("logical_name", r)
	}
	if r, ok := object.LunStorage(); ok {
		XMLHostStorageWriteOne(writer, r, "lun_storage")
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.OpenstackVolumeType(); ok {
		XMLOpenStackVolumeTypeWriteOne(writer, r, "openstack_volume_type")
	}
	if r, ok := object.Permissions(); ok {
		XMLPermissionWriteMany(writer, r, "permissions", "permission")
	}
	if r, ok := object.PropagateErrors(); ok {
		writer.WriteBool("propagate_errors", r)
	}
	if r, ok := object.ProvisionedSize(); ok {
		writer.WriteInt64("provisioned_size", r)
	}
	if r, ok := object.QcowVersion(); ok {
		XMLQcowVersionWriteOne(writer, r, "qcow_version")
	}
	if r, ok := object.Quota(); ok {
		XMLQuotaWriteOne(writer, r, "quota")
	}
	if r, ok := object.ReadOnly(); ok {
		writer.WriteBool("read_only", r)
	}
	if r, ok := object.Sgio(); ok {
		XMLScsiGenericIOWriteOne(writer, r, "sgio")
	}
	if r, ok := object.Shareable(); ok {
		writer.WriteBool("shareable", r)
	}
	if r, ok := object.Snapshot(); ok {
		XMLSnapshotWriteOne(writer, r, "snapshot")
	}
	if r, ok := object.Sparse(); ok {
		writer.WriteBool("sparse", r)
	}
	if r, ok := object.Statistics(); ok {
		XMLStatisticWriteMany(writer, r, "statistics", "statistic")
	}
	if r, ok := object.Status(); ok {
		XMLDiskStatusWriteOne(writer, r, "status")
	}
	if r, ok := object.StorageDomain(); ok {
		XMLStorageDomainWriteOne(writer, r, "storage_domain")
	}
	if r, ok := object.StorageDomains(); ok {
		XMLStorageDomainWriteMany(writer, r, "storage_domains", "storage_domain")
	}
	if r, ok := object.StorageType(); ok {
		XMLDiskStorageTypeWriteOne(writer, r, "storage_type")
	}
	if r, ok := object.Template(); ok {
		XMLTemplateWriteOne(writer, r, "template")
	}
	if r, ok := object.UsesScsiReservation(); ok {
		writer.WriteBool("uses_scsi_reservation", r)
	}
	if r, ok := object.Vm(); ok {
		XMLVmWriteOne(writer, r, "vm")
	}
	if r, ok := object.Vms(); ok {
		XMLVmWriteMany(writer, r, "vms", "vm")
	}
	if r, ok := object.WipeAfterDelete(); ok {
		writer.WriteBool("wipe_after_delete", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLDiskSnapshotWriteMany(writer *XMLWriter, structSlice *DiskSnapshotSlice, plural, singular string) error {
	if plural == "" {
		plural = "disk_snapshots"
	}
	if singular == "" {
		singular = "disk_snapshot"
	}
	writer.WriteStart("", "disk_snapshots", nil)
	for _, o := range structSlice.Slice() {
		XMLDiskSnapshotWriteOne(writer, &o, "disk_snapshot")
	}
	writer.WriteEnd("disk_snapshots")
	return nil
}

func XMLDiskWriteOne(writer *XMLWriter, object *Disk, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "disk"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Active(); ok {
		writer.WriteBool("active", r)
	}
	if r, ok := object.ActualSize(); ok {
		writer.WriteInt64("actual_size", r)
	}
	if r, ok := object.Alias(); ok {
		writer.WriteCharacter("alias", r)
	}
	if r, ok := object.Bootable(); ok {
		writer.WriteBool("bootable", r)
	}
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.DiskProfile(); ok {
		XMLDiskProfileWriteOne(writer, r, "disk_profile")
	}
	if r, ok := object.Format(); ok {
		XMLDiskFormatWriteOne(writer, r, "format")
	}
	if r, ok := object.ImageId(); ok {
		writer.WriteCharacter("image_id", r)
	}
	if r, ok := object.InitialSize(); ok {
		writer.WriteInt64("initial_size", r)
	}
	if r, ok := object.InstanceType(); ok {
		XMLInstanceTypeWriteOne(writer, r, "instance_type")
	}
	if r, ok := object.Interface(); ok {
		XMLDiskInterfaceWriteOne(writer, r, "interface")
	}
	if r, ok := object.LogicalName(); ok {
		writer.WriteCharacter("logical_name", r)
	}
	if r, ok := object.LunStorage(); ok {
		XMLHostStorageWriteOne(writer, r, "lun_storage")
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.OpenstackVolumeType(); ok {
		XMLOpenStackVolumeTypeWriteOne(writer, r, "openstack_volume_type")
	}
	if r, ok := object.Permissions(); ok {
		XMLPermissionWriteMany(writer, r, "permissions", "permission")
	}
	if r, ok := object.PropagateErrors(); ok {
		writer.WriteBool("propagate_errors", r)
	}
	if r, ok := object.ProvisionedSize(); ok {
		writer.WriteInt64("provisioned_size", r)
	}
	if r, ok := object.QcowVersion(); ok {
		XMLQcowVersionWriteOne(writer, r, "qcow_version")
	}
	if r, ok := object.Quota(); ok {
		XMLQuotaWriteOne(writer, r, "quota")
	}
	if r, ok := object.ReadOnly(); ok {
		writer.WriteBool("read_only", r)
	}
	if r, ok := object.Sgio(); ok {
		XMLScsiGenericIOWriteOne(writer, r, "sgio")
	}
	if r, ok := object.Shareable(); ok {
		writer.WriteBool("shareable", r)
	}
	if r, ok := object.Snapshot(); ok {
		XMLSnapshotWriteOne(writer, r, "snapshot")
	}
	if r, ok := object.Sparse(); ok {
		writer.WriteBool("sparse", r)
	}
	if r, ok := object.Statistics(); ok {
		XMLStatisticWriteMany(writer, r, "statistics", "statistic")
	}
	if r, ok := object.Status(); ok {
		XMLDiskStatusWriteOne(writer, r, "status")
	}
	if r, ok := object.StorageDomain(); ok {
		XMLStorageDomainWriteOne(writer, r, "storage_domain")
	}
	if r, ok := object.StorageDomains(); ok {
		XMLStorageDomainWriteMany(writer, r, "storage_domains", "storage_domain")
	}
	if r, ok := object.StorageType(); ok {
		XMLDiskStorageTypeWriteOne(writer, r, "storage_type")
	}
	if r, ok := object.Template(); ok {
		XMLTemplateWriteOne(writer, r, "template")
	}
	if r, ok := object.UsesScsiReservation(); ok {
		writer.WriteBool("uses_scsi_reservation", r)
	}
	if r, ok := object.Vm(); ok {
		XMLVmWriteOne(writer, r, "vm")
	}
	if r, ok := object.Vms(); ok {
		XMLVmWriteMany(writer, r, "vms", "vm")
	}
	if r, ok := object.WipeAfterDelete(); ok {
		writer.WriteBool("wipe_after_delete", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLDiskWriteMany(writer *XMLWriter, structSlice *DiskSlice, plural, singular string) error {
	if plural == "" {
		plural = "disks"
	}
	if singular == "" {
		singular = "disk"
	}
	writer.WriteStart("", "disks", nil)
	for _, o := range structSlice.Slice() {
		XMLDiskWriteOne(writer, &o, "disk")
	}
	writer.WriteEnd("disks")
	return nil
}

func XMLVirtualNumaNodeWriteOne(writer *XMLWriter, object *VirtualNumaNode, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "vm_numa_node"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Cpu(); ok {
		XMLCpuWriteOne(writer, r, "cpu")
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Host(); ok {
		XMLHostWriteOne(writer, r, "host")
	}
	if r, ok := object.Index(); ok {
		writer.WriteInt64("index", r)
	}
	if r, ok := object.Memory(); ok {
		writer.WriteInt64("memory", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.NodeDistance(); ok {
		writer.WriteCharacter("node_distance", r)
	}
	if r, ok := object.NumaNodePins(); ok {
		XMLNumaNodePinWriteMany(writer, r, "numa_node_pins", "numa_node_pin")
	}
	if r, ok := object.Statistics(); ok {
		XMLStatisticWriteMany(writer, r, "statistics", "statistic")
	}
	if r, ok := object.Vm(); ok {
		XMLVmWriteOne(writer, r, "vm")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLVirtualNumaNodeWriteMany(writer *XMLWriter, structSlice *VirtualNumaNodeSlice, plural, singular string) error {
	if plural == "" {
		plural = "vm_numa_nodes"
	}
	if singular == "" {
		singular = "vm_numa_node"
	}
	writer.WriteStart("", "vm_numa_nodes", nil)
	for _, o := range structSlice.Slice() {
		XMLVirtualNumaNodeWriteOne(writer, &o, "vm_numa_node")
	}
	writer.WriteEnd("vm_numa_nodes")
	return nil
}

func XMLHardwareInformationWriteOne(writer *XMLWriter, object *HardwareInformation, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "hardware_information"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Family(); ok {
		writer.WriteCharacter("family", r)
	}
	if r, ok := object.Manufacturer(); ok {
		writer.WriteCharacter("manufacturer", r)
	}
	if r, ok := object.ProductName(); ok {
		writer.WriteCharacter("product_name", r)
	}
	if r, ok := object.SerialNumber(); ok {
		writer.WriteCharacter("serial_number", r)
	}
	if r, ok := object.SupportedRngSources(); ok {
		XMLRngSourceWriteMany(writer, r, "supported_rng_sources", "rng_source")
	}
	if r, ok := object.Uuid(); ok {
		writer.WriteCharacter("uuid", r)
	}
	if r, ok := object.Version(); ok {
		writer.WriteCharacter("version", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLHardwareInformationWriteMany(writer *XMLWriter, structSlice *HardwareInformationSlice, plural, singular string) error {
	if plural == "" {
		plural = "hardware_informations"
	}
	if singular == "" {
		singular = "hardware_information"
	}
	writer.WriteStart("", "hardware_informations", nil)
	for _, o := range structSlice.Slice() {
		XMLHardwareInformationWriteOne(writer, &o, "hardware_information")
	}
	writer.WriteEnd("hardware_informations")
	return nil
}

func XMLOpenStackNetworkWriteOne(writer *XMLWriter, object *OpenStackNetwork, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "openstack_network"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.OpenstackNetworkProvider(); ok {
		XMLOpenStackNetworkProviderWriteOne(writer, r, "openstack_network_provider")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLOpenStackNetworkWriteMany(writer *XMLWriter, structSlice *OpenStackNetworkSlice, plural, singular string) error {
	if plural == "" {
		plural = "openstack_networks"
	}
	if singular == "" {
		singular = "openstack_network"
	}
	writer.WriteStart("", "openstack_networks", nil)
	for _, o := range structSlice.Slice() {
		XMLOpenStackNetworkWriteOne(writer, &o, "openstack_network")
	}
	writer.WriteEnd("openstack_networks")
	return nil
}

func XMLApiWriteOne(writer *XMLWriter, object *Api, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "api"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.ProductInfo(); ok {
		XMLProductInfoWriteOne(writer, r, "product_info")
	}
	if r, ok := object.SpecialObjects(); ok {
		XMLSpecialObjectsWriteOne(writer, r, "special_objects")
	}
	if r, ok := object.Summary(); ok {
		XMLApiSummaryWriteOne(writer, r, "summary")
	}
	if r, ok := object.Time(); ok {
		writer.WriteDate("time", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLApiWriteMany(writer *XMLWriter, structSlice *ApiSlice, plural, singular string) error {
	if plural == "" {
		plural = "apis"
	}
	if singular == "" {
		singular = "api"
	}
	writer.WriteStart("", "apis", nil)
	for _, o := range structSlice.Slice() {
		XMLApiWriteOne(writer, &o, "api")
	}
	writer.WriteEnd("apis")
	return nil
}

func XMLQosWriteOne(writer *XMLWriter, object *Qos, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "qos"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.CpuLimit(); ok {
		writer.WriteInt64("cpu_limit", r)
	}
	if r, ok := object.DataCenter(); ok {
		XMLDataCenterWriteOne(writer, r, "data_center")
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.InboundAverage(); ok {
		writer.WriteInt64("inbound_average", r)
	}
	if r, ok := object.InboundBurst(); ok {
		writer.WriteInt64("inbound_burst", r)
	}
	if r, ok := object.InboundPeak(); ok {
		writer.WriteInt64("inbound_peak", r)
	}
	if r, ok := object.MaxIops(); ok {
		writer.WriteInt64("max_iops", r)
	}
	if r, ok := object.MaxReadIops(); ok {
		writer.WriteInt64("max_read_iops", r)
	}
	if r, ok := object.MaxReadThroughput(); ok {
		writer.WriteInt64("max_read_throughput", r)
	}
	if r, ok := object.MaxThroughput(); ok {
		writer.WriteInt64("max_throughput", r)
	}
	if r, ok := object.MaxWriteIops(); ok {
		writer.WriteInt64("max_write_iops", r)
	}
	if r, ok := object.MaxWriteThroughput(); ok {
		writer.WriteInt64("max_write_throughput", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.OutboundAverage(); ok {
		writer.WriteInt64("outbound_average", r)
	}
	if r, ok := object.OutboundAverageLinkshare(); ok {
		writer.WriteInt64("outbound_average_linkshare", r)
	}
	if r, ok := object.OutboundAverageRealtime(); ok {
		writer.WriteInt64("outbound_average_realtime", r)
	}
	if r, ok := object.OutboundAverageUpperlimit(); ok {
		writer.WriteInt64("outbound_average_upperlimit", r)
	}
	if r, ok := object.OutboundBurst(); ok {
		writer.WriteInt64("outbound_burst", r)
	}
	if r, ok := object.OutboundPeak(); ok {
		writer.WriteInt64("outbound_peak", r)
	}
	if r, ok := object.Type(); ok {
		XMLQosTypeWriteOne(writer, r, "type")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLQosWriteMany(writer *XMLWriter, structSlice *QosSlice, plural, singular string) error {
	if plural == "" {
		plural = "qoss"
	}
	if singular == "" {
		singular = "qos"
	}
	writer.WriteStart("", "qoss", nil)
	for _, o := range structSlice.Slice() {
		XMLQosWriteOne(writer, &o, "qos")
	}
	writer.WriteEnd("qoss")
	return nil
}

func XMLBootMenuWriteOne(writer *XMLWriter, object *BootMenu, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "boot_menu"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Enabled(); ok {
		writer.WriteBool("enabled", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLBootMenuWriteMany(writer *XMLWriter, structSlice *BootMenuSlice, plural, singular string) error {
	if plural == "" {
		plural = "boot_menus"
	}
	if singular == "" {
		singular = "boot_menu"
	}
	writer.WriteStart("", "boot_menus", nil)
	for _, o := range structSlice.Slice() {
		XMLBootMenuWriteOne(writer, &o, "boot_menu")
	}
	writer.WriteEnd("boot_menus")
	return nil
}

func XMLExternalDiscoveredHostWriteOne(writer *XMLWriter, object *ExternalDiscoveredHost, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "external_discovered_host"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.ExternalHostProvider(); ok {
		XMLExternalHostProviderWriteOne(writer, r, "external_host_provider")
	}
	if r, ok := object.Ip(); ok {
		writer.WriteCharacter("ip", r)
	}
	if r, ok := object.LastReport(); ok {
		writer.WriteCharacter("last_report", r)
	}
	if r, ok := object.Mac(); ok {
		writer.WriteCharacter("mac", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.SubnetName(); ok {
		writer.WriteCharacter("subnet_name", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLExternalDiscoveredHostWriteMany(writer *XMLWriter, structSlice *ExternalDiscoveredHostSlice, plural, singular string) error {
	if plural == "" {
		plural = "external_discovered_hosts"
	}
	if singular == "" {
		singular = "external_discovered_host"
	}
	writer.WriteStart("", "external_discovered_hosts", nil)
	for _, o := range structSlice.Slice() {
		XMLExternalDiscoveredHostWriteOne(writer, &o, "external_discovered_host")
	}
	writer.WriteEnd("external_discovered_hosts")
	return nil
}

func XMLImageTransferWriteOne(writer *XMLWriter, object *ImageTransfer, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "image_transfer"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Direction(); ok {
		XMLImageTransferDirectionWriteOne(writer, r, "direction")
	}
	if r, ok := object.Host(); ok {
		XMLHostWriteOne(writer, r, "host")
	}
	if r, ok := object.Image(); ok {
		XMLImageWriteOne(writer, r, "image")
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Phase(); ok {
		XMLImageTransferPhaseWriteOne(writer, r, "phase")
	}
	if r, ok := object.ProxyUrl(); ok {
		writer.WriteCharacter("proxy_url", r)
	}
	if r, ok := object.SignedTicket(); ok {
		writer.WriteCharacter("signed_ticket", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLImageTransferWriteMany(writer *XMLWriter, structSlice *ImageTransferSlice, plural, singular string) error {
	if plural == "" {
		plural = "image_transfers"
	}
	if singular == "" {
		singular = "image_transfer"
	}
	writer.WriteStart("", "image_transfers", nil)
	for _, o := range structSlice.Slice() {
		XMLImageTransferWriteOne(writer, &o, "image_transfer")
	}
	writer.WriteEnd("image_transfers")
	return nil
}

func XMLFencingPolicyWriteOne(writer *XMLWriter, object *FencingPolicy, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "fencing_policy"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Enabled(); ok {
		writer.WriteBool("enabled", r)
	}
	if r, ok := object.SkipIfConnectivityBroken(); ok {
		XMLSkipIfConnectivityBrokenWriteOne(writer, r, "skip_if_connectivity_broken")
	}
	if r, ok := object.SkipIfGlusterBricksUp(); ok {
		writer.WriteBool("skip_if_gluster_bricks_up", r)
	}
	if r, ok := object.SkipIfGlusterQuorumNotMet(); ok {
		writer.WriteBool("skip_if_gluster_quorum_not_met", r)
	}
	if r, ok := object.SkipIfSdActive(); ok {
		XMLSkipIfSdActiveWriteOne(writer, r, "skip_if_sd_active")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLFencingPolicyWriteMany(writer *XMLWriter, structSlice *FencingPolicySlice, plural, singular string) error {
	if plural == "" {
		plural = "fencing_policies"
	}
	if singular == "" {
		singular = "fencing_policy"
	}
	writer.WriteStart("", "fencing_policies", nil)
	for _, o := range structSlice.Slice() {
		XMLFencingPolicyWriteOne(writer, &o, "fencing_policy")
	}
	writer.WriteEnd("fencing_policies")
	return nil
}

func XMLOpenStackVolumeProviderWriteOne(writer *XMLWriter, object *OpenStackVolumeProvider, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "openstack_volume_provider"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.AuthenticationKeys(); ok {
		XMLOpenstackVolumeAuthenticationKeyWriteMany(writer, r, "authentication_keys", "openstack_volume_authentication_key")
	}
	if r, ok := object.AuthenticationUrl(); ok {
		writer.WriteCharacter("authentication_url", r)
	}
	if r, ok := object.Certificates(); ok {
		XMLCertificateWriteMany(writer, r, "certificates", "certificate")
	}
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.DataCenter(); ok {
		XMLDataCenterWriteOne(writer, r, "data_center")
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Password(); ok {
		writer.WriteCharacter("password", r)
	}
	if r, ok := object.Properties(); ok {
		XMLPropertyWriteMany(writer, r, "properties", "property")
	}
	if r, ok := object.RequiresAuthentication(); ok {
		writer.WriteBool("requires_authentication", r)
	}
	if r, ok := object.TenantName(); ok {
		writer.WriteCharacter("tenant_name", r)
	}
	if r, ok := object.Url(); ok {
		writer.WriteCharacter("url", r)
	}
	if r, ok := object.Username(); ok {
		writer.WriteCharacter("username", r)
	}
	if r, ok := object.VolumeTypes(); ok {
		XMLOpenStackVolumeTypeWriteMany(writer, r, "volume_types", "open_stack_volume_type")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLOpenStackVolumeProviderWriteMany(writer *XMLWriter, structSlice *OpenStackVolumeProviderSlice, plural, singular string) error {
	if plural == "" {
		plural = "openstack_volume_providers"
	}
	if singular == "" {
		singular = "openstack_volume_provider"
	}
	writer.WriteStart("", "openstack_volume_providers", nil)
	for _, o := range structSlice.Slice() {
		XMLOpenStackVolumeProviderWriteOne(writer, &o, "openstack_volume_provider")
	}
	writer.WriteEnd("openstack_volume_providers")
	return nil
}

func XMLFileWriteOne(writer *XMLWriter, object *File, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "file"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Content(); ok {
		writer.WriteCharacter("content", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.StorageDomain(); ok {
		XMLStorageDomainWriteOne(writer, r, "storage_domain")
	}
	if r, ok := object.Type(); ok {
		writer.WriteCharacter("type", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLFileWriteMany(writer *XMLWriter, structSlice *FileSlice, plural, singular string) error {
	if plural == "" {
		plural = "files"
	}
	if singular == "" {
		singular = "file"
	}
	writer.WriteStart("", "files", nil)
	for _, o := range structSlice.Slice() {
		XMLFileWriteOne(writer, &o, "file")
	}
	writer.WriteEnd("files")
	return nil
}

func XMLReportedDeviceWriteOne(writer *XMLWriter, object *ReportedDevice, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "reported_device"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Ips(); ok {
		XMLIpWriteMany(writer, r, "ips", "ip")
	}
	if r, ok := object.Mac(); ok {
		XMLMacWriteOne(writer, r, "mac")
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Type(); ok {
		XMLReportedDeviceTypeWriteOne(writer, r, "type")
	}
	if r, ok := object.Vm(); ok {
		XMLVmWriteOne(writer, r, "vm")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLReportedDeviceWriteMany(writer *XMLWriter, structSlice *ReportedDeviceSlice, plural, singular string) error {
	if plural == "" {
		plural = "reported_devices"
	}
	if singular == "" {
		singular = "reported_device"
	}
	writer.WriteStart("", "reported_devices", nil)
	for _, o := range structSlice.Slice() {
		XMLReportedDeviceWriteOne(writer, &o, "reported_device")
	}
	writer.WriteEnd("reported_devices")
	return nil
}

func XMLSchedulingPolicyUnitWriteOne(writer *XMLWriter, object *SchedulingPolicyUnit, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "scheduling_policy_unit"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Enabled(); ok {
		writer.WriteBool("enabled", r)
	}
	if r, ok := object.Internal(); ok {
		writer.WriteBool("internal", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Properties(); ok {
		XMLPropertyWriteMany(writer, r, "properties", "property")
	}
	if r, ok := object.Type(); ok {
		XMLPolicyUnitTypeWriteOne(writer, r, "type")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLSchedulingPolicyUnitWriteMany(writer *XMLWriter, structSlice *SchedulingPolicyUnitSlice, plural, singular string) error {
	if plural == "" {
		plural = "scheduling_policy_units"
	}
	if singular == "" {
		singular = "scheduling_policy_unit"
	}
	writer.WriteStart("", "scheduling_policy_units", nil)
	for _, o := range structSlice.Slice() {
		XMLSchedulingPolicyUnitWriteOne(writer, &o, "scheduling_policy_unit")
	}
	writer.WriteEnd("scheduling_policy_units")
	return nil
}

func XMLProxyTicketWriteOne(writer *XMLWriter, object *ProxyTicket, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "proxy_ticket"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Value(); ok {
		writer.WriteCharacter("value", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLProxyTicketWriteMany(writer *XMLWriter, structSlice *ProxyTicketSlice, plural, singular string) error {
	if plural == "" {
		plural = "proxy_tickets"
	}
	if singular == "" {
		singular = "proxy_ticket"
	}
	writer.WriteStart("", "proxy_tickets", nil)
	for _, o := range structSlice.Slice() {
		XMLProxyTicketWriteOne(writer, &o, "proxy_ticket")
	}
	writer.WriteEnd("proxy_tickets")
	return nil
}

func XMLMigrationBandwidthWriteOne(writer *XMLWriter, object *MigrationBandwidth, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "migration_bandwidth"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.AssignmentMethod(); ok {
		XMLMigrationBandwidthAssignmentMethodWriteOne(writer, r, "assignment_method")
	}
	if r, ok := object.CustomValue(); ok {
		writer.WriteInt64("custom_value", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLMigrationBandwidthWriteMany(writer *XMLWriter, structSlice *MigrationBandwidthSlice, plural, singular string) error {
	if plural == "" {
		plural = "migration_bandwidths"
	}
	if singular == "" {
		singular = "migration_bandwidth"
	}
	writer.WriteStart("", "migration_bandwidths", nil)
	for _, o := range structSlice.Slice() {
		XMLMigrationBandwidthWriteOne(writer, &o, "migration_bandwidth")
	}
	writer.WriteEnd("migration_bandwidths")
	return nil
}

func XMLInitializationWriteOne(writer *XMLWriter, object *Initialization, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "initialization"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.ActiveDirectoryOu(); ok {
		writer.WriteCharacter("active_directory_ou", r)
	}
	if r, ok := object.AuthorizedSshKeys(); ok {
		writer.WriteCharacter("authorized_ssh_keys", r)
	}
	if r, ok := object.CloudInit(); ok {
		XMLCloudInitWriteOne(writer, r, "cloud_init")
	}
	if r, ok := object.Configuration(); ok {
		XMLConfigurationWriteOne(writer, r, "configuration")
	}
	if r, ok := object.CustomScript(); ok {
		writer.WriteCharacter("custom_script", r)
	}
	if r, ok := object.DnsSearch(); ok {
		writer.WriteCharacter("dns_search", r)
	}
	if r, ok := object.DnsServers(); ok {
		writer.WriteCharacter("dns_servers", r)
	}
	if r, ok := object.Domain(); ok {
		writer.WriteCharacter("domain", r)
	}
	if r, ok := object.HostName(); ok {
		writer.WriteCharacter("host_name", r)
	}
	if r, ok := object.InputLocale(); ok {
		writer.WriteCharacter("input_locale", r)
	}
	if r, ok := object.NicConfigurations(); ok {
		XMLNicConfigurationWriteMany(writer, r, "nic_configurations", "nic_configuration")
	}
	if r, ok := object.OrgName(); ok {
		writer.WriteCharacter("org_name", r)
	}
	if r, ok := object.RegenerateIds(); ok {
		writer.WriteBool("regenerate_ids", r)
	}
	if r, ok := object.RegenerateSshKeys(); ok {
		writer.WriteBool("regenerate_ssh_keys", r)
	}
	if r, ok := object.RootPassword(); ok {
		writer.WriteCharacter("root_password", r)
	}
	if r, ok := object.SystemLocale(); ok {
		writer.WriteCharacter("system_locale", r)
	}
	if r, ok := object.Timezone(); ok {
		writer.WriteCharacter("timezone", r)
	}
	if r, ok := object.UiLanguage(); ok {
		writer.WriteCharacter("ui_language", r)
	}
	if r, ok := object.UserLocale(); ok {
		writer.WriteCharacter("user_locale", r)
	}
	if r, ok := object.UserName(); ok {
		writer.WriteCharacter("user_name", r)
	}
	if r, ok := object.WindowsLicenseKey(); ok {
		writer.WriteCharacter("windows_license_key", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLInitializationWriteMany(writer *XMLWriter, structSlice *InitializationSlice, plural, singular string) error {
	if plural == "" {
		plural = "initializations"
	}
	if singular == "" {
		singular = "initialization"
	}
	writer.WriteStart("", "initializations", nil)
	for _, o := range structSlice.Slice() {
		XMLInitializationWriteOne(writer, &o, "initialization")
	}
	writer.WriteEnd("initializations")
	return nil
}

func XMLTransparentHugePagesWriteOne(writer *XMLWriter, object *TransparentHugePages, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "transparent_hugepages"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Enabled(); ok {
		writer.WriteBool("enabled", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLTransparentHugePagesWriteMany(writer *XMLWriter, structSlice *TransparentHugePagesSlice, plural, singular string) error {
	if plural == "" {
		plural = "transparent_huge_pagess"
	}
	if singular == "" {
		singular = "transparent_hugepages"
	}
	writer.WriteStart("", "transparent_huge_pagess", nil)
	for _, o := range structSlice.Slice() {
		XMLTransparentHugePagesWriteOne(writer, &o, "transparent_hugepages")
	}
	writer.WriteEnd("transparent_huge_pagess")
	return nil
}

func XMLVolumeGroupWriteOne(writer *XMLWriter, object *VolumeGroup, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "volume_group"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.LogicalUnits(); ok {
		XMLLogicalUnitWriteMany(writer, r, "logical_units", "logical_unit")
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLVolumeGroupWriteMany(writer *XMLWriter, structSlice *VolumeGroupSlice, plural, singular string) error {
	if plural == "" {
		plural = "volume_groups"
	}
	if singular == "" {
		singular = "volume_group"
	}
	writer.WriteStart("", "volume_groups", nil)
	for _, o := range structSlice.Slice() {
		XMLVolumeGroupWriteOne(writer, &o, "volume_group")
	}
	writer.WriteEnd("volume_groups")
	return nil
}

func XMLNicWriteOne(writer *XMLWriter, object *Nic, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "nic"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.BootProtocol(); ok {
		XMLBootProtocolWriteOne(writer, r, "boot_protocol")
	}
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.InstanceType(); ok {
		XMLInstanceTypeWriteOne(writer, r, "instance_type")
	}
	if r, ok := object.Interface(); ok {
		XMLNicInterfaceWriteOne(writer, r, "interface")
	}
	if r, ok := object.Linked(); ok {
		writer.WriteBool("linked", r)
	}
	if r, ok := object.Mac(); ok {
		XMLMacWriteOne(writer, r, "mac")
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Network(); ok {
		XMLNetworkWriteOne(writer, r, "network")
	}
	if r, ok := object.NetworkAttachments(); ok {
		XMLNetworkAttachmentWriteMany(writer, r, "network_attachments", "network_attachment")
	}
	if r, ok := object.NetworkFilterParameters(); ok {
		XMLNetworkFilterParameterWriteMany(writer, r, "network_filter_parameters", "network_filter_parameter")
	}
	if r, ok := object.NetworkLabels(); ok {
		XMLNetworkLabelWriteMany(writer, r, "network_labels", "network_label")
	}
	if r, ok := object.OnBoot(); ok {
		writer.WriteBool("on_boot", r)
	}
	if r, ok := object.Plugged(); ok {
		writer.WriteBool("plugged", r)
	}
	if r, ok := object.ReportedDevices(); ok {
		XMLReportedDeviceWriteMany(writer, r, "reported_devices", "reported_device")
	}
	if r, ok := object.Statistics(); ok {
		XMLStatisticWriteMany(writer, r, "statistics", "statistic")
	}
	if r, ok := object.Template(); ok {
		XMLTemplateWriteOne(writer, r, "template")
	}
	if r, ok := object.VirtualFunctionAllowedLabels(); ok {
		XMLNetworkLabelWriteMany(writer, r, "virtual_function_allowed_labels", "network_label")
	}
	if r, ok := object.VirtualFunctionAllowedNetworks(); ok {
		XMLNetworkWriteMany(writer, r, "virtual_function_allowed_networks", "network")
	}
	if r, ok := object.Vm(); ok {
		XMLVmWriteOne(writer, r, "vm")
	}
	if r, ok := object.Vms(); ok {
		XMLVmWriteMany(writer, r, "vms", "vm")
	}
	if r, ok := object.VnicProfile(); ok {
		XMLVnicProfileWriteOne(writer, r, "vnic_profile")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLNicWriteMany(writer *XMLWriter, structSlice *NicSlice, plural, singular string) error {
	if plural == "" {
		plural = "nics"
	}
	if singular == "" {
		singular = "nic"
	}
	writer.WriteStart("", "nics", nil)
	for _, o := range structSlice.Slice() {
		XMLNicWriteOne(writer, &o, "nic")
	}
	writer.WriteEnd("nics")
	return nil
}

func XMLNetworkFilterParameterWriteOne(writer *XMLWriter, object *NetworkFilterParameter, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "network_filter_parameter"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Nic(); ok {
		XMLNicWriteOne(writer, r, "nic")
	}
	if r, ok := object.Value(); ok {
		writer.WriteCharacter("value", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLNetworkFilterParameterWriteMany(writer *XMLWriter, structSlice *NetworkFilterParameterSlice, plural, singular string) error {
	if plural == "" {
		plural = "network_filter_parameters"
	}
	if singular == "" {
		singular = "network_filter_parameter"
	}
	writer.WriteStart("", "network_filter_parameters", nil)
	for _, o := range structSlice.Slice() {
		XMLNetworkFilterParameterWriteOne(writer, &o, "network_filter_parameter")
	}
	writer.WriteEnd("network_filter_parameters")
	return nil
}

func XMLCdromWriteOne(writer *XMLWriter, object *Cdrom, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "cdrom"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.File(); ok {
		XMLFileWriteOne(writer, r, "file")
	}
	if r, ok := object.InstanceType(); ok {
		XMLInstanceTypeWriteOne(writer, r, "instance_type")
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Template(); ok {
		XMLTemplateWriteOne(writer, r, "template")
	}
	if r, ok := object.Vm(); ok {
		XMLVmWriteOne(writer, r, "vm")
	}
	if r, ok := object.Vms(); ok {
		XMLVmWriteMany(writer, r, "vms", "vm")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLCdromWriteMany(writer *XMLWriter, structSlice *CdromSlice, plural, singular string) error {
	if plural == "" {
		plural = "cdroms"
	}
	if singular == "" {
		singular = "cdrom"
	}
	writer.WriteStart("", "cdroms", nil)
	for _, o := range structSlice.Slice() {
		XMLCdromWriteOne(writer, &o, "cdrom")
	}
	writer.WriteEnd("cdroms")
	return nil
}

func XMLTagWriteOne(writer *XMLWriter, object *Tag, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "tag"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Group(); ok {
		XMLGroupWriteOne(writer, r, "group")
	}
	if r, ok := object.Host(); ok {
		XMLHostWriteOne(writer, r, "host")
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Parent(); ok {
		XMLTagWriteOne(writer, r, "parent")
	}
	if r, ok := object.Template(); ok {
		XMLTemplateWriteOne(writer, r, "template")
	}
	if r, ok := object.User(); ok {
		XMLUserWriteOne(writer, r, "user")
	}
	if r, ok := object.Vm(); ok {
		XMLVmWriteOne(writer, r, "vm")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLTagWriteMany(writer *XMLWriter, structSlice *TagSlice, plural, singular string) error {
	if plural == "" {
		plural = "tags"
	}
	if singular == "" {
		singular = "tag"
	}
	writer.WriteStart("", "tags", nil)
	for _, o := range structSlice.Slice() {
		XMLTagWriteOne(writer, &o, "tag")
	}
	writer.WriteEnd("tags")
	return nil
}

func XMLSessionWriteOne(writer *XMLWriter, object *Session, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "session"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.ConsoleUser(); ok {
		writer.WriteBool("console_user", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Ip(); ok {
		XMLIpWriteOne(writer, r, "ip")
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Protocol(); ok {
		writer.WriteCharacter("protocol", r)
	}
	if r, ok := object.User(); ok {
		XMLUserWriteOne(writer, r, "user")
	}
	if r, ok := object.Vm(); ok {
		XMLVmWriteOne(writer, r, "vm")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLSessionWriteMany(writer *XMLWriter, structSlice *SessionSlice, plural, singular string) error {
	if plural == "" {
		plural = "sessions"
	}
	if singular == "" {
		singular = "session"
	}
	writer.WriteStart("", "sessions", nil)
	for _, o := range structSlice.Slice() {
		XMLSessionWriteOne(writer, &o, "session")
	}
	writer.WriteEnd("sessions")
	return nil
}

func XMLBalanceWriteOne(writer *XMLWriter, object *Balance, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "balance"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.SchedulingPolicy(); ok {
		XMLSchedulingPolicyWriteOne(writer, r, "scheduling_policy")
	}
	if r, ok := object.SchedulingPolicyUnit(); ok {
		XMLSchedulingPolicyUnitWriteOne(writer, r, "scheduling_policy_unit")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLBalanceWriteMany(writer *XMLWriter, structSlice *BalanceSlice, plural, singular string) error {
	if plural == "" {
		plural = "balances"
	}
	if singular == "" {
		singular = "balance"
	}
	writer.WriteStart("", "balances", nil)
	for _, o := range structSlice.Slice() {
		XMLBalanceWriteOne(writer, &o, "balance")
	}
	writer.WriteEnd("balances")
	return nil
}

func XMLExternalVmImportWriteOne(writer *XMLWriter, object *ExternalVmImport, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "external_vm_import"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Cluster(); ok {
		XMLClusterWriteOne(writer, r, "cluster")
	}
	if r, ok := object.CpuProfile(); ok {
		XMLCpuProfileWriteOne(writer, r, "cpu_profile")
	}
	if r, ok := object.DriversIso(); ok {
		XMLFileWriteOne(writer, r, "drivers_iso")
	}
	if r, ok := object.Host(); ok {
		XMLHostWriteOne(writer, r, "host")
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Password(); ok {
		writer.WriteCharacter("password", r)
	}
	if r, ok := object.Provider(); ok {
		XMLExternalVmProviderTypeWriteOne(writer, r, "provider")
	}
	if r, ok := object.Quota(); ok {
		XMLQuotaWriteOne(writer, r, "quota")
	}
	if r, ok := object.Sparse(); ok {
		writer.WriteBool("sparse", r)
	}
	if r, ok := object.StorageDomain(); ok {
		XMLStorageDomainWriteOne(writer, r, "storage_domain")
	}
	if r, ok := object.Url(); ok {
		writer.WriteCharacter("url", r)
	}
	if r, ok := object.Username(); ok {
		writer.WriteCharacter("username", r)
	}
	if r, ok := object.Vm(); ok {
		XMLVmWriteOne(writer, r, "vm")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLExternalVmImportWriteMany(writer *XMLWriter, structSlice *ExternalVmImportSlice, plural, singular string) error {
	if plural == "" {
		plural = "external_vm_imports"
	}
	if singular == "" {
		singular = "external_vm_import"
	}
	writer.WriteStart("", "external_vm_imports", nil)
	for _, o := range structSlice.Slice() {
		XMLExternalVmImportWriteOne(writer, &o, "external_vm_import")
	}
	writer.WriteEnd("external_vm_imports")
	return nil
}

func XMLOpenStackVolumeTypeWriteOne(writer *XMLWriter, object *OpenStackVolumeType, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "open_stack_volume_type"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.OpenstackVolumeProvider(); ok {
		XMLOpenStackVolumeProviderWriteOne(writer, r, "openstack_volume_provider")
	}
	if r, ok := object.Properties(); ok {
		XMLPropertyWriteMany(writer, r, "properties", "property")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLOpenStackVolumeTypeWriteMany(writer *XMLWriter, structSlice *OpenStackVolumeTypeSlice, plural, singular string) error {
	if plural == "" {
		plural = "open_stack_volume_types"
	}
	if singular == "" {
		singular = "open_stack_volume_type"
	}
	writer.WriteStart("", "open_stack_volume_types", nil)
	for _, o := range structSlice.Slice() {
		XMLOpenStackVolumeTypeWriteOne(writer, &o, "open_stack_volume_type")
	}
	writer.WriteEnd("open_stack_volume_types")
	return nil
}

func XMLMacWriteOne(writer *XMLWriter, object *Mac, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "mac"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Address(); ok {
		writer.WriteCharacter("address", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLMacWriteMany(writer *XMLWriter, structSlice *MacSlice, plural, singular string) error {
	if plural == "" {
		plural = "macs"
	}
	if singular == "" {
		singular = "mac"
	}
	writer.WriteStart("", "macs", nil)
	for _, o := range structSlice.Slice() {
		XMLMacWriteOne(writer, &o, "mac")
	}
	writer.WriteEnd("macs")
	return nil
}

func XMLGlusterBrickMemoryInfoWriteOne(writer *XMLWriter, object *GlusterBrickMemoryInfo, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "brick_memoryinfo"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.MemoryPools(); ok {
		XMLGlusterMemoryPoolWriteMany(writer, r, "memory_pools", "memory_pool")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLGlusterBrickMemoryInfoWriteMany(writer *XMLWriter, structSlice *GlusterBrickMemoryInfoSlice, plural, singular string) error {
	if plural == "" {
		plural = "gluster_brick_memory_infos"
	}
	if singular == "" {
		singular = "brick_memoryinfo"
	}
	writer.WriteStart("", "gluster_brick_memory_infos", nil)
	for _, o := range structSlice.Slice() {
		XMLGlusterBrickMemoryInfoWriteOne(writer, &o, "brick_memoryinfo")
	}
	writer.WriteEnd("gluster_brick_memory_infos")
	return nil
}

func XMLNumaNodeWriteOne(writer *XMLWriter, object *NumaNode, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "host_numa_node"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Cpu(); ok {
		XMLCpuWriteOne(writer, r, "cpu")
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Host(); ok {
		XMLHostWriteOne(writer, r, "host")
	}
	if r, ok := object.Index(); ok {
		writer.WriteInt64("index", r)
	}
	if r, ok := object.Memory(); ok {
		writer.WriteInt64("memory", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.NodeDistance(); ok {
		writer.WriteCharacter("node_distance", r)
	}
	if r, ok := object.Statistics(); ok {
		XMLStatisticWriteMany(writer, r, "statistics", "statistic")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLNumaNodeWriteMany(writer *XMLWriter, structSlice *NumaNodeSlice, plural, singular string) error {
	if plural == "" {
		plural = "host_numa_nodes"
	}
	if singular == "" {
		singular = "host_numa_node"
	}
	writer.WriteStart("", "host_numa_nodes", nil)
	for _, o := range structSlice.Slice() {
		XMLNumaNodeWriteOne(writer, &o, "host_numa_node")
	}
	writer.WriteEnd("host_numa_nodes")
	return nil
}

func XMLOpenStackImageProviderWriteOne(writer *XMLWriter, object *OpenStackImageProvider, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "openstack_image_provider"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.AuthenticationUrl(); ok {
		writer.WriteCharacter("authentication_url", r)
	}
	if r, ok := object.Certificates(); ok {
		XMLCertificateWriteMany(writer, r, "certificates", "certificate")
	}
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Images(); ok {
		XMLOpenStackImageWriteMany(writer, r, "images", "openstack_image")
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Password(); ok {
		writer.WriteCharacter("password", r)
	}
	if r, ok := object.Properties(); ok {
		XMLPropertyWriteMany(writer, r, "properties", "property")
	}
	if r, ok := object.RequiresAuthentication(); ok {
		writer.WriteBool("requires_authentication", r)
	}
	if r, ok := object.TenantName(); ok {
		writer.WriteCharacter("tenant_name", r)
	}
	if r, ok := object.Url(); ok {
		writer.WriteCharacter("url", r)
	}
	if r, ok := object.Username(); ok {
		writer.WriteCharacter("username", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLOpenStackImageProviderWriteMany(writer *XMLWriter, structSlice *OpenStackImageProviderSlice, plural, singular string) error {
	if plural == "" {
		plural = "openstack_image_providers"
	}
	if singular == "" {
		singular = "openstack_image_provider"
	}
	writer.WriteStart("", "openstack_image_providers", nil)
	for _, o := range structSlice.Slice() {
		XMLOpenStackImageProviderWriteOne(writer, &o, "openstack_image_provider")
	}
	writer.WriteEnd("openstack_image_providers")
	return nil
}

func XMLHostStorageWriteOne(writer *XMLWriter, object *HostStorage, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "host_storage"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Address(); ok {
		writer.WriteCharacter("address", r)
	}
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Host(); ok {
		XMLHostWriteOne(writer, r, "host")
	}
	if r, ok := object.LogicalUnits(); ok {
		XMLLogicalUnitWriteMany(writer, r, "logical_units", "logical_unit")
	}
	if r, ok := object.MountOptions(); ok {
		writer.WriteCharacter("mount_options", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.NfsRetrans(); ok {
		writer.WriteInt64("nfs_retrans", r)
	}
	if r, ok := object.NfsTimeo(); ok {
		writer.WriteInt64("nfs_timeo", r)
	}
	if r, ok := object.NfsVersion(); ok {
		XMLNfsVersionWriteOne(writer, r, "nfs_version")
	}
	if r, ok := object.OverrideLuns(); ok {
		writer.WriteBool("override_luns", r)
	}
	if r, ok := object.Password(); ok {
		writer.WriteCharacter("password", r)
	}
	if r, ok := object.Path(); ok {
		writer.WriteCharacter("path", r)
	}
	if r, ok := object.Port(); ok {
		writer.WriteInt64("port", r)
	}
	if r, ok := object.Portal(); ok {
		writer.WriteCharacter("portal", r)
	}
	if r, ok := object.Target(); ok {
		writer.WriteCharacter("target", r)
	}
	if r, ok := object.Type(); ok {
		XMLStorageTypeWriteOne(writer, r, "type")
	}
	if r, ok := object.Username(); ok {
		writer.WriteCharacter("username", r)
	}
	if r, ok := object.VfsType(); ok {
		writer.WriteCharacter("vfs_type", r)
	}
	if r, ok := object.VolumeGroup(); ok {
		XMLVolumeGroupWriteOne(writer, r, "volume_group")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLHostStorageWriteMany(writer *XMLWriter, structSlice *HostStorageSlice, plural, singular string) error {
	if plural == "" {
		plural = "host_storages"
	}
	if singular == "" {
		singular = "host_storage"
	}
	writer.WriteStart("", "host_storages", nil)
	for _, o := range structSlice.Slice() {
		XMLHostStorageWriteOne(writer, &o, "host_storage")
	}
	writer.WriteEnd("host_storages")
	return nil
}

func XMLStorageDomainLeaseWriteOne(writer *XMLWriter, object *StorageDomainLease, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "storage_domain_lease"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.StorageDomain(); ok {
		XMLStorageDomainWriteOne(writer, r, "storage_domain")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLStorageDomainLeaseWriteMany(writer *XMLWriter, structSlice *StorageDomainLeaseSlice, plural, singular string) error {
	if plural == "" {
		plural = "storage_domain_leases"
	}
	if singular == "" {
		singular = "storage_domain_lease"
	}
	writer.WriteStart("", "storage_domain_leases", nil)
	for _, o := range structSlice.Slice() {
		XMLStorageDomainLeaseWriteOne(writer, &o, "storage_domain_lease")
	}
	writer.WriteEnd("storage_domain_leases")
	return nil
}

func XMLIpAddressAssignmentWriteOne(writer *XMLWriter, object *IpAddressAssignment, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "ip_address_assignment"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.AssignmentMethod(); ok {
		XMLBootProtocolWriteOne(writer, r, "assignment_method")
	}
	if r, ok := object.Ip(); ok {
		XMLIpWriteOne(writer, r, "ip")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLIpAddressAssignmentWriteMany(writer *XMLWriter, structSlice *IpAddressAssignmentSlice, plural, singular string) error {
	if plural == "" {
		plural = "ip_address_assignments"
	}
	if singular == "" {
		singular = "ip_address_assignment"
	}
	writer.WriteStart("", "ip_address_assignments", nil)
	for _, o := range structSlice.Slice() {
		XMLIpAddressAssignmentWriteOne(writer, &o, "ip_address_assignment")
	}
	writer.WriteEnd("ip_address_assignments")
	return nil
}

func XMLNfsProfileDetailWriteOne(writer *XMLWriter, object *NfsProfileDetail, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "nfs_profile_detail"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.NfsServerIp(); ok {
		writer.WriteCharacter("nfs_server_ip", r)
	}
	if r, ok := object.ProfileDetails(); ok {
		XMLProfileDetailWriteMany(writer, r, "profile_details", "profile_detail")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLNfsProfileDetailWriteMany(writer *XMLWriter, structSlice *NfsProfileDetailSlice, plural, singular string) error {
	if plural == "" {
		plural = "nfs_profile_details"
	}
	if singular == "" {
		singular = "nfs_profile_detail"
	}
	writer.WriteStart("", "nfs_profile_details", nil)
	for _, o := range structSlice.Slice() {
		XMLNfsProfileDetailWriteOne(writer, &o, "nfs_profile_detail")
	}
	writer.WriteEnd("nfs_profile_details")
	return nil
}

func XMLQuotaStorageLimitWriteOne(writer *XMLWriter, object *QuotaStorageLimit, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "quota_storage_limit"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Limit(); ok {
		writer.WriteInt64("limit", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Quota(); ok {
		XMLQuotaWriteOne(writer, r, "quota")
	}
	if r, ok := object.StorageDomain(); ok {
		XMLStorageDomainWriteOne(writer, r, "storage_domain")
	}
	if r, ok := object.Usage(); ok {
		writer.WriteFloat64("usage", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLQuotaStorageLimitWriteMany(writer *XMLWriter, structSlice *QuotaStorageLimitSlice, plural, singular string) error {
	if plural == "" {
		plural = "quota_storage_limits"
	}
	if singular == "" {
		singular = "quota_storage_limit"
	}
	writer.WriteStart("", "quota_storage_limits", nil)
	for _, o := range structSlice.Slice() {
		XMLQuotaStorageLimitWriteOne(writer, &o, "quota_storage_limit")
	}
	writer.WriteEnd("quota_storage_limits")
	return nil
}

func XMLFilterWriteOne(writer *XMLWriter, object *Filter, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "filter"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Position(); ok {
		writer.WriteInt64("position", r)
	}
	if r, ok := object.SchedulingPolicyUnit(); ok {
		XMLSchedulingPolicyUnitWriteOne(writer, r, "scheduling_policy_unit")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLFilterWriteMany(writer *XMLWriter, structSlice *FilterSlice, plural, singular string) error {
	if plural == "" {
		plural = "filters"
	}
	if singular == "" {
		singular = "filter"
	}
	writer.WriteStart("", "filters", nil)
	for _, o := range structSlice.Slice() {
		XMLFilterWriteOne(writer, &o, "filter")
	}
	writer.WriteEnd("filters")
	return nil
}

func XMLSkipIfSdActiveWriteOne(writer *XMLWriter, object *SkipIfSdActive, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "skip_if_sd_active"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Enabled(); ok {
		writer.WriteBool("enabled", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLSkipIfSdActiveWriteMany(writer *XMLWriter, structSlice *SkipIfSdActiveSlice, plural, singular string) error {
	if plural == "" {
		plural = "skip_if_sd_actives"
	}
	if singular == "" {
		singular = "skip_if_sd_active"
	}
	writer.WriteStart("", "skip_if_sd_actives", nil)
	for _, o := range structSlice.Slice() {
		XMLSkipIfSdActiveWriteOne(writer, &o, "skip_if_sd_active")
	}
	writer.WriteEnd("skip_if_sd_actives")
	return nil
}

func XMLApiSummaryWriteOne(writer *XMLWriter, object *ApiSummary, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "api_summary"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Hosts(); ok {
		XMLApiSummaryItemWriteOne(writer, r, "hosts")
	}
	if r, ok := object.StorageDomains(); ok {
		XMLApiSummaryItemWriteOne(writer, r, "storage_domains")
	}
	if r, ok := object.Users(); ok {
		XMLApiSummaryItemWriteOne(writer, r, "users")
	}
	if r, ok := object.Vms(); ok {
		XMLApiSummaryItemWriteOne(writer, r, "vms")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLApiSummaryWriteMany(writer *XMLWriter, structSlice *ApiSummarySlice, plural, singular string) error {
	if plural == "" {
		plural = "api_summaries"
	}
	if singular == "" {
		singular = "api_summary"
	}
	writer.WriteStart("", "api_summaries", nil)
	for _, o := range structSlice.Slice() {
		XMLApiSummaryWriteOne(writer, &o, "api_summary")
	}
	writer.WriteEnd("api_summaries")
	return nil
}

func XMLGlusterHookWriteOne(writer *XMLWriter, object *GlusterHook, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "gluster_hook"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.Checksum(); ok {
		writer.WriteCharacter("checksum", r)
	}
	if r, ok := object.Cluster(); ok {
		XMLClusterWriteOne(writer, r, "cluster")
	}
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.ConflictStatus(); ok {
		writer.WriteInt64("conflict_status", r)
	}
	if r, ok := object.Conflicts(); ok {
		writer.WriteCharacter("conflicts", r)
	}
	if r, ok := object.Content(); ok {
		writer.WriteCharacter("content", r)
	}
	if r, ok := object.ContentType(); ok {
		XMLHookContentTypeWriteOne(writer, r, "content_type")
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.GlusterCommand(); ok {
		writer.WriteCharacter("gluster_command", r)
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.ServerHooks(); ok {
		XMLGlusterServerHookWriteMany(writer, r, "server_hooks", "server_hook")
	}
	if r, ok := object.Stage(); ok {
		XMLHookStageWriteOne(writer, r, "stage")
	}
	if r, ok := object.Status(); ok {
		XMLGlusterHookStatusWriteOne(writer, r, "status")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLGlusterHookWriteMany(writer *XMLWriter, structSlice *GlusterHookSlice, plural, singular string) error {
	if plural == "" {
		plural = "gluster_hooks"
	}
	if singular == "" {
		singular = "gluster_hook"
	}
	writer.WriteStart("", "gluster_hooks", nil)
	for _, o := range structSlice.Slice() {
		XMLGlusterHookWriteOne(writer, &o, "gluster_hook")
	}
	writer.WriteEnd("gluster_hooks")
	return nil
}

func XMLMigrationOptionsWriteOne(writer *XMLWriter, object *MigrationOptions, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "migration"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.AutoConverge(); ok {
		XMLInheritableBooleanWriteOne(writer, r, "auto_converge")
	}
	if r, ok := object.Bandwidth(); ok {
		XMLMigrationBandwidthWriteOne(writer, r, "bandwidth")
	}
	if r, ok := object.Compressed(); ok {
		XMLInheritableBooleanWriteOne(writer, r, "compressed")
	}
	if r, ok := object.Policy(); ok {
		XMLMigrationPolicyWriteOne(writer, r, "policy")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLMigrationOptionsWriteMany(writer *XMLWriter, structSlice *MigrationOptionsSlice, plural, singular string) error {
	if plural == "" {
		plural = "migration_optionss"
	}
	if singular == "" {
		singular = "migration"
	}
	writer.WriteStart("", "migration_optionss", nil)
	for _, o := range structSlice.Slice() {
		XMLMigrationOptionsWriteOne(writer, &o, "migration")
	}
	writer.WriteEnd("migration_optionss")
	return nil
}

func XMLFaultWriteOne(writer *XMLWriter, object *Fault, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "fault"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Detail(); ok {
		writer.WriteCharacter("detail", r)
	}
	if r, ok := object.Reason(); ok {
		writer.WriteCharacter("reason", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLFaultWriteMany(writer *XMLWriter, structSlice *FaultSlice, plural, singular string) error {
	if plural == "" {
		plural = "faults"
	}
	if singular == "" {
		singular = "fault"
	}
	writer.WriteStart("", "faults", nil)
	for _, o := range structSlice.Slice() {
		XMLFaultWriteOne(writer, &o, "fault")
	}
	writer.WriteEnd("faults")
	return nil
}

func XMLGracePeriodWriteOne(writer *XMLWriter, object *GracePeriod, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "grace_period"
	}
	writer.WriteStart("", tag, nil)
	if r, ok := object.Expiry(); ok {
		writer.WriteInt64("expiry", r)
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLGracePeriodWriteMany(writer *XMLWriter, structSlice *GracePeriodSlice, plural, singular string) error {
	if plural == "" {
		plural = "grace_periods"
	}
	if singular == "" {
		singular = "grace_period"
	}
	writer.WriteStart("", "grace_periods", nil)
	for _, o := range structSlice.Slice() {
		XMLGracePeriodWriteOne(writer, &o, "grace_period")
	}
	writer.WriteEnd("grace_periods")
	return nil
}

func XMLActionWriteOne(writer *XMLWriter, object *Action, tag string) error {
	if object == nil {
		return fmt.Errorf("input object pointer is nil")
	}
	if tag == "" {
		tag = "action"
	}
	var attrs map[string]string
	if r, ok := object.Id(); ok {
		if attrs == nil {
			attrs = make(map[string]string)
		}
		attrs["id"] = r
	}
	writer.WriteStart("", tag, attrs)
	if r, ok := object.AllowPartialImport(); ok {
		writer.WriteBool("allow_partial_import", r)
	}
	if r, ok := object.Async(); ok {
		writer.WriteBool("async", r)
	}
	if r, ok := object.Bricks(); ok {
		XMLGlusterBrickWriteMany(writer, r, "bricks", "brick")
	}
	if r, ok := object.Certificates(); ok {
		XMLCertificateWriteMany(writer, r, "certificates", "certificate")
	}
	if r, ok := object.CheckConnectivity(); ok {
		writer.WriteBool("check_connectivity", r)
	}
	if r, ok := object.Clone(); ok {
		writer.WriteBool("clone", r)
	}
	if r, ok := object.Cluster(); ok {
		XMLClusterWriteOne(writer, r, "cluster")
	}
	if r, ok := object.CollapseSnapshots(); ok {
		writer.WriteBool("collapse_snapshots", r)
	}
	if r, ok := object.Comment(); ok {
		writer.WriteCharacter("comment", r)
	}
	if r, ok := object.ConnectivityTimeout(); ok {
		writer.WriteInt64("connectivity_timeout", r)
	}
	if r, ok := object.DataCenter(); ok {
		XMLDataCenterWriteOne(writer, r, "data_center")
	}
	if r, ok := object.DeployHostedEngine(); ok {
		writer.WriteBool("deploy_hosted_engine", r)
	}
	if r, ok := object.Description(); ok {
		writer.WriteCharacter("description", r)
	}
	if r, ok := object.Details(); ok {
		XMLGlusterVolumeProfileDetailsWriteOne(writer, r, "details")
	}
	if r, ok := object.DiscardSnapshots(); ok {
		writer.WriteBool("discard_snapshots", r)
	}
	if r, ok := object.Disk(); ok {
		XMLDiskWriteOne(writer, r, "disk")
	}
	if r, ok := object.Disks(); ok {
		XMLDiskWriteMany(writer, r, "disks", "disk")
	}
	if r, ok := object.Exclusive(); ok {
		writer.WriteBool("exclusive", r)
	}
	if r, ok := object.Fault(); ok {
		XMLFaultWriteOne(writer, r, "fault")
	}
	if r, ok := object.FenceType(); ok {
		writer.WriteCharacter("fence_type", r)
	}
	if r, ok := object.Filter(); ok {
		writer.WriteBool("filter", r)
	}
	if r, ok := object.FixLayout(); ok {
		writer.WriteBool("fix_layout", r)
	}
	if r, ok := object.Force(); ok {
		writer.WriteBool("force", r)
	}
	if r, ok := object.GracePeriod(); ok {
		XMLGracePeriodWriteOne(writer, r, "grace_period")
	}
	if r, ok := object.Host(); ok {
		XMLHostWriteOne(writer, r, "host")
	}
	if r, ok := object.Image(); ok {
		writer.WriteCharacter("image", r)
	}
	if r, ok := object.ImportAsTemplate(); ok {
		writer.WriteBool("import_as_template", r)
	}
	if r, ok := object.IsAttached(); ok {
		writer.WriteBool("is_attached", r)
	}
	if r, ok := object.Iscsi(); ok {
		XMLIscsiDetailsWriteOne(writer, r, "iscsi")
	}
	if r, ok := object.IscsiTargets(); ok {
		writer.WriteCharacters("iscsi_targets", r)
	}
	if r, ok := object.Job(); ok {
		XMLJobWriteOne(writer, r, "job")
	}
	if r, ok := object.LogicalUnits(); ok {
		XMLLogicalUnitWriteMany(writer, r, "logical_units", "logical_unit")
	}
	if r, ok := object.MaintenanceEnabled(); ok {
		writer.WriteBool("maintenance_enabled", r)
	}
	if r, ok := object.ModifiedBonds(); ok {
		XMLHostNicWriteMany(writer, r, "modified_bonds", "host_nic")
	}
	if r, ok := object.ModifiedLabels(); ok {
		XMLNetworkLabelWriteMany(writer, r, "modified_labels", "network_label")
	}
	if r, ok := object.ModifiedNetworkAttachments(); ok {
		XMLNetworkAttachmentWriteMany(writer, r, "modified_network_attachments", "network_attachment")
	}
	if r, ok := object.Name(); ok {
		writer.WriteCharacter("name", r)
	}
	if r, ok := object.Option(); ok {
		XMLOptionWriteOne(writer, r, "option")
	}
	if r, ok := object.Pause(); ok {
		writer.WriteBool("pause", r)
	}
	if r, ok := object.PowerManagement(); ok {
		XMLPowerManagementWriteOne(writer, r, "power_management")
	}
	if r, ok := object.ProxyTicket(); ok {
		XMLProxyTicketWriteOne(writer, r, "proxy_ticket")
	}
	if r, ok := object.Reason(); ok {
		writer.WriteCharacter("reason", r)
	}
	if r, ok := object.ReassignBadMacs(); ok {
		writer.WriteBool("reassign_bad_macs", r)
	}
	if r, ok := object.RemoteViewerConnectionFile(); ok {
		writer.WriteCharacter("remote_viewer_connection_file", r)
	}
	if r, ok := object.RemovedBonds(); ok {
		XMLHostNicWriteMany(writer, r, "removed_bonds", "host_nic")
	}
	if r, ok := object.RemovedLabels(); ok {
		XMLNetworkLabelWriteMany(writer, r, "removed_labels", "network_label")
	}
	if r, ok := object.RemovedNetworkAttachments(); ok {
		XMLNetworkAttachmentWriteMany(writer, r, "removed_network_attachments", "network_attachment")
	}
	if r, ok := object.ResolutionType(); ok {
		writer.WriteCharacter("resolution_type", r)
	}
	if r, ok := object.RestoreMemory(); ok {
		writer.WriteBool("restore_memory", r)
	}
	if r, ok := object.RootPassword(); ok {
		writer.WriteCharacter("root_password", r)
	}
	if r, ok := object.Snapshot(); ok {
		XMLSnapshotWriteOne(writer, r, "snapshot")
	}
	if r, ok := object.Ssh(); ok {
		XMLSshWriteOne(writer, r, "ssh")
	}
	if r, ok := object.Status(); ok {
		writer.WriteCharacter("status", r)
	}
	if r, ok := object.StopGlusterService(); ok {
		writer.WriteBool("stop_gluster_service", r)
	}
	if r, ok := object.StorageDomain(); ok {
		XMLStorageDomainWriteOne(writer, r, "storage_domain")
	}
	if r, ok := object.StorageDomains(); ok {
		XMLStorageDomainWriteMany(writer, r, "storage_domains", "storage_domain")
	}
	if r, ok := object.Succeeded(); ok {
		writer.WriteBool("succeeded", r)
	}
	if r, ok := object.SynchronizedNetworkAttachments(); ok {
		XMLNetworkAttachmentWriteMany(writer, r, "synchronized_network_attachments", "network_attachment")
	}
	if r, ok := object.Template(); ok {
		XMLTemplateWriteOne(writer, r, "template")
	}
	if r, ok := object.Ticket(); ok {
		XMLTicketWriteOne(writer, r, "ticket")
	}
	if r, ok := object.UndeployHostedEngine(); ok {
		writer.WriteBool("undeploy_hosted_engine", r)
	}
	if r, ok := object.UseCloudInit(); ok {
		writer.WriteBool("use_cloud_init", r)
	}
	if r, ok := object.UseSysprep(); ok {
		writer.WriteBool("use_sysprep", r)
	}
	if r, ok := object.VirtualFunctionsConfiguration(); ok {
		XMLHostNicVirtualFunctionsConfigurationWriteOne(writer, r, "virtual_functions_configuration")
	}
	if r, ok := object.Vm(); ok {
		XMLVmWriteOne(writer, r, "vm")
	}
	if r, ok := object.VnicProfileMappings(); ok {
		XMLVnicProfileMappingWriteMany(writer, r, "vnic_profile_mappings", "vnic_profile_mapping")
	}
	writer.WriteEnd(tag)
	return nil
}

func XMLActionWriteMany(writer *XMLWriter, structSlice *ActionSlice, plural, singular string) error {
	if plural == "" {
		plural = "actions"
	}
	if singular == "" {
		singular = "action"
	}
	writer.WriteStart("", "actions", nil)
	for _, o := range structSlice.Slice() {
		XMLActionWriteOne(writer, &o, "action")
	}
	writer.WriteEnd("actions")
	return nil
}

func XMLGlusterVolumeStatusWriteOne(writer *XMLWriter, enum GlusterVolumeStatus, tag string) {
	if tag == "" {
		tag = "gluster_volume_status"
	}
	writer.WriteCharacter("gluster_volume_status", string(enum))
}

func XMLGlusterVolumeStatusWriteMany(writer *XMLWriter, enums []GlusterVolumeStatus, plural, singular string) error {
	if plural == "" {
		plural = "gluster_volume_statuss"
	}
	if singular == "" {
		singular = "gluster_volume_status"
	}
	writer.WriteStart("", "gluster_volume_statuss", nil)
	for _, e := range enums {
		writer.WriteCharacter("gluster_volume_status", string(e))
	}
	writer.WriteEnd("gluster_volume_statuss")
	return nil
}

func XMLGlusterBrickStatusWriteOne(writer *XMLWriter, enum GlusterBrickStatus, tag string) {
	if tag == "" {
		tag = "gluster_brick_status"
	}
	writer.WriteCharacter("gluster_brick_status", string(enum))
}

func XMLGlusterBrickStatusWriteMany(writer *XMLWriter, enums []GlusterBrickStatus, plural, singular string) error {
	if plural == "" {
		plural = "gluster_brick_statuss"
	}
	if singular == "" {
		singular = "gluster_brick_status"
	}
	writer.WriteStart("", "gluster_brick_statuss", nil)
	for _, e := range enums {
		writer.WriteCharacter("gluster_brick_status", string(e))
	}
	writer.WriteEnd("gluster_brick_statuss")
	return nil
}

func XMLGlusterHookStatusWriteOne(writer *XMLWriter, enum GlusterHookStatus, tag string) {
	if tag == "" {
		tag = "gluster_hook_status"
	}
	writer.WriteCharacter("gluster_hook_status", string(enum))
}

func XMLGlusterHookStatusWriteMany(writer *XMLWriter, enums []GlusterHookStatus, plural, singular string) error {
	if plural == "" {
		plural = "gluster_hook_statuss"
	}
	if singular == "" {
		singular = "gluster_hook_status"
	}
	writer.WriteStart("", "gluster_hook_statuss", nil)
	for _, e := range enums {
		writer.WriteCharacter("gluster_hook_status", string(e))
	}
	writer.WriteEnd("gluster_hook_statuss")
	return nil
}

func XMLStorageDomainTypeWriteOne(writer *XMLWriter, enum StorageDomainType, tag string) {
	if tag == "" {
		tag = "storage_domain_type"
	}
	writer.WriteCharacter("storage_domain_type", string(enum))
}

func XMLStorageDomainTypeWriteMany(writer *XMLWriter, enums []StorageDomainType, plural, singular string) error {
	if plural == "" {
		plural = "storage_domain_types"
	}
	if singular == "" {
		singular = "storage_domain_type"
	}
	writer.WriteStart("", "storage_domain_types", nil)
	for _, e := range enums {
		writer.WriteCharacter("storage_domain_type", string(e))
	}
	writer.WriteEnd("storage_domain_types")
	return nil
}

func XMLAutoNumaStatusWriteOne(writer *XMLWriter, enum AutoNumaStatus, tag string) {
	if tag == "" {
		tag = "auto_numa_status"
	}
	writer.WriteCharacter("auto_numa_status", string(enum))
}

func XMLAutoNumaStatusWriteMany(writer *XMLWriter, enums []AutoNumaStatus, plural, singular string) error {
	if plural == "" {
		plural = "auto_numa_statuss"
	}
	if singular == "" {
		singular = "auto_numa_status"
	}
	writer.WriteStart("", "auto_numa_statuss", nil)
	for _, e := range enums {
		writer.WriteCharacter("auto_numa_status", string(e))
	}
	writer.WriteEnd("auto_numa_statuss")
	return nil
}

func XMLPolicyUnitTypeWriteOne(writer *XMLWriter, enum PolicyUnitType, tag string) {
	if tag == "" {
		tag = "policy_unit_type"
	}
	writer.WriteCharacter("policy_unit_type", string(enum))
}

func XMLPolicyUnitTypeWriteMany(writer *XMLWriter, enums []PolicyUnitType, plural, singular string) error {
	if plural == "" {
		plural = "policy_unit_types"
	}
	if singular == "" {
		singular = "policy_unit_type"
	}
	writer.WriteStart("", "policy_unit_types", nil)
	for _, e := range enums {
		writer.WriteCharacter("policy_unit_type", string(e))
	}
	writer.WriteEnd("policy_unit_types")
	return nil
}

func XMLNicInterfaceWriteOne(writer *XMLWriter, enum NicInterface, tag string) {
	if tag == "" {
		tag = "nic_interface"
	}
	writer.WriteCharacter("nic_interface", string(enum))
}

func XMLNicInterfaceWriteMany(writer *XMLWriter, enums []NicInterface, plural, singular string) error {
	if plural == "" {
		plural = "nic_interfaces"
	}
	if singular == "" {
		singular = "nic_interface"
	}
	writer.WriteStart("", "nic_interfaces", nil)
	for _, e := range enums {
		writer.WriteCharacter("nic_interface", string(e))
	}
	writer.WriteEnd("nic_interfaces")
	return nil
}

func XMLQcowVersionWriteOne(writer *XMLWriter, enum QcowVersion, tag string) {
	if tag == "" {
		tag = "qcow_version"
	}
	writer.WriteCharacter("qcow_version", string(enum))
}

func XMLQcowVersionWriteMany(writer *XMLWriter, enums []QcowVersion, plural, singular string) error {
	if plural == "" {
		plural = "qcow_versions"
	}
	if singular == "" {
		singular = "qcow_version"
	}
	writer.WriteStart("", "qcow_versions", nil)
	for _, e := range enums {
		writer.WriteCharacter("qcow_version", string(e))
	}
	writer.WriteEnd("qcow_versions")
	return nil
}

func XMLSerialNumberPolicyWriteOne(writer *XMLWriter, enum SerialNumberPolicy, tag string) {
	if tag == "" {
		tag = "serial_number_policy"
	}
	writer.WriteCharacter("serial_number_policy", string(enum))
}

func XMLSerialNumberPolicyWriteMany(writer *XMLWriter, enums []SerialNumberPolicy, plural, singular string) error {
	if plural == "" {
		plural = "serial_number_policies"
	}
	if singular == "" {
		singular = "serial_number_policy"
	}
	writer.WriteStart("", "serial_number_policies", nil)
	for _, e := range enums {
		writer.WriteCharacter("serial_number_policy", string(e))
	}
	writer.WriteEnd("serial_number_policies")
	return nil
}

func XMLHookStatusWriteOne(writer *XMLWriter, enum HookStatus, tag string) {
	if tag == "" {
		tag = "hook_status"
	}
	writer.WriteCharacter("hook_status", string(enum))
}

func XMLHookStatusWriteMany(writer *XMLWriter, enums []HookStatus, plural, singular string) error {
	if plural == "" {
		plural = "hook_statuss"
	}
	if singular == "" {
		singular = "hook_status"
	}
	writer.WriteStart("", "hook_statuss", nil)
	for _, e := range enums {
		writer.WriteCharacter("hook_status", string(e))
	}
	writer.WriteEnd("hook_statuss")
	return nil
}

func XMLUsbTypeWriteOne(writer *XMLWriter, enum UsbType, tag string) {
	if tag == "" {
		tag = "usb_type"
	}
	writer.WriteCharacter("usb_type", string(enum))
}

func XMLUsbTypeWriteMany(writer *XMLWriter, enums []UsbType, plural, singular string) error {
	if plural == "" {
		plural = "usb_types"
	}
	if singular == "" {
		singular = "usb_type"
	}
	writer.WriteStart("", "usb_types", nil)
	for _, e := range enums {
		writer.WriteCharacter("usb_type", string(e))
	}
	writer.WriteEnd("usb_types")
	return nil
}

func XMLSeLinuxModeWriteOne(writer *XMLWriter, enum SeLinuxMode, tag string) {
	if tag == "" {
		tag = "se_linux_mode"
	}
	writer.WriteCharacter("se_linux_mode", string(enum))
}

func XMLSeLinuxModeWriteMany(writer *XMLWriter, enums []SeLinuxMode, plural, singular string) error {
	if plural == "" {
		plural = "se_linux_modes"
	}
	if singular == "" {
		singular = "se_linux_mode"
	}
	writer.WriteStart("", "se_linux_modes", nil)
	for _, e := range enums {
		writer.WriteCharacter("se_linux_mode", string(e))
	}
	writer.WriteEnd("se_linux_modes")
	return nil
}

func XMLExternalVmProviderTypeWriteOne(writer *XMLWriter, enum ExternalVmProviderType, tag string) {
	if tag == "" {
		tag = "external_vm_provider_type"
	}
	writer.WriteCharacter("external_vm_provider_type", string(enum))
}

func XMLExternalVmProviderTypeWriteMany(writer *XMLWriter, enums []ExternalVmProviderType, plural, singular string) error {
	if plural == "" {
		plural = "external_vm_provider_types"
	}
	if singular == "" {
		singular = "external_vm_provider_type"
	}
	writer.WriteStart("", "external_vm_provider_types", nil)
	for _, e := range enums {
		writer.WriteCharacter("external_vm_provider_type", string(e))
	}
	writer.WriteEnd("external_vm_provider_types")
	return nil
}

func XMLDiskInterfaceWriteOne(writer *XMLWriter, enum DiskInterface, tag string) {
	if tag == "" {
		tag = "disk_interface"
	}
	writer.WriteCharacter("disk_interface", string(enum))
}

func XMLDiskInterfaceWriteMany(writer *XMLWriter, enums []DiskInterface, plural, singular string) error {
	if plural == "" {
		plural = "disk_interfaces"
	}
	if singular == "" {
		singular = "disk_interface"
	}
	writer.WriteStart("", "disk_interfaces", nil)
	for _, e := range enums {
		writer.WriteCharacter("disk_interface", string(e))
	}
	writer.WriteEnd("disk_interfaces")
	return nil
}

func XMLLunStatusWriteOne(writer *XMLWriter, enum LunStatus, tag string) {
	if tag == "" {
		tag = "lun_status"
	}
	writer.WriteCharacter("lun_status", string(enum))
}

func XMLLunStatusWriteMany(writer *XMLWriter, enums []LunStatus, plural, singular string) error {
	if plural == "" {
		plural = "lun_statuss"
	}
	if singular == "" {
		singular = "lun_status"
	}
	writer.WriteStart("", "lun_statuss", nil)
	for _, e := range enums {
		writer.WriteCharacter("lun_status", string(e))
	}
	writer.WriteEnd("lun_statuss")
	return nil
}

func XMLMessageBrokerTypeWriteOne(writer *XMLWriter, enum MessageBrokerType, tag string) {
	if tag == "" {
		tag = "message_broker_type"
	}
	writer.WriteCharacter("message_broker_type", string(enum))
}

func XMLMessageBrokerTypeWriteMany(writer *XMLWriter, enums []MessageBrokerType, plural, singular string) error {
	if plural == "" {
		plural = "message_broker_types"
	}
	if singular == "" {
		singular = "message_broker_type"
	}
	writer.WriteStart("", "message_broker_types", nil)
	for _, e := range enums {
		writer.WriteCharacter("message_broker_type", string(e))
	}
	writer.WriteEnd("message_broker_types")
	return nil
}

func XMLVnicPassThroughModeWriteOne(writer *XMLWriter, enum VnicPassThroughMode, tag string) {
	if tag == "" {
		tag = "vnic_pass_through_mode"
	}
	writer.WriteCharacter("vnic_pass_through_mode", string(enum))
}

func XMLVnicPassThroughModeWriteMany(writer *XMLWriter, enums []VnicPassThroughMode, plural, singular string) error {
	if plural == "" {
		plural = "vnic_pass_through_modes"
	}
	if singular == "" {
		singular = "vnic_pass_through_mode"
	}
	writer.WriteStart("", "vnic_pass_through_modes", nil)
	for _, e := range enums {
		writer.WriteCharacter("vnic_pass_through_mode", string(e))
	}
	writer.WriteEnd("vnic_pass_through_modes")
	return nil
}

func XMLDiskStorageTypeWriteOne(writer *XMLWriter, enum DiskStorageType, tag string) {
	if tag == "" {
		tag = "disk_storage_type"
	}
	writer.WriteCharacter("disk_storage_type", string(enum))
}

func XMLDiskStorageTypeWriteMany(writer *XMLWriter, enums []DiskStorageType, plural, singular string) error {
	if plural == "" {
		plural = "disk_storage_types"
	}
	if singular == "" {
		singular = "disk_storage_type"
	}
	writer.WriteStart("", "disk_storage_types", nil)
	for _, e := range enums {
		writer.WriteCharacter("disk_storage_type", string(e))
	}
	writer.WriteEnd("disk_storage_types")
	return nil
}

func XMLVmPoolTypeWriteOne(writer *XMLWriter, enum VmPoolType, tag string) {
	if tag == "" {
		tag = "vm_pool_type"
	}
	writer.WriteCharacter("vm_pool_type", string(enum))
}

func XMLVmPoolTypeWriteMany(writer *XMLWriter, enums []VmPoolType, plural, singular string) error {
	if plural == "" {
		plural = "vm_pool_types"
	}
	if singular == "" {
		singular = "vm_pool_type"
	}
	writer.WriteStart("", "vm_pool_types", nil)
	for _, e := range enums {
		writer.WriteCharacter("vm_pool_type", string(e))
	}
	writer.WriteEnd("vm_pool_types")
	return nil
}

func XMLResolutionTypeWriteOne(writer *XMLWriter, enum ResolutionType, tag string) {
	if tag == "" {
		tag = "resolution_type"
	}
	writer.WriteCharacter("resolution_type", string(enum))
}

func XMLResolutionTypeWriteMany(writer *XMLWriter, enums []ResolutionType, plural, singular string) error {
	if plural == "" {
		plural = "resolution_types"
	}
	if singular == "" {
		singular = "resolution_type"
	}
	writer.WriteStart("", "resolution_types", nil)
	for _, e := range enums {
		writer.WriteCharacter("resolution_type", string(e))
	}
	writer.WriteEnd("resolution_types")
	return nil
}

func XMLReportedDeviceTypeWriteOne(writer *XMLWriter, enum ReportedDeviceType, tag string) {
	if tag == "" {
		tag = "reported_device_type"
	}
	writer.WriteCharacter("reported_device_type", string(enum))
}

func XMLReportedDeviceTypeWriteMany(writer *XMLWriter, enums []ReportedDeviceType, plural, singular string) error {
	if plural == "" {
		plural = "reported_device_types"
	}
	if singular == "" {
		singular = "reported_device_type"
	}
	writer.WriteStart("", "reported_device_types", nil)
	for _, e := range enums {
		writer.WriteCharacter("reported_device_type", string(e))
	}
	writer.WriteEnd("reported_device_types")
	return nil
}

func XMLMigrateOnErrorWriteOne(writer *XMLWriter, enum MigrateOnError, tag string) {
	if tag == "" {
		tag = "migrate_on_error"
	}
	writer.WriteCharacter("migrate_on_error", string(enum))
}

func XMLMigrateOnErrorWriteMany(writer *XMLWriter, enums []MigrateOnError, plural, singular string) error {
	if plural == "" {
		plural = "migrate_on_errors"
	}
	if singular == "" {
		singular = "migrate_on_error"
	}
	writer.WriteStart("", "migrate_on_errors", nil)
	for _, e := range enums {
		writer.WriteCharacter("migrate_on_error", string(e))
	}
	writer.WriteEnd("migrate_on_errors")
	return nil
}

func XMLStorageTypeWriteOne(writer *XMLWriter, enum StorageType, tag string) {
	if tag == "" {
		tag = "storage_type"
	}
	writer.WriteCharacter("storage_type", string(enum))
}

func XMLStorageTypeWriteMany(writer *XMLWriter, enums []StorageType, plural, singular string) error {
	if plural == "" {
		plural = "storage_types"
	}
	if singular == "" {
		singular = "storage_type"
	}
	writer.WriteStart("", "storage_types", nil)
	for _, e := range enums {
		writer.WriteCharacter("storage_type", string(e))
	}
	writer.WriteEnd("storage_types")
	return nil
}

func XMLAccessProtocolWriteOne(writer *XMLWriter, enum AccessProtocol, tag string) {
	if tag == "" {
		tag = "access_protocol"
	}
	writer.WriteCharacter("access_protocol", string(enum))
}

func XMLAccessProtocolWriteMany(writer *XMLWriter, enums []AccessProtocol, plural, singular string) error {
	if plural == "" {
		plural = "access_protocols"
	}
	if singular == "" {
		singular = "access_protocol"
	}
	writer.WriteStart("", "access_protocols", nil)
	for _, e := range enums {
		writer.WriteCharacter("access_protocol", string(e))
	}
	writer.WriteEnd("access_protocols")
	return nil
}

func XMLStepStatusWriteOne(writer *XMLWriter, enum StepStatus, tag string) {
	if tag == "" {
		tag = "step_status"
	}
	writer.WriteCharacter("step_status", string(enum))
}

func XMLStepStatusWriteMany(writer *XMLWriter, enums []StepStatus, plural, singular string) error {
	if plural == "" {
		plural = "step_statuss"
	}
	if singular == "" {
		singular = "step_status"
	}
	writer.WriteStart("", "step_statuss", nil)
	for _, e := range enums {
		writer.WriteCharacter("step_status", string(e))
	}
	writer.WriteEnd("step_statuss")
	return nil
}

func XMLDiskStatusWriteOne(writer *XMLWriter, enum DiskStatus, tag string) {
	if tag == "" {
		tag = "disk_status"
	}
	writer.WriteCharacter("disk_status", string(enum))
}

func XMLDiskStatusWriteMany(writer *XMLWriter, enums []DiskStatus, plural, singular string) error {
	if plural == "" {
		plural = "disk_statuss"
	}
	if singular == "" {
		singular = "disk_status"
	}
	writer.WriteStart("", "disk_statuss", nil)
	for _, e := range enums {
		writer.WriteCharacter("disk_status", string(e))
	}
	writer.WriteEnd("disk_statuss")
	return nil
}

func XMLNetworkStatusWriteOne(writer *XMLWriter, enum NetworkStatus, tag string) {
	if tag == "" {
		tag = "network_status"
	}
	writer.WriteCharacter("network_status", string(enum))
}

func XMLNetworkStatusWriteMany(writer *XMLWriter, enums []NetworkStatus, plural, singular string) error {
	if plural == "" {
		plural = "network_statuss"
	}
	if singular == "" {
		singular = "network_status"
	}
	writer.WriteStart("", "network_statuss", nil)
	for _, e := range enums {
		writer.WriteCharacter("network_status", string(e))
	}
	writer.WriteEnd("network_statuss")
	return nil
}

func XMLRoleTypeWriteOne(writer *XMLWriter, enum RoleType, tag string) {
	if tag == "" {
		tag = "role_type"
	}
	writer.WriteCharacter("role_type", string(enum))
}

func XMLRoleTypeWriteMany(writer *XMLWriter, enums []RoleType, plural, singular string) error {
	if plural == "" {
		plural = "role_types"
	}
	if singular == "" {
		singular = "role_type"
	}
	writer.WriteStart("", "role_types", nil)
	for _, e := range enums {
		writer.WriteCharacter("role_type", string(e))
	}
	writer.WriteEnd("role_types")
	return nil
}

func XMLPayloadEncodingWriteOne(writer *XMLWriter, enum PayloadEncoding, tag string) {
	if tag == "" {
		tag = "payload_encoding"
	}
	writer.WriteCharacter("payload_encoding", string(enum))
}

func XMLPayloadEncodingWriteMany(writer *XMLWriter, enums []PayloadEncoding, plural, singular string) error {
	if plural == "" {
		plural = "payload_encodings"
	}
	if singular == "" {
		singular = "payload_encoding"
	}
	writer.WriteStart("", "payload_encodings", nil)
	for _, e := range enums {
		writer.WriteCharacter("payload_encoding", string(e))
	}
	writer.WriteEnd("payload_encodings")
	return nil
}

func XMLQosTypeWriteOne(writer *XMLWriter, enum QosType, tag string) {
	if tag == "" {
		tag = "qos_type"
	}
	writer.WriteCharacter("qos_type", string(enum))
}

func XMLQosTypeWriteMany(writer *XMLWriter, enums []QosType, plural, singular string) error {
	if plural == "" {
		plural = "qos_types"
	}
	if singular == "" {
		singular = "qos_type"
	}
	writer.WriteStart("", "qos_types", nil)
	for _, e := range enums {
		writer.WriteCharacter("qos_type", string(e))
	}
	writer.WriteEnd("qos_types")
	return nil
}

func XMLPmProxyTypeWriteOne(writer *XMLWriter, enum PmProxyType, tag string) {
	if tag == "" {
		tag = "pm_proxy_type"
	}
	writer.WriteCharacter("pm_proxy_type", string(enum))
}

func XMLPmProxyTypeWriteMany(writer *XMLWriter, enums []PmProxyType, plural, singular string) error {
	if plural == "" {
		plural = "pm_proxy_types"
	}
	if singular == "" {
		singular = "pm_proxy_type"
	}
	writer.WriteStart("", "pm_proxy_types", nil)
	for _, e := range enums {
		writer.WriteCharacter("pm_proxy_type", string(e))
	}
	writer.WriteEnd("pm_proxy_types")
	return nil
}

func XMLBootDeviceWriteOne(writer *XMLWriter, enum BootDevice, tag string) {
	if tag == "" {
		tag = "boot_device"
	}
	writer.WriteCharacter("boot_device", string(enum))
}

func XMLBootDeviceWriteMany(writer *XMLWriter, enums []BootDevice, plural, singular string) error {
	if plural == "" {
		plural = "boot_devices"
	}
	if singular == "" {
		singular = "boot_device"
	}
	writer.WriteStart("", "boot_devices", nil)
	for _, e := range enums {
		writer.WriteCharacter("boot_device", string(e))
	}
	writer.WriteEnd("boot_devices")
	return nil
}

func XMLConfigurationTypeWriteOne(writer *XMLWriter, enum ConfigurationType, tag string) {
	if tag == "" {
		tag = "configuration_type"
	}
	writer.WriteCharacter("configuration_type", string(enum))
}

func XMLConfigurationTypeWriteMany(writer *XMLWriter, enums []ConfigurationType, plural, singular string) error {
	if plural == "" {
		plural = "configuration_types"
	}
	if singular == "" {
		singular = "configuration_type"
	}
	writer.WriteStart("", "configuration_types", nil)
	for _, e := range enums {
		writer.WriteCharacter("configuration_type", string(e))
	}
	writer.WriteEnd("configuration_types")
	return nil
}

func XMLVmDeviceTypeWriteOne(writer *XMLWriter, enum VmDeviceType, tag string) {
	if tag == "" {
		tag = "vm_device_type"
	}
	writer.WriteCharacter("vm_device_type", string(enum))
}

func XMLVmDeviceTypeWriteMany(writer *XMLWriter, enums []VmDeviceType, plural, singular string) error {
	if plural == "" {
		plural = "vm_device_types"
	}
	if singular == "" {
		singular = "vm_device_type"
	}
	writer.WriteStart("", "vm_device_types", nil)
	for _, e := range enums {
		writer.WriteCharacter("vm_device_type", string(e))
	}
	writer.WriteEnd("vm_device_types")
	return nil
}

func XMLHookStageWriteOne(writer *XMLWriter, enum HookStage, tag string) {
	if tag == "" {
		tag = "hook_stage"
	}
	writer.WriteCharacter("hook_stage", string(enum))
}

func XMLHookStageWriteMany(writer *XMLWriter, enums []HookStage, plural, singular string) error {
	if plural == "" {
		plural = "hook_stages"
	}
	if singular == "" {
		singular = "hook_stage"
	}
	writer.WriteStart("", "hook_stages", nil)
	for _, e := range enums {
		writer.WriteCharacter("hook_stage", string(e))
	}
	writer.WriteEnd("hook_stages")
	return nil
}

func XMLHostProtocolWriteOne(writer *XMLWriter, enum HostProtocol, tag string) {
	if tag == "" {
		tag = "host_protocol"
	}
	writer.WriteCharacter("host_protocol", string(enum))
}

func XMLHostProtocolWriteMany(writer *XMLWriter, enums []HostProtocol, plural, singular string) error {
	if plural == "" {
		plural = "host_protocols"
	}
	if singular == "" {
		singular = "host_protocol"
	}
	writer.WriteStart("", "host_protocols", nil)
	for _, e := range enums {
		writer.WriteCharacter("host_protocol", string(e))
	}
	writer.WriteEnd("host_protocols")
	return nil
}

func XMLQuotaModeTypeWriteOne(writer *XMLWriter, enum QuotaModeType, tag string) {
	if tag == "" {
		tag = "quota_mode_type"
	}
	writer.WriteCharacter("quota_mode_type", string(enum))
}

func XMLQuotaModeTypeWriteMany(writer *XMLWriter, enums []QuotaModeType, plural, singular string) error {
	if plural == "" {
		plural = "quota_mode_types"
	}
	if singular == "" {
		singular = "quota_mode_type"
	}
	writer.WriteStart("", "quota_mode_types", nil)
	for _, e := range enums {
		writer.WriteCharacter("quota_mode_type", string(e))
	}
	writer.WriteEnd("quota_mode_types")
	return nil
}

func XMLSpmStatusWriteOne(writer *XMLWriter, enum SpmStatus, tag string) {
	if tag == "" {
		tag = "spm_status"
	}
	writer.WriteCharacter("spm_status", string(enum))
}

func XMLSpmStatusWriteMany(writer *XMLWriter, enums []SpmStatus, plural, singular string) error {
	if plural == "" {
		plural = "spm_statuss"
	}
	if singular == "" {
		singular = "spm_status"
	}
	writer.WriteStart("", "spm_statuss", nil)
	for _, e := range enums {
		writer.WriteCharacter("spm_status", string(e))
	}
	writer.WriteEnd("spm_statuss")
	return nil
}

func XMLImageTransferPhaseWriteOne(writer *XMLWriter, enum ImageTransferPhase, tag string) {
	if tag == "" {
		tag = "image_transfer_phase"
	}
	writer.WriteCharacter("image_transfer_phase", string(enum))
}

func XMLImageTransferPhaseWriteMany(writer *XMLWriter, enums []ImageTransferPhase, plural, singular string) error {
	if plural == "" {
		plural = "image_transfer_phases"
	}
	if singular == "" {
		singular = "image_transfer_phase"
	}
	writer.WriteStart("", "image_transfer_phases", nil)
	for _, e := range enums {
		writer.WriteCharacter("image_transfer_phase", string(e))
	}
	writer.WriteEnd("image_transfer_phases")
	return nil
}

func XMLStatisticUnitWriteOne(writer *XMLWriter, enum StatisticUnit, tag string) {
	if tag == "" {
		tag = "statistic_unit"
	}
	writer.WriteCharacter("statistic_unit", string(enum))
}

func XMLStatisticUnitWriteMany(writer *XMLWriter, enums []StatisticUnit, plural, singular string) error {
	if plural == "" {
		plural = "statistic_units"
	}
	if singular == "" {
		singular = "statistic_unit"
	}
	writer.WriteStart("", "statistic_units", nil)
	for _, e := range enums {
		writer.WriteCharacter("statistic_unit", string(e))
	}
	writer.WriteEnd("statistic_units")
	return nil
}

func XMLBootProtocolWriteOne(writer *XMLWriter, enum BootProtocol, tag string) {
	if tag == "" {
		tag = "boot_protocol"
	}
	writer.WriteCharacter("boot_protocol", string(enum))
}

func XMLBootProtocolWriteMany(writer *XMLWriter, enums []BootProtocol, plural, singular string) error {
	if plural == "" {
		plural = "boot_protocols"
	}
	if singular == "" {
		singular = "boot_protocol"
	}
	writer.WriteStart("", "boot_protocols", nil)
	for _, e := range enums {
		writer.WriteCharacter("boot_protocol", string(e))
	}
	writer.WriteEnd("boot_protocols")
	return nil
}

func XMLVmTypeWriteOne(writer *XMLWriter, enum VmType, tag string) {
	if tag == "" {
		tag = "vm_type"
	}
	writer.WriteCharacter("vm_type", string(enum))
}

func XMLVmTypeWriteMany(writer *XMLWriter, enums []VmType, plural, singular string) error {
	if plural == "" {
		plural = "vm_types"
	}
	if singular == "" {
		singular = "vm_type"
	}
	writer.WriteStart("", "vm_types", nil)
	for _, e := range enums {
		writer.WriteCharacter("vm_type", string(e))
	}
	writer.WriteEnd("vm_types")
	return nil
}

func XMLArchitectureWriteOne(writer *XMLWriter, enum Architecture, tag string) {
	if tag == "" {
		tag = "architecture"
	}
	writer.WriteCharacter("architecture", string(enum))
}

func XMLArchitectureWriteMany(writer *XMLWriter, enums []Architecture, plural, singular string) error {
	if plural == "" {
		plural = "architectures"
	}
	if singular == "" {
		singular = "architecture"
	}
	writer.WriteStart("", "architectures", nil)
	for _, e := range enums {
		writer.WriteCharacter("architecture", string(e))
	}
	writer.WriteEnd("architectures")
	return nil
}

func XMLStorageFormatWriteOne(writer *XMLWriter, enum StorageFormat, tag string) {
	if tag == "" {
		tag = "storage_format"
	}
	writer.WriteCharacter("storage_format", string(enum))
}

func XMLStorageFormatWriteMany(writer *XMLWriter, enums []StorageFormat, plural, singular string) error {
	if plural == "" {
		plural = "storage_formats"
	}
	if singular == "" {
		singular = "storage_format"
	}
	writer.WriteStart("", "storage_formats", nil)
	for _, e := range enums {
		writer.WriteCharacter("storage_format", string(e))
	}
	writer.WriteEnd("storage_formats")
	return nil
}

func XMLStorageDomainStatusWriteOne(writer *XMLWriter, enum StorageDomainStatus, tag string) {
	if tag == "" {
		tag = "storage_domain_status"
	}
	writer.WriteCharacter("storage_domain_status", string(enum))
}

func XMLStorageDomainStatusWriteMany(writer *XMLWriter, enums []StorageDomainStatus, plural, singular string) error {
	if plural == "" {
		plural = "storage_domain_statuss"
	}
	if singular == "" {
		singular = "storage_domain_status"
	}
	writer.WriteStart("", "storage_domain_statuss", nil)
	for _, e := range enums {
		writer.WriteCharacter("storage_domain_status", string(e))
	}
	writer.WriteEnd("storage_domain_statuss")
	return nil
}

func XMLOpenStackNetworkProviderTypeWriteOne(writer *XMLWriter, enum OpenStackNetworkProviderType, tag string) {
	if tag == "" {
		tag = "open_stack_network_provider_type"
	}
	writer.WriteCharacter("open_stack_network_provider_type", string(enum))
}

func XMLOpenStackNetworkProviderTypeWriteMany(writer *XMLWriter, enums []OpenStackNetworkProviderType, plural, singular string) error {
	if plural == "" {
		plural = "open_stack_network_provider_types"
	}
	if singular == "" {
		singular = "open_stack_network_provider_type"
	}
	writer.WriteStart("", "open_stack_network_provider_types", nil)
	for _, e := range enums {
		writer.WriteCharacter("open_stack_network_provider_type", string(e))
	}
	writer.WriteEnd("open_stack_network_provider_types")
	return nil
}

func XMLKdumpStatusWriteOne(writer *XMLWriter, enum KdumpStatus, tag string) {
	if tag == "" {
		tag = "kdump_status"
	}
	writer.WriteCharacter("kdump_status", string(enum))
}

func XMLKdumpStatusWriteMany(writer *XMLWriter, enums []KdumpStatus, plural, singular string) error {
	if plural == "" {
		plural = "kdump_statuss"
	}
	if singular == "" {
		singular = "kdump_status"
	}
	writer.WriteStart("", "kdump_statuss", nil)
	for _, e := range enums {
		writer.WriteCharacter("kdump_status", string(e))
	}
	writer.WriteEnd("kdump_statuss")
	return nil
}

func XMLSshAuthenticationMethodWriteOne(writer *XMLWriter, enum SshAuthenticationMethod, tag string) {
	if tag == "" {
		tag = "ssh_authentication_method"
	}
	writer.WriteCharacter("ssh_authentication_method", string(enum))
}

func XMLSshAuthenticationMethodWriteMany(writer *XMLWriter, enums []SshAuthenticationMethod, plural, singular string) error {
	if plural == "" {
		plural = "ssh_authentication_methods"
	}
	if singular == "" {
		singular = "ssh_authentication_method"
	}
	writer.WriteStart("", "ssh_authentication_methods", nil)
	for _, e := range enums {
		writer.WriteCharacter("ssh_authentication_method", string(e))
	}
	writer.WriteEnd("ssh_authentication_methods")
	return nil
}

func XMLVmStatusWriteOne(writer *XMLWriter, enum VmStatus, tag string) {
	if tag == "" {
		tag = "vm_status"
	}
	writer.WriteCharacter("vm_status", string(enum))
}

func XMLVmStatusWriteMany(writer *XMLWriter, enums []VmStatus, plural, singular string) error {
	if plural == "" {
		plural = "vm_statuss"
	}
	if singular == "" {
		singular = "vm_status"
	}
	writer.WriteStart("", "vm_statuss", nil)
	for _, e := range enums {
		writer.WriteCharacter("vm_status", string(e))
	}
	writer.WriteEnd("vm_statuss")
	return nil
}

func XMLOsTypeWriteOne(writer *XMLWriter, enum OsType, tag string) {
	if tag == "" {
		tag = "os_type"
	}
	writer.WriteCharacter("os_type", string(enum))
}

func XMLOsTypeWriteMany(writer *XMLWriter, enums []OsType, plural, singular string) error {
	if plural == "" {
		plural = "os_types"
	}
	if singular == "" {
		singular = "os_type"
	}
	writer.WriteStart("", "os_types", nil)
	for _, e := range enums {
		writer.WriteCharacter("os_type", string(e))
	}
	writer.WriteEnd("os_types")
	return nil
}

func XMLNfsVersionWriteOne(writer *XMLWriter, enum NfsVersion, tag string) {
	if tag == "" {
		tag = "nfs_version"
	}
	writer.WriteCharacter("nfs_version", string(enum))
}

func XMLNfsVersionWriteMany(writer *XMLWriter, enums []NfsVersion, plural, singular string) error {
	if plural == "" {
		plural = "nfs_versions"
	}
	if singular == "" {
		singular = "nfs_version"
	}
	writer.WriteStart("", "nfs_versions", nil)
	for _, e := range enums {
		writer.WriteCharacter("nfs_version", string(e))
	}
	writer.WriteEnd("nfs_versions")
	return nil
}

func XMLGlusterStateWriteOne(writer *XMLWriter, enum GlusterState, tag string) {
	if tag == "" {
		tag = "gluster_state"
	}
	writer.WriteCharacter("gluster_state", string(enum))
}

func XMLGlusterStateWriteMany(writer *XMLWriter, enums []GlusterState, plural, singular string) error {
	if plural == "" {
		plural = "gluster_states"
	}
	if singular == "" {
		singular = "gluster_state"
	}
	writer.WriteStart("", "gluster_states", nil)
	for _, e := range enums {
		writer.WriteCharacter("gluster_state", string(e))
	}
	writer.WriteEnd("gluster_states")
	return nil
}

func XMLIpVersionWriteOne(writer *XMLWriter, enum IpVersion, tag string) {
	if tag == "" {
		tag = "ip_version"
	}
	writer.WriteCharacter("ip_version", string(enum))
}

func XMLIpVersionWriteMany(writer *XMLWriter, enums []IpVersion, plural, singular string) error {
	if plural == "" {
		plural = "ip_versions"
	}
	if singular == "" {
		singular = "ip_version"
	}
	writer.WriteStart("", "ip_versions", nil)
	for _, e := range enums {
		writer.WriteCharacter("ip_version", string(e))
	}
	writer.WriteEnd("ip_versions")
	return nil
}

func XMLLogSeverityWriteOne(writer *XMLWriter, enum LogSeverity, tag string) {
	if tag == "" {
		tag = "log_severity"
	}
	writer.WriteCharacter("log_severity", string(enum))
}

func XMLLogSeverityWriteMany(writer *XMLWriter, enums []LogSeverity, plural, singular string) error {
	if plural == "" {
		plural = "log_severities"
	}
	if singular == "" {
		singular = "log_severity"
	}
	writer.WriteStart("", "log_severities", nil)
	for _, e := range enums {
		writer.WriteCharacter("log_severity", string(e))
	}
	writer.WriteEnd("log_severities")
	return nil
}

func XMLScsiGenericIOWriteOne(writer *XMLWriter, enum ScsiGenericIO, tag string) {
	if tag == "" {
		tag = "scsi_generic_i_o"
	}
	writer.WriteCharacter("scsi_generic_i_o", string(enum))
}

func XMLScsiGenericIOWriteMany(writer *XMLWriter, enums []ScsiGenericIO, plural, singular string) error {
	if plural == "" {
		plural = "scsi_generic_i_os"
	}
	if singular == "" {
		singular = "scsi_generic_i_o"
	}
	writer.WriteStart("", "scsi_generic_i_os", nil)
	for _, e := range enums {
		writer.WriteCharacter("scsi_generic_i_o", string(e))
	}
	writer.WriteEnd("scsi_generic_i_os")
	return nil
}

func XMLJobStatusWriteOne(writer *XMLWriter, enum JobStatus, tag string) {
	if tag == "" {
		tag = "job_status"
	}
	writer.WriteCharacter("job_status", string(enum))
}

func XMLJobStatusWriteMany(writer *XMLWriter, enums []JobStatus, plural, singular string) error {
	if plural == "" {
		plural = "job_statuss"
	}
	if singular == "" {
		singular = "job_status"
	}
	writer.WriteStart("", "job_statuss", nil)
	for _, e := range enums {
		writer.WriteCharacter("job_status", string(e))
	}
	writer.WriteEnd("job_statuss")
	return nil
}

func XMLSnapshotStatusWriteOne(writer *XMLWriter, enum SnapshotStatus, tag string) {
	if tag == "" {
		tag = "snapshot_status"
	}
	writer.WriteCharacter("snapshot_status", string(enum))
}

func XMLSnapshotStatusWriteMany(writer *XMLWriter, enums []SnapshotStatus, plural, singular string) error {
	if plural == "" {
		plural = "snapshot_statuss"
	}
	if singular == "" {
		singular = "snapshot_status"
	}
	writer.WriteStart("", "snapshot_statuss", nil)
	for _, e := range enums {
		writer.WriteCharacter("snapshot_status", string(e))
	}
	writer.WriteEnd("snapshot_statuss")
	return nil
}

func XMLExternalStatusWriteOne(writer *XMLWriter, enum ExternalStatus, tag string) {
	if tag == "" {
		tag = "external_status"
	}
	writer.WriteCharacter("external_status", string(enum))
}

func XMLExternalStatusWriteMany(writer *XMLWriter, enums []ExternalStatus, plural, singular string) error {
	if plural == "" {
		plural = "external_statuss"
	}
	if singular == "" {
		singular = "external_status"
	}
	writer.WriteStart("", "external_statuss", nil)
	for _, e := range enums {
		writer.WriteCharacter("external_status", string(e))
	}
	writer.WriteEnd("external_statuss")
	return nil
}

func XMLSwitchTypeWriteOne(writer *XMLWriter, enum SwitchType, tag string) {
	if tag == "" {
		tag = "switch_type"
	}
	writer.WriteCharacter("switch_type", string(enum))
}

func XMLSwitchTypeWriteMany(writer *XMLWriter, enums []SwitchType, plural, singular string) error {
	if plural == "" {
		plural = "switch_types"
	}
	if singular == "" {
		singular = "switch_type"
	}
	writer.WriteStart("", "switch_types", nil)
	for _, e := range enums {
		writer.WriteCharacter("switch_type", string(e))
	}
	writer.WriteEnd("switch_types")
	return nil
}

func XMLStepEnumWriteOne(writer *XMLWriter, enum StepEnum, tag string) {
	if tag == "" {
		tag = "step_enum"
	}
	writer.WriteCharacter("step_enum", string(enum))
}

func XMLStepEnumWriteMany(writer *XMLWriter, enums []StepEnum, plural, singular string) error {
	if plural == "" {
		plural = "step_enums"
	}
	if singular == "" {
		singular = "step_enum"
	}
	writer.WriteStart("", "step_enums", nil)
	for _, e := range enums {
		writer.WriteCharacter("step_enum", string(e))
	}
	writer.WriteEnd("step_enums")
	return nil
}

func XMLWatchdogModelWriteOne(writer *XMLWriter, enum WatchdogModel, tag string) {
	if tag == "" {
		tag = "watchdog_model"
	}
	writer.WriteCharacter("watchdog_model", string(enum))
}

func XMLWatchdogModelWriteMany(writer *XMLWriter, enums []WatchdogModel, plural, singular string) error {
	if plural == "" {
		plural = "watchdog_models"
	}
	if singular == "" {
		singular = "watchdog_model"
	}
	writer.WriteStart("", "watchdog_models", nil)
	for _, e := range enums {
		writer.WriteCharacter("watchdog_model", string(e))
	}
	writer.WriteEnd("watchdog_models")
	return nil
}

func XMLEntityExternalStatusWriteOne(writer *XMLWriter, enum EntityExternalStatus, tag string) {
	if tag == "" {
		tag = "entity_external_status"
	}
	writer.WriteCharacter("entity_external_status", string(enum))
}

func XMLEntityExternalStatusWriteMany(writer *XMLWriter, enums []EntityExternalStatus, plural, singular string) error {
	if plural == "" {
		plural = "entity_external_statuss"
	}
	if singular == "" {
		singular = "entity_external_status"
	}
	writer.WriteStart("", "entity_external_statuss", nil)
	for _, e := range enums {
		writer.WriteCharacter("entity_external_status", string(e))
	}
	writer.WriteEnd("entity_external_statuss")
	return nil
}

func XMLFenceTypeWriteOne(writer *XMLWriter, enum FenceType, tag string) {
	if tag == "" {
		tag = "fence_type"
	}
	writer.WriteCharacter("fence_type", string(enum))
}

func XMLFenceTypeWriteMany(writer *XMLWriter, enums []FenceType, plural, singular string) error {
	if plural == "" {
		plural = "fence_types"
	}
	if singular == "" {
		singular = "fence_type"
	}
	writer.WriteStart("", "fence_types", nil)
	for _, e := range enums {
		writer.WriteCharacter("fence_type", string(e))
	}
	writer.WriteEnd("fence_types")
	return nil
}

func XMLTemplateStatusWriteOne(writer *XMLWriter, enum TemplateStatus, tag string) {
	if tag == "" {
		tag = "template_status"
	}
	writer.WriteCharacter("template_status", string(enum))
}

func XMLTemplateStatusWriteMany(writer *XMLWriter, enums []TemplateStatus, plural, singular string) error {
	if plural == "" {
		plural = "template_statuss"
	}
	if singular == "" {
		singular = "template_status"
	}
	writer.WriteStart("", "template_statuss", nil)
	for _, e := range enums {
		writer.WriteCharacter("template_status", string(e))
	}
	writer.WriteEnd("template_statuss")
	return nil
}

func XMLDiskFormatWriteOne(writer *XMLWriter, enum DiskFormat, tag string) {
	if tag == "" {
		tag = "disk_format"
	}
	writer.WriteCharacter("disk_format", string(enum))
}

func XMLDiskFormatWriteMany(writer *XMLWriter, enums []DiskFormat, plural, singular string) error {
	if plural == "" {
		plural = "disk_formats"
	}
	if singular == "" {
		singular = "disk_format"
	}
	writer.WriteStart("", "disk_formats", nil)
	for _, e := range enums {
		writer.WriteCharacter("disk_format", string(e))
	}
	writer.WriteEnd("disk_formats")
	return nil
}

func XMLPowerManagementStatusWriteOne(writer *XMLWriter, enum PowerManagementStatus, tag string) {
	if tag == "" {
		tag = "power_management_status"
	}
	writer.WriteCharacter("power_management_status", string(enum))
}

func XMLPowerManagementStatusWriteMany(writer *XMLWriter, enums []PowerManagementStatus, plural, singular string) error {
	if plural == "" {
		plural = "power_management_statuss"
	}
	if singular == "" {
		singular = "power_management_status"
	}
	writer.WriteStart("", "power_management_statuss", nil)
	for _, e := range enums {
		writer.WriteCharacter("power_management_status", string(e))
	}
	writer.WriteEnd("power_management_statuss")
	return nil
}

func XMLDataCenterStatusWriteOne(writer *XMLWriter, enum DataCenterStatus, tag string) {
	if tag == "" {
		tag = "data_center_status"
	}
	writer.WriteCharacter("data_center_status", string(enum))
}

func XMLDataCenterStatusWriteMany(writer *XMLWriter, enums []DataCenterStatus, plural, singular string) error {
	if plural == "" {
		plural = "data_center_statuss"
	}
	if singular == "" {
		singular = "data_center_status"
	}
	writer.WriteStart("", "data_center_statuss", nil)
	for _, e := range enums {
		writer.WriteCharacter("data_center_status", string(e))
	}
	writer.WriteEnd("data_center_statuss")
	return nil
}

func XMLCpuModeWriteOne(writer *XMLWriter, enum CpuMode, tag string) {
	if tag == "" {
		tag = "cpu_mode"
	}
	writer.WriteCharacter("cpu_mode", string(enum))
}

func XMLCpuModeWriteMany(writer *XMLWriter, enums []CpuMode, plural, singular string) error {
	if plural == "" {
		plural = "cpu_modes"
	}
	if singular == "" {
		singular = "cpu_mode"
	}
	writer.WriteStart("", "cpu_modes", nil)
	for _, e := range enums {
		writer.WriteCharacter("cpu_mode", string(e))
	}
	writer.WriteEnd("cpu_modes")
	return nil
}

func XMLVmAffinityWriteOne(writer *XMLWriter, enum VmAffinity, tag string) {
	if tag == "" {
		tag = "vm_affinity"
	}
	writer.WriteCharacter("vm_affinity", string(enum))
}

func XMLVmAffinityWriteMany(writer *XMLWriter, enums []VmAffinity, plural, singular string) error {
	if plural == "" {
		plural = "vm_affinities"
	}
	if singular == "" {
		singular = "vm_affinity"
	}
	writer.WriteStart("", "vm_affinities", nil)
	for _, e := range enums {
		writer.WriteCharacter("vm_affinity", string(e))
	}
	writer.WriteEnd("vm_affinities")
	return nil
}

func XMLDiskTypeWriteOne(writer *XMLWriter, enum DiskType, tag string) {
	if tag == "" {
		tag = "disk_type"
	}
	writer.WriteCharacter("disk_type", string(enum))
}

func XMLDiskTypeWriteMany(writer *XMLWriter, enums []DiskType, plural, singular string) error {
	if plural == "" {
		plural = "disk_types"
	}
	if singular == "" {
		singular = "disk_type"
	}
	writer.WriteStart("", "disk_types", nil)
	for _, e := range enums {
		writer.WriteCharacter("disk_type", string(e))
	}
	writer.WriteEnd("disk_types")
	return nil
}

func XMLDisplayTypeWriteOne(writer *XMLWriter, enum DisplayType, tag string) {
	if tag == "" {
		tag = "display_type"
	}
	writer.WriteCharacter("display_type", string(enum))
}

func XMLDisplayTypeWriteMany(writer *XMLWriter, enums []DisplayType, plural, singular string) error {
	if plural == "" {
		plural = "display_types"
	}
	if singular == "" {
		singular = "display_type"
	}
	writer.WriteStart("", "display_types", nil)
	for _, e := range enums {
		writer.WriteCharacter("display_type", string(e))
	}
	writer.WriteEnd("display_types")
	return nil
}

func XMLGlusterVolumeTypeWriteOne(writer *XMLWriter, enum GlusterVolumeType, tag string) {
	if tag == "" {
		tag = "gluster_volume_type"
	}
	writer.WriteCharacter("gluster_volume_type", string(enum))
}

func XMLGlusterVolumeTypeWriteMany(writer *XMLWriter, enums []GlusterVolumeType, plural, singular string) error {
	if plural == "" {
		plural = "gluster_volume_types"
	}
	if singular == "" {
		singular = "gluster_volume_type"
	}
	writer.WriteStart("", "gluster_volume_types", nil)
	for _, e := range enums {
		writer.WriteCharacter("gluster_volume_type", string(e))
	}
	writer.WriteEnd("gluster_volume_types")
	return nil
}

func XMLStatisticKindWriteOne(writer *XMLWriter, enum StatisticKind, tag string) {
	if tag == "" {
		tag = "statistic_kind"
	}
	writer.WriteCharacter("statistic_kind", string(enum))
}

func XMLStatisticKindWriteMany(writer *XMLWriter, enums []StatisticKind, plural, singular string) error {
	if plural == "" {
		plural = "statistic_kinds"
	}
	if singular == "" {
		singular = "statistic_kind"
	}
	writer.WriteStart("", "statistic_kinds", nil)
	for _, e := range enums {
		writer.WriteCharacter("statistic_kind", string(e))
	}
	writer.WriteEnd("statistic_kinds")
	return nil
}

func XMLNumaTuneModeWriteOne(writer *XMLWriter, enum NumaTuneMode, tag string) {
	if tag == "" {
		tag = "numa_tune_mode"
	}
	writer.WriteCharacter("numa_tune_mode", string(enum))
}

func XMLNumaTuneModeWriteMany(writer *XMLWriter, enums []NumaTuneMode, plural, singular string) error {
	if plural == "" {
		plural = "numa_tune_modes"
	}
	if singular == "" {
		singular = "numa_tune_mode"
	}
	writer.WriteStart("", "numa_tune_modes", nil)
	for _, e := range enums {
		writer.WriteCharacter("numa_tune_mode", string(e))
	}
	writer.WriteEnd("numa_tune_modes")
	return nil
}

func XMLNetworkUsageWriteOne(writer *XMLWriter, enum NetworkUsage, tag string) {
	if tag == "" {
		tag = "network_usage"
	}
	writer.WriteCharacter("network_usage", string(enum))
}

func XMLNetworkUsageWriteMany(writer *XMLWriter, enums []NetworkUsage, plural, singular string) error {
	if plural == "" {
		plural = "network_usages"
	}
	if singular == "" {
		singular = "network_usage"
	}
	writer.WriteStart("", "network_usages", nil)
	for _, e := range enums {
		writer.WriteCharacter("network_usage", string(e))
	}
	writer.WriteEnd("network_usages")
	return nil
}

func XMLOpenstackVolumeAuthenticationKeyUsageTypeWriteOne(writer *XMLWriter, enum OpenstackVolumeAuthenticationKeyUsageType, tag string) {
	if tag == "" {
		tag = "openstack_volume_authentication_key_usage_type"
	}
	writer.WriteCharacter("openstack_volume_authentication_key_usage_type", string(enum))
}

func XMLOpenstackVolumeAuthenticationKeyUsageTypeWriteMany(writer *XMLWriter, enums []OpenstackVolumeAuthenticationKeyUsageType, plural, singular string) error {
	if plural == "" {
		plural = "openstack_volume_authentication_key_usage_types"
	}
	if singular == "" {
		singular = "openstack_volume_authentication_key_usage_type"
	}
	writer.WriteStart("", "openstack_volume_authentication_key_usage_types", nil)
	for _, e := range enums {
		writer.WriteCharacter("openstack_volume_authentication_key_usage_type", string(e))
	}
	writer.WriteEnd("openstack_volume_authentication_key_usage_types")
	return nil
}

func XMLSsoMethodWriteOne(writer *XMLWriter, enum SsoMethod, tag string) {
	if tag == "" {
		tag = "sso_method"
	}
	writer.WriteCharacter("sso_method", string(enum))
}

func XMLSsoMethodWriteMany(writer *XMLWriter, enums []SsoMethod, plural, singular string) error {
	if plural == "" {
		plural = "sso_methods"
	}
	if singular == "" {
		singular = "sso_method"
	}
	writer.WriteStart("", "sso_methods", nil)
	for _, e := range enums {
		writer.WriteCharacter("sso_method", string(e))
	}
	writer.WriteEnd("sso_methods")
	return nil
}

func XMLHostTypeWriteOne(writer *XMLWriter, enum HostType, tag string) {
	if tag == "" {
		tag = "host_type"
	}
	writer.WriteCharacter("host_type", string(enum))
}

func XMLHostTypeWriteMany(writer *XMLWriter, enums []HostType, plural, singular string) error {
	if plural == "" {
		plural = "host_types"
	}
	if singular == "" {
		singular = "host_type"
	}
	writer.WriteStart("", "host_types", nil)
	for _, e := range enums {
		writer.WriteCharacter("host_type", string(e))
	}
	writer.WriteEnd("host_types")
	return nil
}

func XMLExternalSystemTypeWriteOne(writer *XMLWriter, enum ExternalSystemType, tag string) {
	if tag == "" {
		tag = "external_system_type"
	}
	writer.WriteCharacter("external_system_type", string(enum))
}

func XMLExternalSystemTypeWriteMany(writer *XMLWriter, enums []ExternalSystemType, plural, singular string) error {
	if plural == "" {
		plural = "external_system_types"
	}
	if singular == "" {
		singular = "external_system_type"
	}
	writer.WriteStart("", "external_system_types", nil)
	for _, e := range enums {
		writer.WriteCharacter("external_system_type", string(e))
	}
	writer.WriteEnd("external_system_types")
	return nil
}

func XMLNetworkPluginTypeWriteOne(writer *XMLWriter, enum NetworkPluginType, tag string) {
	if tag == "" {
		tag = "network_plugin_type"
	}
	writer.WriteCharacter("network_plugin_type", string(enum))
}

func XMLNetworkPluginTypeWriteMany(writer *XMLWriter, enums []NetworkPluginType, plural, singular string) error {
	if plural == "" {
		plural = "network_plugin_types"
	}
	if singular == "" {
		singular = "network_plugin_type"
	}
	writer.WriteStart("", "network_plugin_types", nil)
	for _, e := range enums {
		writer.WriteCharacter("network_plugin_type", string(e))
	}
	writer.WriteEnd("network_plugin_types")
	return nil
}

func XMLInheritableBooleanWriteOne(writer *XMLWriter, enum InheritableBoolean, tag string) {
	if tag == "" {
		tag = "inheritable_boolean"
	}
	writer.WriteCharacter("inheritable_boolean", string(enum))
}

func XMLInheritableBooleanWriteMany(writer *XMLWriter, enums []InheritableBoolean, plural, singular string) error {
	if plural == "" {
		plural = "inheritable_booleans"
	}
	if singular == "" {
		singular = "inheritable_boolean"
	}
	writer.WriteStart("", "inheritable_booleans", nil)
	for _, e := range enums {
		writer.WriteCharacter("inheritable_boolean", string(e))
	}
	writer.WriteEnd("inheritable_booleans")
	return nil
}

func XMLCreationStatusWriteOne(writer *XMLWriter, enum CreationStatus, tag string) {
	if tag == "" {
		tag = "creation_status"
	}
	writer.WriteCharacter("creation_status", string(enum))
}

func XMLCreationStatusWriteMany(writer *XMLWriter, enums []CreationStatus, plural, singular string) error {
	if plural == "" {
		plural = "creation_statuss"
	}
	if singular == "" {
		singular = "creation_status"
	}
	writer.WriteStart("", "creation_statuss", nil)
	for _, e := range enums {
		writer.WriteCharacter("creation_status", string(e))
	}
	writer.WriteEnd("creation_statuss")
	return nil
}

func XMLHookContentTypeWriteOne(writer *XMLWriter, enum HookContentType, tag string) {
	if tag == "" {
		tag = "hook_content_type"
	}
	writer.WriteCharacter("hook_content_type", string(enum))
}

func XMLHookContentTypeWriteMany(writer *XMLWriter, enums []HookContentType, plural, singular string) error {
	if plural == "" {
		plural = "hook_content_types"
	}
	if singular == "" {
		singular = "hook_content_type"
	}
	writer.WriteStart("", "hook_content_types", nil)
	for _, e := range enums {
		writer.WriteCharacter("hook_content_type", string(e))
	}
	writer.WriteEnd("hook_content_types")
	return nil
}

func XMLValueTypeWriteOne(writer *XMLWriter, enum ValueType, tag string) {
	if tag == "" {
		tag = "value_type"
	}
	writer.WriteCharacter("value_type", string(enum))
}

func XMLValueTypeWriteMany(writer *XMLWriter, enums []ValueType, plural, singular string) error {
	if plural == "" {
		plural = "value_types"
	}
	if singular == "" {
		singular = "value_type"
	}
	writer.WriteStart("", "value_types", nil)
	for _, e := range enums {
		writer.WriteCharacter("value_type", string(e))
	}
	writer.WriteEnd("value_types")
	return nil
}

func XMLGraphicsTypeWriteOne(writer *XMLWriter, enum GraphicsType, tag string) {
	if tag == "" {
		tag = "graphics_type"
	}
	writer.WriteCharacter("graphics_type", string(enum))
}

func XMLGraphicsTypeWriteMany(writer *XMLWriter, enums []GraphicsType, plural, singular string) error {
	if plural == "" {
		plural = "graphics_types"
	}
	if singular == "" {
		singular = "graphics_type"
	}
	writer.WriteStart("", "graphics_types", nil)
	for _, e := range enums {
		writer.WriteCharacter("graphics_type", string(e))
	}
	writer.WriteEnd("graphics_types")
	return nil
}

func XMLNicStatusWriteOne(writer *XMLWriter, enum NicStatus, tag string) {
	if tag == "" {
		tag = "nic_status"
	}
	writer.WriteCharacter("nic_status", string(enum))
}

func XMLNicStatusWriteMany(writer *XMLWriter, enums []NicStatus, plural, singular string) error {
	if plural == "" {
		plural = "nic_statuss"
	}
	if singular == "" {
		singular = "nic_status"
	}
	writer.WriteStart("", "nic_statuss", nil)
	for _, e := range enums {
		writer.WriteCharacter("nic_status", string(e))
	}
	writer.WriteEnd("nic_statuss")
	return nil
}

func XMLImageTransferDirectionWriteOne(writer *XMLWriter, enum ImageTransferDirection, tag string) {
	if tag == "" {
		tag = "image_transfer_direction"
	}
	writer.WriteCharacter("image_transfer_direction", string(enum))
}

func XMLImageTransferDirectionWriteMany(writer *XMLWriter, enums []ImageTransferDirection, plural, singular string) error {
	if plural == "" {
		plural = "image_transfer_directions"
	}
	if singular == "" {
		singular = "image_transfer_direction"
	}
	writer.WriteStart("", "image_transfer_directions", nil)
	for _, e := range enums {
		writer.WriteCharacter("image_transfer_direction", string(e))
	}
	writer.WriteEnd("image_transfer_directions")
	return nil
}

func XMLTransportTypeWriteOne(writer *XMLWriter, enum TransportType, tag string) {
	if tag == "" {
		tag = "transport_type"
	}
	writer.WriteCharacter("transport_type", string(enum))
}

func XMLTransportTypeWriteMany(writer *XMLWriter, enums []TransportType, plural, singular string) error {
	if plural == "" {
		plural = "transport_types"
	}
	if singular == "" {
		singular = "transport_type"
	}
	writer.WriteStart("", "transport_types", nil)
	for _, e := range enums {
		writer.WriteCharacter("transport_type", string(e))
	}
	writer.WriteEnd("transport_types")
	return nil
}

func XMLWatchdogActionWriteOne(writer *XMLWriter, enum WatchdogAction, tag string) {
	if tag == "" {
		tag = "watchdog_action"
	}
	writer.WriteCharacter("watchdog_action", string(enum))
}

func XMLWatchdogActionWriteMany(writer *XMLWriter, enums []WatchdogAction, plural, singular string) error {
	if plural == "" {
		plural = "watchdog_actions"
	}
	if singular == "" {
		singular = "watchdog_action"
	}
	writer.WriteStart("", "watchdog_actions", nil)
	for _, e := range enums {
		writer.WriteCharacter("watchdog_action", string(e))
	}
	writer.WriteEnd("watchdog_actions")
	return nil
}

func XMLMigrationBandwidthAssignmentMethodWriteOne(writer *XMLWriter, enum MigrationBandwidthAssignmentMethod, tag string) {
	if tag == "" {
		tag = "migration_bandwidth_assignment_method"
	}
	writer.WriteCharacter("migration_bandwidth_assignment_method", string(enum))
}

func XMLMigrationBandwidthAssignmentMethodWriteMany(writer *XMLWriter, enums []MigrationBandwidthAssignmentMethod, plural, singular string) error {
	if plural == "" {
		plural = "migration_bandwidth_assignment_methods"
	}
	if singular == "" {
		singular = "migration_bandwidth_assignment_method"
	}
	writer.WriteStart("", "migration_bandwidth_assignment_methods", nil)
	for _, e := range enums {
		writer.WriteCharacter("migration_bandwidth_assignment_method", string(e))
	}
	writer.WriteEnd("migration_bandwidth_assignment_methods")
	return nil
}

func XMLHostStatusWriteOne(writer *XMLWriter, enum HostStatus, tag string) {
	if tag == "" {
		tag = "host_status"
	}
	writer.WriteCharacter("host_status", string(enum))
}

func XMLHostStatusWriteMany(writer *XMLWriter, enums []HostStatus, plural, singular string) error {
	if plural == "" {
		plural = "host_statuss"
	}
	if singular == "" {
		singular = "host_status"
	}
	writer.WriteStart("", "host_statuss", nil)
	for _, e := range enums {
		writer.WriteCharacter("host_status", string(e))
	}
	writer.WriteEnd("host_statuss")
	return nil
}

func XMLRngSourceWriteOne(writer *XMLWriter, enum RngSource, tag string) {
	if tag == "" {
		tag = "rng_source"
	}
	writer.WriteCharacter("rng_source", string(enum))
}

func XMLRngSourceWriteMany(writer *XMLWriter, enums []RngSource, plural, singular string) error {
	if plural == "" {
		plural = "rng_sources"
	}
	if singular == "" {
		singular = "rng_source"
	}
	writer.WriteStart("", "rng_sources", nil)
	for _, e := range enums {
		writer.WriteCharacter("rng_source", string(e))
	}
	writer.WriteEnd("rng_sources")
	return nil
}

func XMLSnapshotTypeWriteOne(writer *XMLWriter, enum SnapshotType, tag string) {
	if tag == "" {
		tag = "snapshot_type"
	}
	writer.WriteCharacter("snapshot_type", string(enum))
}

func XMLSnapshotTypeWriteMany(writer *XMLWriter, enums []SnapshotType, plural, singular string) error {
	if plural == "" {
		plural = "snapshot_types"
	}
	if singular == "" {
		singular = "snapshot_type"
	}
	writer.WriteStart("", "snapshot_types", nil)
	for _, e := range enums {
		writer.WriteCharacter("snapshot_type", string(e))
	}
	writer.WriteEnd("snapshot_types")
	return nil
}
