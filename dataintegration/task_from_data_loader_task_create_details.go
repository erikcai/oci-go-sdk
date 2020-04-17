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

// TaskFromDataLoaderTaskCreateDetails auto generated description
type TaskFromDataLoaderTaskCreateDetails struct {

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value can be edited by the user and it is restricted to 1000 characters
	Name *string `mandatory:"true" json:"name"`

	// Value can only contain upper case letters, underscore and numbers. It should begin with upper case letter or underscore. The value can be edited by the user.
	Identifier *string `mandatory:"true" json:"identifier"`

	// Generated key that can be used in API calls to identify task. On scenarios where reference to the task is needed, a value can be passed in create.
	Key *string `mandatory:"false" json:"key"`

	// modelVersion
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Descriptive text for the object.
	Description *string `mandatory:"false" json:"description"`

	// Status of object, can set this to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// inputPorts
	InputPorts []InputPort `mandatory:"false" json:"inputPorts"`

	// outputPorts
	OutputPorts []OutputPort `mandatory:"false" json:"outputPorts"`

	// parameters
	Parameters []Parameter `mandatory:"false" json:"parameters"`

	OpConfigValues *ConfigValues `mandatory:"false" json:"opConfigValues"`

	ConfigProviderDelegate *ConfigProvider `mandatory:"false" json:"configProviderDelegate"`

	DataFlow *DataFlowDetails `mandatory:"false" json:"dataFlow"`

	SourceConnection *ConnectionRefInfo `mandatory:"false" json:"sourceConnection"`

	TargetConnection *ConnectionRefInfo `mandatory:"false" json:"targetConnection"`

	// entityInfo
	EntityInfo []EntityInfo `mandatory:"false" json:"entityInfo"`

	DataSetting *DataLoaderDataSetting `mandatory:"false" json:"dataSetting"`
}

//GetKey returns Key
func (m TaskFromDataLoaderTaskCreateDetails) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m TaskFromDataLoaderTaskCreateDetails) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m TaskFromDataLoaderTaskCreateDetails) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetName returns Name
func (m TaskFromDataLoaderTaskCreateDetails) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m TaskFromDataLoaderTaskCreateDetails) GetDescription() *string {
	return m.Description
}

//GetObjectStatus returns ObjectStatus
func (m TaskFromDataLoaderTaskCreateDetails) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetIdentifier returns Identifier
func (m TaskFromDataLoaderTaskCreateDetails) GetIdentifier() *string {
	return m.Identifier
}

//GetInputPorts returns InputPorts
func (m TaskFromDataLoaderTaskCreateDetails) GetInputPorts() []InputPort {
	return m.InputPorts
}

//GetOutputPorts returns OutputPorts
func (m TaskFromDataLoaderTaskCreateDetails) GetOutputPorts() []OutputPort {
	return m.OutputPorts
}

//GetParameters returns Parameters
func (m TaskFromDataLoaderTaskCreateDetails) GetParameters() []Parameter {
	return m.Parameters
}

//GetOpConfigValues returns OpConfigValues
func (m TaskFromDataLoaderTaskCreateDetails) GetOpConfigValues() *ConfigValues {
	return m.OpConfigValues
}

//GetConfigProviderDelegate returns ConfigProviderDelegate
func (m TaskFromDataLoaderTaskCreateDetails) GetConfigProviderDelegate() *ConfigProvider {
	return m.ConfigProviderDelegate
}

func (m TaskFromDataLoaderTaskCreateDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m TaskFromDataLoaderTaskCreateDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTaskFromDataLoaderTaskCreateDetails TaskFromDataLoaderTaskCreateDetails
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeTaskFromDataLoaderTaskCreateDetails
	}{
		"DATA_LOADER_TASK",
		(MarshalTypeTaskFromDataLoaderTaskCreateDetails)(m),
	}

	return json.Marshal(&s)
}
