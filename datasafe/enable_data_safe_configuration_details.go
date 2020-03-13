// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Data Safe APIs
//
// APIs for using Data Safe
//

package datasafe

import (
	"github.com/oracle/oci-go-sdk/common"
)

// EnableDataSafeConfigurationDetails The information needed to enalbe data safe configuration for the tenant
type EnableDataSafeConfigurationDetails struct {

	// whether the configuration is enabled
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m EnableDataSafeConfigurationDetails) String() string {
	return common.PointerString(m)
}
