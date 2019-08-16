// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// OceInstance API
//
// A description of the OceInstance API
//

package oce

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateOceInstanceDetails The information to be updated.
type UpdateOceInstanceDetails struct {

	// OceInstance Identifier
	OceInstanceType *string `mandatory:"true" json:"oceInstanceType"`

	// OceInstance Identifier
	DisplayName *string `mandatory:"false" json:"displayName"`

	// OceInstance description
	Description *string `mandatory:"false" json:"description"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateOceInstanceDetails) String() string {
	return common.PointerString(m)
}
