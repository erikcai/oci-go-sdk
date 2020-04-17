// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
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

// RuleTypeConfig auto generated description
type RuleTypeConfig struct {

	// Object key
	Key *string `mandatory:"false" json:"key"`

	// modelVersion
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// reference to a typed object
	Scope *string `mandatory:"false" json:"scope"`

	// orderByRule
	IsOrderByRule *bool `mandatory:"false" json:"isOrderByRule"`

	// projectionRules
	ProjectionRules []ProjectionRule `mandatory:"false" json:"projectionRules"`

	ConfigValues *ConfigValues `mandatory:"false" json:"configValues"`

	// Status of object, can set this to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// classType
	ClassType *string `mandatory:"false" json:"classType"`
}

func (m RuleTypeConfig) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m RuleTypeConfig) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeRuleTypeConfig RuleTypeConfig
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeRuleTypeConfig
	}{
		"RULE_TYPE_CONFIGS",
		(MarshalTypeRuleTypeConfig)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *RuleTypeConfig) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Key             *string          `json:"key"`
		ModelVersion    *string          `json:"modelVersion"`
		ParentRef       *ParentReference `json:"parentRef"`
		Scope           *string          `json:"scope"`
		IsOrderByRule   *bool            `json:"isOrderByRule"`
		ProjectionRules []projectionrule `json:"projectionRules"`
		ConfigValues    *ConfigValues    `json:"configValues"`
		ObjectStatus    *int             `json:"objectStatus"`
		ClassType       *string          `json:"classType"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Key = model.Key

	m.ModelVersion = model.ModelVersion

	m.ParentRef = model.ParentRef

	m.Scope = model.Scope

	m.IsOrderByRule = model.IsOrderByRule

	m.ProjectionRules = make([]ProjectionRule, len(model.ProjectionRules))
	for i, n := range model.ProjectionRules {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.ProjectionRules[i] = nn.(ProjectionRule)
		} else {
			m.ProjectionRules[i] = nil
		}
	}

	m.ConfigValues = model.ConfigValues

	m.ObjectStatus = model.ObjectStatus

	m.ClassType = model.ClassType
	return
}
