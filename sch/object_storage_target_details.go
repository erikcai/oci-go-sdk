// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Connectors API
//
// A description of the Connectors API
//

package sch

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// ObjectStorageTargetDetails The object storage target.
type ObjectStorageTargetDetails struct {

	// The name of the bucket.
	BucketName *string `mandatory:"true" json:"bucketName"`

	// The prefix of the objects.
	ObjectNamePrefix *string `mandatory:"false" json:"objectNamePrefix"`
}

func (m ObjectStorageTargetDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ObjectStorageTargetDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeObjectStorageTargetDetails ObjectStorageTargetDetails
	s := struct {
		DiscriminatorParam string `json:"kind"`
		MarshalTypeObjectStorageTargetDetails
	}{
		"objectStorage",
		(MarshalTypeObjectStorageTargetDetails)(m),
	}

	return json.Marshal(&s)
}
