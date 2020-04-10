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

// CreateConfigurationDetails The details required to create a new Configuration.
type CreateConfigurationDetails struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name of the associated Shape.
	ShapeName *string `mandatory:"true" json:"shapeName"`

	Variables *ConfigurationVariables `mandatory:"true" json:"variables"`

	// User-provided data about the Configuration.
	Description *string `mandatory:"false" json:"description"`

	// The display name of the Configuration.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The OCID of the Configuration from which the new Configuration is derived. The values in CreateConfigurationDetails.options supersede the options of the parent Configuration.
	ParentConfigurationId *string `mandatory:"false" json:"parentConfigurationId"`

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
