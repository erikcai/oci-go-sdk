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

// PrivateAccessGatewaySummary A summary of Private Access Gateways (PGW) for calls that return a list of PGWs. More details about a resource
// are returned when a request is made to retrieve that specific resource.
type PrivateAccessGatewaySummary struct {

	// The Private Access Gateway's Oracle ID (OCID) (/Content/General/Concepts/identifiers.htm).
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name. Does not have to be unique.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`
}

func (m PrivateAccessGatewaySummary) String() string {
	return common.PointerString(m)
}
