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

// UpdateRunJobOnComputeInstanceGroupStageDetails Specifies Run job on the compute instance groups stage.
type UpdateRunJobOnComputeInstanceGroupStageDetails struct {

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

	// An instance group environment ocid.
	ComputeInstanceGroupEnvironmentId *string `mandatory:"false" json:"computeInstanceGroupEnvironmentId"`

	// The OCID of the artifact that contains the deployment specification.
	PrimaryArtifactId *string `mandatory:"false" json:"primaryArtifactId"`

	// File artifact ocids, there should be atleast one artifact containing deployment specifications.
	AdditionalArtifactIds []string `mandatory:"false" json:"additionalArtifactIds"`

	RolloutPolicy *UpdateComputeInstanceGroupRollingDeploymentRollOutPolicyDetails `mandatory:"false" json:"rolloutPolicy"`

	// The path to the deployment specification file in the primary artifact for this stage's Environment. The default location if not specified is deployment_spec.yaml
	DeploymentSpecFile *string `mandatory:"false" json:"deploymentSpecFile"`
}

//GetDisplayName returns DisplayName
func (m UpdateRunJobOnComputeInstanceGroupStageDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetTimeoutInSeconds returns TimeoutInSeconds
func (m UpdateRunJobOnComputeInstanceGroupStageDetails) GetTimeoutInSeconds() *int {
	return m.TimeoutInSeconds
}

//GetStagePredecessorCollection returns StagePredecessorCollection
func (m UpdateRunJobOnComputeInstanceGroupStageDetails) GetStagePredecessorCollection() *StagePredecessorCollection {
	return m.StagePredecessorCollection
}

//GetFreeformTags returns FreeformTags
func (m UpdateRunJobOnComputeInstanceGroupStageDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m UpdateRunJobOnComputeInstanceGroupStageDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m UpdateRunJobOnComputeInstanceGroupStageDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m UpdateRunJobOnComputeInstanceGroupStageDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateRunJobOnComputeInstanceGroupStageDetails UpdateRunJobOnComputeInstanceGroupStageDetails
	s := struct {
		DiscriminatorParam string `json:"stageType"`
		MarshalTypeUpdateRunJobOnComputeInstanceGroupStageDetails
	}{
		"RUN_JOB_ON_COMPUTE_INSTANCE_GROUP",
		(MarshalTypeUpdateRunJobOnComputeInstanceGroupStageDetails)(m),
	}

	return json.Marshal(&s)
}
