// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// MySQL as a Service API
//
// The API for the MySQL Service
//

package mysql

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateBackupDetails Complete information for a Backup.
type CreateBackupDetails struct {

	// The type of backup.
	BackupType CreateBackupDetailsBackupTypeEnum `mandatory:"true" json:"backupType"`

	// If the backup was created automatically, or by a manual request.
	CreationType BackupCreationTypeEnum `mandatory:"true" json:"creationType"`

	// The OCID of the DB System the Backup is associated to.
	DbSystemId *string `mandatory:"true" json:"dbSystemId"`

	// A user-supplied display name for the backup.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A user-supplied description for the backup.
	Description *string `mandatory:"false" json:"description"`

	// Number of days to retain this backup.
	RetentionInDays *int `mandatory:"false" json:"retentionInDays"`

	// Simple key-value pair applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateBackupDetails) String() string {
	return common.PointerString(m)
}

// CreateBackupDetailsBackupTypeEnum Enum with underlying type: string
type CreateBackupDetailsBackupTypeEnum string

// Set of constants representing the allowable values for CreateBackupDetailsBackupTypeEnum
const (
	CreateBackupDetailsBackupTypeFull        CreateBackupDetailsBackupTypeEnum = "FULL"
	CreateBackupDetailsBackupTypeIncremental CreateBackupDetailsBackupTypeEnum = "INCREMENTAL"
)

var mappingCreateBackupDetailsBackupType = map[string]CreateBackupDetailsBackupTypeEnum{
	"FULL":        CreateBackupDetailsBackupTypeFull,
	"INCREMENTAL": CreateBackupDetailsBackupTypeIncremental,
}

// GetCreateBackupDetailsBackupTypeEnumValues Enumerates the set of values for CreateBackupDetailsBackupTypeEnum
func GetCreateBackupDetailsBackupTypeEnumValues() []CreateBackupDetailsBackupTypeEnum {
	values := make([]CreateBackupDetailsBackupTypeEnum, 0)
	for _, v := range mappingCreateBackupDetailsBackupType {
		values = append(values, v)
	}
	return values
}
