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

// DeleteOceInstanceDetails The information about the resource to be deleted.
type DeleteOceInstanceDetails struct {

	// IDCS Proof of Stripe Token
	IdcsAt *string `mandatory:"true" json:"idcsAt"`
}

func (m DeleteOceInstanceDetails) String() string {
	return common.PointerString(m)
}
