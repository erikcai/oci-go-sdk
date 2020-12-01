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

// ArtifactSummary Summary of the Artifact.
type ArtifactSummary struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// Application Identifier
	ApplicationId *string `mandatory:"true" json:"applicationId"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Type of the Artifact
	ArtifactType *string `mandatory:"true" json:"artifactType"`

	ArtifactSource ArtifactSource `mandatory:"true" json:"artifactSource"`

	// Artifact Identifier, can be renamed
	DisplayName *string `mandatory:"false" json:"displayName"`

	ArtifactParameters *ArtifactParameterCollection `mandatory:"false" json:"artifactParameters"`

	// The time the the Artifact was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the Artifact was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the Artifact.
	LifecycleState ArtifactLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m ArtifactSummary) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *ArtifactSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName        *string                           `json:"displayName"`
		ArtifactParameters *ArtifactParameterCollection      `json:"artifactParameters"`
		TimeCreated        *common.SDKTime                   `json:"timeCreated"`
		TimeUpdated        *common.SDKTime                   `json:"timeUpdated"`
		LifecycleState     ArtifactLifecycleStateEnum        `json:"lifecycleState"`
		LifecycleDetails   *string                           `json:"lifecycleDetails"`
		FreeformTags       map[string]string                 `json:"freeformTags"`
		DefinedTags        map[string]map[string]interface{} `json:"definedTags"`
		SystemTags         map[string]map[string]interface{} `json:"systemTags"`
		Id                 *string                           `json:"id"`
		ApplicationId      *string                           `json:"applicationId"`
		CompartmentId      *string                           `json:"compartmentId"`
		ArtifactType       *string                           `json:"artifactType"`
		ArtifactSource     artifactsource                    `json:"artifactSource"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.ArtifactParameters = model.ArtifactParameters

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleState = model.LifecycleState

	m.LifecycleDetails = model.LifecycleDetails

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.ApplicationId = model.ApplicationId

	m.CompartmentId = model.CompartmentId

	m.ArtifactType = model.ArtifactType

	nn, e = model.ArtifactSource.UnmarshalPolymorphicJSON(model.ArtifactSource.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ArtifactSource = nn.(ArtifactSource)
	} else {
		m.ArtifactSource = nil
	}

	return
}
