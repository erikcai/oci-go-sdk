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

// UpdatePrivateIpDetails The representation of UpdatePrivateIpDetails
type UpdatePrivateIpDetails struct {

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "foo-value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// User friendly name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Hostname for the private IP. Only the hostname label, not the FQDN.
	HostnameLabel *string `mandatory:"false" json:"hostnameLabel"`

	// The OCID of the VNIC, which would cause the IP to be reassigned to a new VNIC.
	// VNIC must be in the subnet of the private IP.
	VnicId *string `mandatory:"false" json:"vnicId"`
}

func (m UpdatePrivateIpDetails) String() string {
	return common.PointerString(m)
}
