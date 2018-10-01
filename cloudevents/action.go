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

// Action Detailed representation of a action.
type Action interface {

	// OCID of the action
	GetId() *string

	// Service-generated message about the current state of this rule
	GetLifecycleMessage() *string

	// specifies current state of the action
	GetLifecycleState() ActionLifecycleStateEnum
}

type action struct {
	JsonData         []byte
	Id               *string                  `mandatory:"true" json:"id"`
	LifecycleMessage *string                  `mandatory:"true" json:"lifecycleMessage"`
	LifecycleState   ActionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
	ActionType       string                   `json:"actionType"`
}

// UnmarshalJSON unmarshals json
func (m *action) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshaleraction action
	s := struct {
		Model Unmarshaleraction
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.LifecycleMessage = s.Model.LifecycleMessage
	m.LifecycleState = s.Model.LifecycleState
	m.ActionType = s.Model.ActionType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *action) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ActionType {
	case "OSS":
		mm := StreamingServiceAction{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OBJECTSTORAGE":
		mm := ObjectStorageServiceAction{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ONS":
		mm := NotificationServiceAction{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MS_EVENT_GRID":
		mm := EventGridAction{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FAAS":
		mm := FaaSAction{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "NOOP":
		mm := NoOpAction{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetId returns Id
func (m action) GetId() *string {
	return m.Id
}

//GetLifecycleMessage returns LifecycleMessage
func (m action) GetLifecycleMessage() *string {
	return m.LifecycleMessage
}

//GetLifecycleState returns LifecycleState
func (m action) GetLifecycleState() ActionLifecycleStateEnum {
	return m.LifecycleState
}

func (m action) String() string {
	return common.PointerString(m)
}

// ActionLifecycleStateEnum Enum with underlying type: string
type ActionLifecycleStateEnum string

// Set of constants representing the allowable values for ActionLifecycleStateEnum
const (
	ActionLifecycleStateCreating ActionLifecycleStateEnum = "CREATING"
	ActionLifecycleStateActive   ActionLifecycleStateEnum = "ACTIVE"
	ActionLifecycleStateInactive ActionLifecycleStateEnum = "INACTIVE"
	ActionLifecycleStateUpdating ActionLifecycleStateEnum = "UPDATING"
	ActionLifecycleStateDeleting ActionLifecycleStateEnum = "DELETING"
	ActionLifecycleStateDeleted  ActionLifecycleStateEnum = "DELETED"
	ActionLifecycleStateFailed   ActionLifecycleStateEnum = "FAILED"
)

var mappingActionLifecycleState = map[string]ActionLifecycleStateEnum{
	"CREATING": ActionLifecycleStateCreating,
	"ACTIVE":   ActionLifecycleStateActive,
	"INACTIVE": ActionLifecycleStateInactive,
	"UPDATING": ActionLifecycleStateUpdating,
	"DELETING": ActionLifecycleStateDeleting,
	"DELETED":  ActionLifecycleStateDeleted,
	"FAILED":   ActionLifecycleStateFailed,
}

// GetActionLifecycleStateEnumValues Enumerates the set of values for ActionLifecycleStateEnum
func GetActionLifecycleStateEnumValues() []ActionLifecycleStateEnum {
	values := make([]ActionLifecycleStateEnum, 0)
	for _, v := range mappingActionLifecycleState {
		values = append(values, v)
	}
	return values
}
