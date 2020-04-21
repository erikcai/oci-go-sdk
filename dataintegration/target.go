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

// Target auto generated description
type Target struct {

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

	// The reference to target data entity
	Entity *string `mandatory:"false" json:"entity"`

	Binding DataEntityBinding `mandatory:"false" json:"binding"`

	// readAccess
	IsReadAccess *bool `mandatory:"false" json:"isReadAccess"`

	// copyFields
	IsCopyFields *bool `mandatory:"false" json:"isCopyFields"`

	// predefinedShape
	IsPredefinedShape *bool `mandatory:"false" json:"isPredefinedShape"`

	WriteOperationConfig *WriteOperationConfig `mandatory:"false" json:"writeOperationConfig"`

	// dataProperty
	DataProperty TargetDataPropertyEnum `mandatory:"false" json:"dataProperty,omitempty"`
}

//GetKey returns Key
func (m Target) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m Target) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m Target) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetName returns Name
func (m Target) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m Target) GetDescription() *string {
	return m.Description
}

//GetObjectVersion returns ObjectVersion
func (m Target) GetObjectVersion() *int {
	return m.ObjectVersion
}

//GetInputPorts returns InputPorts
func (m Target) GetInputPorts() []InputPort {
	return m.InputPorts
}

//GetOutputPorts returns OutputPorts
func (m Target) GetOutputPorts() []OutputPort {
	return m.OutputPorts
}

//GetObjectStatus returns ObjectStatus
func (m Target) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetIdentifier returns Identifier
func (m Target) GetIdentifier() *string {
	return m.Identifier
}

//GetParameters returns Parameters
func (m Target) GetParameters() []Parameter {
	return m.Parameters
}

//GetOpConfigValues returns OpConfigValues
func (m Target) GetOpConfigValues() *ConfigValues {
	return m.OpConfigValues
}

func (m Target) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m Target) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTarget Target
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeTarget
	}{
		"TARGET_OPERATOR",
		(MarshalTypeTarget)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *Target) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Key                  *string                `json:"key"`
		ModelVersion         *string                `json:"modelVersion"`
		ParentRef            *ParentReference       `json:"parentRef"`
		Name                 *string                `json:"name"`
		Description          *string                `json:"description"`
		ObjectVersion        *int                   `json:"objectVersion"`
		InputPorts           []InputPort            `json:"inputPorts"`
		OutputPorts          []OutputPort           `json:"outputPorts"`
		ObjectStatus         *int                   `json:"objectStatus"`
		Identifier           *string                `json:"identifier"`
		Parameters           []Parameter            `json:"parameters"`
		OpConfigValues       *ConfigValues          `json:"opConfigValues"`
		Entity               *string                `json:"entity"`
		Binding              dataentitybinding      `json:"binding"`
		IsReadAccess         *bool                  `json:"isReadAccess"`
		IsCopyFields         *bool                  `json:"isCopyFields"`
		IsPredefinedShape    *bool                  `json:"isPredefinedShape"`
		DataProperty         TargetDataPropertyEnum `json:"dataProperty"`
		WriteOperationConfig *WriteOperationConfig  `json:"writeOperationConfig"`
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

	m.DataProperty = model.DataProperty

	m.WriteOperationConfig = model.WriteOperationConfig
	return
}

// TargetDataPropertyEnum Enum with underlying type: string
type TargetDataPropertyEnum string

// Set of constants representing the allowable values for TargetDataPropertyEnum
const (
	TargetDataPropertyTruncate  TargetDataPropertyEnum = "TRUNCATE"
	TargetDataPropertyMerge     TargetDataPropertyEnum = "MERGE"
	TargetDataPropertyBackup    TargetDataPropertyEnum = "BACKUP"
	TargetDataPropertyOverwrite TargetDataPropertyEnum = "OVERWRITE"
	TargetDataPropertyAppend    TargetDataPropertyEnum = "APPEND"
	TargetDataPropertyIgnore    TargetDataPropertyEnum = "IGNORE"
)

var mappingTargetDataProperty = map[string]TargetDataPropertyEnum{
	"TRUNCATE":  TargetDataPropertyTruncate,
	"MERGE":     TargetDataPropertyMerge,
	"BACKUP":    TargetDataPropertyBackup,
	"OVERWRITE": TargetDataPropertyOverwrite,
	"APPEND":    TargetDataPropertyAppend,
	"IGNORE":    TargetDataPropertyIgnore,
}

// GetTargetDataPropertyEnumValues Enumerates the set of values for TargetDataPropertyEnum
func GetTargetDataPropertyEnumValues() []TargetDataPropertyEnum {
	values := make([]TargetDataPropertyEnum, 0)
	for _, v := range mappingTargetDataProperty {
		values = append(values, v)
	}
	return values
}