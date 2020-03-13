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

// BlockchainPlatformComponentDetails Blockchain Platform component details.
type BlockchainPlatformComponentDetails struct {

	// List of OSNs
	Osns []OsnDetails `mandatory:"false" json:"osns"`

	Replicas *ReplicaInfo `mandatory:"false" json:"replicas"`

	// List of Peers
	Peers []PeerDetails `mandatory:"false" json:"peers"`
}

func (m BlockchainPlatformComponentDetails) String() string {
	return common.PointerString(m)
}
