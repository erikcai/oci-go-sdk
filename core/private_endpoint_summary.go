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

// PrivateEndpointSummary A summary of Private Endpoints (PE) for calls that return a list of PEs. More details about a resource
// are returned when a request is made to retrieve that specific resource.
type PrivateEndpointSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Private Endpoint.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Endpoint Service OCID that sits behind this Private Endpoint.
	EndpointServiceId *string `mandatory:"true" json:"endpointServiceId"`
}

func (m PrivateEndpointSummary) String() string {
	return common.PointerString(m)
}
