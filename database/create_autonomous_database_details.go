// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// CreateAutonomousDatabaseDetails Details to create an Oracle Autonomous Database.
type CreateAutonomousDatabaseDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment of the autonomous database.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The database name. The name must begin with an alphabetic character and can contain a maximum of 14 alphanumeric characters. Special characters are not permitted. The database name must be unique in the tenancy.
	DbName *string `mandatory:"true" json:"dbName"`

	// The number of CPU Cores to be made available to the database.
	CpuCoreCount *int `mandatory:"true" json:"cpuCoreCount"`

	// The size, in terabytes, of the data volume that will be created and attached to the database. This storage can later be scaled up if needed.
	DataStorageSizeInTBs *int `mandatory:"true" json:"dataStorageSizeInTBs"`

	// The password must be between 12 and 30 characters long, and must contain at least 1 uppercase, 1 lowercase, and 1 numeric character. It cannot contain the double quote symbol (") or the username "admin", regardless of casing.
	AdminPassword *string `mandatory:"true" json:"adminPassword"`

	// Indicates if this is an Always Free resource. The default value is false. Note that Always Free Autonomous Databases have 1 CPU and 20GB memory. For Always Free databases, memory and CPU cannot be scaled.
	IsFreeTier *bool `mandatory:"false" json:"isFreeTier"`

	// The user-friendly name for the Autonomous Database. The name does not have to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// If set to true, indicates that an Autonomous Database preview version is being provisioned, and that the preview version's terms of service have been accepted. Note that preview version software is only available for serverless deployments (https://docs.cloud.oracle.com/Content/Database/Concepts/adboverview.htm#AEI).
	IsPreviewVersionWithServiceTermsAccepted *bool `mandatory:"false" json:"isPreviewVersionWithServiceTermsAccepted"`

	// Indicates if auto scaling is enabled for the Autonomous Database CPU core count. The default value is false. Note that auto scaling is available for serverless deployments (https://docs.cloud.oracle.com/Content/Database/Concepts/adboverview.htm#AEI) only.
	IsAutoScalingEnabled *bool `mandatory:"false" json:"isAutoScalingEnabled"`

	// True if the database uses the dedicated deployment (https://docs.cloud.oracle.com/Content/Database/Concepts/adbddoverview.htm) option.
	IsDedicated *bool `mandatory:"false" json:"isDedicated"`

	// The Autonomous Container Database OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	AutonomousContainerDatabaseId *string `mandatory:"false" json:"autonomousContainerDatabaseId"`

	// The client IP access control list (ACL). This feature is available for serverless deployments (https://docs.cloud.oracle.com/Content/Database/Concepts/adboverview.htm#AEI) only.
	// Only clients connecting from an IP address included in the ACL may access the Autonomous Database instance. This is an array of CIDR (Classless Inter-Domain Routing) notations for a subnet or VCN OCID.
	// To add the whitelist VCN specific subnet or IP, use a semicoln ';' as a deliminator to add the VCN specific subnets or IPs.
	// Example: `["1.1.1.1","1.1.1.0/24","ocid1.vcn.oc1.sea.aaaaaaaard2hfx2nn3e5xeo6j6o62jga44xjizkw","ocid1.vcn.oc1.sea.aaaaaaaard2hfx2nn3e5xeo6j6o62jga44xjizkw;1.1.1.1","ocid1.vcn.oc1.sea.aaaaaaaard2hfx2nn3e5xeo6j6o62jga44xjizkw;1.1.0.0/16"]`
	WhitelistedIps []string `mandatory:"false" json:"whitelistedIps"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the subnet the resource is associated with.
	// **Subnet Restrictions:**
	// - For bare metal DB systems and for single node virtual machine DB systems, do not use a subnet that overlaps with 192.168.16.16/28.
	// - For Exadata and virtual machine 2-node RAC DB systems, do not use a subnet that overlaps with 192.168.128.0/20.
	// - For Autonomous Database, setting this will disable public secure access to the database.
	// These subnets are used by the Oracle Clusterware private interconnect on the database instance.
	// Specifying an overlapping subnet will cause the private interconnect to malfunction.
	// This restriction applies to both the client subnet and the backup subnet.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// A list of the OCIDs (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that this DB system belongs to. Setting this to an empty array after the list is created removes the resource from all NSGs. For more information about NSGs, see Security Rules (https://docs.cloud.oracle.com/Content/Network/Concepts/securityrules.htm).
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// The private endpoint label for the resource.
	PrivateEndpointLabel *string `mandatory:"false" json:"privateEndpointLabel"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The autonomous database workload type. OLTP indicates an Autonomous Transaction Processing database and DW indicates an Autonomous Data Warehouse. The default is OLTP.
	DbWorkload CreateAutonomousDatabaseBaseDbWorkloadEnum `mandatory:"false" json:"dbWorkload,omitempty"`

	// The Oracle license model that applies to the Oracle Autonomous Database. The default for Autonomous Database using the shared deployment is BRING_YOUR_OWN_LICENSE. Note that when provisioning an Autonomous Database using the dedicated deployment (https://docs.cloud.oracle.com/Content/Database/Concepts/adbddoverview.htm) option, this attribute must be null because the attribute is already set on Autonomous Exadata Infrastructure level.
	LicenseModel CreateAutonomousDatabaseBaseLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`
}

//GetCompartmentId returns CompartmentId
func (m CreateAutonomousDatabaseDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetDbName returns DbName
func (m CreateAutonomousDatabaseDetails) GetDbName() *string {
	return m.DbName
}

//GetCpuCoreCount returns CpuCoreCount
func (m CreateAutonomousDatabaseDetails) GetCpuCoreCount() *int {
	return m.CpuCoreCount
}

//GetDbWorkload returns DbWorkload
func (m CreateAutonomousDatabaseDetails) GetDbWorkload() CreateAutonomousDatabaseBaseDbWorkloadEnum {
	return m.DbWorkload
}

//GetDataStorageSizeInTBs returns DataStorageSizeInTBs
func (m CreateAutonomousDatabaseDetails) GetDataStorageSizeInTBs() *int {
	return m.DataStorageSizeInTBs
}

//GetIsFreeTier returns IsFreeTier
func (m CreateAutonomousDatabaseDetails) GetIsFreeTier() *bool {
	return m.IsFreeTier
}

//GetAdminPassword returns AdminPassword
func (m CreateAutonomousDatabaseDetails) GetAdminPassword() *string {
	return m.AdminPassword
}

//GetDisplayName returns DisplayName
func (m CreateAutonomousDatabaseDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetLicenseModel returns LicenseModel
func (m CreateAutonomousDatabaseDetails) GetLicenseModel() CreateAutonomousDatabaseBaseLicenseModelEnum {
	return m.LicenseModel
}

//GetIsPreviewVersionWithServiceTermsAccepted returns IsPreviewVersionWithServiceTermsAccepted
func (m CreateAutonomousDatabaseDetails) GetIsPreviewVersionWithServiceTermsAccepted() *bool {
	return m.IsPreviewVersionWithServiceTermsAccepted
}

//GetIsAutoScalingEnabled returns IsAutoScalingEnabled
func (m CreateAutonomousDatabaseDetails) GetIsAutoScalingEnabled() *bool {
	return m.IsAutoScalingEnabled
}

//GetIsDedicated returns IsDedicated
func (m CreateAutonomousDatabaseDetails) GetIsDedicated() *bool {
	return m.IsDedicated
}

//GetAutonomousContainerDatabaseId returns AutonomousContainerDatabaseId
func (m CreateAutonomousDatabaseDetails) GetAutonomousContainerDatabaseId() *string {
	return m.AutonomousContainerDatabaseId
}

//GetWhitelistedIps returns WhitelistedIps
func (m CreateAutonomousDatabaseDetails) GetWhitelistedIps() []string {
	return m.WhitelistedIps
}

//GetSubnetId returns SubnetId
func (m CreateAutonomousDatabaseDetails) GetSubnetId() *string {
	return m.SubnetId
}

//GetNsgIds returns NsgIds
func (m CreateAutonomousDatabaseDetails) GetNsgIds() []string {
	return m.NsgIds
}

//GetPrivateEndpointLabel returns PrivateEndpointLabel
func (m CreateAutonomousDatabaseDetails) GetPrivateEndpointLabel() *string {
	return m.PrivateEndpointLabel
}

//GetFreeformTags returns FreeformTags
func (m CreateAutonomousDatabaseDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m CreateAutonomousDatabaseDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateAutonomousDatabaseDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateAutonomousDatabaseDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateAutonomousDatabaseDetails CreateAutonomousDatabaseDetails
	s := struct {
		DiscriminatorParam string `json:"source"`
		MarshalTypeCreateAutonomousDatabaseDetails
	}{
		"NONE",
		(MarshalTypeCreateAutonomousDatabaseDetails)(m),
	}

	return json.Marshal(&s)
}
