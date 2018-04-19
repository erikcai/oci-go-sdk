// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateBootVolumeBackupDetails The representation of CreateBootVolumeBackupDetails
type CreateBootVolumeBackupDetails struct {

	// The OCID of the boot volume that needs to be backed up.
	BootVolumeId *string `mandatory:"true" json:"bootVolumeId"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "foo-value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name for the boot volume backup. Does not have to be unique and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The type of backup to create. If omitted, defaults to incremental.
	Type CreateBootVolumeBackupDetailsTypeEnum `mandatory:"false" json:"type,omitempty"`
}

func (m CreateBootVolumeBackupDetails) String() string {
	return common.PointerString(m)
}

// CreateBootVolumeBackupDetailsTypeEnum Enum with underlying type: string
type CreateBootVolumeBackupDetailsTypeEnum string

// Set of constants representing the allowable values for CreateBootVolumeBackupDetailsType
const (
	CreateBootVolumeBackupDetailsTypeFull        CreateBootVolumeBackupDetailsTypeEnum = "FULL"
	CreateBootVolumeBackupDetailsTypeIncremental CreateBootVolumeBackupDetailsTypeEnum = "INCREMENTAL"
)

var mappingCreateBootVolumeBackupDetailsType = map[string]CreateBootVolumeBackupDetailsTypeEnum{
	"FULL":        CreateBootVolumeBackupDetailsTypeFull,
	"INCREMENTAL": CreateBootVolumeBackupDetailsTypeIncremental,
}

// GetCreateBootVolumeBackupDetailsTypeEnumValues Enumerates the set of values for CreateBootVolumeBackupDetailsType
func GetCreateBootVolumeBackupDetailsTypeEnumValues() []CreateBootVolumeBackupDetailsTypeEnum {
	values := make([]CreateBootVolumeBackupDetailsTypeEnum, 0)
	for _, v := range mappingCreateBootVolumeBackupDetailsType {
		values = append(values, v)
	}
	return values
}
