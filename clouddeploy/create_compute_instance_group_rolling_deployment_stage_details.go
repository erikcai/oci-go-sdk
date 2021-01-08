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

// CreateComputeInstanceGroupRollingDeploymentStageDetails Specifies the Instance group rolling Deployment Stage.
type CreateComputeInstanceGroupRollingDeploymentStageDetails struct {

	// Stage Identifier
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Pipeline Identifier
	PipelineId *string `mandatory:"true" json:"pipelineId"`

	StagePredecessorCollection *StagePredecessorCollection `mandatory:"true" json:"stagePredecessorCollection"`

	// An compute instance group environment ocid for rolling deployment.
	ComputeInstanceGroupEnvironmentId *string `mandatory:"true" json:"computeInstanceGroupEnvironmentId"`

	// The OCID of the artifact that contains the deployment specification.
	PrimaryArtifactId *string `mandatory:"true" json:"primaryArtifactId"`

	RolloutPolicy *CreateComputeInstanceGroupRollingDeploymentRollOutPolicyDetails `mandatory:"true" json:"rolloutPolicy"`

	// A timeout in seconds, this stage may take to return.
	TimeoutInSeconds *int `mandatory:"false" json:"timeoutInSeconds"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The path to the deployment specification file in the primary artifact. By default, it should be deployment_spec.yaml
	DeploymentSpecFile *string `mandatory:"false" json:"deploymentSpecFile"`

	// Additional file artifact ocids.
	AdditionalArtifactIds []string `mandatory:"false" json:"additionalArtifactIds"`

	// An optional Load balancer ocid.
	LoadBalancerId *string `mandatory:"false" json:"loadBalancerId"`
}

//GetDisplayName returns DisplayName
func (m CreateComputeInstanceGroupRollingDeploymentStageDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetPipelineId returns PipelineId
func (m CreateComputeInstanceGroupRollingDeploymentStageDetails) GetPipelineId() *string {
	return m.PipelineId
}

//GetTimeoutInSeconds returns TimeoutInSeconds
func (m CreateComputeInstanceGroupRollingDeploymentStageDetails) GetTimeoutInSeconds() *int {
	return m.TimeoutInSeconds
}

//GetStagePredecessorCollection returns StagePredecessorCollection
func (m CreateComputeInstanceGroupRollingDeploymentStageDetails) GetStagePredecessorCollection() *StagePredecessorCollection {
	return m.StagePredecessorCollection
}

//GetFreeformTags returns FreeformTags
func (m CreateComputeInstanceGroupRollingDeploymentStageDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m CreateComputeInstanceGroupRollingDeploymentStageDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateComputeInstanceGroupRollingDeploymentStageDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateComputeInstanceGroupRollingDeploymentStageDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateComputeInstanceGroupRollingDeploymentStageDetails CreateComputeInstanceGroupRollingDeploymentStageDetails
	s := struct {
		DiscriminatorParam string `json:"stageType"`
		MarshalTypeCreateComputeInstanceGroupRollingDeploymentStageDetails
	}{
		"COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT",
		(MarshalTypeCreateComputeInstanceGroupRollingDeploymentStageDetails)(m),
	}

	return json.Marshal(&s)
}
