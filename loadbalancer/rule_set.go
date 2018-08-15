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

// RuleSet The representation of RuleSet
type RuleSet struct {
	Items []Rule `mandatory:"true" json:"items"`

	// The name for this set of rules. It must be unique and it cannot be changed. Avoid entering
	// confidential information.
	// Example: `example_rule_set`
	Name *string `mandatory:"true" json:"name"`
}

func (m RuleSet) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *RuleSet) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Items []rule  `json:"items"`
		Name  *string `json:"name"`
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
		if nn != nil {
			m.Items[i] = nn.(Rule)
		} else {
			m.Items[i] = nil
		}
	}
	m.Name = model.Name
	return
}
