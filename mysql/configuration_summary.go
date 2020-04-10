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

// ConfigurationSummary The general details of a Configuration such as its id, displayName, type, and shape association.
type ConfigurationSummary struct {

	// The OCID of the Configuration.
	Id *string `mandatory:"true" json:"id"`

	// OCID of the Compartment the Configuration exists in.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name of the associated Shape.
	ShapeName *string `mandatory:"true" json:"shapeName"`

	// The Configuration type, DEFAULT or CUSTOM
	Type ConfigurationTypeEnum `mandatory:"true" json:"type"`

	// The current state of the Configuration.
	LifecycleState ConfigurationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// User-provided data about the Configuration.
	Description *string `mandatory:"false" json:"description"`

	// The display name of the Configuration.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The date and time the Configuration was created, as described by RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the Configuration was last updated, as described by RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Simple key-value pair applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m ConfigurationSummary) String() string {
	return common.PointerString(m)
}
