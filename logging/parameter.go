// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// PublicLoggingControlplane API
//
// PublicLoggingControlplane API specification
//

package logging

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Parameter Parameters that a category of resource supports.
type Parameter struct {

	// Parameter name.
	Name *string `mandatory:"true" json:"name"`

	// Parameter type. One of integer, string, boolean.
	Type ParameterTypeEnum `mandatory:"true" json:"type"`

	// Java regex pattern to validate parameter value.
	Pattern *string `mandatory:"false" json:"pattern"`
}

func (m Parameter) String() string {
	return common.PointerString(m)
}

// ParameterTypeEnum Enum with underlying type: string
type ParameterTypeEnum string

// Set of constants representing the allowable values for ParameterTypeEnum
const (
	ParameterTypeInteger ParameterTypeEnum = "integer"
	ParameterTypeString  ParameterTypeEnum = "string"
	ParameterTypeBoolean ParameterTypeEnum = "boolean"
)

var mappingParameterType = map[string]ParameterTypeEnum{
	"integer": ParameterTypeInteger,
	"string":  ParameterTypeString,
	"boolean": ParameterTypeBoolean,
}

// GetParameterTypeEnumValues Enumerates the set of values for ParameterTypeEnum
func GetParameterTypeEnumValues() []ParameterTypeEnum {
	values := make([]ParameterTypeEnum, 0)
	for _, v := range mappingParameterType {
		values = append(values, v)
	}
	return values
}
