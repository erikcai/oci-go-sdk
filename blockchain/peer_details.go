// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Blockchain Platform Control Plane API
//
// Blockchain Platform Plane API
//

package blockchain

import (
	"github.com/oracle/oci-go-sdk/common"
)

// PeerDetails A Peer details
type PeerDetails struct {

	// Peer role
	Role PeerRoleRoleEnum `mandatory:"false" json:"role,omitempty"`

	// Peer Log level
	LogLevel CommonLogLevelLogLevelEnum `mandatory:"false" json:"logLevel,omitempty"`

	// peer identifier
	PeerName *string `mandatory:"false" json:"peerName"`

	// organization name
	Organization *string `mandatory:"false" json:"organization"`

	// peer alias
	Alias *string `mandatory:"false" json:"alias"`

	OcpuAllocationParam *OcpuAllocationNumberParam `mandatory:"false" json:"ocpuAllocationParam"`

	// Host on which the Peer exists
	Host *string `mandatory:"false" json:"host"`

	// Availability Domain of peer
	Ad AvailabilityDomainAdsEnum `mandatory:"false" json:"ad,omitempty"`
}

func (m PeerDetails) String() string {
	return common.PointerString(m)
}
