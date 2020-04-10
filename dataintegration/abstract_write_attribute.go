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

// AbstractWriteAttribute The abstract write attribute.
type AbstractWriteAttribute interface {
}

type abstractwriteattribute struct {
	JsonData  []byte
	ModelType string `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *abstractwriteattribute) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerabstractwriteattribute abstractwriteattribute
	s := struct {
		Model Unmarshalerabstractwriteattribute
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *abstractwriteattribute) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "ORACLEADWCWRITEATTRIBUTE":
		mm := OracleAdwcWriteAttribute{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLEWRITEATTRIBUTE":
		mm := OracleWriteAttribute{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLEATPWRITEATTRIBUTE":
		mm := OracleAtpWriteAttribute{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m abstractwriteattribute) String() string {
	return common.PointerString(m)
}

// AbstractWriteAttributeModelTypeEnum Enum with underlying type: string
type AbstractWriteAttributeModelTypeEnum string

// Set of constants representing the allowable values for AbstractWriteAttributeModelTypeEnum
const (
	AbstractWriteAttributeModelTypeOraclewriteattribute     AbstractWriteAttributeModelTypeEnum = "ORACLEWRITEATTRIBUTE"
	AbstractWriteAttributeModelTypeOracleatpwriteattribute  AbstractWriteAttributeModelTypeEnum = "ORACLEATPWRITEATTRIBUTE"
	AbstractWriteAttributeModelTypeOracleadwcwriteattribute AbstractWriteAttributeModelTypeEnum = "ORACLEADWCWRITEATTRIBUTE"
)

var mappingAbstractWriteAttributeModelType = map[string]AbstractWriteAttributeModelTypeEnum{
	"ORACLEWRITEATTRIBUTE":     AbstractWriteAttributeModelTypeOraclewriteattribute,
	"ORACLEATPWRITEATTRIBUTE":  AbstractWriteAttributeModelTypeOracleatpwriteattribute,
	"ORACLEADWCWRITEATTRIBUTE": AbstractWriteAttributeModelTypeOracleadwcwriteattribute,
}

// GetAbstractWriteAttributeModelTypeEnumValues Enumerates the set of values for AbstractWriteAttributeModelTypeEnum
func GetAbstractWriteAttributeModelTypeEnumValues() []AbstractWriteAttributeModelTypeEnum {
	values := make([]AbstractWriteAttributeModelTypeEnum, 0)
	for _, v := range mappingAbstractWriteAttributeModelType {
		values = append(values, v)
	}
	return values
}
