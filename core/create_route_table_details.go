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

// CreateRouteTableDetails The representation of CreateRouteTableDetails
type CreateRouteTableDetails struct {

	// The OCID of the compartment to contain the route table.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The collection of rules used for routing destination IPs to network devices.
	RouteRules []RouteRule `mandatory:"true" json:"routeRules"`

	// The OCID of the VCN the route table belongs to.
	VcnId *string `mandatory:"true" json:"vcnId"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "foo-value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`
}

func (m CreateRouteTableDetails) String() string {
	return common.PointerString(m)
}
