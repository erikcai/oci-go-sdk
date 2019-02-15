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

// CreateIpv6Details Attributes available for IPv6 creation. VNIC to be attached to is the only required attribute.
// User can choose to define a specific address and have the address automatically assigned.
type CreateIpv6Details struct {

	// The OCID of the VNIC to assign the IPv6 to. The VNIC and IPv6 will be in the same subnet.
	VnicId *string `mandatory:"true" json:"vnicId"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid
	// entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see
	// Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// An IPv6 address of your choice. Must be an available IP address within
	// the subnet's CIDR. If you don't specify a value, Oracle automatically
	// assigns an IPv6 address from the subnet.
	// Example: `2001:0db8:0123:4567:89ab:cdef:1234:5678`
	IpAddress *string `mandatory:"false" json:"ipAddress"`

	// Whether IPv6 is usable for internet communication. Internet access via IPv6 will not be allowed for
	// private subnet the same way as IPv4. Internet access will be enabled by default for a public subnet.
	// If VCN has IPv6 enabled with a custom IPv6 CIDR, a different public IPv6 address will be assigned
	// for a particular IPv6.
	IsInternetAccessAllowed *bool `mandatory:"false" json:"isInternetAccessAllowed"`
}

func (m CreateIpv6Details) String() string {
	return common.PointerString(m)
}
