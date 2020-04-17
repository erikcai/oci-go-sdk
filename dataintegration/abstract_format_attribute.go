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

// AbstractFormatAttribute The abstract format attribute.
type AbstractFormatAttribute interface {
}

type abstractformatattribute struct {
	JsonData  []byte
	ModelType string `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *abstractformatattribute) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerabstractformatattribute abstractformatattribute
	s := struct {
		Model Unmarshalerabstractformatattribute
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *abstractformatattribute) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "AVROFORMAT":
		mm := AvroFormatAttribute{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CSVFORMAT":
		mm := CsvFormatAttribute{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TEXTFILEFORMAT":
		mm := TextFileFormatAttribute{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m abstractformatattribute) String() string {
	return common.PointerString(m)
}

// AbstractFormatAttributeModelTypeEnum Enum with underlying type: string
type AbstractFormatAttributeModelTypeEnum string

// Set of constants representing the allowable values for AbstractFormatAttributeModelTypeEnum
const (
	AbstractFormatAttributeModelTypeTextfileformat AbstractFormatAttributeModelTypeEnum = "TEXTFILEFORMAT"
	AbstractFormatAttributeModelTypeAvroformat     AbstractFormatAttributeModelTypeEnum = "AVROFORMAT"
	AbstractFormatAttributeModelTypeCsvformat      AbstractFormatAttributeModelTypeEnum = "CSVFORMAT"
)

var mappingAbstractFormatAttributeModelType = map[string]AbstractFormatAttributeModelTypeEnum{
	"TEXTFILEFORMAT": AbstractFormatAttributeModelTypeTextfileformat,
	"AVROFORMAT":     AbstractFormatAttributeModelTypeAvroformat,
	"CSVFORMAT":      AbstractFormatAttributeModelTypeCsvformat,
}

// GetAbstractFormatAttributeModelTypeEnumValues Enumerates the set of values for AbstractFormatAttributeModelTypeEnum
func GetAbstractFormatAttributeModelTypeEnumValues() []AbstractFormatAttributeModelTypeEnum {
	values := make([]AbstractFormatAttributeModelTypeEnum, 0)
	for _, v := range mappingAbstractFormatAttributeModelType {
		values = append(values, v)
	}
	return values
}
