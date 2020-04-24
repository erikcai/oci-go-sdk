// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// API covering the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services. Use this API
// to manage resources such as virtual cloud networks (VCNs), compute instances, and
// block storage volumes.
//

package core

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ModifyReverseConnectionsDetails Details for modifying reverse connections configuration for the specified private endpoint.
type ModifyReverseConnectionsDetails struct {

	// List of DNS zones to exclude from the default DNS resolution context.
	ExcludedDnsZones []string `mandatory:"false" json:"excludedDnsZones"`

	// A list of the OCIDs of the network security groups that the reverse connection's VNIC belongs to.
	// For more information about NSGs, see
	// NetworkSecurityGroup.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// Number of customer endpoints that the service provider expects to establish connections to using this RCE. The default is 0.
	// When non-zero value is specified, reverse connection configuration will be allocated with a list of CIDRs, from
	// which NAT IP addresses will be allocated. These list of CIDRs will not be shared by other reverse
	// connection enabled private endpoints.
	// When zero is specified, reverse connection configuration will get NAT IP addresses from common pool of CIDRs,
	// which will be shared with other reverse connection enabled private endpoints.
	// If the private endpoint was enabled with reverse connection with 0 already, the field is not updatable.
	// The size may not be updated with smaller number than previously specified value, but may be increased.
	CustomerEndpointsSize *int `mandatory:"false" json:"customerEndpointsSize"`

	// Layer 4 transport protocol to be used when resolving DNS queries within the default DNS resolution context.
	DefaultDnsContextTransport ModifyReverseConnectionsDetailsDefaultDnsContextTransportEnum `mandatory:"false" json:"defaultDnsContextTransport,omitempty"`
}

func (m ModifyReverseConnectionsDetails) String() string {
	return common.PointerString(m)
}

// ModifyReverseConnectionsDetailsDefaultDnsContextTransportEnum Enum with underlying type: string
type ModifyReverseConnectionsDetailsDefaultDnsContextTransportEnum string

// Set of constants representing the allowable values for ModifyReverseConnectionsDetailsDefaultDnsContextTransportEnum
const (
	ModifyReverseConnectionsDetailsDefaultDnsContextTransportTcp ModifyReverseConnectionsDetailsDefaultDnsContextTransportEnum = "TCP"
	ModifyReverseConnectionsDetailsDefaultDnsContextTransportUdp ModifyReverseConnectionsDetailsDefaultDnsContextTransportEnum = "UDP"
)

var mappingModifyReverseConnectionsDetailsDefaultDnsContextTransport = map[string]ModifyReverseConnectionsDetailsDefaultDnsContextTransportEnum{
	"TCP": ModifyReverseConnectionsDetailsDefaultDnsContextTransportTcp,
	"UDP": ModifyReverseConnectionsDetailsDefaultDnsContextTransportUdp,
}

// GetModifyReverseConnectionsDetailsDefaultDnsContextTransportEnumValues Enumerates the set of values for ModifyReverseConnectionsDetailsDefaultDnsContextTransportEnum
func GetModifyReverseConnectionsDetailsDefaultDnsContextTransportEnumValues() []ModifyReverseConnectionsDetailsDefaultDnsContextTransportEnum {
	values := make([]ModifyReverseConnectionsDetailsDefaultDnsContextTransportEnum, 0)
	for _, v := range mappingModifyReverseConnectionsDetailsDefaultDnsContextTransport {
		values = append(values, v)
	}
	return values
}
