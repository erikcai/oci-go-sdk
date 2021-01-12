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

// CreateInlineArtifactSourceDetails Specifies the Inline Artifact Source Details
type CreateInlineArtifactSourceDetails struct {

	// base64 Encoded String
	Base64EncodedContent []byte `mandatory:"true" json:"base64EncodedContent"`
}

func (m CreateInlineArtifactSourceDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateInlineArtifactSourceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateInlineArtifactSourceDetails CreateInlineArtifactSourceDetails
	s := struct {
		DiscriminatorParam string `json:"artifactSourceType"`
		MarshalTypeCreateInlineArtifactSourceDetails
	}{
		"INLINE",
		(MarshalTypeCreateInlineArtifactSourceDetails)(m),
	}

	return json.Marshal(&s)
}
