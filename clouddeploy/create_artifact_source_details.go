// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DeploymentService API
//
// A description of the DeploymentService API
//

package clouddeploy

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v29/common"
)

// CreateArtifactSourceDetails Specifies source of an artifact
type CreateArtifactSourceDetails interface {
}

type createartifactsourcedetails struct {
	JsonData           []byte
	ArtifactSourceType string `json:"artifactSourceType"`
}

// UnmarshalJSON unmarshals json
func (m *createartifactsourcedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreateartifactsourcedetails createartifactsourcedetails
	s := struct {
		Model Unmarshalercreateartifactsourcedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ArtifactSourceType = s.Model.ArtifactSourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createartifactsourcedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ArtifactSourceType {
	case "INLINE":
		mm := CreateInlineArtifactSourceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OBJECT_STORAGE":
		mm := CreateObjectStorageArtifactSourceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCIR":
		mm := CreateOcirArtifactSourceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m createartifactsourcedetails) String() string {
	return common.PointerString(m)
}
