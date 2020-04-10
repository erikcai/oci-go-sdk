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

// TaskDetailsForCreate The details for the task.
type TaskDetailsForCreate interface {

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value can be edited by the user and it is restricted to 1000 characters
	GetName() *string

	// Value can only contain upper case letters, underscore and numbers. It should begin with upper case letter or underscore. The value can be edited by the user.
	GetIdentifier() *string

	// Generated key that can be used in API calls to identify task. On scenarios where reference to the task is needed, a value can be passed in create.
	GetKey() *string

	// modelVersion
	GetModelVersion() *string

	GetParentRef() *ParentReference

	// Descriptive text for the object.
	GetDescription() *string

	// Status of object, can set this to value 1 for shallow references across objects, other values reserved.
	GetObjectStatus() *int

	// inputPorts
	GetInputPorts() []InputPort

	// outputPorts
	GetOutputPorts() []OutputPort

	// parameters
	GetParameters() []Parameter

	GetOpConfigValues() *ConfigValues

	GetConfigProviderDelegate() *ConfigProvider
}

type taskdetailsforcreate struct {
	JsonData               []byte
	Name                   *string          `mandatory:"true" json:"name"`
	Identifier             *string          `mandatory:"true" json:"identifier"`
	Key                    *string          `mandatory:"false" json:"key"`
	ModelVersion           *string          `mandatory:"false" json:"modelVersion"`
	ParentRef              *ParentReference `mandatory:"false" json:"parentRef"`
	Description            *string          `mandatory:"false" json:"description"`
	ObjectStatus           *int             `mandatory:"false" json:"objectStatus"`
	InputPorts             []InputPort      `mandatory:"false" json:"inputPorts"`
	OutputPorts            []OutputPort     `mandatory:"false" json:"outputPorts"`
	Parameters             []Parameter      `mandatory:"false" json:"parameters"`
	OpConfigValues         *ConfigValues    `mandatory:"false" json:"opConfigValues"`
	ConfigProviderDelegate *ConfigProvider  `mandatory:"false" json:"configProviderDelegate"`
	ModelType              string           `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *taskdetailsforcreate) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalertaskdetailsforcreate taskdetailsforcreate
	s := struct {
		Model Unmarshalertaskdetailsforcreate
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Name = s.Model.Name
	m.Identifier = s.Model.Identifier
	m.Key = s.Model.Key
	m.ModelVersion = s.Model.ModelVersion
	m.ParentRef = s.Model.ParentRef
	m.Description = s.Model.Description
	m.ObjectStatus = s.Model.ObjectStatus
	m.InputPorts = s.Model.InputPorts
	m.OutputPorts = s.Model.OutputPorts
	m.Parameters = s.Model.Parameters
	m.OpConfigValues = s.Model.OpConfigValues
	m.ConfigProviderDelegate = s.Model.ConfigProviderDelegate
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *taskdetailsforcreate) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "INTEGRATION_TASK":
		mm := TaskFromIntegrationTaskCreateDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DATA_LOADER_TASK":
		mm := TaskFromDataLoaderTaskCreateDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetName returns Name
func (m taskdetailsforcreate) GetName() *string {
	return m.Name
}

//GetIdentifier returns Identifier
func (m taskdetailsforcreate) GetIdentifier() *string {
	return m.Identifier
}

//GetKey returns Key
func (m taskdetailsforcreate) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m taskdetailsforcreate) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m taskdetailsforcreate) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetDescription returns Description
func (m taskdetailsforcreate) GetDescription() *string {
	return m.Description
}

//GetObjectStatus returns ObjectStatus
func (m taskdetailsforcreate) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetInputPorts returns InputPorts
func (m taskdetailsforcreate) GetInputPorts() []InputPort {
	return m.InputPorts
}

//GetOutputPorts returns OutputPorts
func (m taskdetailsforcreate) GetOutputPorts() []OutputPort {
	return m.OutputPorts
}

//GetParameters returns Parameters
func (m taskdetailsforcreate) GetParameters() []Parameter {
	return m.Parameters
}

//GetOpConfigValues returns OpConfigValues
func (m taskdetailsforcreate) GetOpConfigValues() *ConfigValues {
	return m.OpConfigValues
}

//GetConfigProviderDelegate returns ConfigProviderDelegate
func (m taskdetailsforcreate) GetConfigProviderDelegate() *ConfigProvider {
	return m.ConfigProviderDelegate
}

func (m taskdetailsforcreate) String() string {
	return common.PointerString(m)
}

// TaskDetailsForCreateModelTypeEnum Enum with underlying type: string
type TaskDetailsForCreateModelTypeEnum string

// Set of constants representing the allowable values for TaskDetailsForCreateModelTypeEnum
const (
	TaskDetailsForCreateModelTypeIntegrationTask TaskDetailsForCreateModelTypeEnum = "INTEGRATION_TASK"
	TaskDetailsForCreateModelTypeDataLoaderTask  TaskDetailsForCreateModelTypeEnum = "DATA_LOADER_TASK"
)

var mappingTaskDetailsForCreateModelType = map[string]TaskDetailsForCreateModelTypeEnum{
	"INTEGRATION_TASK": TaskDetailsForCreateModelTypeIntegrationTask,
	"DATA_LOADER_TASK": TaskDetailsForCreateModelTypeDataLoaderTask,
}

// GetTaskDetailsForCreateModelTypeEnumValues Enumerates the set of values for TaskDetailsForCreateModelTypeEnum
func GetTaskDetailsForCreateModelTypeEnumValues() []TaskDetailsForCreateModelTypeEnum {
	values := make([]TaskDetailsForCreateModelTypeEnum, 0)
	for _, v := range mappingTaskDetailsForCreateModelType {
		values = append(values, v)
	}
	return values
}
