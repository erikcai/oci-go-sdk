// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// ONS Gateway API
//
// A description of the ONS Gateway API
//

package ons

import (
	"github.com/oracle/oci-go-sdk/common"
)

// SubscriptionConfirmation Encapsulates a subscription confirmation object.
type SubscriptionConfirmation struct {

	// The OCID of the subscription.
	Id *string `mandatory:"true" json:"id"`

	// A string containing the authToken for the subscription confirmation from the endpoint side.
	Url *string `mandatory:"true" json:"url"`
}

func (m SubscriptionConfirmation) String() string {
	return common.PointerString(m)
}
