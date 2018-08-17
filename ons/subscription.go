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

// Subscription subscription
type Subscription struct {

	// The OCID of the subscription.
	Id *string `mandatory:"false" json:"id"`

	// The subscription protocol. Valid values include EMAIL, or HTTPS.
	Protocol *string `mandatory:"false" json:"protocol"`

	// The endpoint for the subscription. Valid values are an email address or an HTTP URL.
	Endpoint *string `mandatory:"false" json:"endpoint"`

	// The lifecycle state of the subscription. Valid values are PENDING, ACTIVE, or DELETED.
	LifecycleState SubscriptionLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The delivery policy of the subscription. Stored as a JSON string.
	DeliverPolicy *string `mandatory:"false" json:"deliverPolicy"`

	// Used for pre-conditional updating.
	Etag *string `mandatory:"false" json:"etag"`
}

func (m Subscription) String() string {
	return common.PointerString(m)
}

// SubscriptionLifecycleStateEnum Enum with underlying type: string
type SubscriptionLifecycleStateEnum string

// Set of constants representing the allowable values for SubscriptionLifecycleState
const (
	SubscriptionLifecycleStatePending SubscriptionLifecycleStateEnum = "PENDING"
	SubscriptionLifecycleStateActive  SubscriptionLifecycleStateEnum = "ACTIVE"
	SubscriptionLifecycleStateDeleted SubscriptionLifecycleStateEnum = "DELETED"
)

var mappingSubscriptionLifecycleState = map[string]SubscriptionLifecycleStateEnum{
	"PENDING": SubscriptionLifecycleStatePending,
	"ACTIVE":  SubscriptionLifecycleStateActive,
	"DELETED": SubscriptionLifecycleStateDeleted,
}

// GetSubscriptionLifecycleStateEnumValues Enumerates the set of values for SubscriptionLifecycleState
func GetSubscriptionLifecycleStateEnumValues() []SubscriptionLifecycleStateEnum {
	values := make([]SubscriptionLifecycleStateEnum, 0)
	for _, v := range mappingSubscriptionLifecycleState {
		values = append(values, v)
	}
	return values
}
