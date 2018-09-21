// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// EventsControlService API
//
// This service exposes APIs to create, update and delete Rules. Rules are used to tap into the Events stream.
//

package cloudevents

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// NotificationServiceAction delivery to an Oracle Notification Service topic
type NotificationServiceAction struct {

	// OCID of the action
	Id *string `mandatory:"true" json:"id"`

	// Service-generated message about the current state of this rule
	LifecycleMessage *string `mandatory:"true" json:"lifecycleMessage"`

	// OCID of the topic to deliver messages to
	TopicId *string `mandatory:"false" json:"topicId"`

	// specifies current state of the action
	LifecycleState ActionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

//GetId returns Id
func (m NotificationServiceAction) GetId() *string {
	return m.Id
}

//GetLifecycleMessage returns LifecycleMessage
func (m NotificationServiceAction) GetLifecycleMessage() *string {
	return m.LifecycleMessage
}

//GetLifecycleState returns LifecycleState
func (m NotificationServiceAction) GetLifecycleState() ActionLifecycleStateEnum {
	return m.LifecycleState
}

func (m NotificationServiceAction) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m NotificationServiceAction) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeNotificationServiceAction NotificationServiceAction
	s := struct {
		DiscriminatorParam string `json:"actionType"`
		MarshalTypeNotificationServiceAction
	}{
		"ONS",
		(MarshalTypeNotificationServiceAction)(m),
	}

	return json.Marshal(&s)
}
