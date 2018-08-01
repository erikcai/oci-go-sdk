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

// CreateCrossRegionBackupDetails The representation of CreateCrossRegionBackupDetails
type CreateCrossRegionBackupDetails struct {

	// The OCID of the compartment the backup is to be copied to.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The size of the backup in MBs.
	SizeInMBs *int64 `mandatory:"true" json:"sizeInMBs"`

	// The OCID of the volume backup from which the backup is copied from.
	SourceBackupId *string `mandatory:"true" json:"sourceBackupId"`

	// The Availability Domain of the source volume from which the source backup is copied from.
	SourceVolumeAvailabilityDomain *string `mandatory:"true" json:"sourceVolumeAvailabilityDomain"`

	// The OCID of the source volume of the source backup from which the backup is copied from.
	SourceVolumeId *string `mandatory:"true" json:"sourceVolumeId"`

	// The volumeTypeId of the source volume (V2_ALL, V2_BOOT)
	SourceVolumeTypeId *string `mandatory:"true" json:"sourceVolumeTypeId"`

	// The tenant the backup is copied to. Also used in principal along with the userSubjectId to validate user permissions.
	TenantId *string `mandatory:"true" json:"tenantId"`

	// The type of a volume backup.
	Type CreateCrossRegionBackupDetailsTypeEnum `mandatory:"true" json:"type"`

	// The subjectId of the copy API caller
	UserSubjectId *string `mandatory:"true" json:"userSubjectId"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// volume encryption keys
	EncryptionKeys *CrossRegionBackupEncryptionKeys `mandatory:"false" json:"encryptionKeys"`

	// The date and time the volume backup will expire and be automatically deleted.
	// Format defined by RFC3339. This parameter will always be present for backups that
	// were created automatically by a scheduled-backup policy. For manually created backups,
	// it will be absent, signifying that there is no expiration time and the backup will
	// last forever until manually deleted.
	ExpirationTime *common.SDKTime `mandatory:"false" json:"expirationTime"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see
	// Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The image OCID used to create the boot volume the backup is taken from.
	ImageId *string `mandatory:"false" json:"imageId"`

	// Specifies whether the backup was created manually, or via scheduled backup policy. Defaults to MANUAL
	SourceType CreateCrossRegionBackupDetailsSourceTypeEnum `mandatory:"false" json:"sourceType,omitempty"`

	// The size of the volume in MB.
	UniqueSizeInMBs *int64 `mandatory:"false" json:"uniqueSizeInMBs"`
}

func (m CreateCrossRegionBackupDetails) String() string {
	return common.PointerString(m)
}

// CreateCrossRegionBackupDetailsSourceTypeEnum Enum with underlying type: string
type CreateCrossRegionBackupDetailsSourceTypeEnum string

// Set of constants representing the allowable values for CreateCrossRegionBackupDetailsSourceType
const (
	CreateCrossRegionBackupDetailsSourceTypeManual    CreateCrossRegionBackupDetailsSourceTypeEnum = "MANUAL"
	CreateCrossRegionBackupDetailsSourceTypeScheduled CreateCrossRegionBackupDetailsSourceTypeEnum = "SCHEDULED"
)

var mappingCreateCrossRegionBackupDetailsSourceType = map[string]CreateCrossRegionBackupDetailsSourceTypeEnum{
	"MANUAL":    CreateCrossRegionBackupDetailsSourceTypeManual,
	"SCHEDULED": CreateCrossRegionBackupDetailsSourceTypeScheduled,
}

// GetCreateCrossRegionBackupDetailsSourceTypeEnumValues Enumerates the set of values for CreateCrossRegionBackupDetailsSourceType
func GetCreateCrossRegionBackupDetailsSourceTypeEnumValues() []CreateCrossRegionBackupDetailsSourceTypeEnum {
	values := make([]CreateCrossRegionBackupDetailsSourceTypeEnum, 0)
	for _, v := range mappingCreateCrossRegionBackupDetailsSourceType {
		values = append(values, v)
	}
	return values
}

// CreateCrossRegionBackupDetailsTypeEnum Enum with underlying type: string
type CreateCrossRegionBackupDetailsTypeEnum string

// Set of constants representing the allowable values for CreateCrossRegionBackupDetailsType
const (
	CreateCrossRegionBackupDetailsTypeFull        CreateCrossRegionBackupDetailsTypeEnum = "FULL"
	CreateCrossRegionBackupDetailsTypeIncremental CreateCrossRegionBackupDetailsTypeEnum = "INCREMENTAL"
)

var mappingCreateCrossRegionBackupDetailsType = map[string]CreateCrossRegionBackupDetailsTypeEnum{
	"FULL":        CreateCrossRegionBackupDetailsTypeFull,
	"INCREMENTAL": CreateCrossRegionBackupDetailsTypeIncremental,
}

// GetCreateCrossRegionBackupDetailsTypeEnumValues Enumerates the set of values for CreateCrossRegionBackupDetailsType
func GetCreateCrossRegionBackupDetailsTypeEnumValues() []CreateCrossRegionBackupDetailsTypeEnum {
	values := make([]CreateCrossRegionBackupDetailsTypeEnum, 0)
	for _, v := range mappingCreateCrossRegionBackupDetailsType {
		values = append(values, v)
	}
	return values
}
