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

// ActionDetails Object used to create a action.
type ActionDetails interface {

	// whether or not this aciton is currently enabled
	GetIsEnabled() *bool
}

type actiondetails struct {
	JsonData   []byte
	IsEnabled  *bool  `mandatory:"true" json:"isEnabled"`
	ActionType string `json:"actionType"`
}

// UnmarshalJSON unmarshals json
func (m *actiondetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshaleractiondetails actiondetails
	s := struct {
		Model Unmarshaleractiondetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.IsEnabled = s.Model.IsEnabled
	m.ActionType = s.Model.ActionType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *actiondetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ActionType {
	case "OSS":
		mm := CreateStreamingServiceActionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FAAS":
		mm := CreateFaaSActionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "NOOP":
		mm := CreateNoOpActionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OBJECTSTORAGE":
		mm := CreateObjectStorageServiceActionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MS_EVENT_GRID":
		mm := CreateEventGridActionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ONS":
		mm := CreateNotificationServiceActionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetIsEnabled returns IsEnabled
func (m actiondetails) GetIsEnabled() *bool {
	return m.IsEnabled
}

func (m actiondetails) String() string {
	return common.PointerString(m)
}
