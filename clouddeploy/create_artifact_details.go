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
	"github.com/erikcai/oci-go-sdk/v33/common"
)

// CreateArtifactDetails The information about new Artifact.
type CreateArtifactDetails struct {

	// Artifact Name
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Type of the Artifact
	ArtifactType *string `mandatory:"true" json:"artifactType"`

	ArtifactSource CreateArtifactSourceDetails `mandatory:"true" json:"artifactSource"`

	// Application Identifier
	ApplicationId *string `mandatory:"true" json:"applicationId"`

	ArtifactParameters *ArtifactParameterCollection `mandatory:"false" json:"artifactParameters"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateArtifactDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *CreateArtifactDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ArtifactParameters *ArtifactParameterCollection      `json:"artifactParameters"`
		FreeformTags       map[string]string                 `json:"freeformTags"`
		DefinedTags        map[string]map[string]interface{} `json:"definedTags"`
		DisplayName        *string                           `json:"displayName"`
		ArtifactType       *string                           `json:"artifactType"`
		ArtifactSource     createartifactsourcedetails       `json:"artifactSource"`
		ApplicationId      *string                           `json:"applicationId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.ArtifactParameters = model.ArtifactParameters

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.DisplayName = model.DisplayName

	m.ArtifactType = model.ArtifactType

	nn, e = model.ArtifactSource.UnmarshalPolymorphicJSON(model.ArtifactSource.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ArtifactSource = nn.(CreateArtifactSourceDetails)
	} else {
		m.ArtifactSource = nil
	}

	m.ApplicationId = model.ApplicationId

	return
}
