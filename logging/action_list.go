// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// PublicLoggingControlplane API
//
// PublicLoggingControlplane API specification
//

package logging

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// ActionList A list of Action objects associated with a Rule.
type ActionList struct {

	// A list of one or more Action objects.
	Actions []ActionDetails `mandatory:"true" json:"actions"`
}

func (m ActionList) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *ActionList) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Actions []actiondetails `json:"actions"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Actions = make([]ActionDetails, len(model.Actions))
	for i, n := range model.Actions {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Actions[i] = nn.(ActionDetails)
		} else {
			m.Actions[i] = nil
		}
	}
	return
}
