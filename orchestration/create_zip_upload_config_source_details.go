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

// CreateZipUploadConfigSourceDetails The representation of CreateZipUploadConfigSourceDetails
type CreateZipUploadConfigSourceDetails struct {
	ZipFileBase64Encoded *string `mandatory:"true" json:"zipFileBase64Encoded"`
}

func (m CreateZipUploadConfigSourceDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateZipUploadConfigSourceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateZipUploadConfigSourceDetails CreateZipUploadConfigSourceDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeCreateZipUploadConfigSourceDetails
	}{
		"ZIPUPLOAD",
		(MarshalTypeCreateZipUploadConfigSourceDetails)(m),
	}

	return json.Marshal(&s)
}
