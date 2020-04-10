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

// DynamicInputField The type representing the dynamic field concept. Dynamic fields have a dynamic type handler to define how to generate the field.
type DynamicInputField struct {

	// Object key
	Key *string `mandatory:"false" json:"key"`

	// modelVersion
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	ConfigValues *ConfigValues `mandatory:"false" json:"configValues"`

	// Status of object, can set this to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value can be edited by the user and it is restricted to 1000 characters
	Name *string `mandatory:"false" json:"name"`

	// Descriptive text for the object.
	Description *string `mandatory:"false" json:"description"`

	Type BaseType `mandatory:"false" json:"type"`

	// Labels are keywords or labels that you can add to data assets, dataflows etc. You can define your own labels and use them to categorize content.
	Labels []string `mandatory:"false" json:"labels"`

	// isPrivate
	IsPrivate *bool `mandatory:"false" json:"isPrivate"`
}

//GetKey returns Key
func (m DynamicInputField) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m DynamicInputField) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m DynamicInputField) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetConfigValues returns ConfigValues
func (m DynamicInputField) GetConfigValues() *ConfigValues {
	return m.ConfigValues
}

//GetObjectStatus returns ObjectStatus
func (m DynamicInputField) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetName returns Name
func (m DynamicInputField) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m DynamicInputField) GetDescription() *string {
	return m.Description
}

func (m DynamicInputField) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m DynamicInputField) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDynamicInputField DynamicInputField
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeDynamicInputField
	}{
		"DYNAMIC_INPUT_FIELD",
		(MarshalTypeDynamicInputField)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *DynamicInputField) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Key          *string          `json:"key"`
		ModelVersion *string          `json:"modelVersion"`
		ParentRef    *ParentReference `json:"parentRef"`
		ConfigValues *ConfigValues    `json:"configValues"`
		ObjectStatus *int             `json:"objectStatus"`
		Name         *string          `json:"name"`
		Description  *string          `json:"description"`
		Type         basetype         `json:"type"`
		Labels       []string         `json:"labels"`
		IsPrivate    *bool            `json:"isPrivate"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Key = model.Key

	m.ModelVersion = model.ModelVersion

	m.ParentRef = model.ParentRef

	m.ConfigValues = model.ConfigValues

	m.ObjectStatus = model.ObjectStatus

	m.Name = model.Name

	m.Description = model.Description

	nn, e = model.Type.UnmarshalPolymorphicJSON(model.Type.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Type = nn.(BaseType)
	} else {
		m.Type = nil
	}

	m.Labels = make([]string, len(model.Labels))
	for i, n := range model.Labels {
		m.Labels[i] = n
	}

	m.IsPrivate = model.IsPrivate
	return
}
