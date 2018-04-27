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

// UpdateVnicDetails The representation of UpdateVnicDetails
type UpdateVnicDetails struct {

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "foo-value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The hostname DNS label used in the FQDN to identify the private IP. Does have to be unique within the subnet.
	HostnameLabel *string `mandatory:"false" json:"hostnameLabel"`

	// Indicates whether Source/Destination check is disabled on the VNIC.
	// Defaults to `false`, in which case we enable Source/Destination check on the VNIC.
	SkipSourceDestCheck *bool `mandatory:"false" json:"skipSourceDestCheck"`
}

func (m UpdateVnicDetails) String() string {
	return common.PointerString(m)
}
