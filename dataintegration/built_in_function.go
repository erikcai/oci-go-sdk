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

// BuiltInFunction auto generated description
type BuiltInFunction struct {

	// Object key
	Key *string `mandatory:"false" json:"key"`

	// modelVersion
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value can be edited by the user and it is restricted to 1000 characters
	Name *string `mandatory:"false" json:"name"`

	// Descriptive text for the object.
	Description *string `mandatory:"false" json:"description"`

	// signatures
	Signatures []FunctionSignature `mandatory:"false" json:"signatures"`

	Expr *Expression `mandatory:"false" json:"expr"`

	// Status of object, can set this to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`
}

func (m BuiltInFunction) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m BuiltInFunction) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeBuiltInFunction BuiltInFunction
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeBuiltInFunction
	}{
		"BUILT_IN_FUNCTIONS",
		(MarshalTypeBuiltInFunction)(m),
	}

	return json.Marshal(&s)
}
