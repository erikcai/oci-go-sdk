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

// DynamicType auto generated description
type DynamicType struct {

	// Object key
	Key *string `mandatory:"false" json:"key"`

	// modelVersion
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value can be edited by the user and it is restricted to 1000 characters
	Name *string `mandatory:"false" json:"name"`

	// Status of object, can set this to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Descriptive text for the object.
	Description *string `mandatory:"false" json:"description"`

	TypeHandler DynamicTypeHandler `mandatory:"false" json:"typeHandler"`

	ConfigDef *ConfigDefinition `mandatory:"false" json:"configDef"`
}

//GetKey returns Key
func (m DynamicType) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m DynamicType) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m DynamicType) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetName returns Name
func (m DynamicType) GetName() *string {
	return m.Name
}

//GetObjectStatus returns ObjectStatus
func (m DynamicType) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetDescription returns Description
func (m DynamicType) GetDescription() *string {
	return m.Description
}

func (m DynamicType) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m DynamicType) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDynamicType DynamicType
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeDynamicType
	}{
		"DYNAMIC_TYPE",
		(MarshalTypeDynamicType)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *DynamicType) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Key          *string            `json:"key"`
		ModelVersion *string            `json:"modelVersion"`
		ParentRef    *ParentReference   `json:"parentRef"`
		Name         *string            `json:"name"`
		ObjectStatus *int               `json:"objectStatus"`
		Description  *string            `json:"description"`
		TypeHandler  dynamictypehandler `json:"typeHandler"`
		ConfigDef    *ConfigDefinition  `json:"configDef"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Key = model.Key

	m.ModelVersion = model.ModelVersion

	m.ParentRef = model.ParentRef

	m.Name = model.Name

	m.ObjectStatus = model.ObjectStatus

	m.Description = model.Description

	nn, e = model.TypeHandler.UnmarshalPolymorphicJSON(model.TypeHandler.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.TypeHandler = nn.(DynamicTypeHandler)
	} else {
		m.TypeHandler = nil
	}

	m.ConfigDef = model.ConfigDef
	return
}
