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

// ActionDetailsList List of ActionDetails object.
type ActionDetailsList struct {

	// List of action.
	Actions []ActionDetails `mandatory:"true" json:"actions"`
}

func (m ActionDetailsList) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *ActionDetailsList) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Actions []actiondetails `json:"actions"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	m.Actions = make([]ActionDetails, len(model.Actions))
	for i, n := range model.Actions {
		nn, err := n.UnmarshalPolymorphicJSON(n.JsonData)
		if err != nil {
			return err
		}
		if nn != nil {
			m.Actions[i] = nn.(ActionDetails)
		} else {
			m.Actions[i] = nil
		}
	}
	return
}
