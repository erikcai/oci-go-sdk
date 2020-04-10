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

// ConfigParamValue This holds the values/objects.
type ConfigParamValue struct {

	// stringValue
	StringValue *string `mandatory:"false" json:"stringValue"`

	// intValue
	IntValue *int `mandatory:"false" json:"intValue"`

	// The root object reference value.
	RefValue *string `mandatory:"false" json:"refValue"`

	ParameterValue *Parameter `mandatory:"false" json:"parameterValue"`

	// configParamTypeName
	ConfigParamTypeName *string `mandatory:"false" json:"configParamTypeName"`

	// configParamDescription
	ConfigParamDescription *string `mandatory:"false" json:"configParamDescription"`

	// isClassFieldValue
	IsClassFieldValue *bool `mandatory:"false" json:"isClassFieldValue"`

	// classFieldName
	ClassFieldName *string `mandatory:"false" json:"classFieldName"`
}

func (m ConfigParamValue) String() string {
	return common.PointerString(m)
}
