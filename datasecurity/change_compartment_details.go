// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Data Security Control Plane API
//
// The API to manage data security instance creation and deletion
//

package datasecurity

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ChangeCompartmentDetails The compartment to be changed to
type ChangeCompartmentDetails struct {

	// The new compartment identifier for the data security instance
	CompartmentId *string `mandatory:"false" json:"compartmentId"`
}

func (m ChangeCompartmentDetails) String() string {
	return common.PointerString(m)
}
