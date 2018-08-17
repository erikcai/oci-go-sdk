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

// SubscriptionSummary Encapsulates a subscription.
type SubscriptionSummary struct {

	// The OCID of the subscription.
	Id *string `mandatory:"false" json:"id"`

	// The OCID of associated topic.
	TopicId *string `mandatory:"false" json:"topicId"`

	// The protocol used for the subscription. Valid values include EMAIL, or HTTPS.
	Protocol *string `mandatory:"false" json:"protocol"`

	// The endpoint for the subscription. Valid values are an email address or an HTTP URL.
	Endpoint *string `mandatory:"false" json:"endpoint"`

	// The lifecycle state of the subscription. Valid values are PENDING, ACTIVE, or DELETED.
	LifecycleState SubscriptionSummaryLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The time when this suscription was created.
	CreatedTime *int64 `mandatory:"false" json:"createdTime"`

	// Used for pre-conditional change.
	Etag *string `mandatory:"false" json:"etag"`
}

func (m SubscriptionSummary) String() string {
	return common.PointerString(m)
}

// SubscriptionSummaryLifecycleStateEnum Enum with underlying type: string
type SubscriptionSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for SubscriptionSummaryLifecycleState
const (
	SubscriptionSummaryLifecycleStatePending SubscriptionSummaryLifecycleStateEnum = "PENDING"
	SubscriptionSummaryLifecycleStateActive  SubscriptionSummaryLifecycleStateEnum = "ACTIVE"
	SubscriptionSummaryLifecycleStateDeleted SubscriptionSummaryLifecycleStateEnum = "DELETED"
)

var mappingSubscriptionSummaryLifecycleState = map[string]SubscriptionSummaryLifecycleStateEnum{
	"PENDING": SubscriptionSummaryLifecycleStatePending,
	"ACTIVE":  SubscriptionSummaryLifecycleStateActive,
	"DELETED": SubscriptionSummaryLifecycleStateDeleted,
}

// GetSubscriptionSummaryLifecycleStateEnumValues Enumerates the set of values for SubscriptionSummaryLifecycleState
func GetSubscriptionSummaryLifecycleStateEnumValues() []SubscriptionSummaryLifecycleStateEnum {
	values := make([]SubscriptionSummaryLifecycleStateEnum, 0)
	for _, v := range mappingSubscriptionSummaryLifecycleState {
		values = append(values, v)
	}
	return values
}
