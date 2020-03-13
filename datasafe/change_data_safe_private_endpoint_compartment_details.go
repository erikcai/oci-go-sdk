// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Data Safe APIs
//
// APIs for using Data Safe
//

package datasafe

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ChangeDataSafePrivateEndpointCompartmentDetails The compartment to be changed to
type ChangeDataSafePrivateEndpointCompartmentDetails struct {

	// The new compartment identifier for the private endpoint
	CompartmentId *string `mandatory:"false" json:"compartmentId"`
}

func (m ChangeDataSafePrivateEndpointCompartmentDetails) String() string {
	return common.PointerString(m)
}
