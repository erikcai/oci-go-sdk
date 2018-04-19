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

// CreateVolumeBackupDetails The representation of CreateVolumeBackupDetails
type CreateVolumeBackupDetails struct {

	// The OCID of the volume that needs to be backed up.
	VolumeId *string `mandatory:"true" json:"volumeId"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "foo-value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name for the volume backup. Does not have to be unique and it's changeable.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The type of backup to create. If omitted, defaults to INCREMENTAL.
	Type CreateVolumeBackupDetailsTypeEnum `mandatory:"false" json:"type,omitempty"`
}

func (m CreateVolumeBackupDetails) String() string {
	return common.PointerString(m)
}

// CreateVolumeBackupDetailsTypeEnum Enum with underlying type: string
type CreateVolumeBackupDetailsTypeEnum string

// Set of constants representing the allowable values for CreateVolumeBackupDetailsType
const (
	CreateVolumeBackupDetailsTypeFull        CreateVolumeBackupDetailsTypeEnum = "FULL"
	CreateVolumeBackupDetailsTypeIncremental CreateVolumeBackupDetailsTypeEnum = "INCREMENTAL"
)

var mappingCreateVolumeBackupDetailsType = map[string]CreateVolumeBackupDetailsTypeEnum{
	"FULL":        CreateVolumeBackupDetailsTypeFull,
	"INCREMENTAL": CreateVolumeBackupDetailsTypeIncremental,
}

// GetCreateVolumeBackupDetailsTypeEnumValues Enumerates the set of values for CreateVolumeBackupDetailsType
func GetCreateVolumeBackupDetailsTypeEnumValues() []CreateVolumeBackupDetailsTypeEnum {
	values := make([]CreateVolumeBackupDetailsTypeEnum, 0)
	for _, v := range mappingCreateVolumeBackupDetailsType {
		values = append(values, v)
	}
	return values
}
