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

// PublicIp A *public IP* is a conceptual term that refers to a public IP address and related properties.
// The `publicIp` object is the API representation of a public IP.
// A public IP has a scope and a lifetime associated with it. Both scope and lifetime cannot be changed after a public IP has been created.
// *scope* defines whether a public IP can be assigned to private IPs within or across availability domains.
// *lifetime* defines when a public IP is deleted and released to Oracle's public IP pool. Valid values are `EPHEMERAL` or `RESERVED`.
// `RESERVED` public IP addresses are only deleted when the customer explicitly issues a delete. These public IP addresses need not be assigned to a private IP.
// `EPHEMERAL` public IP addresses are always assigned to a private IP and will be deleted and released when:
// 1. The private IP is deleted.
// 2. A VNIC is detached or deleted. In this case all `EPHEMERAL` public IP addresses assigned to private IPs in the VNIC are deleted.
// 3. Instance is stopped.
// `EPHEMERAL` public IP addresses belong to the same compartment and availability domain as the private IP it is assigned to.
// Currently the following types of public IP addresses are supported:
// 1. (scope: `REGION` and lifetime: `RESERVED`) and
// 2. (scope: `AVAILABILITY_DOMAIN` and lifetime: `EPHEMERAL`).
type PublicIp struct {

	// The public IP's Availability Domain. This property is set only when the scope of the public IP is set to
	// AVAILABILITY_DOMAIN.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// The OCID of the compartment containing the public IP.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid
	// entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The public IP's Oracle ID (OCID).
	Id *string `mandatory:"false" json:"id"`

	// The public IP address of the `publicIp` object.
	// Example: `129.146.2.1`
	IpAddress *string `mandatory:"false" json:"ipAddress"`

	// The public IP's current state.
	LifecycleState PublicIpLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Defines when the public IP is deleted and released back to Oracle's public IP pool.
	Lifetime PublicIpLifetimeEnum `mandatory:"false" json:"lifetime,omitempty"`

	// The OCID of the private IP that the public IP can be optionally assigned to. The private IP and public IP
	// can be in different compartments.
	PrivateIpId *string `mandatory:"false" json:"privateIpId"`

	// The scope in which the public IP can be assigned to a private IP.
	// `REGION`: The public IP can be assigned to private IP in any availability domain of the region.
	// `AVAILABILITY_DOMAIN`: The public IP can be assigned to private IP belonging to the availability domain as
	// defined in the `availabilityDomain` property of the public IP object.
	Scope PublicIpScopeEnum `mandatory:"false" json:"scope,omitempty"`

	// The date and time the public IP was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m PublicIp) String() string {
	return common.PointerString(m)
}

// PublicIpLifecycleStateEnum Enum with underlying type: string
type PublicIpLifecycleStateEnum string

// Set of constants representing the allowable values for PublicIpLifecycleState
const (
	PublicIpLifecycleStateProvisioning PublicIpLifecycleStateEnum = "PROVISIONING"
	PublicIpLifecycleStateAvailable    PublicIpLifecycleStateEnum = "AVAILABLE"
	PublicIpLifecycleStateAssigning    PublicIpLifecycleStateEnum = "ASSIGNING"
	PublicIpLifecycleStateAssigned     PublicIpLifecycleStateEnum = "ASSIGNED"
	PublicIpLifecycleStateUnassigning  PublicIpLifecycleStateEnum = "UNASSIGNING"
	PublicIpLifecycleStateUnassigned   PublicIpLifecycleStateEnum = "UNASSIGNED"
	PublicIpLifecycleStateTerminating  PublicIpLifecycleStateEnum = "TERMINATING"
	PublicIpLifecycleStateTerminated   PublicIpLifecycleStateEnum = "TERMINATED"
)

var mappingPublicIpLifecycleState = map[string]PublicIpLifecycleStateEnum{
	"PROVISIONING": PublicIpLifecycleStateProvisioning,
	"AVAILABLE":    PublicIpLifecycleStateAvailable,
	"ASSIGNING":    PublicIpLifecycleStateAssigning,
	"ASSIGNED":     PublicIpLifecycleStateAssigned,
	"UNASSIGNING":  PublicIpLifecycleStateUnassigning,
	"UNASSIGNED":   PublicIpLifecycleStateUnassigned,
	"TERMINATING":  PublicIpLifecycleStateTerminating,
	"TERMINATED":   PublicIpLifecycleStateTerminated,
}

// GetPublicIpLifecycleStateEnumValues Enumerates the set of values for PublicIpLifecycleState
func GetPublicIpLifecycleStateEnumValues() []PublicIpLifecycleStateEnum {
	values := make([]PublicIpLifecycleStateEnum, 0)
	for _, v := range mappingPublicIpLifecycleState {
		values = append(values, v)
	}
	return values
}

// PublicIpLifetimeEnum Enum with underlying type: string
type PublicIpLifetimeEnum string

// Set of constants representing the allowable values for PublicIpLifetime
const (
	PublicIpLifetimeEphemeral PublicIpLifetimeEnum = "EPHEMERAL"
	PublicIpLifetimeReserved  PublicIpLifetimeEnum = "RESERVED"
)

var mappingPublicIpLifetime = map[string]PublicIpLifetimeEnum{
	"EPHEMERAL": PublicIpLifetimeEphemeral,
	"RESERVED":  PublicIpLifetimeReserved,
}

// GetPublicIpLifetimeEnumValues Enumerates the set of values for PublicIpLifetime
func GetPublicIpLifetimeEnumValues() []PublicIpLifetimeEnum {
	values := make([]PublicIpLifetimeEnum, 0)
	for _, v := range mappingPublicIpLifetime {
		values = append(values, v)
	}
	return values
}

// PublicIpScopeEnum Enum with underlying type: string
type PublicIpScopeEnum string

// Set of constants representing the allowable values for PublicIpScope
const (
	PublicIpScopeRegion             PublicIpScopeEnum = "REGION"
	PublicIpScopeAvailabilityDomain PublicIpScopeEnum = "AVAILABILITY_DOMAIN"
)

var mappingPublicIpScope = map[string]PublicIpScopeEnum{
	"REGION":              PublicIpScopeRegion,
	"AVAILABILITY_DOMAIN": PublicIpScopeAvailabilityDomain,
}

// GetPublicIpScopeEnumValues Enumerates the set of values for PublicIpScope
func GetPublicIpScopeEnumValues() []PublicIpScopeEnum {
	values := make([]PublicIpScopeEnum, 0)
	for _, v := range mappingPublicIpScope {
		values = append(values, v)
	}
	return values
}
