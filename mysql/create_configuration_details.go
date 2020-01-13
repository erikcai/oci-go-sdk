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

// CreateConfigurationDetails The details required to create a new MySQLaaS Configuration.
type CreateConfigurationDetails struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name of the associated MySQLaaS Instance Shape.
	ShapeName *string `mandatory:"true" json:"shapeName"`

	// The version of MySQLaaS that this Configuration and its Options are relevant to.
	MysqlVersion *string `mandatory:"true" json:"mysqlVersion"`

	// User-provided data about the MySQLaaS Instance.
	Description *string `mandatory:"false" json:"description"`

	// The display name of the MySQLaaS Configuration.
	DisplayName *string `mandatory:"false" json:"displayName"`

	Options []Option `mandatory:"false" json:"options"`

	// Simple key-value pair applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateConfigurationDetails) String() string {
	return common.PointerString(m)
}
