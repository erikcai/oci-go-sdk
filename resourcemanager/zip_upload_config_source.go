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

// ZipUploadConfigSource The representation of ZipUploadConfigSource
type ZipUploadConfigSource struct {

	// The path of the directory from which to run terraform. If not specified the the root will be used.
	WorkingDirectory *string `mandatory:"false" json:"workingDirectory"`
}

//GetWorkingDirectory returns WorkingDirectory
func (m ZipUploadConfigSource) GetWorkingDirectory() *string {
	return m.WorkingDirectory
}

func (m ZipUploadConfigSource) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ZipUploadConfigSource) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeZipUploadConfigSource ZipUploadConfigSource
	s := struct {
		DiscriminatorParam string `json:"configSourceType"`
		MarshalTypeZipUploadConfigSource
	}{
		"ZIP_UPLOAD",
		(MarshalTypeZipUploadConfigSource)(m),
	}

	return json.Marshal(&s)
}
