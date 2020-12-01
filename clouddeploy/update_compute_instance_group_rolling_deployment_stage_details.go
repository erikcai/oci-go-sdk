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

// UpdateComputeInstanceGroupRollingDeploymentStageDetails Specifies the Instance group rolling Deployment Stage.
type UpdateComputeInstanceGroupRollingDeploymentStageDetails struct {

	// Stage Identifier
	DisplayName *string `mandatory:"true" json:"displayName"`

	// A timeout in seconds, this stage may take to return.
	TimeoutInSeconds *int `mandatory:"false" json:"timeoutInSeconds"`

	StagePredecessorCollection *StagePredecessorCollection `mandatory:"false" json:"stagePredecessorCollection"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// An compute instance group environment ocid for rolling deployment.
	ComputeInstanceGroupEnvironmentId *string `mandatory:"false" json:"computeInstanceGroupEnvironmentId"`

	// The OCID of the artifact that contains the deployment specification.
	PrimaryArtifactId *string `mandatory:"false" json:"primaryArtifactId"`

	// The path to the deployment specification file in the primary artifact. By default, it should be deployment_spec.yaml
	DeploymentSpecFile *string `mandatory:"false" json:"deploymentSpecFile"`

	// Additional file artifact ocids.
	AdditionalArtifactIds []string `mandatory:"false" json:"additionalArtifactIds"`

	RolloutPolicy *UpdateComputeInstanceGroupRollingDeploymentRollOutPolicyDetails `mandatory:"false" json:"rolloutPolicy"`

	// An optional Load balancer ocid.
	LoadBalancerId *string `mandatory:"false" json:"loadBalancerId"`
}

//GetDisplayName returns DisplayName
func (m UpdateComputeInstanceGroupRollingDeploymentStageDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetTimeoutInSeconds returns TimeoutInSeconds
func (m UpdateComputeInstanceGroupRollingDeploymentStageDetails) GetTimeoutInSeconds() *int {
	return m.TimeoutInSeconds
}

//GetStagePredecessorCollection returns StagePredecessorCollection
func (m UpdateComputeInstanceGroupRollingDeploymentStageDetails) GetStagePredecessorCollection() *StagePredecessorCollection {
	return m.StagePredecessorCollection
}

//GetFreeformTags returns FreeformTags
func (m UpdateComputeInstanceGroupRollingDeploymentStageDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m UpdateComputeInstanceGroupRollingDeploymentStageDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m UpdateComputeInstanceGroupRollingDeploymentStageDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m UpdateComputeInstanceGroupRollingDeploymentStageDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateComputeInstanceGroupRollingDeploymentStageDetails UpdateComputeInstanceGroupRollingDeploymentStageDetails
	s := struct {
		DiscriminatorParam string `json:"stageType"`
		MarshalTypeUpdateComputeInstanceGroupRollingDeploymentStageDetails
	}{
		"COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT",
		(MarshalTypeUpdateComputeInstanceGroupRollingDeploymentStageDetails)(m),
	}

	return json.Marshal(&s)
}
