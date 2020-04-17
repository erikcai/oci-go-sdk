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

// DataEntityBinding The data entity binding object.
type DataEntityBinding interface {
}

type dataentitybinding struct {
	JsonData  []byte
	ModelType string `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *dataentitybinding) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdataentitybinding dataentitybinding
	s := struct {
		Model Unmarshalerdataentitybinding
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *dataentitybinding) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "PHYSICAL_BINDING":
		mm := PhysicalBinding{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m dataentitybinding) String() string {
	return common.PointerString(m)
}

// DataEntityBindingModelTypeEnum Enum with underlying type: string
type DataEntityBindingModelTypeEnum string

// Set of constants representing the allowable values for DataEntityBindingModelTypeEnum
const (
	DataEntityBindingModelTypePhysicalBinding DataEntityBindingModelTypeEnum = "PHYSICAL_BINDING"
)

var mappingDataEntityBindingModelType = map[string]DataEntityBindingModelTypeEnum{
	"PHYSICAL_BINDING": DataEntityBindingModelTypePhysicalBinding,
}

// GetDataEntityBindingModelTypeEnumValues Enumerates the set of values for DataEntityBindingModelTypeEnum
func GetDataEntityBindingModelTypeEnumValues() []DataEntityBindingModelTypeEnum {
	values := make([]DataEntityBindingModelTypeEnum, 0)
	for _, v := range mappingDataEntityBindingModelType {
		values = append(values, v)
	}
	return values
}
