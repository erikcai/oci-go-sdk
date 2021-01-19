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
	"github.com/oracle/oci-go-sdk/v33/common"
)

// ArtifactSource Specifies source of an artifact
type ArtifactSource interface {
}

type artifactsource struct {
	JsonData           []byte
	ArtifactSourceType string `json:"artifactSourceType"`
}

// UnmarshalJSON unmarshals json
func (m *artifactsource) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerartifactsource artifactsource
	s := struct {
		Model Unmarshalerartifactsource
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ArtifactSourceType = s.Model.ArtifactSourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *artifactsource) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ArtifactSourceType {
	case "INLINE":
		mm := InlineArtifactSource{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OBJECT_STORAGE":
		mm := ObjectStorageArtifactSource{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCIR":
		mm := OcirArtifactSource{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m artifactsource) String() string {
	return common.PointerString(m)
}

// ArtifactSourceArtifactSourceTypeEnum Enum with underlying type: string
type ArtifactSourceArtifactSourceTypeEnum string

// Set of constants representing the allowable values for ArtifactSourceArtifactSourceTypeEnum
const (
	ArtifactSourceArtifactSourceTypeObjectStorage ArtifactSourceArtifactSourceTypeEnum = "OBJECT_STORAGE"
	ArtifactSourceArtifactSourceTypeInline        ArtifactSourceArtifactSourceTypeEnum = "INLINE"
	ArtifactSourceArtifactSourceTypeOcir          ArtifactSourceArtifactSourceTypeEnum = "OCIR"
)

var mappingArtifactSourceArtifactSourceType = map[string]ArtifactSourceArtifactSourceTypeEnum{
	"OBJECT_STORAGE": ArtifactSourceArtifactSourceTypeObjectStorage,
	"INLINE":         ArtifactSourceArtifactSourceTypeInline,
	"OCIR":           ArtifactSourceArtifactSourceTypeOcir,
}

// GetArtifactSourceArtifactSourceTypeEnumValues Enumerates the set of values for ArtifactSourceArtifactSourceTypeEnum
func GetArtifactSourceArtifactSourceTypeEnumValues() []ArtifactSourceArtifactSourceTypeEnum {
	values := make([]ArtifactSourceArtifactSourceTypeEnum, 0)
	for _, v := range mappingArtifactSourceArtifactSourceType {
		values = append(values, v)
	}
	return values
}
