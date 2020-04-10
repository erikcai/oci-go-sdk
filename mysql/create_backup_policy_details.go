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

// CreateBackupPolicyDetails Backup policy as optionally used for DB System Creation.
type CreateBackupPolicyDetails struct {

	// The time at which you want the 30-minute backup window to begin.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// # @fixme copy
	WindowStartTime *string `mandatory:"false" json:"windowStartTime"`

	// Number of days to retain a scheduled backup.
	RetentionInDays *int `mandatory:"false" json:"retentionInDays"`

	// Simple key-value pair applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Tags defined here will be copied verbatim as tags on the Backup resource created by this BackupPolicy.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateBackupPolicyDetails) String() string {
	return common.PointerString(m)
}
