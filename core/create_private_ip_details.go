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

// CreatePrivateIpDetails The representation of CreatePrivateIpDetails
type CreatePrivateIpDetails struct {

	// The OCID of the VNIC. VNIC must be in the subnet of the private IP.
	VnicId *string `mandatory:"true" json:"vnicId"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "foo-value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Hostname for the private IP. Only the hostname label, not the FQDN.
	HostnameLabel *string `mandatory:"false" json:"hostnameLabel"`

	// The address of private IP. System will select an IP address from available IP addresses in subnet if none is
	// specified.
	IpAddress *string `mandatory:"false" json:"ipAddress"`
}

func (m CreatePrivateIpDetails) String() string {
	return common.PointerString(m)
}
