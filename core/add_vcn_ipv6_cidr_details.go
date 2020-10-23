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
	"github.com/oracle/oci-go-sdk/v27/common"
)

// AddVcnIpv6CidrDetails Contains an optional field which might be specified if private and public IPv6 CIDR block need to be different
type AddVcnIpv6CidrDetails struct {

	// This is not a required field and should only be specified if a custom private IPv6 CIDR
	// is desired for VCN's private IP address space.
	// The VCN size is always /48. If you don't provide a value, Oracle provides one and uses
	// that *same* CIDR for the `ipv6PublicCidrBlock`. If you do provide a value, Oracle provides
	// a *different* CIDR for the `ipv6PublicCidrBlock`. If provided, the new CIDR must be valid
	// and correctly formatted.
	// SeeIPv6 Addresses (https://docs.cloud.oracle.com/Content/Network/Concepts/ipv6.htm).
	// Example: `2001:0db8:0123::/48`
	Ipv6CidrBlock *string `mandatory:"false" json:"ipv6CidrBlock"`
}

func (m AddVcnIpv6CidrDetails) String() string {
	return common.PointerString(m)
}
