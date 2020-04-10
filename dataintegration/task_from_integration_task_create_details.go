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

// TaskFromIntegrationTaskCreateDetails auto generated description
type TaskFromIntegrationTaskCreateDetails struct {

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
}

//GetKey returns Key
func (m TaskFromIntegrationTaskCreateDetails) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m TaskFromIntegrationTaskCreateDetails) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m TaskFromIntegrationTaskCreateDetails) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetName returns Name
func (m TaskFromIntegrationTaskCreateDetails) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m TaskFromIntegrationTaskCreateDetails) GetDescription() *string {
	return m.Description
}

//GetObjectStatus returns ObjectStatus
func (m TaskFromIntegrationTaskCreateDetails) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetIdentifier returns Identifier
func (m TaskFromIntegrationTaskCreateDetails) GetIdentifier() *string {
	return m.Identifier
}

//GetInputPorts returns InputPorts
func (m TaskFromIntegrationTaskCreateDetails) GetInputPorts() []InputPort {
	return m.InputPorts
}

//GetOutputPorts returns OutputPorts
func (m TaskFromIntegrationTaskCreateDetails) GetOutputPorts() []OutputPort {
	return m.OutputPorts
}

//GetParameters returns Parameters
func (m TaskFromIntegrationTaskCreateDetails) GetParameters() []Parameter {
	return m.Parameters
}

//GetOpConfigValues returns OpConfigValues
func (m TaskFromIntegrationTaskCreateDetails) GetOpConfigValues() *ConfigValues {
	return m.OpConfigValues
}

//GetConfigProviderDelegate returns ConfigProviderDelegate
func (m TaskFromIntegrationTaskCreateDetails) GetConfigProviderDelegate() *ConfigProvider {
	return m.ConfigProviderDelegate
}

func (m TaskFromIntegrationTaskCreateDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m TaskFromIntegrationTaskCreateDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTaskFromIntegrationTaskCreateDetails TaskFromIntegrationTaskCreateDetails
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeTaskFromIntegrationTaskCreateDetails
	}{
		"INTEGRATION_TASK",
		(MarshalTypeTaskFromIntegrationTaskCreateDetails)(m),
	}

	return json.Marshal(&s)
}
