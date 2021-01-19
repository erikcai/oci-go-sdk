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
	"github.com/oracle/oci-go-sdk/v33/common"
)

// CreateUpdateFunctionApplicationStageDetails Specifies the creation of an Update Function Application Stage, either by updating the value of a property or by providing the name of an existing image in the Docker registry that you now want to base the function on, instead of the previously specified image.
type CreateUpdateFunctionApplicationStageDetails struct {

	// Stage Identifier
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Pipeline Identifier
	PipelineId *string `mandatory:"true" json:"pipelineId"`

	StagePredecessorCollection *StagePredecessorCollection `mandatory:"true" json:"stagePredecessorCollection"`

	// Function environment ocid.
	FunctionEnvironmentId *string `mandatory:"true" json:"functionEnvironmentId"`

	// A timeout in seconds, this stage may take to return.
	TimeoutInSeconds *int `mandatory:"false" json:"timeoutInSeconds"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A docker image artifact ocid.
	DockerImageArtifactId *string `mandatory:"false" json:"dockerImageArtifactId"`

	// User provided key/value pair configuration which is assigned through constants or parameter.
	Config map[string]string `mandatory:"false" json:"config"`
}

//GetDisplayName returns DisplayName
func (m CreateUpdateFunctionApplicationStageDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetPipelineId returns PipelineId
func (m CreateUpdateFunctionApplicationStageDetails) GetPipelineId() *string {
	return m.PipelineId
}

//GetTimeoutInSeconds returns TimeoutInSeconds
func (m CreateUpdateFunctionApplicationStageDetails) GetTimeoutInSeconds() *int {
	return m.TimeoutInSeconds
}

//GetStagePredecessorCollection returns StagePredecessorCollection
func (m CreateUpdateFunctionApplicationStageDetails) GetStagePredecessorCollection() *StagePredecessorCollection {
	return m.StagePredecessorCollection
}

//GetFreeformTags returns FreeformTags
func (m CreateUpdateFunctionApplicationStageDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m CreateUpdateFunctionApplicationStageDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateUpdateFunctionApplicationStageDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateUpdateFunctionApplicationStageDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateUpdateFunctionApplicationStageDetails CreateUpdateFunctionApplicationStageDetails
	s := struct {
		DiscriminatorParam string `json:"stageType"`
		MarshalTypeCreateUpdateFunctionApplicationStageDetails
	}{
		"UPDATE_FUNCTION_APPLICATION",
		(MarshalTypeCreateUpdateFunctionApplicationStageDetails)(m),
	}

	return json.Marshal(&s)
}
