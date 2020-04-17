// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Blockchain Platform Control Plane API
//
// Blockchain Platform Plane API
//

package blockchain

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ScaleOutPeerDetails The Peer details to be added
type ScaleOutPeerDetails struct {

	// organization name
	Organization *string `mandatory:"false" json:"organization"`

	// Peer role
	Role PeerRoleRoleEnum `mandatory:"false" json:"role,omitempty"`

	// Peer Log level
	LogLevel CommonLogLevelLogLevelEnum `mandatory:"false" json:"logLevel,omitempty"`

	// peer alias
	Alias *string `mandatory:"false" json:"alias"`

	OcpuAllocationParam *OcpuAllocationNumberParam `mandatory:"false" json:"ocpuAllocationParam"`

	// Availability Domain to place new peer
	Ad AvailabilityDomainAdsEnum `mandatory:"false" json:"ad,omitempty"`
}

func (m ScaleOutPeerDetails) String() string {
	return common.PointerString(m)
}
