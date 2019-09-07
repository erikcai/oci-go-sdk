// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// PublicLoggingControlplane API
//
// PublicLoggingControlplane API specification
//

package publiclogging

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// Source The source the log object comes from.
type Source interface {
}

type source struct {
	JsonData   []byte
	SourceType string `json:"sourceType"`
}

// UnmarshalJSON unmarshals json
func (m *source) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalersource source
	s := struct {
		Model Unmarshalersource
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.SourceType = s.Model.SourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *source) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SourceType {
	case "OCISERVICE":
		mm := OciService{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m source) String() string {
	return common.PointerString(m)
}
