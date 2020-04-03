// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ChangeDataSafePrivateEndpointCompartmentDetails The details used to change the compartment of a Data Safe private endpoint.
type ChangeDataSafePrivateEndpointCompartmentDetails struct {

	// The OCID of the new compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`
}

func (m ChangeDataSafePrivateEndpointCompartmentDetails) String() string {
	return common.PointerString(m)
}
