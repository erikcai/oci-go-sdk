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

// Key The key object.
type Key interface {
}

type key struct {
	JsonData  []byte
	ModelType string `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *key) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerkey key
	s := struct {
		Model Unmarshalerkey
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *key) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "UNIQUE_KEY":
		mm := UniqueKey{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FOREIGN_KEY":
		mm := ForeignKey{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m key) String() string {
	return common.PointerString(m)
}

// KeyModelTypeEnum Enum with underlying type: string
type KeyModelTypeEnum string

// Set of constants representing the allowable values for KeyModelTypeEnum
const (
	KeyModelTypeForeignKey KeyModelTypeEnum = "FOREIGN_KEY"
	KeyModelTypePrimaryKey KeyModelTypeEnum = "PRIMARY_KEY"
	KeyModelTypeUniqueKey  KeyModelTypeEnum = "UNIQUE_KEY"
)

var mappingKeyModelType = map[string]KeyModelTypeEnum{
	"FOREIGN_KEY": KeyModelTypeForeignKey,
	"PRIMARY_KEY": KeyModelTypePrimaryKey,
	"UNIQUE_KEY":  KeyModelTypeUniqueKey,
}

// GetKeyModelTypeEnumValues Enumerates the set of values for KeyModelTypeEnum
func GetKeyModelTypeEnumValues() []KeyModelTypeEnum {
	values := make([]KeyModelTypeEnum, 0)
	for _, v := range mappingKeyModelType {
		values = append(values, v)
	}
	return values
}
