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

// NoOpAction delivery to logs
type NoOpAction struct {

	// OCID of the action
	Id *string `mandatory:"true" json:"id"`

	// Service-generated message about the current state of this rule
	LifecycleMessage *string `mandatory:"true" json:"lifecycleMessage"`

	// name of the noop action; will appear in the logs
	DisplayName *string `mandatory:"false" json:"displayName"`

	// specifies current state of the action
	LifecycleState ActionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

//GetId returns Id
func (m NoOpAction) GetId() *string {
	return m.Id
}

//GetLifecycleMessage returns LifecycleMessage
func (m NoOpAction) GetLifecycleMessage() *string {
	return m.LifecycleMessage
}

//GetLifecycleState returns LifecycleState
func (m NoOpAction) GetLifecycleState() ActionLifecycleStateEnum {
	return m.LifecycleState
}

func (m NoOpAction) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m NoOpAction) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeNoOpAction NoOpAction
	s := struct {
		DiscriminatorParam string `json:"actionType"`
		MarshalTypeNoOpAction
	}{
		"NOOP",
		(MarshalTypeNoOpAction)(m),
	}

	return json.Marshal(&s)
}
