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

// BlockchainPlatformCollection Result of a platform list or search. Contains BlockchainPlatformSummary
type BlockchainPlatformCollection struct {

	// Collection of BlockchainPlatformSummary
	Items []BlockchainPlatformSummary `mandatory:"true" json:"items"`
}

func (m BlockchainPlatformCollection) String() string {
	return common.PointerString(m)
}