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

// CreateEventGridActionDetails Delivery Microsoft EventGrid topic
type CreateEventGridActionDetails struct {

	// whether or not this aciton is currently enabled
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// Key for the topic
	TopicKey *string `mandatory:"false" json:"topicKey"`

	// URL of the topic to deliver to
	TopicUrl *string `mandatory:"false" json:"topicUrl"`
}

//GetIsEnabled returns IsEnabled
func (m CreateEventGridActionDetails) GetIsEnabled() *bool {
	return m.IsEnabled
}

func (m CreateEventGridActionDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateEventGridActionDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateEventGridActionDetails CreateEventGridActionDetails
	s := struct {
		DiscriminatorParam string `json:"actionType"`
		MarshalTypeCreateEventGridActionDetails
	}{
		"MS_EVENT_GRID",
		(MarshalTypeCreateEventGridActionDetails)(m),
	}

	return json.Marshal(&s)
}
