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

// OkeDeploymentStage Specifies the OKEDeployment stage.
type OkeDeploymentStage struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// Application Identifier
	ApplicationId *string `mandatory:"true" json:"applicationId"`

	// Pipeline Identifier
	PipelineId *string `mandatory:"true" json:"pipelineId"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// OkeCluster environment ocid for deployment.
	OkeClusterEnvironmentId *string `mandatory:"true" json:"okeClusterEnvironmentId"`

	// KubernetesManifest artifact ocid, this manifest should not include any Job resource.
	KubernetesManifestArtifactId *string `mandatory:"true" json:"kubernetesManifestArtifactId"`

	// Default Namespace to be used for Kubernetes Deployment when not specified in the manifest
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

	// The current state of the Stage.
	LifecycleState StageLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

//GetId returns Id
func (m OkeDeploymentStage) GetId() *string {
	return m.Id
}

//GetDisplayName returns DisplayName
func (m OkeDeploymentStage) GetDisplayName() *string {
	return m.DisplayName
}

//GetApplicationId returns ApplicationId
func (m OkeDeploymentStage) GetApplicationId() *string {
	return m.ApplicationId
}

//GetPipelineId returns PipelineId
func (m OkeDeploymentStage) GetPipelineId() *string {
	return m.PipelineId
}

//GetCompartmentId returns CompartmentId
func (m OkeDeploymentStage) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetTimeCreated returns TimeCreated
func (m OkeDeploymentStage) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m OkeDeploymentStage) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetLifecycleState returns LifecycleState
func (m OkeDeploymentStage) GetLifecycleState() StageLifecycleStateEnum {
	return m.LifecycleState
}

//GetLifecyleDetails returns LifecyleDetails
func (m OkeDeploymentStage) GetLifecyleDetails() *string {
	return m.LifecyleDetails
}

//GetTimeoutInSeconds returns TimeoutInSeconds
func (m OkeDeploymentStage) GetTimeoutInSeconds() *int {
	return m.TimeoutInSeconds
}

//GetStagePredecessorCollection returns StagePredecessorCollection
func (m OkeDeploymentStage) GetStagePredecessorCollection() *StagePredecessorCollection {
	return m.StagePredecessorCollection
}

//GetFreeformTags returns FreeformTags
func (m OkeDeploymentStage) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m OkeDeploymentStage) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetSystemTags returns SystemTags
func (m OkeDeploymentStage) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m OkeDeploymentStage) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m OkeDeploymentStage) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOkeDeploymentStage OkeDeploymentStage
	s := struct {
		DiscriminatorParam string `json:"stageType"`
		MarshalTypeOkeDeploymentStage
	}{
		"OKE_DEPLOYMENT",
		(MarshalTypeOkeDeploymentStage)(m),
	}

	return json.Marshal(&s)
}
