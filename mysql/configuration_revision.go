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

// ConfigurationRevision The MySQLaaS Configuration Revision.
// A Revision is a specific snapshot of the values of the system variables
// and options in the life of a Configuration.
type ConfigurationRevision struct {

	// An identifier for this this specific revision of the Configuration
	// within the identifier space of this Configuration. Currently an
	// auto-increment integer, but not guaranteed to be one; treat as an
	// opaque string identifier.
	Id *string `mandatory:"true" json:"id"`

	// Revision creation time, as described by RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The status of this revision.
	Status ConfigurationApplicationStateEnum `mandatory:"true" json:"status"`

	// A user-provided description for this revision.
	Description *string `mandatory:"true" json:"description"`

	// The names of the option values specifically changed in this Revision of the Configuration.
	ChangedOptions []string `mandatory:"true" json:"changedOptions"`

	// The entire collection of Option values in this Revision of the Configuration.
	Options []Option `mandatory:"true" json:"options"`
}

func (m ConfigurationRevision) String() string {
	return common.PointerString(m)
}
