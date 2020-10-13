// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataintegration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v27/common"
)

// FlowPort Each operator owns a set of `InputPort` and `OutputPort` objects (can scale to zero), which represent the ports that can be connected to/from the operator.
type FlowPort struct {

	// The key of the object.
	Key *string `mandatory:"false" json:"key"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	ConfigValues *ConfigValues `mandatory:"false" json:"configValues"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// Detailed description for the object.
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
