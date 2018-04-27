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

// UpdatePublicIpDetails The representation of UpdatePublicIpDetails
type UpdatePublicIpDetails struct {

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "foo-value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid
	// entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The OCID of the private IP to assign the public IP to.
	// - If the public IP was assigned to a different private IP, it will be unsassigned and the reassigned to
	//  the given private IP.
	// - If this field is set to an empty string then the public IP will be unassigned from the private IP it was assigned to.
	PrivateIpId *string `mandatory:"false" json:"privateIpId"`
}

func (m UpdatePublicIpDetails) String() string {
	return common.PointerString(m)
}
