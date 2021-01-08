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

// PrivateEndpointAssociation A summary of private endpoint information. This object is returned when listing the private
// endpoints associated with a given endpoint service.
type PrivateEndpointAssociation struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the private endpoint.
	Id *string `mandatory:"false" json:"id"`

	// The private IP address (in the customer's VCN) that represents the access point for the
	// associated endpoint service.
	PrivateEndpointIp *string `mandatory:"false" json:"privateEndpointIp"`

	// The three-label FQDN to use for the private endpoint. The customer VCN's DNS records use
	// this FQDN.
	// For important information about how this attribute is used, see the discussion
	// of DNS and FQDNs in PrivateEndpoint.
	// Example: `xyz.oraclecloud.com`
	EndpointFqdn *string `mandatory:"false" json:"endpointFqdn"`

	// A list of additional FQDNs that you can provide along with endpointFqdn. These FQDNs are added to the
	// customer VCN's DNS record. For more information, see the discussion of DNS and FQDNs in
	// PrivateEndpoint.
	AdditionalFqdns []string `mandatory:"false" json:"additionalFqdns"`

	ReverseConnectionConfiguration *ReverseConnectionConfiguration `mandatory:"false" json:"reverseConnectionConfiguration"`
}

func (m PrivateEndpointAssociation) String() string {
	return common.PointerString(m)
}
