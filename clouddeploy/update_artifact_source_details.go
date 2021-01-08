// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DeploymentService API
//
// A description of the DeploymentService API
//

package clouddeploy

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v31/common"
)

// UpdateArtifactSourceDetails Specifies source of an artifact
type UpdateArtifactSourceDetails interface {
}

type updateartifactsourcedetails struct {
	JsonData           []byte
	ArtifactSourceType string `json:"artifactSourceType"`
}

// UnmarshalJSON unmarshals json
func (m *updateartifactsourcedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdateartifactsourcedetails updateartifactsourcedetails
	s := struct {
		Model Unmarshalerupdateartifactsourcedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ArtifactSourceType = s.Model.ArtifactSourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updateartifactsourcedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ArtifactSourceType {
	case "INLINE":
		mm := UpdateInlineArtifactSourceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCIR":
		mm := UpdateOcirArtifactSourceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OBJECT_STORAGE":
		mm := UpdateObjectStorageArtifactSourceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m updateartifactsourcedetails) String() string {
	return common.PointerString(m)
}
