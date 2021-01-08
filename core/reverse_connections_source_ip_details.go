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

// ReverseConnectionsSourceIpDetails IP information for reverse connection configuration. Returned as part of the `PrivateEndpoint` object.
type ReverseConnectionsSourceIpDetails struct {

	// The IP address in the customer's VCN to be used as the source IP for reverse connection packets
	// traveling from the customer's VCN to the service's VCN.
	// Example: `10.0.4.9`
	SourceIp *string `mandatory:"false" json:"sourceIp"`
}

func (m ReverseConnectionsSourceIpDetails) String() string {
	return common.PointerString(m)
}
