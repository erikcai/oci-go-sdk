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

// UpdateConfigurationDetails The details required to update a MySQLaaS Configuration.
type UpdateConfigurationDetails struct {

	// The OCID of the compartment to move the MySQLaaS Configuration to.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// User-provided data about the MySQLaaS Instance.
	Description *string `mandatory:"false" json:"description"`

	// A new display name for the MySQLaaS Configuration.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The name of the shape the MySQLaaS Configuration is related to.
	ShapeName *string `mandatory:"false" json:"shapeName"`

	// The version of MySQLaaS that this Configuration and its Options are relevant to.
	MysqlVersion *string `mandatory:"false" json:"mysqlVersion"`

	Options []Option `mandatory:"false" json:"options"`

	// Simple key-value pair applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateConfigurationDetails) String() string {
	return common.PointerString(m)
}
