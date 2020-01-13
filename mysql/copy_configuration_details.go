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

// CopyConfigurationDetails The details required to create a copy of a MySQLaaS Configuration.
type CopyConfigurationDetails struct {

	// User-provided data about the MySQLaaS Instance.
	Description *string `mandatory:"false" json:"description"`

	// The display name of the MySQLaaS Configuration copy.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The version of MySQLaaS that this Configuration and its options are relevant to.
	MysqlVersion *string `mandatory:"false" json:"mysqlVersion"`

	// Optionally, the OCID of the compartment which the MySQLaaS
	// Configuration should be copied to. If this is not present, the
	// Configuration is copied to the same compartment as the original.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Simple key-value pair applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CopyConfigurationDetails) String() string {
	return common.PointerString(m)
}
