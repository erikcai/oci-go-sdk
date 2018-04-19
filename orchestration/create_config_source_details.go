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

// CreateConfigSourceDetails The representation of CreateConfigSourceDetails
type CreateConfigSourceDetails interface {
}

type createconfigsourcedetails struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *createconfigsourcedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreateconfigsourcedetails createconfigsourcedetails
	s := struct {
		Model Unmarshalercreateconfigsourcedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createconfigsourcedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {
	var err error
	switch m.Type {
	case "ZIPUPLOAD":
		mm := CreateZipUploadConfigSourceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return m, nil
	}
}

func (m createconfigsourcedetails) String() string {
	return common.PointerString(m)
}
