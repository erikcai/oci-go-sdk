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

// Function auto generated description
type Function interface {
}

type function struct {
	JsonData  []byte
	ModelType string `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *function) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerfunction function
	s := struct {
		Model Unmarshalerfunction
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *function) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "USER_DEFINED_FUNCTIONS":
		mm := UserDefinedFunction{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "BUILT_IN_FUNCTIONS":
		mm := BuiltInFunction{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m function) String() string {
	return common.PointerString(m)
}

// FunctionModelTypeEnum Enum with underlying type: string
type FunctionModelTypeEnum string

// Set of constants representing the allowable values for FunctionModelTypeEnum
const (
	FunctionModelTypeBuiltInFunctions     FunctionModelTypeEnum = "BUILT_IN_FUNCTIONS"
	FunctionModelTypeUserDefinedFunctions FunctionModelTypeEnum = "USER_DEFINED_FUNCTIONS"
)

var mappingFunctionModelType = map[string]FunctionModelTypeEnum{
	"BUILT_IN_FUNCTIONS":     FunctionModelTypeBuiltInFunctions,
	"USER_DEFINED_FUNCTIONS": FunctionModelTypeUserDefinedFunctions,
}

// GetFunctionModelTypeEnumValues Enumerates the set of values for FunctionModelTypeEnum
func GetFunctionModelTypeEnumValues() []FunctionModelTypeEnum {
	values := make([]FunctionModelTypeEnum, 0)
	for _, v := range mappingFunctionModelType {
		values = append(values, v)
	}
	return values
}
