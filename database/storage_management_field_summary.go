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

// StorageManagementFieldSummary The DB system storage management field summary.
type StorageManagementFieldSummary struct {

	// The DB system's storage management option name.
	Name StorageManagementFieldSummaryNameEnum `mandatory:"true" json:"name"`

	// True if the enumerated storage management option (LVM or ASM) is allowed.
	IsAllowed *bool `mandatory:"true" json:"isAllowed"`
}

func (m StorageManagementFieldSummary) String() string {
	return common.PointerString(m)
}

// StorageManagementFieldSummaryNameEnum Enum with underlying type: string
type StorageManagementFieldSummaryNameEnum string

// Set of constants representing the allowable values for StorageManagementFieldSummaryNameEnum
const (
	StorageManagementFieldSummaryNameAsm StorageManagementFieldSummaryNameEnum = "ASM"
	StorageManagementFieldSummaryNameLvm StorageManagementFieldSummaryNameEnum = "LVM"
)

var mappingStorageManagementFieldSummaryName = map[string]StorageManagementFieldSummaryNameEnum{
	"ASM": StorageManagementFieldSummaryNameAsm,
	"LVM": StorageManagementFieldSummaryNameLvm,
}

// GetStorageManagementFieldSummaryNameEnumValues Enumerates the set of values for StorageManagementFieldSummaryNameEnum
func GetStorageManagementFieldSummaryNameEnumValues() []StorageManagementFieldSummaryNameEnum {
	values := make([]StorageManagementFieldSummaryNameEnum, 0)
	for _, v := range mappingStorageManagementFieldSummaryName {
		values = append(values, v)
	}
	return values
}
