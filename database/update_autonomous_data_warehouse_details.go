// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateAutonomousDataWarehouseDetails Details to update an Oracle Autonomous Data Warehouse.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type UpdateAutonomousDataWarehouseDetails struct {

	// The number of CPU cores to be made available to the database.
	CpuCoreCount *int `mandatory:"false" json:"cpuCoreCount"`

	// Size, in terabytes, of the data volume that will be attached to the database.
	DataStorageSizeInTBs *int `mandatory:"false" json:"dataStorageSizeInTBs"`

	// The user-friendly name for the Autonomous Data Warehouse. The name does not have to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The password must be between 12 and 30 characters long, and must contain at least 1 uppercase, 1 lowercase, and 1 numeric character. It cannot contain the double quote symbol (") or the username "admin", regardless of casing. It must be different from the last four passwords and it must not be a password used within the last 24 hours.
	AdminPassword *string `mandatory:"false" json:"adminPassword"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The new Oracle license model that applies to the Oracle Autonomous Data Warehouse.
	LicenseModel UpdateAutonomousDataWarehouseDetailsLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`
}

func (m UpdateAutonomousDataWarehouseDetails) String() string {
	return common.PointerString(m)
}

// UpdateAutonomousDataWarehouseDetailsLicenseModelEnum Enum with underlying type: string
type UpdateAutonomousDataWarehouseDetailsLicenseModelEnum string

// Set of constants representing the allowable values for UpdateAutonomousDataWarehouseDetailsLicenseModelEnum
const (
	UpdateAutonomousDataWarehouseDetailsLicenseModelLicenseIncluded     UpdateAutonomousDataWarehouseDetailsLicenseModelEnum = "LICENSE_INCLUDED"
	UpdateAutonomousDataWarehouseDetailsLicenseModelBringYourOwnLicense UpdateAutonomousDataWarehouseDetailsLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingUpdateAutonomousDataWarehouseDetailsLicenseModel = map[string]UpdateAutonomousDataWarehouseDetailsLicenseModelEnum{
	"LICENSE_INCLUDED":       UpdateAutonomousDataWarehouseDetailsLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": UpdateAutonomousDataWarehouseDetailsLicenseModelBringYourOwnLicense,
}

// GetUpdateAutonomousDataWarehouseDetailsLicenseModelEnumValues Enumerates the set of values for UpdateAutonomousDataWarehouseDetailsLicenseModelEnum
func GetUpdateAutonomousDataWarehouseDetailsLicenseModelEnumValues() []UpdateAutonomousDataWarehouseDetailsLicenseModelEnum {
	values := make([]UpdateAutonomousDataWarehouseDetailsLicenseModelEnum, 0)
	for _, v := range mappingUpdateAutonomousDataWarehouseDetailsLicenseModel {
		values = append(values, v)
	}
	return values
}
