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

// UpdateOcirArtifactSourceDetails Specifies OCIR details - optionally include tag
type UpdateOcirArtifactSourceDetails struct {

	// Specifies OCIR Image Path - optionally include tag
	ImageUri *string `mandatory:"false" json:"imageUri"`

	// Specifies image digest for the version of the image
	ImageDigest *string `mandatory:"false" json:"imageDigest"`
}

func (m UpdateOcirArtifactSourceDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m UpdateOcirArtifactSourceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateOcirArtifactSourceDetails UpdateOcirArtifactSourceDetails
	s := struct {
		DiscriminatorParam string `json:"artifactSourceType"`
		MarshalTypeUpdateOcirArtifactSourceDetails
	}{
		"OCIR",
		(MarshalTypeUpdateOcirArtifactSourceDetails)(m),
	}

	return json.Marshal(&s)
}
