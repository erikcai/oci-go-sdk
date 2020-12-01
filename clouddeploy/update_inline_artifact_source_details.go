// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DeploymentService API
//
// A description of the DeploymentService API
//

package clouddeploy

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v30/common"
)

// UpdateInlineArtifactSourceDetails Specifies the Inline Artifact Source Details
type UpdateInlineArtifactSourceDetails struct {

	// base64 Encoded String
	Base64EncodedContent []byte `mandatory:"false" json:"base64EncodedContent"`
}

func (m UpdateInlineArtifactSourceDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m UpdateInlineArtifactSourceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateInlineArtifactSourceDetails UpdateInlineArtifactSourceDetails
	s := struct {
		DiscriminatorParam string `json:"artifactSourceType"`
		MarshalTypeUpdateInlineArtifactSourceDetails
	}{
		"INLINE",
		(MarshalTypeUpdateInlineArtifactSourceDetails)(m),
	}

	return json.Marshal(&s)
}
