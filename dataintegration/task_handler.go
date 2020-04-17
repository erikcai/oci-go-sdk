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

// TaskHandler auto generated description
type TaskHandler interface {
}

type taskhandler struct {
	JsonData  []byte
	ModelType string `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *taskhandler) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalertaskhandler taskhandler
	s := struct {
		Model Unmarshalertaskhandler
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *taskhandler) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "DECOMPOSABLE_TASK_HANDLER":
		mm := DecomposableTaskHandler{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m taskhandler) String() string {
	return common.PointerString(m)
}

// TaskHandlerModelTypeEnum Enum with underlying type: string
type TaskHandlerModelTypeEnum string

// Set of constants representing the allowable values for TaskHandlerModelTypeEnum
const (
	TaskHandlerModelTypeIntegrationTaskHandler TaskHandlerModelTypeEnum = "INTEGRATION_TASK_HANDLER"
	TaskHandlerModelTypeDataLoaderTaskHandler  TaskHandlerModelTypeEnum = "DATA_LOADER_TASK_HANDLER"
)

var mappingTaskHandlerModelType = map[string]TaskHandlerModelTypeEnum{
	"INTEGRATION_TASK_HANDLER": TaskHandlerModelTypeIntegrationTaskHandler,
	"DATA_LOADER_TASK_HANDLER": TaskHandlerModelTypeDataLoaderTaskHandler,
}

// GetTaskHandlerModelTypeEnumValues Enumerates the set of values for TaskHandlerModelTypeEnum
func GetTaskHandlerModelTypeEnumValues() []TaskHandlerModelTypeEnum {
	values := make([]TaskHandlerModelTypeEnum, 0)
	for _, v := range mappingTaskHandlerModelType {
		values = append(values, v)
	}
	return values
}
