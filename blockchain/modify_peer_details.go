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

// ModifyPeerDetails peer to modify ocpu allocation
type ModifyPeerDetails struct {

	// peer identifier
	PeerName *string `mandatory:"false" json:"peerName"`

	OcpuAllocationParam *OcpuAllocationNumberParam `mandatory:"false" json:"ocpuAllocationParam"`
}

func (m ModifyPeerDetails) String() string {
	return common.PointerString(m)
}
