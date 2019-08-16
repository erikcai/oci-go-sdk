// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Notifications API
//
// Use the Notifications API to broadcast messages to distributed components by topic, using a publish-subscribe pattern.
// For information about managing topics, subscriptions, and messages, see Notifications Overview (https://docs.cloud.oracle.com/iaas/Content/Notification/Concepts/notificationoverview.htm).
//

package ons

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Subscription The subscription's configuration.
type Subscription struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated topic.
	TopicId *string `mandatory:"true" json:"topicId"`

	// The protocol used for the subscription. Valid values: EMAIL, HTTPS.
	Protocol *string `mandatory:"true" json:"protocol"`

	// The endpoint of the subscription. Valid values depend on the protocol.
	// For EMAIL, only an email address is valid. For HTTPS, only a URL or address corresponding to the provider is valid. A URL cannot exceed 512 characters.
	Endpoint *string `mandatory:"true" json:"endpoint"`

	// The lifecycle state of the subscription. The status of a new subscription is PENDING; when confirmed, the subscription status changes to ACTIVE.
	LifecycleState SubscriptionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for the subscription.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The time when this suscription was created.
	CreatedTime *int64 `mandatory:"false" json:"createdTime"`

	// The delivery policy of the subscription. Stored as a JSON string.
	DeliverPolicy *string `mandatory:"false" json:"deliverPolicy"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `mandatory:"false" json:"etag"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m Subscription) String() string {
	return common.PointerString(m)
}

// SubscriptionLifecycleStateEnum Enum with underlying type: string
type SubscriptionLifecycleStateEnum string

// Set of constants representing the allowable values for SubscriptionLifecycleStateEnum
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

// GetSubscriptionLifecycleStateEnumValues Enumerates the set of values for SubscriptionLifecycleStateEnum
func GetSubscriptionLifecycleStateEnumValues() []SubscriptionLifecycleStateEnum {
	values := make([]SubscriptionLifecycleStateEnum, 0)
	for _, v := range mappingSubscriptionLifecycleState {
		values = append(values, v)
	}
	return values
}
