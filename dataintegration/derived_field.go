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

// DerivedField The type representing the derived field concept. Derived fields have an expression to define how to derive the field.
type DerivedField struct {

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

	Expr *Expression `mandatory:"false" json:"expr"`

	Type BaseType `mandatory:"false" json:"type"`

	// Labels are keywords or labels that you can add to data assets, dataflows etc. You can define your own labels and use them to categorize content.
	Labels []string `mandatory:"false" json:"labels"`

	// isPrivate
	IsPrivate *bool `mandatory:"false" json:"isPrivate"`
}

//GetKey returns Key
func (m DerivedField) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m DerivedField) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m DerivedField) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetConfigValues returns ConfigValues
func (m DerivedField) GetConfigValues() *ConfigValues {
	return m.ConfigValues
}

//GetObjectStatus returns ObjectStatus
func (m DerivedField) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetName returns Name
func (m DerivedField) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m DerivedField) GetDescription() *string {
	return m.Description
}

func (m DerivedField) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m DerivedField) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDerivedField DerivedField
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeDerivedField
	}{
		"DERIVED_FIELD",
		(MarshalTypeDerivedField)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *DerivedField) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Key          *string          `json:"key"`
		ModelVersion *string          `json:"modelVersion"`
		ParentRef    *ParentReference `json:"parentRef"`
		ConfigValues *ConfigValues    `json:"configValues"`
		ObjectStatus *int             `json:"objectStatus"`
		Name         *string          `json:"name"`
		Description  *string          `json:"description"`
		Expr         *Expression      `json:"expr"`
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

	m.Expr = model.Expr

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
