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

// Backup The representation of Backup
type Backup struct {

	// OCID of the backup itself
	Id *string `mandatory:"true" json:"id"`

	// The time the backup record was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time at which the backup was updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The state of the backup.
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Additional information about the current lifecycleState.
	LifecycleDetails *string `mandatory:"true" json:"lifecycleDetails"`

	// The type of backup.
	BackupType BackupTypeEnum `mandatory:"true" json:"backupType"`

	// If the backup was created automatically, or by a manual request.
	CreationType BackupCreationTypeEnum `mandatory:"true" json:"creationType"`

	// The size of the backup in base-2 (IEC) mebibytes. (MiB)
	SizeInMBs *int `mandatory:"true" json:"sizeInMBs"`

	// A user-supplied display name for the backup.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A user-supplied description for the backup.
	Description *string `mandatory:"false" json:"description"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The OCID of the DbSystem the Backup is associated to.
	DbSystemId *string `mandatory:"false" json:"dbSystemId"`

	// Number of days to retain this backup.
	RetentionDays *int `mandatory:"false" json:"retentionDays"`

	// Simple key-value pair applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m Backup) String() string {
	return common.PointerString(m)
}
