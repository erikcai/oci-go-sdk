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
	"github.com/erikcai/oci-go-sdk/v33/common"
)

// CreatePrivateIpNextHopDetails The data to create private IP's nextHop configuration.
type CreatePrivateIpNextHopDetails struct {

	// Details of nextHop targets.
	Targets []PrivateIpNextHopTarget `mandatory:"true" json:"targets"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// VNICaaS will flow-hash traffic that matches the service protocol and port.
	ServiceProtocolPorts []PrivateIpNextHopProtocolPort `mandatory:"false" json:"serviceProtocolPorts"`
}

func (m CreatePrivateIpNextHopDetails) String() string {
	return common.PointerString(m)
}
