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

// RunJobOnComputeInstanceStage Specifies Run job on an compute instance stage.
type RunJobOnComputeInstanceStage struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// Application Identifier
	ApplicationId *string `mandatory:"true" json:"applicationId"`

	// Pipeline Identifier
	PipelineId *string `mandatory:"true" json:"pipelineId"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Instance OCID on which this stage runs the job.
	ComputeInstanceId *string `mandatory:"true" json:"computeInstanceId"`

	// The OCID of the artifact that contains the deployment specification.
	PrimaryArtifactId *string `mandatory:"true" json:"primaryArtifactId"`

	// Stage Identifier, can be renamed
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The time the the Stage was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the Stage was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecyleDetails *string `mandatory:"false" json:"lifecyleDetails"`

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

	// File artifact ocids.
	AdditionalArtifactIds []string `mandatory:"false" json:"additionalArtifactIds"`

	// The path to the deployment specification file in the primary artifact for this stage's Environment. The default location if not specified is deployment_spec.yaml
	DeploymentSpecFile *string `mandatory:"false" json:"deploymentSpecFile"`

	// The current state of the Stage.
	LifecycleState StageLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

//GetId returns Id
func (m RunJobOnComputeInstanceStage) GetId() *string {
	return m.Id
}

//GetDisplayName returns DisplayName
func (m RunJobOnComputeInstanceStage) GetDisplayName() *string {
	return m.DisplayName
}

//GetApplicationId returns ApplicationId
func (m RunJobOnComputeInstanceStage) GetApplicationId() *string {
	return m.ApplicationId
}

//GetPipelineId returns PipelineId
func (m RunJobOnComputeInstanceStage) GetPipelineId() *string {
	return m.PipelineId
}

//GetCompartmentId returns CompartmentId
func (m RunJobOnComputeInstanceStage) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetTimeCreated returns TimeCreated
func (m RunJobOnComputeInstanceStage) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m RunJobOnComputeInstanceStage) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetLifecycleState returns LifecycleState
func (m RunJobOnComputeInstanceStage) GetLifecycleState() StageLifecycleStateEnum {
	return m.LifecycleState
}

//GetLifecyleDetails returns LifecyleDetails
func (m RunJobOnComputeInstanceStage) GetLifecyleDetails() *string {
	return m.LifecyleDetails
}

//GetTimeoutInSeconds returns TimeoutInSeconds
func (m RunJobOnComputeInstanceStage) GetTimeoutInSeconds() *int {
	return m.TimeoutInSeconds
}

//GetStagePredecessorCollection returns StagePredecessorCollection
func (m RunJobOnComputeInstanceStage) GetStagePredecessorCollection() *StagePredecessorCollection {
	return m.StagePredecessorCollection
}

//GetFreeformTags returns FreeformTags
func (m RunJobOnComputeInstanceStage) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m RunJobOnComputeInstanceStage) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetSystemTags returns SystemTags
func (m RunJobOnComputeInstanceStage) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m RunJobOnComputeInstanceStage) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m RunJobOnComputeInstanceStage) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeRunJobOnComputeInstanceStage RunJobOnComputeInstanceStage
	s := struct {
		DiscriminatorParam string `json:"stageType"`
		MarshalTypeRunJobOnComputeInstanceStage
	}{
		"RUN_JOB_ON_COMPUTE_INSTANCE",
		(MarshalTypeRunJobOnComputeInstanceStage)(m),
	}

	return json.Marshal(&s)
}
