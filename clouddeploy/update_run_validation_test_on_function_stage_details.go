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

// UpdateRunValidationTestOnFunctionStageDetails Specifies Run validation test on Function to validate on going deployment.
type UpdateRunValidationTestOnFunctionStageDetails struct {

	// Stage Identifier
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Function environment ocid.
	FunctionEnvironmentId *string `mandatory:"true" json:"functionEnvironmentId"`

	// A timeout in seconds, this stage may take to return.
	TimeoutInSeconds *int `mandatory:"false" json:"timeoutInSeconds"`

	StagePredecessorCollection *StagePredecessorCollection `mandatory:"false" json:"stagePredecessorCollection"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Optional binary artifact user may provide to this stage.
	BinaryArtifactId *string `mandatory:"false" json:"binaryArtifactId"`
}

//GetDisplayName returns DisplayName
func (m UpdateRunValidationTestOnFunctionStageDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetTimeoutInSeconds returns TimeoutInSeconds
func (m UpdateRunValidationTestOnFunctionStageDetails) GetTimeoutInSeconds() *int {
	return m.TimeoutInSeconds
}

//GetStagePredecessorCollection returns StagePredecessorCollection
func (m UpdateRunValidationTestOnFunctionStageDetails) GetStagePredecessorCollection() *StagePredecessorCollection {
	return m.StagePredecessorCollection
}

//GetFreeformTags returns FreeformTags
func (m UpdateRunValidationTestOnFunctionStageDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m UpdateRunValidationTestOnFunctionStageDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m UpdateRunValidationTestOnFunctionStageDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m UpdateRunValidationTestOnFunctionStageDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateRunValidationTestOnFunctionStageDetails UpdateRunValidationTestOnFunctionStageDetails
	s := struct {
		DiscriminatorParam string `json:"stageType"`
		MarshalTypeUpdateRunValidationTestOnFunctionStageDetails
	}{
		"RUN_VALIDATION_TEST_ON_FUNCTION",
		(MarshalTypeUpdateRunValidationTestOnFunctionStageDetails)(m),
	}

	return json.Marshal(&s)
}
