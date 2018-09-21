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

// FaaSAction Delivery to a Function-as-a-service endpoint
type FaaSAction struct {

	// OCID of the action
	Id *string `mandatory:"true" json:"id"`

	// Service-generated message about the current state of this rule
	LifecycleMessage *string `mandatory:"true" json:"lifecycleMessage"`

	// OCID of the function to call
	FunctionId *string `mandatory:"false" json:"functionId"`

	// specifies current state of the action
	LifecycleState ActionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

//GetId returns Id
func (m FaaSAction) GetId() *string {
	return m.Id
}

//GetLifecycleMessage returns LifecycleMessage
func (m FaaSAction) GetLifecycleMessage() *string {
	return m.LifecycleMessage
}

//GetLifecycleState returns LifecycleState
func (m FaaSAction) GetLifecycleState() ActionLifecycleStateEnum {
	return m.LifecycleState
}

func (m FaaSAction) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m FaaSAction) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeFaaSAction FaaSAction
	s := struct {
		DiscriminatorParam string `json:"actionType"`
		MarshalTypeFaaSAction
	}{
		"FAAS",
		(MarshalTypeFaaSAction)(m),
	}

	return json.Marshal(&s)
}
