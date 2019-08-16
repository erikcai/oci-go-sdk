// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// OceInstance API
//
// A description of the OceInstance API
//

package oce

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ChangeOceInstanceCompartmentDetails The information about compartment details.
type ChangeOceInstanceCompartmentDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment
	// into which the OceInstance should be moved.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`
}

func (m ChangeOceInstanceCompartmentDetails) String() string {
	return common.PointerString(m)
}
