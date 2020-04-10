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

// DirectNamedFieldMap auto generated description
type DirectNamedFieldMap struct {

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

	// sourceFieldName
	SourceFieldName *string `mandatory:"false" json:"sourceFieldName"`

	// targetFieldName
	TargetFieldName *string `mandatory:"false" json:"targetFieldName"`

	// Status of object, can set this to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`
}

//GetDescription returns Description
func (m DirectNamedFieldMap) GetDescription() *string {
	return m.Description
}

func (m DirectNamedFieldMap) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m DirectNamedFieldMap) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDirectNamedFieldMap DirectNamedFieldMap
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeDirectNamedFieldMap
	}{
		"DIRECT_NAMED_FIELD_MAP",
		(MarshalTypeDirectNamedFieldMap)(m),
	}

	return json.Marshal(&s)
}
