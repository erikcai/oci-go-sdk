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

// ConfigParamDefinition This object represents the configurable properties for an object type.
type ConfigParamDefinition struct {
	ParamType BaseType `mandatory:"false" json:"paramType"`

	// paramName
	ParamName *string `mandatory:"false" json:"paramName"`

	// Descriptive text for the object.
	Description *string `mandatory:"false" json:"description"`

	// auto generated description
	DefaultValue *interface{} `mandatory:"false" json:"defaultValue"`

	// classFieldName
	ClassFieldName *string `mandatory:"false" json:"classFieldName"`

	// static
	IsStatic *bool `mandatory:"false" json:"isStatic"`

	// classFieldValue
	IsClassFieldValue *bool `mandatory:"false" json:"isClassFieldValue"`
}

func (m ConfigParamDefinition) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *ConfigParamDefinition) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ParamType         basetype     `json:"paramType"`
		ParamName         *string      `json:"paramName"`
		Description       *string      `json:"description"`
		DefaultValue      *interface{} `json:"defaultValue"`
		ClassFieldName    *string      `json:"classFieldName"`
		IsStatic          *bool        `json:"isStatic"`
		IsClassFieldValue *bool        `json:"isClassFieldValue"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.ParamType.UnmarshalPolymorphicJSON(model.ParamType.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ParamType = nn.(BaseType)
	} else {
		m.ParamType = nil
	}

	m.ParamName = model.ParamName

	m.Description = model.Description

	m.DefaultValue = model.DefaultValue

	m.ClassFieldName = model.ClassFieldName

	m.IsStatic = model.IsStatic

	m.IsClassFieldValue = model.IsClassFieldValue
	return
}
