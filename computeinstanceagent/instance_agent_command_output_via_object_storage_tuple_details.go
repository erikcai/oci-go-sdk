// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// InstanceAgentService API
//
// Instance Agent Service API
//

package computeinstanceagent

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v28/common"
)

// InstanceAgentCommandOutputViaObjectStorageTupleDetails Command output via object storage tuple.
type InstanceAgentCommandOutputViaObjectStorageTupleDetails struct {

	// The Object Storage bucket for the command output.
	BucketName *string `mandatory:"true" json:"bucketName"`

	// The Object Storage namespace for the command output.
	NamespaceName *string `mandatory:"true" json:"namespaceName"`

	// The Object Storage name for the command output.
	ObjectName *string `mandatory:"true" json:"objectName"`
}

func (m InstanceAgentCommandOutputViaObjectStorageTupleDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m InstanceAgentCommandOutputViaObjectStorageTupleDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeInstanceAgentCommandOutputViaObjectStorageTupleDetails InstanceAgentCommandOutputViaObjectStorageTupleDetails
	s := struct {
		DiscriminatorParam string `json:"outputType"`
		MarshalTypeInstanceAgentCommandOutputViaObjectStorageTupleDetails
	}{
		"OBJECT_STORAGE_TUPLE",
		(MarshalTypeInstanceAgentCommandOutputViaObjectStorageTupleDetails)(m),
	}

	return json.Marshal(&s)
}
