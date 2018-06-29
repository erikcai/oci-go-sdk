// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Oracle Resource Manager
//
// Oracle Resource Manager API
//

package resourcemanager

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// CreateConfigSourceDetails The representation of CreateConfigSourceDetails
type CreateConfigSourceDetails interface {

	// The path of the directory from which to run terraform. If not specified the the root will be used.
	GetWorkingDirectory() *string
}

type createconfigsourcedetails struct {
	JsonData         []byte
	WorkingDirectory *string `mandatory:"false" json:"workingDirectory"`
	ConfigSourceType string  `json:"configSourceType"`
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
	m.WorkingDirectory = s.Model.WorkingDirectory
	m.ConfigSourceType = s.Model.ConfigSourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createconfigsourcedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {
	var err error
	switch m.ConfigSourceType {
	case "ZIP_UPLOAD":
		mm := CreateZipUploadConfigSourceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return m, nil
	}
}

//GetWorkingDirectory returns WorkingDirectory
func (m createconfigsourcedetails) GetWorkingDirectory() *string {
	return m.WorkingDirectory
}

func (m createconfigsourcedetails) String() string {
	return common.PointerString(m)
}
