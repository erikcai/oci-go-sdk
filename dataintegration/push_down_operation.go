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

// PushDownOperation auto generated description
type PushDownOperation interface {
}

type pushdownoperation struct {
	JsonData  []byte
	ModelType string `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *pushdownoperation) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerpushdownoperation pushdownoperation
	s := struct {
		Model Unmarshalerpushdownoperation
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *pushdownoperation) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "QUERY":
		mm := Query{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SELECT":
		mm := ModelSelect{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "JOIN":
		mm := Join{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SORT":
		mm := Sort{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FILTER":
		mm := FilterPush{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m pushdownoperation) String() string {
	return common.PointerString(m)
}

// PushDownOperationModelTypeEnum Enum with underlying type: string
type PushDownOperationModelTypeEnum string

// Set of constants representing the allowable values for PushDownOperationModelTypeEnum
const (
	PushDownOperationModelTypeFilter PushDownOperationModelTypeEnum = "FILTER"
	PushDownOperationModelTypeJoin   PushDownOperationModelTypeEnum = "JOIN"
	PushDownOperationModelTypeSelect PushDownOperationModelTypeEnum = "SELECT"
	PushDownOperationModelTypeSort   PushDownOperationModelTypeEnum = "SORT"
	PushDownOperationModelTypeQuery  PushDownOperationModelTypeEnum = "QUERY"
)

var mappingPushDownOperationModelType = map[string]PushDownOperationModelTypeEnum{
	"FILTER": PushDownOperationModelTypeFilter,
	"JOIN":   PushDownOperationModelTypeJoin,
	"SELECT": PushDownOperationModelTypeSelect,
	"SORT":   PushDownOperationModelTypeSort,
	"QUERY":  PushDownOperationModelTypeQuery,
}

// GetPushDownOperationModelTypeEnumValues Enumerates the set of values for PushDownOperationModelTypeEnum
func GetPushDownOperationModelTypeEnumValues() []PushDownOperationModelTypeEnum {
	values := make([]PushDownOperationModelTypeEnum, 0)
	for _, v := range mappingPushDownOperationModelType {
		values = append(values, v)
	}
	return values
}
