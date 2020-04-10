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

// ReadOperationConfig auto generated description
type ReadOperationConfig struct {

	// Object key
	Key *string `mandatory:"false" json:"key"`

	// modelVersion
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value can be edited by the user and it is restricted to 1000 characters
	Name *string `mandatory:"false" json:"name"`

	// operations
	Operations []PushDownOperation `mandatory:"false" json:"operations"`

	DataFormat *DataFormat `mandatory:"false" json:"dataFormat"`

	PartitionConfig PartitionConfig `mandatory:"false" json:"partitionConfig"`

	ReadAttribute AbstractReadAttribute `mandatory:"false" json:"readAttribute"`

	// Status of object, can set this to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`
}

func (m ReadOperationConfig) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ReadOperationConfig) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeReadOperationConfig ReadOperationConfig
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeReadOperationConfig
	}{
		"READ_OPERATION_CONFIG",
		(MarshalTypeReadOperationConfig)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *ReadOperationConfig) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Key             *string               `json:"key"`
		ModelVersion    *string               `json:"modelVersion"`
		ParentRef       *ParentReference      `json:"parentRef"`
		Name            *string               `json:"name"`
		Operations      []pushdownoperation   `json:"operations"`
		DataFormat      *DataFormat           `json:"dataFormat"`
		PartitionConfig partitionconfig       `json:"partitionConfig"`
		ReadAttribute   abstractreadattribute `json:"readAttribute"`
		ObjectStatus    *int                  `json:"objectStatus"`
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

	m.Operations = make([]PushDownOperation, len(model.Operations))
	for i, n := range model.Operations {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Operations[i] = nn.(PushDownOperation)
		} else {
			m.Operations[i] = nil
		}
	}

	m.DataFormat = model.DataFormat

	nn, e = model.PartitionConfig.UnmarshalPolymorphicJSON(model.PartitionConfig.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.PartitionConfig = nn.(PartitionConfig)
	} else {
		m.PartitionConfig = nil
	}

	nn, e = model.ReadAttribute.UnmarshalPolymorphicJSON(model.ReadAttribute.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ReadAttribute = nn.(AbstractReadAttribute)
	} else {
		m.ReadAttribute = nil
	}

	m.ObjectStatus = model.ObjectStatus
	return
}
