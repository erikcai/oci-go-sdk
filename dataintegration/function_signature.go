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

// FunctionSignature auto generated description
type FunctionSignature struct {

	// Object key
	Key *string `mandatory:"false" json:"key"`

	// The type of the object.
	ModelType *string `mandatory:"false" json:"modelType"`

	// modelVersion
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	RetType *ConfiguredType `mandatory:"false" json:"retType"`

	// arguments
	Arguments []TypedObject `mandatory:"false" json:"arguments"`

	// The key of the owning function.
	OwningFunction *string `mandatory:"false" json:"owningFunction"`

	// Status of object, can set this to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`
}

func (m FunctionSignature) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *FunctionSignature) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Key            *string          `json:"key"`
		ModelType      *string          `json:"modelType"`
		ModelVersion   *string          `json:"modelVersion"`
		ParentRef      *ParentReference `json:"parentRef"`
		RetType        *ConfiguredType  `json:"retType"`
		Arguments      []typedobject    `json:"arguments"`
		OwningFunction *string          `json:"owningFunction"`
		ObjectStatus   *int             `json:"objectStatus"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Key = model.Key

	m.ModelType = model.ModelType

	m.ModelVersion = model.ModelVersion

	m.ParentRef = model.ParentRef

	m.RetType = model.RetType

	m.Arguments = make([]TypedObject, len(model.Arguments))
	for i, n := range model.Arguments {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Arguments[i] = nn.(TypedObject)
		} else {
			m.Arguments[i] = nil
		}
	}

	m.OwningFunction = model.OwningFunction

	m.ObjectStatus = model.ObjectStatus
	return
}
