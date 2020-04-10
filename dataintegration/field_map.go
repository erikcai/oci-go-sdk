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

// FieldMap A field map is a way to map a source row shape to a target row shape that may be different.
type FieldMap interface {

	// Descriptive text for the object.
	GetDescription() *string
}

type fieldmap struct {
	JsonData    []byte
	Description *string `mandatory:"false" json:"description"`
	ModelType   string  `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *fieldmap) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerfieldmap fieldmap
	s := struct {
		Model Unmarshalerfieldmap
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Description = s.Model.Description
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *fieldmap) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "RULE_BASED_FIELD_MAP":
		mm := RuleBasedFieldMap{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DIRECT_FIELD_MAP":
		mm := DirectFieldMap{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "COMPOSITE_FIELD_MAP":
		mm := CompositeFieldMap{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DIRECT_NAMED_FIELD_MAP":
		mm := DirectNamedFieldMap{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetDescription returns Description
func (m fieldmap) GetDescription() *string {
	return m.Description
}

func (m fieldmap) String() string {
	return common.PointerString(m)
}

// FieldMapModelTypeEnum Enum with underlying type: string
type FieldMapModelTypeEnum string

// Set of constants representing the allowable values for FieldMapModelTypeEnum
const (
	FieldMapModelTypeDirectNamedFieldMap FieldMapModelTypeEnum = "DIRECT_NAMED_FIELD_MAP"
	FieldMapModelTypeCompositeFieldMap   FieldMapModelTypeEnum = "COMPOSITE_FIELD_MAP"
	FieldMapModelTypeDirectFieldMap      FieldMapModelTypeEnum = "DIRECT_FIELD_MAP"
	FieldMapModelTypeRuleBasedFieldMap   FieldMapModelTypeEnum = "RULE_BASED_FIELD_MAP"
)

var mappingFieldMapModelType = map[string]FieldMapModelTypeEnum{
	"DIRECT_NAMED_FIELD_MAP": FieldMapModelTypeDirectNamedFieldMap,
	"COMPOSITE_FIELD_MAP":    FieldMapModelTypeCompositeFieldMap,
	"DIRECT_FIELD_MAP":       FieldMapModelTypeDirectFieldMap,
	"RULE_BASED_FIELD_MAP":   FieldMapModelTypeRuleBasedFieldMap,
}

// GetFieldMapModelTypeEnumValues Enumerates the set of values for FieldMapModelTypeEnum
func GetFieldMapModelTypeEnumValues() []FieldMapModelTypeEnum {
	values := make([]FieldMapModelTypeEnum, 0)
	for _, v := range mappingFieldMapModelType {
		values = append(values, v)
	}
	return values
}
