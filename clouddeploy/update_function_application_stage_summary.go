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
	"github.com/oracle/oci-go-sdk/v29/common"
)

// UpdateFunctionApplicationStageSummary Specifies the summary of an Update Function Application Stage, either by updating the value of a property or by providing the name of an existing image in the Docker registry that you now want to base the function on, instead of the previously specified image.
type UpdateFunctionApplicationStageSummary struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// Application Identifier
	ApplicationId *string `mandatory:"true" json:"applicationId"`

	// Pipeline Identifier
	PipelineId *string `mandatory:"true" json:"pipelineId"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Function environment ocid.
	FunctionEnvironmentId *string `mandatory:"true" json:"functionEnvironmentId"`

	// Stage Identifier, can be renamed
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The time the the Stage was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the Stage was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// A timeout in seconds, this stage may take to return.
	TimeoutInSeconds *int `mandatory:"false" json:"timeoutInSeconds"`

	StagePredecessorCollection *StagePredecessorCollection `mandatory:"false" json:"stagePredecessorCollection"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// A docker image artifact ocid.
	DockerImageArtifactId *string `mandatory:"false" json:"dockerImageArtifactId"`

	// User provided key/value pair configuration which is assigned through constants or parameter.
	Config map[string]string `mandatory:"false" json:"config"`

	// The current state of the Stage.
	LifecycleState StageLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

//GetId returns Id
func (m UpdateFunctionApplicationStageSummary) GetId() *string {
	return m.Id
}

//GetDisplayName returns DisplayName
func (m UpdateFunctionApplicationStageSummary) GetDisplayName() *string {
	return m.DisplayName
}

//GetApplicationId returns ApplicationId
func (m UpdateFunctionApplicationStageSummary) GetApplicationId() *string {
	return m.ApplicationId
}

//GetPipelineId returns PipelineId
func (m UpdateFunctionApplicationStageSummary) GetPipelineId() *string {
	return m.PipelineId
}

//GetCompartmentId returns CompartmentId
func (m UpdateFunctionApplicationStageSummary) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetTimeCreated returns TimeCreated
func (m UpdateFunctionApplicationStageSummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m UpdateFunctionApplicationStageSummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetLifecycleState returns LifecycleState
func (m UpdateFunctionApplicationStageSummary) GetLifecycleState() StageLifecycleStateEnum {
	return m.LifecycleState
}

//GetLifecycleDetails returns LifecycleDetails
func (m UpdateFunctionApplicationStageSummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

//GetTimeoutInSeconds returns TimeoutInSeconds
func (m UpdateFunctionApplicationStageSummary) GetTimeoutInSeconds() *int {
	return m.TimeoutInSeconds
}

//GetStagePredecessorCollection returns StagePredecessorCollection
func (m UpdateFunctionApplicationStageSummary) GetStagePredecessorCollection() *StagePredecessorCollection {
	return m.StagePredecessorCollection
}

//GetFreeformTags returns FreeformTags
func (m UpdateFunctionApplicationStageSummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m UpdateFunctionApplicationStageSummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetSystemTags returns SystemTags
func (m UpdateFunctionApplicationStageSummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m UpdateFunctionApplicationStageSummary) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m UpdateFunctionApplicationStageSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateFunctionApplicationStageSummary UpdateFunctionApplicationStageSummary
	s := struct {
		DiscriminatorParam string `json:"stageType"`
		MarshalTypeUpdateFunctionApplicationStageSummary
	}{
		"UPDATE_FUNCTION_APPLICATION",
		(MarshalTypeUpdateFunctionApplicationStageSummary)(m),
	}

	return json.Marshal(&s)
}
