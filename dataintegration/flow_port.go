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

// FlowPort Each operator owns a set of InputPort and OutputPort objects (can scale to zero), which represent the ports that can be connected to/from the Operator.
type FlowPort struct {

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
}

//GetKey returns Key
func (m FlowPort) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m FlowPort) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m FlowPort) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetConfigValues returns ConfigValues
func (m FlowPort) GetConfigValues() *ConfigValues {
	return m.ConfigValues
}

//GetObjectStatus returns ObjectStatus
func (m FlowPort) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetName returns Name
func (m FlowPort) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m FlowPort) GetDescription() *string {
	return m.Description
}

func (m FlowPort) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m FlowPort) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeFlowPort FlowPort
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeFlowPort
	}{
		"FLOW_PORT",
		(MarshalTypeFlowPort)(m),
	}

	return json.Marshal(&s)
}
