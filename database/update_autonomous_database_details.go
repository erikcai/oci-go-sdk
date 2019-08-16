// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateAutonomousDatabaseDetails Details to update an Oracle Autonomous Database.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type UpdateAutonomousDatabaseDetails struct {

	// The number of CPU cores to be made available to the database.
	CpuCoreCount *int `mandatory:"false" json:"cpuCoreCount"`

	// The size, in terabytes, of the data volume that will be attached to the database.
	DataStorageSizeInTBs *int `mandatory:"false" json:"dataStorageSizeInTBs"`

	// The user-friendly name for the Autonomous Database. The name does not have to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Indicates if this is an Always Free resource. The default value is false. Note that Always Free Autonomous Databases have 1 CPU and 20GB memory. For Always Free databases, memory and CPU cannot be scaled.
	IsFreeTier *bool `mandatory:"false" json:"isFreeTier"`

	// The password must be between 12 and 30 characters long, and must contain at least 1 uppercase, 1 lowercase, and 1 numeric character. It cannot contain the double quote symbol (") or the username "admin", regardless of casing. It must be different from the last four passwords and it must not be a password used within the last 24 hours.
	AdminPassword *string `mandatory:"false" json:"adminPassword"`

	// New name for this Autonomous Database. It must begin with an alphabetic character and can contain a
	// maximum of eight alphanumeric characters. Special characters are not permitted. This is valid only
	// for dedicated databases.
	DbName *string `mandatory:"false" json:"dbName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The Oracle license model that applies to the Oracle Autonomous Database. The default for Autonomous Database using the shared deployment is BRING_YOUR_OWN_LICENSE. Note that when provisioning an Autonomous Database using the dedicated deployment (https://docs.cloud.oracle.com/Content/Database/Concepts/adbddoverview.htm) option, this attribute must be null because the attribute is already set on Autonomous Exadata Infrastructure level.
	LicenseModel UpdateAutonomousDatabaseDetailsLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// The client IP access control list (ACL). This feature is available for serverless deployments (https://docs.cloud.oracle.com/Content/Database/Concepts/adboverview.htm#AEI) only.
	// Only clients connecting from an IP address included in the ACL may access the Autonomous Database instance. This is an array of CIDR (Classless Inter-Domain Routing) notations for a subnet or VCN OCID. To delete all the existing white listed IP’s, use an array with a single empty string entry.
	// To add the whitelist VCN specific subnet or IP, use a semicoln ';' as a deliminator to add the VCN specific subnets or IPs.
	// Example: `["1.1.1.1","1.1.1.0/24","ocid1.vcn.oc1.sea.aaaaaaaard2hfx2nn3e5xeo6j6o62jga44xjizkw","ocid1.vcn.oc1.sea.aaaaaaaard2hfx2nn3e5xeo6j6o62jga44xjizkw;1.1.1.1","ocid1.vcn.oc1.sea.aaaaaaaard2hfx2nn3e5xeo6j6o62jga44xjizkw"]`
	WhitelistedIps []string `mandatory:"false" json:"whitelistedIps"`

	// Indicates if auto scaling is enabled for the Autonomous Database CPU core count. The default value is false. Note that auto scaling is available for serverless deployments (https://docs.cloud.oracle.com/Content/Database/Concepts/adboverview.htm#AEI) only.
	IsAutoScalingEnabled *bool `mandatory:"false" json:"isAutoScalingEnabled"`
}

func (m UpdateAutonomousDatabaseDetails) String() string {
	return common.PointerString(m)
}

// UpdateAutonomousDatabaseDetailsLicenseModelEnum Enum with underlying type: string
type UpdateAutonomousDatabaseDetailsLicenseModelEnum string

// Set of constants representing the allowable values for UpdateAutonomousDatabaseDetailsLicenseModelEnum
const (
	UpdateAutonomousDatabaseDetailsLicenseModelLicenseIncluded     UpdateAutonomousDatabaseDetailsLicenseModelEnum = "LICENSE_INCLUDED"
	UpdateAutonomousDatabaseDetailsLicenseModelBringYourOwnLicense UpdateAutonomousDatabaseDetailsLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingUpdateAutonomousDatabaseDetailsLicenseModel = map[string]UpdateAutonomousDatabaseDetailsLicenseModelEnum{
	"LICENSE_INCLUDED":       UpdateAutonomousDatabaseDetailsLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": UpdateAutonomousDatabaseDetailsLicenseModelBringYourOwnLicense,
}

// GetUpdateAutonomousDatabaseDetailsLicenseModelEnumValues Enumerates the set of values for UpdateAutonomousDatabaseDetailsLicenseModelEnum
func GetUpdateAutonomousDatabaseDetailsLicenseModelEnumValues() []UpdateAutonomousDatabaseDetailsLicenseModelEnum {
	values := make([]UpdateAutonomousDatabaseDetailsLicenseModelEnum, 0)
	for _, v := range mappingUpdateAutonomousDatabaseDetailsLicenseModel {
		values = append(values, v)
	}
	return values
}
