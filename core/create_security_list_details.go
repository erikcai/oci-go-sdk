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

// CreateSecurityListDetails The representation of CreateSecurityListDetails
type CreateSecurityListDetails struct {

	// The OCID of the compartment to contain the security list.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Rules for allowing egress IP packets.
	EgressSecurityRules []EgressSecurityRule `mandatory:"true" json:"egressSecurityRules"`

	// Rules for allowing ingress IP packets.
	IngressSecurityRules []IngressSecurityRule `mandatory:"true" json:"ingressSecurityRules"`

	// The OCID of the VCN the security list belongs to.
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

func (m CreateSecurityListDetails) String() string {
	return common.PointerString(m)
}
