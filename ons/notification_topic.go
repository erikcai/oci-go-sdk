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

// NotificationTopic notification topic
type NotificationTopic struct {

	// The name of the topic.
	Name *string `mandatory:"true" json:"name"`

	// The OCID of the compartment for the topic.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The lifecycle state of the topic. Valid values are ACTIVE, DELETING, or CREATING.
	LifecycleState NotificationTopicLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID of the topic.
	TopicOcid *string `mandatory:"false" json:"topicOcid"`

	// The description of the topic.
	Description *string `mandatory:"false" json:"description"`

	// The time the topic was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Used for pre-conditional update/delete.
	Etag *string `mandatory:"false" json:"etag"`
}

func (m NotificationTopic) String() string {
	return common.PointerString(m)
}

// NotificationTopicLifecycleStateEnum Enum with underlying type: string
type NotificationTopicLifecycleStateEnum string

// Set of constants representing the allowable values for NotificationTopicLifecycleState
const (
	NotificationTopicLifecycleStateActive   NotificationTopicLifecycleStateEnum = "ACTIVE"
	NotificationTopicLifecycleStateDeleting NotificationTopicLifecycleStateEnum = "DELETING"
	NotificationTopicLifecycleStateCreating NotificationTopicLifecycleStateEnum = "CREATING"
)

var mappingNotificationTopicLifecycleState = map[string]NotificationTopicLifecycleStateEnum{
	"ACTIVE":   NotificationTopicLifecycleStateActive,
	"DELETING": NotificationTopicLifecycleStateDeleting,
	"CREATING": NotificationTopicLifecycleStateCreating,
}

// GetNotificationTopicLifecycleStateEnumValues Enumerates the set of values for NotificationTopicLifecycleState
func GetNotificationTopicLifecycleStateEnumValues() []NotificationTopicLifecycleStateEnum {
	values := make([]NotificationTopicLifecycleStateEnum, 0)
	for _, v := range mappingNotificationTopicLifecycleState {
		values = append(values, v)
	}
	return values
}
