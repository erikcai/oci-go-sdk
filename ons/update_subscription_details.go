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

// UpdateSubscriptionDetails Contains details about a subscription.
type UpdateSubscriptionDetails struct {

	// The delivery policy of the subscription. Stored as a JSON string.
	DeliveryPolicy *DeliveryPolicy `mandatory:"false" json:"deliveryPolicy"`
}

func (m UpdateSubscriptionDetails) String() string {
	return common.PointerString(m)
}
