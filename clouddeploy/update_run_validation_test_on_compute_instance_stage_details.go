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

// UpdateRunValidationTestOnComputeInstanceStageDetails Specifies run Validation test on an compute instance stage.
type UpdateRunValidationTestOnComputeInstanceStageDetails struct {

	// Stage Identifier
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Instance OCID on which this stage runs validation.
	ComputeInstanceId *string `mandatory:"true" json:"computeInstanceId"`

	// File artifact ocids, there should be atleast one artifact containing deployment specifications.
	ArtifactIds []string `mandatory:"true" json:"artifactIds"`

	// A timeout in seconds, this stage may take to return.
	TimeoutInSeconds *int `mandatory:"false" json:"timeoutInSeconds"`

	StagePredecessorCollection *StagePredecessorCollection `mandatory:"false" json:"stagePredecessorCollection"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

//GetDisplayName returns DisplayName
func (m UpdateRunValidationTestOnComputeInstanceStageDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetTimeoutInSeconds returns TimeoutInSeconds
func (m UpdateRunValidationTestOnComputeInstanceStageDetails) GetTimeoutInSeconds() *int {
	return m.TimeoutInSeconds
}

//GetStagePredecessorCollection returns StagePredecessorCollection
func (m UpdateRunValidationTestOnComputeInstanceStageDetails) GetStagePredecessorCollection() *StagePredecessorCollection {
	return m.StagePredecessorCollection
}

//GetFreeformTags returns FreeformTags
func (m UpdateRunValidationTestOnComputeInstanceStageDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m UpdateRunValidationTestOnComputeInstanceStageDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m UpdateRunValidationTestOnComputeInstanceStageDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m UpdateRunValidationTestOnComputeInstanceStageDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateRunValidationTestOnComputeInstanceStageDetails UpdateRunValidationTestOnComputeInstanceStageDetails
	s := struct {
		DiscriminatorParam string `json:"stageType"`
		MarshalTypeUpdateRunValidationTestOnComputeInstanceStageDetails
	}{
		"RUN_VALIDATION_TEST_ON_COMPUTE_INSTANCE",
		(MarshalTypeUpdateRunValidationTestOnComputeInstanceStageDetails)(m),
	}

	return json.Marshal(&s)
}
