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

// UpdateZipUploadConfigSourceDetails The representation of UpdateZipUploadConfigSourceDetails
type UpdateZipUploadConfigSourceDetails struct {

	// The path of the directory from which to run terraform. If not specified the the root will be used.
	WorkingDirectory *string `mandatory:"false" json:"workingDirectory"`

	ZipFileBase64Encoded *string `mandatory:"false" json:"zipFileBase64Encoded"`
}

//GetWorkingDirectory returns WorkingDirectory
func (m UpdateZipUploadConfigSourceDetails) GetWorkingDirectory() *string {
	return m.WorkingDirectory
}

func (m UpdateZipUploadConfigSourceDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m UpdateZipUploadConfigSourceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateZipUploadConfigSourceDetails UpdateZipUploadConfigSourceDetails
	s := struct {
		DiscriminatorParam string `json:"configSourceType"`
		MarshalTypeUpdateZipUploadConfigSourceDetails
	}{
		"ZIP_UPLOAD",
		(MarshalTypeUpdateZipUploadConfigSourceDetails)(m),
	}

	return json.Marshal(&s)
}
