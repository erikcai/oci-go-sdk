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

// Ipv6 An *IPv6* is a conceptual term that refers to an IPv6 address and related properties.
// The `IPv6` object is the API representation of an IPv6.
// IPv6 can be created and assigned to any particular VNICs which are inside an IPv6 enabled subnet.
type Ipv6 struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment containing the IPv6.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid
	// entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The IPv6's OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	Id *string `mandatory:"true" json:"id"`

	// The IPv6 address of the `IPv6` object. The address is within the IPv6 CIDR
	// of the VNIC's subnet.
	// Example: `2001:0db8:0123:4567:89ab:cdef:1234:5678`
	IpAddress *string `mandatory:"true" json:"ipAddress"`

	// The IPv6's current state.
	LifecycleState Ipv6LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID of the subnet the VNIC is in.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The date and time the IPv6 was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The OCID of the VNIC the IPv6 is assigned to. The VNIC and IPv6 must be in the same subnet.
	VnicId *string `mandatory:"true" json:"vnicId"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see
	// Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Whether IPv6 is usable for intenet communication. Internet access via IPv6 will not be allowed for
	// private subnet the same way as IPv4. Internet access will be enabled by default for a public subnet.
	// If VCN has IPv6 enabled with a custom IPv6 CIDR, a different public IPv6 address will be assigned
	// for a particular IPv6.
	IsInternetAccessAllowed *bool `mandatory:"false" json:"isInternetAccessAllowed"`

	// The IPv6 address to be used for an internet communication. The address is within the IPv6 public CIDR
	// of the VNIC's subnet. The `publicIpAddress` is always the same `ipAddress` if the VCN does not have
	// a custom IPv6 CIDR (i.e. use Oracle provided CIDR). If VCN has a custom IPv6 CIDR, the `publicIpAddress`
	// is drawn from the subnet public IPv6 CIDR by translating the IPv6 address prefix to the public IPv6 CIDR.
	// If `publicIpAddress` is not available, internet access is not permitted for this particular IPv6.
	PublicIpAddress *string `mandatory:"false" json:"publicIpAddress"`
}

func (m Ipv6) String() string {
	return common.PointerString(m)
}

// Ipv6LifecycleStateEnum Enum with underlying type: string
type Ipv6LifecycleStateEnum string

// Set of constants representing the allowable values for Ipv6LifecycleStateEnum
const (
	Ipv6LifecycleStateProvisioning Ipv6LifecycleStateEnum = "PROVISIONING"
	Ipv6LifecycleStateAvailable    Ipv6LifecycleStateEnum = "AVAILABLE"
	Ipv6LifecycleStateTerminating  Ipv6LifecycleStateEnum = "TERMINATING"
	Ipv6LifecycleStateTerminated   Ipv6LifecycleStateEnum = "TERMINATED"
)

var mappingIpv6LifecycleState = map[string]Ipv6LifecycleStateEnum{
	"PROVISIONING": Ipv6LifecycleStateProvisioning,
	"AVAILABLE":    Ipv6LifecycleStateAvailable,
	"TERMINATING":  Ipv6LifecycleStateTerminating,
	"TERMINATED":   Ipv6LifecycleStateTerminated,
}

// GetIpv6LifecycleStateEnumValues Enumerates the set of values for Ipv6LifecycleStateEnum
func GetIpv6LifecycleStateEnumValues() []Ipv6LifecycleStateEnum {
	values := make([]Ipv6LifecycleStateEnum, 0)
	for _, v := range mappingIpv6LifecycleState {
		values = append(values, v)
	}
	return values
}
