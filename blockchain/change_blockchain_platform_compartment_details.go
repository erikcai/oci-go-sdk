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

// ChangeBlockchainPlatformCompartmentDetails Input payload to change a resource's compartment.
type ChangeBlockchainPlatformCompartmentDetails struct {

	// The OCID of the new compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`
}

func (m ChangeBlockchainPlatformCompartmentDetails) String() string {
	return common.PointerString(m)
}
