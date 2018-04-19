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

// CreatePublicIpDetails The representation of CreatePublicIpDetails
type CreatePublicIpDetails struct {

	// The OCID of the compartment to contain the public IP. For `EPHEMERAL` public IP addresses, the compartmentId has to
	// be the same as for the private IP.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Defines when the public IP is deleted and released back to Oracle's public IP pool.
	Lifetime CreatePublicIpDetailsLifetimeEnum `mandatory:"true" json:"lifetime"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid
	// entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The OCID of the private IP to assign the public IP to.
	// For `EPHEMERAL` public IP addresses this parameter is required since `EPHEMERAL` public IP addresses are always assigned to a private IP.
	// For `RESERVED` public IP this is an optional parameter. If not provided, the public IP will be created
	// and can later be assigned to a private IP via UpdatePublicIp.
	PrivateIpId *string `mandatory:"false" json:"privateIpId"`
}

func (m CreatePublicIpDetails) String() string {
	return common.PointerString(m)
}

// CreatePublicIpDetailsLifetimeEnum Enum with underlying type: string
type CreatePublicIpDetailsLifetimeEnum string

// Set of constants representing the allowable values for CreatePublicIpDetailsLifetime
const (
	CreatePublicIpDetailsLifetimeEphemeral CreatePublicIpDetailsLifetimeEnum = "EPHEMERAL"
	CreatePublicIpDetailsLifetimeReserved  CreatePublicIpDetailsLifetimeEnum = "RESERVED"
)

var mappingCreatePublicIpDetailsLifetime = map[string]CreatePublicIpDetailsLifetimeEnum{
	"EPHEMERAL": CreatePublicIpDetailsLifetimeEphemeral,
	"RESERVED":  CreatePublicIpDetailsLifetimeReserved,
}

// GetCreatePublicIpDetailsLifetimeEnumValues Enumerates the set of values for CreatePublicIpDetailsLifetime
func GetCreatePublicIpDetailsLifetimeEnumValues() []CreatePublicIpDetailsLifetimeEnum {
	values := make([]CreatePublicIpDetailsLifetimeEnum, 0)
	for _, v := range mappingCreatePublicIpDetailsLifetime {
		values = append(values, v)
	}
	return values
}
