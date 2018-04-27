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

// CreateCpeDetails The representation of CreateCpeDetails
type CreateCpeDetails struct {

	// The OCID of the compartment to contain the CPE.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The public IP address of the on-premise router.
	// Example: `143.19.23.16`
	IpAddress *string `mandatory:"true" json:"ipAddress"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "foo-value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`
}

func (m CreateCpeDetails) String() string {
	return common.PointerString(m)
}
