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

// DeliveryPolicy Subscription delivery policy
type DeliveryPolicy struct {
	BackoffRetryPolicy *BackoffRetryPolicy `mandatory:"false" json:"backoffRetryPolicy"`
}

func (m DeliveryPolicy) String() string {
	return common.PointerString(m)
}
