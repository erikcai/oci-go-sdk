// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Data Integration Service API Specification
//
// Data Integration Service API Specification
//

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Expression auto generated description
type Expression struct {

	// Object key
	Key *string `mandatory:"false" json:"key"`

	// The type of the object.
	ModelType *string `mandatory:"false" json:"modelType"`

	// modelVersion
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// exprString
	ExprString *string `mandatory:"false" json:"exprString"`

	ConfigValues *ConfigValues `mandatory:"false" json:"configValues"`

	// Status of object, can set this to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	Root *RootNode `mandatory:"false" json:"root"`
}

func (m Expression) String() string {
	return common.PointerString(m)
}
