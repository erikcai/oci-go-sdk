// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management Service APIs.
//
// This file contains the customer facing APIs for Database Management service.
//

package databasemanagement

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v30/common"
)

// JobExecutionResultLocation Location of the job executions result.
type JobExecutionResultLocation interface {
}

type jobexecutionresultlocation struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *jobexecutionresultlocation) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerjobexecutionresultlocation jobexecutionresultlocation
	s := struct {
		Model Unmarshalerjobexecutionresultlocation
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *jobexecutionresultlocation) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "OBJECT_STORAGE":
		mm := ObjectStorageJobExecutionResultLocation{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m jobexecutionresultlocation) String() string {
	return common.PointerString(m)
}

// JobExecutionResultLocationTypeEnum Enum with underlying type: string
type JobExecutionResultLocationTypeEnum string

// Set of constants representing the allowable values for JobExecutionResultLocationTypeEnum
const (
	JobExecutionResultLocationTypeObjectStorage JobExecutionResultLocationTypeEnum = "OBJECT_STORAGE"
)

var mappingJobExecutionResultLocationType = map[string]JobExecutionResultLocationTypeEnum{
	"OBJECT_STORAGE": JobExecutionResultLocationTypeObjectStorage,
}

// GetJobExecutionResultLocationTypeEnumValues Enumerates the set of values for JobExecutionResultLocationTypeEnum
func GetJobExecutionResultLocationTypeEnumValues() []JobExecutionResultLocationTypeEnum {
	values := make([]JobExecutionResultLocationTypeEnum, 0)
	for _, v := range mappingJobExecutionResultLocationType {
		values = append(values, v)
	}
	return values
}
