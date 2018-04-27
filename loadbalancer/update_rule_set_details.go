// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Load Balancing Service API
//
// API for the Load Balancing Service
//

package loadbalancer

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateRuleSetDetails The representation of UpdateRuleSetDetails
type UpdateRuleSetDetails struct {
	Items []Rule `mandatory:"true" json:"items"`
}

func (m UpdateRuleSetDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *UpdateRuleSetDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Items []rule `json:"items"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	m.Items = make([]Rule, len(model.Items))
	for i, n := range model.Items {
		nn, err := n.UnmarshalPolymorphicJSON(n.JsonData)
		if err != nil {
			return err
		}
		m.Items[i] = nn.(Rule)
	}
	return
}
