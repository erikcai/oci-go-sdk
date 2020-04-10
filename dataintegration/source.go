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

// Source auto generated description
type Source struct {

	// Object key
	Key *string `mandatory:"false" json:"key"`

	// modelVersion
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value can be edited by the user and it is restricted to 1000 characters
	Name *string `mandatory:"false" json:"name"`

	// Descriptive text for the object.
	Description *string `mandatory:"false" json:"description"`

	// The version of the object, to track changes in the object instance
	ObjectVersion *int `mandatory:"false" json:"objectVersion"`

	// inputPorts
	InputPorts []InputPort `mandatory:"false" json:"inputPorts"`

	// outputPorts
	OutputPorts []OutputPort `mandatory:"false" json:"outputPorts"`

	// Status of object, can set this to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Value can only contain upper case letters, underscore and numbers. It should begin with upper case letter or underscore. The value can be edited by the user.
	Identifier *string `mandatory:"false" json:"identifier"`

	// parameters
	Parameters []Parameter `mandatory:"false" json:"parameters"`

	OpConfigValues *ConfigValues `mandatory:"false" json:"opConfigValues"`

	// The referenced data entity
	Entity *string `mandatory:"false" json:"entity"`

	Binding DataEntityBinding `mandatory:"false" json:"binding"`

	// readAccess
	IsReadAccess *bool `mandatory:"false" json:"isReadAccess"`

	// copyFields
	IsCopyFields *bool `mandatory:"false" json:"isCopyFields"`

	// predefinedShape
	IsPredefinedShape *bool `mandatory:"false" json:"isPredefinedShape"`
}

//GetKey returns Key
func (m Source) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m Source) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m Source) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetName returns Name
func (m Source) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m Source) GetDescription() *string {
	return m.Description
}

//GetObjectVersion returns ObjectVersion
func (m Source) GetObjectVersion() *int {
	return m.ObjectVersion
}

//GetInputPorts returns InputPorts
func (m Source) GetInputPorts() []InputPort {
	return m.InputPorts
}

//GetOutputPorts returns OutputPorts
func (m Source) GetOutputPorts() []OutputPort {
	return m.OutputPorts
}

//GetObjectStatus returns ObjectStatus
func (m Source) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetIdentifier returns Identifier
func (m Source) GetIdentifier() *string {
	return m.Identifier
}

//GetParameters returns Parameters
func (m Source) GetParameters() []Parameter {
	return m.Parameters
}

//GetOpConfigValues returns OpConfigValues
func (m Source) GetOpConfigValues() *ConfigValues {
	return m.OpConfigValues
}

func (m Source) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m Source) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSource Source
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeSource
	}{
		"SOURCE_OPERATOR",
		(MarshalTypeSource)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *Source) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Key               *string           `json:"key"`
		ModelVersion      *string           `json:"modelVersion"`
		ParentRef         *ParentReference  `json:"parentRef"`
		Name              *string           `json:"name"`
		Description       *string           `json:"description"`
		ObjectVersion     *int              `json:"objectVersion"`
		InputPorts        []InputPort       `json:"inputPorts"`
		OutputPorts       []OutputPort      `json:"outputPorts"`
		ObjectStatus      *int              `json:"objectStatus"`
		Identifier        *string           `json:"identifier"`
		Parameters        []Parameter       `json:"parameters"`
		OpConfigValues    *ConfigValues     `json:"opConfigValues"`
		Entity            *string           `json:"entity"`
		Binding           dataentitybinding `json:"binding"`
		IsReadAccess      *bool             `json:"isReadAccess"`
		IsCopyFields      *bool             `json:"isCopyFields"`
		IsPredefinedShape *bool             `json:"isPredefinedShape"`
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

	m.Description = model.Description

	m.ObjectVersion = model.ObjectVersion

	m.InputPorts = make([]InputPort, len(model.InputPorts))
	for i, n := range model.InputPorts {
		m.InputPorts[i] = n
	}

	m.OutputPorts = make([]OutputPort, len(model.OutputPorts))
	for i, n := range model.OutputPorts {
		m.OutputPorts[i] = n
	}

	m.ObjectStatus = model.ObjectStatus

	m.Identifier = model.Identifier

	m.Parameters = make([]Parameter, len(model.Parameters))
	for i, n := range model.Parameters {
		m.Parameters[i] = n
	}

	m.OpConfigValues = model.OpConfigValues

	m.Entity = model.Entity

	nn, e = model.Binding.UnmarshalPolymorphicJSON(model.Binding.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Binding = nn.(DataEntityBinding)
	} else {
		m.Binding = nil
	}

	m.IsReadAccess = model.IsReadAccess

	m.IsCopyFields = model.IsCopyFields

	m.IsPredefinedShape = model.IsPredefinedShape
	return
}
