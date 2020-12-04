// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management Service APIs.
//
// This file contains the customer facing APIs for Database Management service.
//

package databasemanagement

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v30/common"
)

// ObjectStorageJobExecutionResultDetails Object Storage job execution result details.
type ObjectStorageJobExecutionResultDetails struct {

	// The Object Storage namespace used for job execution result storage.
	NamespaceName *string `mandatory:"false" json:"namespaceName"`

	// Name of the bucket used for job execution result storage.
	BucketName *string `mandatory:"false" json:"bucketName"`

	// Name of the object containing the job execution result.
	ObjectName *string `mandatory:"false" json:"objectName"`

	// Number of rows return in the result. Applicable only for QUERY SqlType.
	RowCount *int `mandatory:"false" json:"rowCount"`
}

func (m ObjectStorageJobExecutionResultDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ObjectStorageJobExecutionResultDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeObjectStorageJobExecutionResultDetails ObjectStorageJobExecutionResultDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeObjectStorageJobExecutionResultDetails
	}{
		"OBJECT_STORAGE",
		(MarshalTypeObjectStorageJobExecutionResultDetails)(m),
	}

	return json.Marshal(&s)
}
