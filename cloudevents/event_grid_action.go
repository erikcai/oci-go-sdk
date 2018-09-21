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

// EventGridAction Delivery Microsoft EventGrid topic
type EventGridAction struct {

	// OCID of the action
	Id *string `mandatory:"true" json:"id"`

	// Service-generated message about the current state of this rule
	LifecycleMessage *string `mandatory:"true" json:"lifecycleMessage"`

	// Key for the topic
	TopicKey *string `mandatory:"false" json:"topicKey"`

	// URL of the topic to deliver to
	TopicUrl *string `mandatory:"false" json:"topicUrl"`

	// specifies current state of the action
	LifecycleState ActionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

//GetId returns Id
func (m EventGridAction) GetId() *string {
	return m.Id
}

//GetLifecycleMessage returns LifecycleMessage
func (m EventGridAction) GetLifecycleMessage() *string {
	return m.LifecycleMessage
}

//GetLifecycleState returns LifecycleState
func (m EventGridAction) GetLifecycleState() ActionLifecycleStateEnum {
	return m.LifecycleState
}

func (m EventGridAction) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m EventGridAction) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeEventGridAction EventGridAction
	s := struct {
		DiscriminatorParam string `json:"actionType"`
		MarshalTypeEventGridAction
	}{
		"MS_EVENT_GRID",
		(MarshalTypeEventGridAction)(m),
	}

	return json.Marshal(&s)
}
