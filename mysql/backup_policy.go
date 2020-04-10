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

// BackupPolicy The Backup policy for the DB System.
type BackupPolicy struct {

	// If automated backups are enabled or disabled.
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// The start of a 30-minute window of time in which daily, automated backups occur.
	// At some point in the window, the system may incur a brief service disruption as the backup is performed.
	// If not defined, a window is selected from the following Region-based time-spans:
	// - eu-frankfurt-1: 20:00 - 04:00 UTC
	// - us-ashburn-1: 03:00 - 11:00 UTC
	// - uk-london-1: 06:00 - 14:00 UTC
	// - ap-tokyo-1: 13:00 - 21:00
	// - us-phoenix-1: 06:00 - 14:00
	// The format of the string is "{day-of-week} {time-of-day}".
	// "{hour-of-week}" is a simple case-insensitive name "mon", "tue", &c.
	WindowStartTime *string `mandatory:"true" json:"windowStartTime"`

	// The number of days automated backups are retained.
	// The default is a function of how the configuration was created: CLI/SDK: 5 days, Console: 21 days.
	RetentionInDays *int `mandatory:"true" json:"retentionInDays"`

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
