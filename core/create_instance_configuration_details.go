// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateInstanceConfigurationDetails An instance configuration that can be used to launch
type CreateInstanceConfigurationDetails struct {

	// The OCID of the compartment containing the instance configuration.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "foo-value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name for the instance configuration
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	InstanceConfigurationInstanceDetails *InstanceConfigurationInstanceDetails `mandatory:"false" json:"instanceConfigurationInstanceDetails"`
}

func (m CreateInstanceConfigurationDetails) String() string {
	return common.PointerString(m)
}