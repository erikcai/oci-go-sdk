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

// CreateAutonomousDataWarehouseDetails Details to create an Oracle Autonomous Data Warehouse.
type CreateAutonomousDataWarehouseDetails struct {

	// A strong password for Admin. The password must be between 8 and 30 characters long and contain at least 1 uppercase, 1 lowercase, 1 number, and 1 special character. The special characters must be _, \#, or -.
	AdminPassword *string `mandatory:"true" json:"adminPassword"`

	// The Oracle Cloud ID (OCID) of the compartment of the DB system.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The number of CPU Cores to be made available to the database.
	CpuCoreCount *int `mandatory:"true" json:"cpuCoreCount"`

	// Size, in terabytes, of the data volume that will be created and attached to the database. This storage can later be scaled up if needed.
	DataStorageSizeInTBs *int `mandatory:"true" json:"dataStorageSizeInTBs"`

	// The database name. The name must begin with an alphabetic character and can contain a maximum of 14 alphanumeric characters. Special characters are not permitted. The database name must be unique in the tenancy.
	DbName *string `mandatory:"true" json:"dbName"`

	// The user-friendly name for the Autonomous Data Warehouse. The name does not have to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The Oracle license model that applies to the Oracle Autonomous Data Warehouse. The default is LICENSE_INCLUDED.
	LicenseModel CreateAutonomousDataWarehouseDetailsLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`
}

func (m CreateAutonomousDataWarehouseDetails) String() string {
	return common.PointerString(m)
}

// CreateAutonomousDataWarehouseDetailsLicenseModelEnum Enum with underlying type: string
type CreateAutonomousDataWarehouseDetailsLicenseModelEnum string

// Set of constants representing the allowable values for CreateAutonomousDataWarehouseDetailsLicenseModel
const (
	CreateAutonomousDataWarehouseDetailsLicenseModelLicenseIncluded     CreateAutonomousDataWarehouseDetailsLicenseModelEnum = "LICENSE_INCLUDED"
	CreateAutonomousDataWarehouseDetailsLicenseModelBringYourOwnLicense CreateAutonomousDataWarehouseDetailsLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingCreateAutonomousDataWarehouseDetailsLicenseModel = map[string]CreateAutonomousDataWarehouseDetailsLicenseModelEnum{
	"LICENSE_INCLUDED":       CreateAutonomousDataWarehouseDetailsLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": CreateAutonomousDataWarehouseDetailsLicenseModelBringYourOwnLicense,
}

// GetCreateAutonomousDataWarehouseDetailsLicenseModelEnumValues Enumerates the set of values for CreateAutonomousDataWarehouseDetailsLicenseModel
func GetCreateAutonomousDataWarehouseDetailsLicenseModelEnumValues() []CreateAutonomousDataWarehouseDetailsLicenseModelEnum {
	values := make([]CreateAutonomousDataWarehouseDetailsLicenseModelEnum, 0)
	for _, v := range mappingCreateAutonomousDataWarehouseDetailsLicenseModel {
		values = append(values, v)
	}
	return values
}
