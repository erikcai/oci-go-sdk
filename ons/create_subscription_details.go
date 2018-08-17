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

// CreateSubscriptionDetails Populates subscription details.
type CreateSubscriptionDetails struct {

	// The OCID of the topic for the subscription.
	TopicOcid *string `mandatory:"true" json:"topicOcid"`

	// The OCID of the compartment for the subscription.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The protocol to use to deliver messages. Valid values include EMAIL, or HTTPS.
	Protocol *string `mandatory:"true" json:"protocol"`

	// The endpoint of the subscription. The URL length can not be more than 512 characters. For the HTTPS protocol, only PagerDuty URLs are valid.
	Endpoint *string `mandatory:"true" json:"endpoint"`

	// Metadata for the subscription.
	Metadata *string `mandatory:"false" json:"metadata"`
}

func (m CreateSubscriptionDetails) String() string {
	return common.PointerString(m)
}
