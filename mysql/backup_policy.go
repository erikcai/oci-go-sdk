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

// BackupPolicy The representation of BackupPolicy
type BackupPolicy struct {

	// If automated backups are enabled or disabled.
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// The start of a (@FIXME TBD: N-minute-long) window of time in which daily, automated backups occur.
	// At some undefined point in the window, the system may incur a brief service disruption as the backup is performed. If the DbSystem does not have HA enabled, then a service interruption is likely. If the DbSystem does have HA enabled, (@FIXME).
	// If not defined, a window will be selected at random from the following Region-based time-spans: [@FIXME: table, docs].
	Window *string `mandatory:"true" json:"window"`

	// The number of days automated backups are retained.
	// The default is a function of how the configuration was created: CLI/SDK: 5 days, Console: 21 days.
	RetentionDays *int `mandatory:"false" json:"retentionDays"`

	// Simple key-value pair applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Tags defined here will be copied verbatim as tags on the Backup resource created by this BackupPolicy.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Tags defined here will be copied verbatim as tags on the Backup resource created by this BackupPolicy.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m BackupPolicy) String() string {
	return common.PointerString(m)
}
