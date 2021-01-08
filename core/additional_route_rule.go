// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// API covering the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm),
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm), and
// Registry (https://docs.cloud.oracle.com/iaas/Content/Registry/Concepts/registryoverview.htm) services.
// Use this API to manage resources such as virtual cloud networks (VCNs),
// compute instances, block storage volumes, and container images.
//

package core

import (
	"github.com/oracle/oci-go-sdk/v31/common"
)

// AdditionalRouteRule A mapping between a destination IP address range and a virtual device to route matching
// packets to (a target). This is used only by a large route table.
type AdditionalRouteRule struct {

	// The OCID for the route rule's target. For information about the type of
	// targets you can specify, see
	// Route Tables (https://docs.cloud.oracle.com/Content/Network/Tasks/managingroutetables.htm).
	NetworkEntityId *string `mandatory:"true" json:"networkEntityId"`

	// An identifier for the route rule. You specify this identifier when you want to delete the rule.
	Id *string `mandatory:"true" json:"id"`

	// Deprecated. Instead use `destination` and `destinationType`. Requests that include both
	// `cidrBlock` and `destination` will be rejected.
	// A destination IP address range in CIDR notation. Matching packets will
	// be routed to the indicated network entity (the target).
	// Cannot be an IPv6 CIDR.
	// Example: `0.0.0.0/0`
	CidrBlock *string `mandatory:"false" json:"cidrBlock"`

	// Conceptually, this is the range of IP addresses used for matching when routing
	// traffic. Required if you provide a `destinationType`.
	// Allowed values:
	//   * IP address range in CIDR notation. Can be an IPv4 or IPv6 CIDR. For example: `192.168.1.0/24`
	//   or `2001:0db8:0123:45::/56`. If you set this to an IPv6 CIDR, the route rule's target
	//   can only be a DRG or internet gateway. Note that IPv6 addressing is currently supported
	//   only in certain regions. See IPv6 Addresses (https://docs.cloud.oracle.com/Content/Network/Concepts/ipv6.htm).
	//   * The `cidrBlock` value for a Service, if you're
	//   setting up a route rule for traffic destined for a particular `Service` through
	//     a service gateway. For example: `oci-phx-objectstorage`.
	Destination *string `mandatory:"false" json:"destination"`

	// Type of destination for the rule. Required if you provide a `destination`.
	//   * `CIDR_BLOCK`: If the rule's `destination` is an IP address range in CIDR notation.
	//   * `SERVICE_CIDR_BLOCK`: If the rule's `destination` is the `cidrBlock` value for a
	//     Service (the rule is for traffic destined for a
	//     particular `Service` through a service gateway).
	DestinationType AdditionalRouteRuleDestinationTypeEnum `mandatory:"false" json:"destinationType,omitempty"`

	// An optional description of your choice for the rule.
	Description *string `mandatory:"false" json:"description"`
}

func (m AdditionalRouteRule) String() string {
	return common.PointerString(m)
}

// AdditionalRouteRuleDestinationTypeEnum Enum with underlying type: string
type AdditionalRouteRuleDestinationTypeEnum string

// Set of constants representing the allowable values for AdditionalRouteRuleDestinationTypeEnum
const (
	AdditionalRouteRuleDestinationTypeCidrBlock        AdditionalRouteRuleDestinationTypeEnum = "CIDR_BLOCK"
	AdditionalRouteRuleDestinationTypeServiceCidrBlock AdditionalRouteRuleDestinationTypeEnum = "SERVICE_CIDR_BLOCK"
)

var mappingAdditionalRouteRuleDestinationType = map[string]AdditionalRouteRuleDestinationTypeEnum{
	"CIDR_BLOCK":         AdditionalRouteRuleDestinationTypeCidrBlock,
	"SERVICE_CIDR_BLOCK": AdditionalRouteRuleDestinationTypeServiceCidrBlock,
}

// GetAdditionalRouteRuleDestinationTypeEnumValues Enumerates the set of values for AdditionalRouteRuleDestinationTypeEnum
func GetAdditionalRouteRuleDestinationTypeEnumValues() []AdditionalRouteRuleDestinationTypeEnum {
	values := make([]AdditionalRouteRuleDestinationTypeEnum, 0)
	for _, v := range mappingAdditionalRouteRuleDestinationType {
		values = append(values, v)
	}
	return values
}
