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

// CrossRegionBackup A point-in-time cross region copy of a volume backup or of a bootVolume backup.
type CrossRegionBackup struct {

	// The OCID of the cross region backup.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the source volume.
	VolumeId *string `mandatory:"true" json:"volumeId"`

	// A user-friendly name for the cross region backup. Does not have to be unique and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The size of the cross region volume, in MBs.
	SizeInMBs *int64 `mandatory:"true" json:"sizeInMBs"`

	// The size used by the backup, in GBs. It is typically smaller than sizeInMBs, depending on the space
	// consumed on the boot volume and whether the backup is full or incremental.
	UniqueSizeInMBs *int64 `mandatory:"true" json:"uniqueSizeInMBs"`

	// The OCID of the compartment that contains the cross region backup.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The date and time the cross region backup was created. This is the time the actual point-in-time image
	// of the volume data was taken. Format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of a cross region backup.
	LifecycleState CrossRegionBackupLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The type of the cross region backup.
	Type CrossRegionBackupTypeEnum `mandatory:"true" json:"type"`

	// Specifies whether the backup was created manually, or via scheduled backup policy.
	SourceType CrossRegionBackupSourceTypeEnum `mandatory:"true" json:"sourceType"`

	// The OCID of the source volume backup (can be Boot or Data volume backup).
	SourceVolumeBackupId *string `mandatory:"false" json:"sourceVolumeBackupId"`

	// The date and time the request to create the cross region backup was received. Format defined by RFC3339.
	TimeRequestReceived *common.SDKTime `mandatory:"false" json:"timeRequestReceived"`

	// The volumeTypeId of the source volume (V2_ALL, V2_BOOT)
	SourceVolumeTypeId *string `mandatory:"false" json:"sourceVolumeTypeId"`

	// The date and time the volume backup will expire and be automatically deleted.
	// Format defined by RFC3339. This parameter will always be present for backups that
	// were created automatically by a scheduled-backup policy. For manually created backups,
	// it will be absent, signifying that there is no expiration time and the backup will
	// last forever until manually deleted.
	ExpirationTime *common.SDKTime `mandatory:"false" json:"expirationTime"`

	// The image OCID used to create the original backup is taken from (if boot volume backup).
	ImageId *string `mandatory:"false" json:"imageId"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see
	// Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CrossRegionBackup) String() string {
	return common.PointerString(m)
}

// CrossRegionBackupLifecycleStateEnum Enum with underlying type: string
type CrossRegionBackupLifecycleStateEnum string

// Set of constants representing the allowable values for CrossRegionBackupLifecycleState
const (
	CrossRegionBackupLifecycleStateRequestReceived CrossRegionBackupLifecycleStateEnum = "REQUEST_RECEIVED"
	CrossRegionBackupLifecycleStateCreating        CrossRegionBackupLifecycleStateEnum = "CREATING"
	CrossRegionBackupLifecycleStateCopyPending     CrossRegionBackupLifecycleStateEnum = "COPY_PENDING"
	CrossRegionBackupLifecycleStateCopying         CrossRegionBackupLifecycleStateEnum = "COPYING"
	CrossRegionBackupLifecycleStateAvailable       CrossRegionBackupLifecycleStateEnum = "AVAILABLE"
	CrossRegionBackupLifecycleStateTerminating     CrossRegionBackupLifecycleStateEnum = "TERMINATING"
	CrossRegionBackupLifecycleStateTerminated      CrossRegionBackupLifecycleStateEnum = "TERMINATED"
	CrossRegionBackupLifecycleStateFaulty          CrossRegionBackupLifecycleStateEnum = "FAULTY"
)

var mappingCrossRegionBackupLifecycleState = map[string]CrossRegionBackupLifecycleStateEnum{
	"REQUEST_RECEIVED": CrossRegionBackupLifecycleStateRequestReceived,
	"CREATING":         CrossRegionBackupLifecycleStateCreating,
	"COPY_PENDING":     CrossRegionBackupLifecycleStateCopyPending,
	"COPYING":          CrossRegionBackupLifecycleStateCopying,
	"AVAILABLE":        CrossRegionBackupLifecycleStateAvailable,
	"TERMINATING":      CrossRegionBackupLifecycleStateTerminating,
	"TERMINATED":       CrossRegionBackupLifecycleStateTerminated,
	"FAULTY":           CrossRegionBackupLifecycleStateFaulty,
}

// GetCrossRegionBackupLifecycleStateEnumValues Enumerates the set of values for CrossRegionBackupLifecycleState
func GetCrossRegionBackupLifecycleStateEnumValues() []CrossRegionBackupLifecycleStateEnum {
	values := make([]CrossRegionBackupLifecycleStateEnum, 0)
	for _, v := range mappingCrossRegionBackupLifecycleState {
		values = append(values, v)
	}
	return values
}

// CrossRegionBackupTypeEnum Enum with underlying type: string
type CrossRegionBackupTypeEnum string

// Set of constants representing the allowable values for CrossRegionBackupType
const (
	CrossRegionBackupTypeFull        CrossRegionBackupTypeEnum = "FULL"
	CrossRegionBackupTypeIncremental CrossRegionBackupTypeEnum = "INCREMENTAL"
)

var mappingCrossRegionBackupType = map[string]CrossRegionBackupTypeEnum{
	"FULL":        CrossRegionBackupTypeFull,
	"INCREMENTAL": CrossRegionBackupTypeIncremental,
}

// GetCrossRegionBackupTypeEnumValues Enumerates the set of values for CrossRegionBackupType
func GetCrossRegionBackupTypeEnumValues() []CrossRegionBackupTypeEnum {
	values := make([]CrossRegionBackupTypeEnum, 0)
	for _, v := range mappingCrossRegionBackupType {
		values = append(values, v)
	}
	return values
}

// CrossRegionBackupSourceTypeEnum Enum with underlying type: string
type CrossRegionBackupSourceTypeEnum string

// Set of constants representing the allowable values for CrossRegionBackupSourceType
const (
	CrossRegionBackupSourceTypeManual    CrossRegionBackupSourceTypeEnum = "MANUAL"
	CrossRegionBackupSourceTypeScheduled CrossRegionBackupSourceTypeEnum = "SCHEDULED"
)

var mappingCrossRegionBackupSourceType = map[string]CrossRegionBackupSourceTypeEnum{
	"MANUAL":    CrossRegionBackupSourceTypeManual,
	"SCHEDULED": CrossRegionBackupSourceTypeScheduled,
}

// GetCrossRegionBackupSourceTypeEnumValues Enumerates the set of values for CrossRegionBackupSourceType
func GetCrossRegionBackupSourceTypeEnumValues() []CrossRegionBackupSourceTypeEnum {
	values := make([]CrossRegionBackupSourceTypeEnum, 0)
	for _, v := range mappingCrossRegionBackupSourceType {
		values = append(values, v)
	}
	return values
}
