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

// NotificationTopicSummary Summary of the notification topic.
type NotificationTopicSummary struct {

	// The name of the notification.
	Name *string `mandatory:"true" json:"name"`

	// The OCID of the topic.
	TopicOcid *string `mandatory:"true" json:"topicOcid"`

	// The OCID of the compartment for the topic.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The lifecycle state of the topic. Valid values are ACTIVE, DELETING, or CREATING.
	LifecycleState NotificationTopicSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The description of the topic.
	Description *string `mandatory:"false" json:"description"`

	// The time the topic was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Used for pre-conditional update/delete.
	Etag *string `mandatory:"false" json:"etag"`
}

func (m NotificationTopicSummary) String() string {
	return common.PointerString(m)
}

// NotificationTopicSummaryLifecycleStateEnum Enum with underlying type: string
type NotificationTopicSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for NotificationTopicSummaryLifecycleState
const (
	NotificationTopicSummaryLifecycleStateActive   NotificationTopicSummaryLifecycleStateEnum = "ACTIVE"
	NotificationTopicSummaryLifecycleStateDeleting NotificationTopicSummaryLifecycleStateEnum = "DELETING"
	NotificationTopicSummaryLifecycleStateCreating NotificationTopicSummaryLifecycleStateEnum = "CREATING"
)

var mappingNotificationTopicSummaryLifecycleState = map[string]NotificationTopicSummaryLifecycleStateEnum{
	"ACTIVE":   NotificationTopicSummaryLifecycleStateActive,
	"DELETING": NotificationTopicSummaryLifecycleStateDeleting,
	"CREATING": NotificationTopicSummaryLifecycleStateCreating,
}

// GetNotificationTopicSummaryLifecycleStateEnumValues Enumerates the set of values for NotificationTopicSummaryLifecycleState
func GetNotificationTopicSummaryLifecycleStateEnumValues() []NotificationTopicSummaryLifecycleStateEnum {
	values := make([]NotificationTopicSummaryLifecycleStateEnum, 0)
	for _, v := range mappingNotificationTopicSummaryLifecycleState {
		values = append(values, v)
	}
	return values
}
