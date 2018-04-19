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

// ZipUploadConfigSource The representation of ZipUploadConfigSource
type ZipUploadConfigSource struct {
	Dummy *string `mandatory:"false" json:"dummy"`
}

func (m ZipUploadConfigSource) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ZipUploadConfigSource) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeZipUploadConfigSource ZipUploadConfigSource
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeZipUploadConfigSource
	}{
		"ZIPUPLOAD",
		(MarshalTypeZipUploadConfigSource)(m),
	}

	return json.Marshal(&s)
}
