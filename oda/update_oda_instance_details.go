// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Digital Assistant Control Plane API
//
// API to create and maintain Digital Assistant (ODA) service instances.
//

package oda

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateOdaInstanceDetails The information to be updated.
type UpdateOdaInstanceDetails struct {

	// ODA instance Identifier
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description of the ODA instance
	Description *string `mandatory:"false" json:"description"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateOdaInstanceDetails) String() string {
	return common.PointerString(m)
}
