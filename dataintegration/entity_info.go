// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Data Integration Service API Specification
//
// Data Integration Service API Specification
//

package dataintegration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// EntityInfo auto generated description
type EntityInfo struct {
	SourceSchema *SchemaDetails `mandatory:"false" json:"sourceSchema"`

	SourceEntity DataEntityDetails `mandatory:"false" json:"sourceEntity"`

	TargetEntity DataEntityDetails `mandatory:"false" json:"targetEntity"`

	// sourceEntityFilterRules
	SourceEntityFilterRules []FilterRule `mandatory:"false" json:"sourceEntityFilterRules"`
}

func (m EntityInfo) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *EntityInfo) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		SourceSchema            *SchemaDetails    `json:"sourceSchema"`
		SourceEntity            dataentitydetails `json:"sourceEntity"`
		TargetEntity            dataentitydetails `json:"targetEntity"`
		SourceEntityFilterRules []FilterRule      `json:"sourceEntityFilterRules"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.SourceSchema = model.SourceSchema

	nn, e = model.SourceEntity.UnmarshalPolymorphicJSON(model.SourceEntity.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.SourceEntity = nn.(DataEntityDetails)
	} else {
		m.SourceEntity = nil
	}

	nn, e = model.TargetEntity.UnmarshalPolymorphicJSON(model.TargetEntity.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.TargetEntity = nn.(DataEntityDetails)
	} else {
		m.TargetEntity = nil
	}

	m.SourceEntityFilterRules = make([]FilterRule, len(model.SourceEntityFilterRules))
	for i, n := range model.SourceEntityFilterRules {
		m.SourceEntityFilterRules[i] = n
	}
	return
}
