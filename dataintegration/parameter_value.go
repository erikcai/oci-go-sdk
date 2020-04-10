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

// ParameterValue auto generated description
type ParameterValue struct {

	// auto generated description
	SimpleValue *interface{} `mandatory:"false" json:"simpleValue"`

	// This can be any object such as a file entity, or a schema or a table.
	RootObjValue *interface{} `mandatory:"false" json:"rootObjValue"`
}

func (m ParameterValue) String() string {
	return common.PointerString(m)
}
