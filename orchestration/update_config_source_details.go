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

	// The path of the directory from which to run terraform. If not specified the the root will be used.
	GetWorkingDirectory() *string
}

type updateconfigsourcedetails struct {
	JsonData         []byte
	WorkingDirectory *string `mandatory:"false" json:"workingDirectory"`
	ConfigSourceType string  `json:"configSourceType"`
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
	m.WorkingDirectory = s.Model.WorkingDirectory
	m.ConfigSourceType = s.Model.ConfigSourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updateconfigsourcedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {
	var err error
	switch m.ConfigSourceType {
	case "ZIP_UPLOAD":
		mm := UpdateZipUploadConfigSourceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return m, nil
	}
}

//GetWorkingDirectory returns WorkingDirectory
func (m updateconfigsourcedetails) GetWorkingDirectory() *string {
	return m.WorkingDirectory
}

func (m updateconfigsourcedetails) String() string {
	return common.PointerString(m)
}
