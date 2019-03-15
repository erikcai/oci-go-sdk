// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// CloudEvents API
//
// API for the CloudEvents Service. Use this API to manage rules and actions that create automation
// in your tenancy. For more information, see Overview of Events (https://docs.cloud.oracle.com/iaas/Content/Events/Concepts/eventsoverview.htm).
//

package cloudevents

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// ActionList A list of Action objects associated with a rule.
type ActionList struct {

	// A list of one or more Action objects.
	Actions []Action `mandatory:"true" json:"actions"`
}

func (m ActionList) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *ActionList) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Actions []action `json:"actions"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	m.Actions = make([]Action, len(model.Actions))
	for i, n := range model.Actions {
		nn, err := n.UnmarshalPolymorphicJSON(n.JsonData)
		if err != nil {
			return err
		}
		if nn != nil {
			m.Actions[i] = nn.(Action)
		} else {
			m.Actions[i] = nil
		}
	}
	return
}
