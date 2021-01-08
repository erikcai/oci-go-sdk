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
	"github.com/oracle/oci-go-sdk/v31/common"
)

// CreateDeployFunctionStageDetails Specifies the DeployFunction Stage.
type CreateDeployFunctionStageDetails struct {

	// Stage Identifier
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Pipeline Identifier
	PipelineId *string `mandatory:"true" json:"pipelineId"`

	StagePredecessorCollection *StagePredecessorCollection `mandatory:"true" json:"stagePredecessorCollection"`

	// Function environment ocid.
	FunctionEnvironmentId *string `mandatory:"true" json:"functionEnvironmentId"`

	// A docker image artifact ocid.
	DockerImageArtifactId *string `mandatory:"true" json:"dockerImageArtifactId"`

	// A timeout in seconds, this stage may take to return.
	TimeoutInSeconds *int `mandatory:"false" json:"timeoutInSeconds"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// User provided key/value pair configuration which assigned through constants or parameter.
	Config map[string]string `mandatory:"false" json:"config"`

	// Maximum usable memory for the function (MiB)
	MaxMemoryInMBs *int64 `mandatory:"false" json:"maxMemoryInMBs"`

	// Timeout for executions of the function. Value in seconds.
	FunctionTimeoutInSeconds *int `mandatory:"false" json:"functionTimeoutInSeconds"`
}

//GetDisplayName returns DisplayName
func (m CreateDeployFunctionStageDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetPipelineId returns PipelineId
func (m CreateDeployFunctionStageDetails) GetPipelineId() *string {
	return m.PipelineId
}

//GetTimeoutInSeconds returns TimeoutInSeconds
func (m CreateDeployFunctionStageDetails) GetTimeoutInSeconds() *int {
	return m.TimeoutInSeconds
}

//GetStagePredecessorCollection returns StagePredecessorCollection
func (m CreateDeployFunctionStageDetails) GetStagePredecessorCollection() *StagePredecessorCollection {
	return m.StagePredecessorCollection
}

//GetFreeformTags returns FreeformTags
func (m CreateDeployFunctionStageDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m CreateDeployFunctionStageDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateDeployFunctionStageDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateDeployFunctionStageDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateDeployFunctionStageDetails CreateDeployFunctionStageDetails
	s := struct {
		DiscriminatorParam string `json:"stageType"`
		MarshalTypeCreateDeployFunctionStageDetails
	}{
		"DEPLOY_FUNCTION",
		(MarshalTypeCreateDeployFunctionStageDetails)(m),
	}

	return json.Marshal(&s)
}
