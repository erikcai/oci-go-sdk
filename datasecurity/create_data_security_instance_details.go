// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Data Security Control Plane API
//
// The API to manage data security instance creation and deletion
//

package datasecurity

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateDataSecurityInstanceDetails The information about new data security instance.
type CreateDataSecurityInstanceDetails struct {

	// data security instance name
	DisplayName *string `mandatory:"true" json:"displayName"`

	// data security instance Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// data security instance Description
	Description *string `mandatory:"false" json:"description"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateDataSecurityInstanceDetails) String() string {
	return common.PointerString(m)
}
