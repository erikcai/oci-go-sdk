// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Orchestration
//
// Orchestration Maestro API
//

package orchestration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateConfigSourceDetails The representation of UpdateConfigSourceDetails
type UpdateConfigSourceDetails interface {
}

type updateconfigsourcedetails struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *updateconfigsourcedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdateconfigsourcedetails updateconfigsourcedetails
	s := struct {
		Model Unmarshalerupdateconfigsourcedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updateconfigsourcedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {
	var err error
	switch m.Type {
	case "ZIPUPLOAD":
		mm := UpdateZipUploadConfigSourceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return m, nil
	}
}

func (m updateconfigsourcedetails) String() string {
	return common.PointerString(m)
}
