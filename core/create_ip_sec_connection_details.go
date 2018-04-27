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

// CreateIpSecConnectionDetails The representation of CreateIpSecConnectionDetails
type CreateIpSecConnectionDetails struct {

	// The OCID of the compartment to contain the IPSec connection.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the CPE.
	CpeId *string `mandatory:"true" json:"cpeId"`

	// The OCID of the DRG.
	DrgId *string `mandatory:"true" json:"drgId"`

	// Static routes to the CPE. At least one route must be included. The CIDR must not be a
	// multicast address or class E address.
	// Example: `10.0.1.0/24`
	StaticRoutes []string `mandatory:"true" json:"staticRoutes"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "foo-value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`
}

func (m CreateIpSecConnectionDetails) String() string {
	return common.PointerString(m)
}
