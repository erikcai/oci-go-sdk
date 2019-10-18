// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
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

// PrivateEndpointAssociation Details of a Private Endpoint using an Endpoint Service
type PrivateEndpointAssociation struct {

	// The Private Endpoint's Oracle ID (OCID) (/Content/General/Concepts/identifiers.htm).
	Id *string `mandatory:"false" json:"id"`

	// IP attached to the PE VNIC, this represents an access point for the Endpoint Service.
	PrivateEndpointIp *string `mandatory:"false" json:"privateEndpointIp"`

	// Service's 3 label FQDN representing the Endpoint Service.
	EndpointFqdn *string `mandatory:"false" json:"endpointFqdn"`

	ReverseConnectionConfiguration *ReverseConnectionConfiguration `mandatory:"false" json:"reverseConnectionConfiguration"`
}

func (m PrivateEndpointAssociation) String() string {
	return common.PointerString(m)
}
