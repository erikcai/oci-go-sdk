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

// PrivateIp A private IP.
type PrivateIp struct {

	// The private IP's Availability Domain.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// The OCID of the compartment containing the private IP.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

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

	// The private IP's OCID.
	Id *string `mandatory:"false" json:"id"`

	// The address of private IP.
	IpAddress *string `mandatory:"false" json:"ipAddress"`

	// Whether the private IP is a primary private IP to a VNIC. Primary private IPs
	// are unassigned and deleted automatically when the instance is terminated.
	IsPrimary *bool `mandatory:"false" json:"isPrimary"`

	// The subnet's OCID.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// Creation time of the entity.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The VNIC's OCID.
	VnicId *string `mandatory:"false" json:"vnicId"`
}

func (m PrivateIp) String() string {
	return common.PointerString(m)
}
