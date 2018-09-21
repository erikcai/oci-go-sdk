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

// StreamingServiceAction delivery an Oracle Stream Service stream
type StreamingServiceAction struct {

	// OCID of the action
	Id *string `mandatory:"true" json:"id"`

	// Service-generated message about the current state of this rule
	LifecycleMessage *string `mandatory:"true" json:"lifecycleMessage"`

	// OCID of the stream to deliver messages to
	StreamId *string `mandatory:"false" json:"streamId"`

	// specifies current state of the action
	LifecycleState ActionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

//GetId returns Id
func (m StreamingServiceAction) GetId() *string {
	return m.Id
}

//GetLifecycleMessage returns LifecycleMessage
func (m StreamingServiceAction) GetLifecycleMessage() *string {
	return m.LifecycleMessage
}

//GetLifecycleState returns LifecycleState
func (m StreamingServiceAction) GetLifecycleState() ActionLifecycleStateEnum {
	return m.LifecycleState
}

func (m StreamingServiceAction) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m StreamingServiceAction) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeStreamingServiceAction StreamingServiceAction
	s := struct {
		DiscriminatorParam string `json:"actionType"`
		MarshalTypeStreamingServiceAction
	}{
		"OSS",
		(MarshalTypeStreamingServiceAction)(m),
	}

	return json.Marshal(&s)
}
