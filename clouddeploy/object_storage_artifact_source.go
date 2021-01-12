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
	"github.com/oracle/oci-go-sdk/v32/common"
)

// ObjectStorageArtifactSource Specifies the Object Storage Artifact Source Details
type ObjectStorageArtifactSource struct {

	// Specifies the object storage bucket
	BucketName *string `mandatory:"true" json:"bucketName"`

	// Specifies the object storage namespace
	NamespaceName *string `mandatory:"true" json:"namespaceName"`

	ObjectNameFilter *ObjectNameFilter `mandatory:"false" json:"objectNameFilter"`
}

func (m ObjectStorageArtifactSource) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ObjectStorageArtifactSource) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeObjectStorageArtifactSource ObjectStorageArtifactSource
	s := struct {
		DiscriminatorParam string `json:"artifactSourceType"`
		MarshalTypeObjectStorageArtifactSource
	}{
		"OBJECT_STORAGE",
		(MarshalTypeObjectStorageArtifactSource)(m),
	}

	return json.Marshal(&s)
}
