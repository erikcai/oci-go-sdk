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

// BackupSummary The representation of BackupSummary
type BackupSummary struct {

	// OCID of the backup itself
	Id *string `mandatory:"true" json:"id"`

	// The time the backup was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The state of the backup.
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The type of backup.
	BackupType BackupTypeEnum `mandatory:"true" json:"backupType"`

	// The size of the backup in base-2 (IEC) mebibytes (aka MiB).
	SizeInMBs *int `mandatory:"true" json:"sizeInMBs"`

	// A user-supplied display name for the backup.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The OCID of the DbSystem the Backup is associated to.
	DbSystemId *string `mandatory:"false" json:"dbSystemId"`

	// Simple key-value pair applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m BackupSummary) String() string {
	return common.PointerString(m)
}
