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

// ConfigSource The representation of ConfigSource
type ConfigSource interface {
}

type configsource struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *configsource) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerconfigsource configsource
	s := struct {
		Model Unmarshalerconfigsource
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *configsource) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {
	var err error
	switch m.Type {
	case "ZIPUPLOAD":
		mm := ZipUploadConfigSource{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return m, nil
	}
}

func (m configsource) String() string {
	return common.PointerString(m)
}
