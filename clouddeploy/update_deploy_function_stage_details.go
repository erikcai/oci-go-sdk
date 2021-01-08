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

// UpdateDeployFunctionStageDetails Specifies the DeployFunction Stage.
type UpdateDeployFunctionStageDetails struct {

	// Stage Identifier
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Function environment ocid.
	FunctionEnvironmentId *string `mandatory:"true" json:"functionEnvironmentId"`

	// A docker image artifact ocid.
	DockerImageArtifactId *string `mandatory:"true" json:"dockerImageArtifactId"`

	// A timeout in seconds, this stage may take to return.
	TimeoutInSeconds *int `mandatory:"false" json:"timeoutInSeconds"`

	StagePredecessorCollection *StagePredecessorCollection `mandatory:"false" json:"stagePredecessorCollection"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// User provided key/value pair configuration which assigned through constants or parameter.
	Config map[string]string `mandatory:"false" json:"config"`

	// Maximum usable memory for the function (MiB)
	MaxMemoryInMBs *int64 `mandatory:"false" json:"maxMemoryInMBs"`

	// Timeout for executions of the function. Value in seconds.
	FunctionTimeoutInSeconds *int `mandatory:"false" json:"functionTimeoutInSeconds"`
}

//GetDisplayName returns DisplayName
func (m UpdateDeployFunctionStageDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetTimeoutInSeconds returns TimeoutInSeconds
func (m UpdateDeployFunctionStageDetails) GetTimeoutInSeconds() *int {
	return m.TimeoutInSeconds
}

//GetStagePredecessorCollection returns StagePredecessorCollection
func (m UpdateDeployFunctionStageDetails) GetStagePredecessorCollection() *StagePredecessorCollection {
	return m.StagePredecessorCollection
}

//GetFreeformTags returns FreeformTags
func (m UpdateDeployFunctionStageDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m UpdateDeployFunctionStageDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m UpdateDeployFunctionStageDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m UpdateDeployFunctionStageDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateDeployFunctionStageDetails UpdateDeployFunctionStageDetails
	s := struct {
		DiscriminatorParam string `json:"stageType"`
		MarshalTypeUpdateDeployFunctionStageDetails
	}{
		"DEPLOY_FUNCTION",
		(MarshalTypeUpdateDeployFunctionStageDetails)(m),
	}

	return json.Marshal(&s)
}
