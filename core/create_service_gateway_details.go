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

// CreateServiceGatewayDetails The representation of CreateServiceGatewayDetails
type CreateServiceGatewayDetails struct {

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the compartment to contain the
	// service gateway.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// List of the services you want to enable on this service gateway. The list can be empty.
	// You can later enable a particular service by using
	// AttachServiceId.
	// For each service, make sure to also set up a route rule with the service's `cidrBlock` as
	// the rule's destination CIDR and the service gateway as the rule's target. See
	// RouteTable.
	Services []ServiceIdRequestDetails `mandatory:"true" json:"services"`

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the VCN the gateway belongs to.
	VcnId *string `mandatory:"true" json:"vcnId"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "foo-value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`
}

func (m CreateServiceGatewayDetails) String() string {
	return common.PointerString(m)
}
