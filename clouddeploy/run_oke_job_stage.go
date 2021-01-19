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

// RunOkeJobStage Specifies Run job on OKE stage.
type RunOkeJobStage struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// Application Identifier
	ApplicationId *string `mandatory:"true" json:"applicationId"`

	// Pipeline Identifier
	PipelineId *string `mandatory:"true" json:"pipelineId"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// OKE Cluster environment ocid.
	OkeClusterEnvironmentId *string `mandatory:"true" json:"okeClusterEnvironmentId"`

	// KubernetesManifest artifact ocid.
	KubernetesManifestArtifactId *string `mandatory:"true" json:"kubernetesManifestArtifactId"`

	// Default Namespace to be used for Kubernetes Job when not specified in the manifest
	Namespace *string `mandatory:"true" json:"namespace"`

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

	// Specifies synchronous or asynchronous mode of execution for this stage.
	ExecutionMode RunOkeJobStageExecutionModeEnum `mandatory:"true" json:"executionMode"`

	// The current state of the Stage.
	LifecycleState StageLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

//GetId returns Id
func (m RunOkeJobStage) GetId() *string {
	return m.Id
}

//GetDisplayName returns DisplayName
func (m RunOkeJobStage) GetDisplayName() *string {
	return m.DisplayName
}

//GetApplicationId returns ApplicationId
func (m RunOkeJobStage) GetApplicationId() *string {
	return m.ApplicationId
}

//GetPipelineId returns PipelineId
func (m RunOkeJobStage) GetPipelineId() *string {
	return m.PipelineId
}

//GetCompartmentId returns CompartmentId
func (m RunOkeJobStage) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetTimeCreated returns TimeCreated
func (m RunOkeJobStage) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m RunOkeJobStage) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetLifecycleState returns LifecycleState
func (m RunOkeJobStage) GetLifecycleState() StageLifecycleStateEnum {
	return m.LifecycleState
}

//GetLifecyleDetails returns LifecyleDetails
func (m RunOkeJobStage) GetLifecyleDetails() *string {
	return m.LifecyleDetails
}

//GetTimeoutInSeconds returns TimeoutInSeconds
func (m RunOkeJobStage) GetTimeoutInSeconds() *int {
	return m.TimeoutInSeconds
}

//GetStagePredecessorCollection returns StagePredecessorCollection
func (m RunOkeJobStage) GetStagePredecessorCollection() *StagePredecessorCollection {
	return m.StagePredecessorCollection
}

//GetFreeformTags returns FreeformTags
func (m RunOkeJobStage) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m RunOkeJobStage) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetSystemTags returns SystemTags
func (m RunOkeJobStage) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m RunOkeJobStage) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m RunOkeJobStage) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeRunOkeJobStage RunOkeJobStage
	s := struct {
		DiscriminatorParam string `json:"stageType"`
		MarshalTypeRunOkeJobStage
	}{
		"RUN_OKE_JOB",
		(MarshalTypeRunOkeJobStage)(m),
	}

	return json.Marshal(&s)
}

// RunOkeJobStageExecutionModeEnum Enum with underlying type: string
type RunOkeJobStageExecutionModeEnum string

// Set of constants representing the allowable values for RunOkeJobStageExecutionModeEnum
const (
	RunOkeJobStageExecutionModeSync  RunOkeJobStageExecutionModeEnum = "SYNC"
	RunOkeJobStageExecutionModeAsync RunOkeJobStageExecutionModeEnum = "ASYNC"
)

var mappingRunOkeJobStageExecutionMode = map[string]RunOkeJobStageExecutionModeEnum{
	"SYNC":  RunOkeJobStageExecutionModeSync,
	"ASYNC": RunOkeJobStageExecutionModeAsync,
}

// GetRunOkeJobStageExecutionModeEnumValues Enumerates the set of values for RunOkeJobStageExecutionModeEnum
func GetRunOkeJobStageExecutionModeEnumValues() []RunOkeJobStageExecutionModeEnum {
	values := make([]RunOkeJobStageExecutionModeEnum, 0)
	for _, v := range mappingRunOkeJobStageExecutionMode {
		values = append(values, v)
	}
	return values
}
