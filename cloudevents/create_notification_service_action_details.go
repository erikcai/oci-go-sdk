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

// CreateNotificationServiceActionDetails delivery to an Oracle Notification Service topic
type CreateNotificationServiceActionDetails struct {

	// whether or not this aciton is currently enabled
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// OCID of the topic to deliver messages to
	TopicId *string `mandatory:"false" json:"topicId"`
}

//GetIsEnabled returns IsEnabled
func (m CreateNotificationServiceActionDetails) GetIsEnabled() *bool {
	return m.IsEnabled
}

func (m CreateNotificationServiceActionDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateNotificationServiceActionDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateNotificationServiceActionDetails CreateNotificationServiceActionDetails
	s := struct {
		DiscriminatorParam string `json:"actionType"`
		MarshalTypeCreateNotificationServiceActionDetails
	}{
		"ONS",
		(MarshalTypeCreateNotificationServiceActionDetails)(m),
	}

	return json.Marshal(&s)
}
