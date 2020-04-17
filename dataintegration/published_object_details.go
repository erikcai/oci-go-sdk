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

// PublishedObjectDetails auto generated description
type PublishedObjectDetails interface {

	// Generated key that can be used in API calls to identify task. On scenarios where reference to the task is needed, a value can be passed in create.
	GetKey() *string

	// modelVersion
	GetModelVersion() *string

	GetParentRef() *ParentReference

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value can be edited by the user and it is restricted to 1000 characters
	GetName() *string

	// Descriptive text for the object.
	GetDescription() *string

	// The version of the object, to track changes in the object instance
	GetObjectVersion() *int

	// Status of object, can set this to value 1 for shallow references across objects, other values reserved.
	GetObjectStatus() *int

	// Value can only contain upper case letters, underscore and numbers. It should begin with upper case letter or underscore. The value can be edited by the user.
	GetIdentifier() *string

	// inputPorts
	GetInputPorts() []InputPort

	// outputPorts
	GetOutputPorts() []OutputPort

	// parameters
	GetParameters() []Parameter

	GetOpConfigValues() *ConfigValues

	GetConfigProviderDelegate() *ConfigProvider
}

type publishedobjectdetails struct {
	JsonData               []byte
	Key                    *string          `mandatory:"false" json:"key"`
	ModelVersion           *string          `mandatory:"false" json:"modelVersion"`
	ParentRef              *ParentReference `mandatory:"false" json:"parentRef"`
	Name                   *string          `mandatory:"false" json:"name"`
	Description            *string          `mandatory:"false" json:"description"`
	ObjectVersion          *int             `mandatory:"false" json:"objectVersion"`
	ObjectStatus           *int             `mandatory:"false" json:"objectStatus"`
	Identifier             *string          `mandatory:"false" json:"identifier"`
	InputPorts             []InputPort      `mandatory:"false" json:"inputPorts"`
	OutputPorts            []OutputPort     `mandatory:"false" json:"outputPorts"`
	Parameters             []Parameter      `mandatory:"false" json:"parameters"`
	OpConfigValues         *ConfigValues    `mandatory:"false" json:"opConfigValues"`
	ConfigProviderDelegate *ConfigProvider  `mandatory:"false" json:"configProviderDelegate"`
	ModelType              string           `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *publishedobjectdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerpublishedobjectdetails publishedobjectdetails
	s := struct {
		Model Unmarshalerpublishedobjectdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Key = s.Model.Key
	m.ModelVersion = s.Model.ModelVersion
	m.ParentRef = s.Model.ParentRef
	m.Name = s.Model.Name
	m.Description = s.Model.Description
	m.ObjectVersion = s.Model.ObjectVersion
	m.ObjectStatus = s.Model.ObjectStatus
	m.Identifier = s.Model.Identifier
	m.InputPorts = s.Model.InputPorts
	m.OutputPorts = s.Model.OutputPorts
	m.Parameters = s.Model.Parameters
	m.OpConfigValues = s.Model.OpConfigValues
	m.ConfigProviderDelegate = s.Model.ConfigProviderDelegate
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *publishedobjectdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "DATA_LOADER_TASK":
		mm := PublishedObjectFromDataLoaderTask{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "INTEGRATION_TASK":
		mm := PublishedObjectFromIntegrationTask{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetKey returns Key
func (m publishedobjectdetails) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m publishedobjectdetails) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m publishedobjectdetails) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetName returns Name
func (m publishedobjectdetails) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m publishedobjectdetails) GetDescription() *string {
	return m.Description
}

//GetObjectVersion returns ObjectVersion
func (m publishedobjectdetails) GetObjectVersion() *int {
	return m.ObjectVersion
}

//GetObjectStatus returns ObjectStatus
func (m publishedobjectdetails) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetIdentifier returns Identifier
func (m publishedobjectdetails) GetIdentifier() *string {
	return m.Identifier
}

//GetInputPorts returns InputPorts
func (m publishedobjectdetails) GetInputPorts() []InputPort {
	return m.InputPorts
}

//GetOutputPorts returns OutputPorts
func (m publishedobjectdetails) GetOutputPorts() []OutputPort {
	return m.OutputPorts
}

//GetParameters returns Parameters
func (m publishedobjectdetails) GetParameters() []Parameter {
	return m.Parameters
}

//GetOpConfigValues returns OpConfigValues
func (m publishedobjectdetails) GetOpConfigValues() *ConfigValues {
	return m.OpConfigValues
}

//GetConfigProviderDelegate returns ConfigProviderDelegate
func (m publishedobjectdetails) GetConfigProviderDelegate() *ConfigProvider {
	return m.ConfigProviderDelegate
}

func (m publishedobjectdetails) String() string {
	return common.PointerString(m)
}

// PublishedObjectDetailsModelTypeEnum Enum with underlying type: string
type PublishedObjectDetailsModelTypeEnum string

// Set of constants representing the allowable values for PublishedObjectDetailsModelTypeEnum
const (
	PublishedObjectDetailsModelTypeIntegrationTask PublishedObjectDetailsModelTypeEnum = "INTEGRATION_TASK"
	PublishedObjectDetailsModelTypeDataLoaderTask  PublishedObjectDetailsModelTypeEnum = "DATA_LOADER_TASK"
)

var mappingPublishedObjectDetailsModelType = map[string]PublishedObjectDetailsModelTypeEnum{
	"INTEGRATION_TASK": PublishedObjectDetailsModelTypeIntegrationTask,
	"DATA_LOADER_TASK": PublishedObjectDetailsModelTypeDataLoaderTask,
}

// GetPublishedObjectDetailsModelTypeEnumValues Enumerates the set of values for PublishedObjectDetailsModelTypeEnum
func GetPublishedObjectDetailsModelTypeEnumValues() []PublishedObjectDetailsModelTypeEnum {
	values := make([]PublishedObjectDetailsModelTypeEnum, 0)
	for _, v := range mappingPublishedObjectDetailsModelType {
		values = append(values, v)
	}
	return values
}
