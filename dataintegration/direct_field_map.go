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

// DirectFieldMap auto generated description
type DirectFieldMap struct {

	// Descriptive text for the object.
	Description *string `mandatory:"false" json:"description"`

	// Object key
	Key *string `mandatory:"false" json:"key"`

	// modelVersion
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	ConfigValues *ConfigValues `mandatory:"false" json:"configValues"`

	// Reference to a typed object
	SourceTypedObject *string `mandatory:"false" json:"sourceTypedObject"`

	// Reference to a typed object
	TargetTypedObject *string `mandatory:"false" json:"targetTypedObject"`

	// Status of object, can set this to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`
}

//GetDescription returns Description
func (m DirectFieldMap) GetDescription() *string {
	return m.Description
}

func (m DirectFieldMap) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m DirectFieldMap) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDirectFieldMap DirectFieldMap
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeDirectFieldMap
	}{
		"DIRECT_FIELD_MAP",
		(MarshalTypeDirectFieldMap)(m),
	}

	return json.Marshal(&s)
}
