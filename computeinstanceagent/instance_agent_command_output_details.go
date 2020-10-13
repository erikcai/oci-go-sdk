// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// InstanceAgentService API
//
// Instance Agent Service API
//

package computeinstanceagent

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v27/common"
)

// InstanceAgentCommandOutputDetails Command output.
type InstanceAgentCommandOutputDetails interface {
}

type instanceagentcommandoutputdetails struct {
	JsonData   []byte
	OutputType string `json:"outputType"`
}

// UnmarshalJSON unmarshals json
func (m *instanceagentcommandoutputdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerinstanceagentcommandoutputdetails instanceagentcommandoutputdetails
	s := struct {
		Model Unmarshalerinstanceagentcommandoutputdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.OutputType = s.Model.OutputType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *instanceagentcommandoutputdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.OutputType {
	case "OBJECT_STORAGE_URI":
		mm := InstanceAgentCommandOutputViaObjectStorageUriDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OBJECT_STORAGE_TUPLE":
		mm := InstanceAgentCommandOutputViaObjectStorageTupleDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TEXT":
		mm := InstanceAgentCommandOutputViaTextDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m instanceagentcommandoutputdetails) String() string {
	return common.PointerString(m)
}

// InstanceAgentCommandOutputDetailsOutputTypeEnum Enum with underlying type: string
type InstanceAgentCommandOutputDetailsOutputTypeEnum string

// Set of constants representing the allowable values for InstanceAgentCommandOutputDetailsOutputTypeEnum
const (
	InstanceAgentCommandOutputDetailsOutputTypeText               InstanceAgentCommandOutputDetailsOutputTypeEnum = "TEXT"
	InstanceAgentCommandOutputDetailsOutputTypeObjectStorageUri   InstanceAgentCommandOutputDetailsOutputTypeEnum = "OBJECT_STORAGE_URI"
	InstanceAgentCommandOutputDetailsOutputTypeObjectStorageTuple InstanceAgentCommandOutputDetailsOutputTypeEnum = "OBJECT_STORAGE_TUPLE"
)

var mappingInstanceAgentCommandOutputDetailsOutputType = map[string]InstanceAgentCommandOutputDetailsOutputTypeEnum{
	"TEXT":                 InstanceAgentCommandOutputDetailsOutputTypeText,
	"OBJECT_STORAGE_URI":   InstanceAgentCommandOutputDetailsOutputTypeObjectStorageUri,
	"OBJECT_STORAGE_TUPLE": InstanceAgentCommandOutputDetailsOutputTypeObjectStorageTuple,
}

// GetInstanceAgentCommandOutputDetailsOutputTypeEnumValues Enumerates the set of values for InstanceAgentCommandOutputDetailsOutputTypeEnum
func GetInstanceAgentCommandOutputDetailsOutputTypeEnumValues() []InstanceAgentCommandOutputDetailsOutputTypeEnum {
	values := make([]InstanceAgentCommandOutputDetailsOutputTypeEnum, 0)
	for _, v := range mappingInstanceAgentCommandOutputDetailsOutputType {
		values = append(values, v)
	}
	return values
}
