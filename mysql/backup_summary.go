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

// BackupSummary Summary information for the Backups that exist.
type BackupSummary struct {

	// OCID of the backup.
	Id *string `mandatory:"true" json:"id"`

	// The time the backup was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The state of the backup.
	LifecycleState BackupLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The type of backup.
	BackupType BackupBackupTypeEnum `mandatory:"true" json:"backupType"`

	// The OCID of the DB System the Backup is associated with.
	DbSystemId *string `mandatory:"true" json:"dbSystemId"`

	// A user-supplied display name for the backup.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A user-supplied description of the backup.
	Description *string `mandatory:"false" json:"description"`

	// Initial size of the data volume in GiBs.
	DataStorageSizeInGBs *int `mandatory:"false" json:"dataStorageSizeInGBs"`

	// The size of the backup in GiBs.
	BackupSizeInGBs *int `mandatory:"false" json:"backupSizeInGBs"`

	// Number of days to retain this backup.
	RetentionInDays *int `mandatory:"false" json:"retentionInDays"`

	// The MySQL server version of the Db Instance used for backup
	MysqlVersion *string `mandatory:"false" json:"mysqlVersion"`

	// The shape of the Db instance used for backup.
	ShapeName *string `mandatory:"false" json:"shapeName"`

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
