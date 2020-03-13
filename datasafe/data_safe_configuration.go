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

// DataSafeConfiguration Object of DataSafeConfiguration
type DataSafeConfiguration struct {

	// whether the configuration is enabled
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// Service URL of the data safe service
	Url *string `mandatory:"false" json:"url"`

	// Tenant identifier used to enable the configuration.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The time the the data safe configuration was enabled. An RFC3339 formatted datetime string
	TimeEnabled *common.SDKTime `mandatory:"false" json:"timeEnabled"`

	// The current state of the data safe configuration.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m DataSafeConfiguration) String() string {
	return common.PointerString(m)
}
