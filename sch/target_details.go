// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Connectors API
//
// A description of the Connectors API
//

package sch

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// TargetDetails The flow target.
type TargetDetails interface {
}

type targetdetails struct {
	JsonData []byte
	Kind     string `json:"kind"`
}

// UnmarshalJSON unmarshals json
func (m *targetdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalertargetdetails targetdetails
	s := struct {
		Model Unmarshalertargetdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Kind = s.Model.Kind

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *targetdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Kind {
	case "notifications":
		mm := NotificationsTargetDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "objectStorage":
		mm := ObjectStorageTargetDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "functions":
		mm := FunctionsTargetDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "streaming":
		mm := StreamingTargetDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "metrics":
		mm := MetricsTargetDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m targetdetails) String() string {
	return common.PointerString(m)
}

// TargetDetailsKindEnum Enum with underlying type: string
type TargetDetailsKindEnum string

// Set of constants representing the allowable values for TargetDetailsKindEnum
const (
	TargetDetailsKindStreaming     TargetDetailsKindEnum = "streaming"
	TargetDetailsKindObjectstorage TargetDetailsKindEnum = "objectStorage"
	TargetDetailsKindMetrics       TargetDetailsKindEnum = "metrics"
	TargetDetailsKindFunctions     TargetDetailsKindEnum = "functions"
	TargetDetailsKindNotifications TargetDetailsKindEnum = "notifications"
)

var mappingTargetDetailsKind = map[string]TargetDetailsKindEnum{
	"streaming":     TargetDetailsKindStreaming,
	"objectStorage": TargetDetailsKindObjectstorage,
	"metrics":       TargetDetailsKindMetrics,
	"functions":     TargetDetailsKindFunctions,
	"notifications": TargetDetailsKindNotifications,
}

// GetTargetDetailsKindEnumValues Enumerates the set of values for TargetDetailsKindEnum
func GetTargetDetailsKindEnumValues() []TargetDetailsKindEnum {
	values := make([]TargetDetailsKindEnum, 0)
	for _, v := range mappingTargetDetailsKind {
		values = append(values, v)
	}
	return values
}
