// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Data Integration Service API Specification
//
// Data Integration Service API Specification
//

package dataintegration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// OracleAdwcWriteAttribute auto generated description
type OracleAdwcWriteAttribute struct {

	// bucketName
	BucketName *string `mandatory:"false" json:"bucketName"`

	// stagingFileName
	StagingFileName *string `mandatory:"false" json:"stagingFileName"`

	StagingDataAsset *DataAsset `mandatory:"false" json:"stagingDataAsset"`

	StagingConnection *Connection `mandatory:"false" json:"stagingConnection"`
}

func (m OracleAdwcWriteAttribute) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m OracleAdwcWriteAttribute) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOracleAdwcWriteAttribute OracleAdwcWriteAttribute
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeOracleAdwcWriteAttribute
	}{
		"ORACLEADWCWRITEATTRIBUTE",
		(MarshalTypeOracleAdwcWriteAttribute)(m),
	}

	return json.Marshal(&s)
}
