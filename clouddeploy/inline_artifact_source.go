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
	"github.com/oracle/oci-go-sdk/v31/common"
)

// InlineArtifactSource Specifies the Inline Artifact Source Details
type InlineArtifactSource struct {

	// base64 Encoded String
	Base64EncodedContent []byte `mandatory:"true" json:"base64EncodedContent"`
}

func (m InlineArtifactSource) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m InlineArtifactSource) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeInlineArtifactSource InlineArtifactSource
	s := struct {
		DiscriminatorParam string `json:"artifactSourceType"`
		MarshalTypeInlineArtifactSource
	}{
		"INLINE",
		(MarshalTypeInlineArtifactSource)(m),
	}

	return json.Marshal(&s)
}
