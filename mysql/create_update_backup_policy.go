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

// CreateUpdateBackupPolicy Backup Configuration object as optionally used for instance
// creation/update. No fields are required, in which case default values
// are determined and used as per the documentation.
type CreateUpdateBackupPolicy struct {

	// # @fixme copy
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// # @fixme copy
	Window *string `mandatory:"false" json:"window"`

	// # @fixme copy
	RetentionDays *int `mandatory:"false" json:"retentionDays"`

	// Simple key-value pair applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Tags defined here will be copied verbatim as tags on the Backup resource created by this BackupPolicy.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateUpdateBackupPolicy) String() string {
	return common.PointerString(m)
}
