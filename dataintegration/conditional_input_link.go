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

// ConditionalInputLink auto generated description
type ConditionalInputLink struct {

	// Object key
	Key *string `mandatory:"false" json:"key"`

	// modelVersion
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Status of object, can set this to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Descriptive text for the object.
	Description *string `mandatory:"false" json:"description"`

	// Key of FlowPort reference
	Port *string `mandatory:"false" json:"port"`

	FromLink *OutputLink `mandatory:"false" json:"fromLink"`

	FieldMap FieldMap `mandatory:"false" json:"fieldMap"`

	Condition *Expression `mandatory:"false" json:"condition"`
}

//GetKey returns Key
func (m ConditionalInputLink) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m ConditionalInputLink) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m ConditionalInputLink) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetObjectStatus returns ObjectStatus
func (m ConditionalInputLink) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetDescription returns Description
func (m ConditionalInputLink) GetDescription() *string {
	return m.Description
}

//GetPort returns Port
func (m ConditionalInputLink) GetPort() *string {
	return m.Port
}

func (m ConditionalInputLink) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ConditionalInputLink) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeConditionalInputLink ConditionalInputLink
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeConditionalInputLink
	}{
		"CONDITIONAL_INPUT_LINK",
		(MarshalTypeConditionalInputLink)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *ConditionalInputLink) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Key          *string          `json:"key"`
		ModelVersion *string          `json:"modelVersion"`
		ParentRef    *ParentReference `json:"parentRef"`
		ObjectStatus *int             `json:"objectStatus"`
		Description  *string          `json:"description"`
		Port         *string          `json:"port"`
		FromLink     *OutputLink      `json:"fromLink"`
		FieldMap     fieldmap         `json:"fieldMap"`
		Condition    *Expression      `json:"condition"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Key = model.Key

	m.ModelVersion = model.ModelVersion

	m.ParentRef = model.ParentRef

	m.ObjectStatus = model.ObjectStatus

	m.Description = model.Description

	m.Port = model.Port

	m.FromLink = model.FromLink

	nn, e = model.FieldMap.UnmarshalPolymorphicJSON(model.FieldMap.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.FieldMap = nn.(FieldMap)
	} else {
		m.FieldMap = nil
	}

	m.Condition = model.Condition
	return
}
