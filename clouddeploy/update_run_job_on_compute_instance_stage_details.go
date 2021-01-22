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

// UpdateRunJobOnComputeInstanceStageDetails Specifies Run job on an instance stage.
type UpdateRunJobOnComputeInstanceStageDetails struct {

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

	// Instance OCID on which this stage runs the job.
	ComputeInstanceId *string `mandatory:"false" json:"computeInstanceId"`

	// File artifact ocids.
	AdditionalArtifactIds []string `mandatory:"false" json:"additionalArtifactIds"`

	// The OCID of the artifact that contains the deployment specification.
	PrimaryArtifactId *string `mandatory:"false" json:"primaryArtifactId"`

	// The path to the deployment specification file in the primary artifact for this stage's Environment. The default location if not specified is deployment_spec.yaml
	DeploymentSpecFile *string `mandatory:"false" json:"deploymentSpecFile"`
}

//GetDisplayName returns DisplayName
func (m UpdateRunJobOnComputeInstanceStageDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetTimeoutInSeconds returns TimeoutInSeconds
func (m UpdateRunJobOnComputeInstanceStageDetails) GetTimeoutInSeconds() *int {
	return m.TimeoutInSeconds
}

//GetStagePredecessorCollection returns StagePredecessorCollection
func (m UpdateRunJobOnComputeInstanceStageDetails) GetStagePredecessorCollection() *StagePredecessorCollection {
	return m.StagePredecessorCollection
}

//GetFreeformTags returns FreeformTags
func (m UpdateRunJobOnComputeInstanceStageDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m UpdateRunJobOnComputeInstanceStageDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m UpdateRunJobOnComputeInstanceStageDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m UpdateRunJobOnComputeInstanceStageDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateRunJobOnComputeInstanceStageDetails UpdateRunJobOnComputeInstanceStageDetails
	s := struct {
		DiscriminatorParam string `json:"stageType"`
		MarshalTypeUpdateRunJobOnComputeInstanceStageDetails
	}{
		"RUN_JOB_ON_COMPUTE_INSTANCE",
		(MarshalTypeUpdateRunJobOnComputeInstanceStageDetails)(m),
	}

	return json.Marshal(&s)
}
