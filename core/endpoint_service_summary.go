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

// EndpointServiceSummary A summary of Endpoint Services (ES) for calls that return a list of ESs. More details about a resource
// are returned when a request is made to retrieve that specific resource.
type EndpointServiceSummary struct {

	// The Endpoint Service's Oracle ID (OCID) (/Content/General/Concepts/identifiers.htm).
	Id *string `mandatory:"true" json:"id"`

	// The OCID  (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)  of the VCN to contain the Endpoint service.
	VcnId *string `mandatory:"true" json:"vcnId"`

	// Name of the Endpoint Service.
	DisplayName *string `mandatory:"true" json:"displayName"`
}

func (m EndpointServiceSummary) String() string {
	return common.PointerString(m)
}
