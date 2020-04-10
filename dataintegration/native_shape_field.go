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

// NativeShapeField The native shape field object.
type NativeShapeField struct {

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value can be edited by the user and it is restricted to 1000 characters
	Name *string `mandatory:"false" json:"name"`

	// Type reference
	Type *string `mandatory:"false" json:"type"`

	ConfigValues *ConfigValues `mandatory:"false" json:"configValues"`

	// sqlDataTypeCode
	SqlDataTypeCode *int `mandatory:"false" json:"sqlDataTypeCode"`

	// position
	Position *int `mandatory:"false" json:"position"`

	// defaultValueString
	DefaultValueString *string `mandatory:"false" json:"defaultValueString"`

	// dataFormat
	DataFormat *string `mandatory:"false" json:"dataFormat"`

	// mandatory
	IsMandatory *bool `mandatory:"false" json:"isMandatory"`
}

func (m NativeShapeField) String() string {
	return common.PointerString(m)
}
