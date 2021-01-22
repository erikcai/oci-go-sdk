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

// CreateRunJobOnComputeInstanceStageDetails Specifies Run job on an instance stage.
type CreateRunJobOnComputeInstanceStageDetails struct {

	// Stage Identifier
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Pipeline Identifier
	PipelineId *string `mandatory:"true" json:"pipelineId"`

	StagePredecessorCollection *StagePredecessorCollection `mandatory:"true" json:"stagePredecessorCollection"`

	// Instance OCID on which this stage runs the job.
	ComputeInstanceId *string `mandatory:"true" json:"computeInstanceId"`

	// The OCID of the artifact that contains the deployment specification.
	PrimaryArtifactId *string `mandatory:"true" json:"primaryArtifactId"`

	// A timeout in seconds, this stage may take to return.
	TimeoutInSeconds *int `mandatory:"false" json:"timeoutInSeconds"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// File artifact ocids.
	AdditionalArtifactIds []string `mandatory:"false" json:"additionalArtifactIds"`

	// The path to the deployment specification file in the primary artifact for this stage's Environment. The default location if not specified is deployment_spec.yaml
	DeploymentSpecFile *string `mandatory:"false" json:"deploymentSpecFile"`
}

//GetDisplayName returns DisplayName
func (m CreateRunJobOnComputeInstanceStageDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetPipelineId returns PipelineId
func (m CreateRunJobOnComputeInstanceStageDetails) GetPipelineId() *string {
	return m.PipelineId
}

//GetTimeoutInSeconds returns TimeoutInSeconds
func (m CreateRunJobOnComputeInstanceStageDetails) GetTimeoutInSeconds() *int {
	return m.TimeoutInSeconds
}

//GetStagePredecessorCollection returns StagePredecessorCollection
func (m CreateRunJobOnComputeInstanceStageDetails) GetStagePredecessorCollection() *StagePredecessorCollection {
	return m.StagePredecessorCollection
}

//GetFreeformTags returns FreeformTags
func (m CreateRunJobOnComputeInstanceStageDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m CreateRunJobOnComputeInstanceStageDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateRunJobOnComputeInstanceStageDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateRunJobOnComputeInstanceStageDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateRunJobOnComputeInstanceStageDetails CreateRunJobOnComputeInstanceStageDetails
	s := struct {
		DiscriminatorParam string `json:"stageType"`
		MarshalTypeCreateRunJobOnComputeInstanceStageDetails
	}{
		"RUN_JOB_ON_COMPUTE_INSTANCE",
		(MarshalTypeCreateRunJobOnComputeInstanceStageDetails)(m),
	}

	return json.Marshal(&s)
}
