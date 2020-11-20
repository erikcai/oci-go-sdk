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
	"github.com/oracle/oci-go-sdk/v29/common"
)

// UpdateRunDeploymentStageDetails Specifies Run Deployment Pipleline stage which runs another pipeline of the application.
type UpdateRunDeploymentStageDetails struct {

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

	// A target Pipeline ocid that will be run in this stage.
	TargetPipelineId *string `mandatory:"false" json:"targetPipelineId"`

	// A boolean flag specifies whether the target deployment is run synchronously or asynchronously.
	IsAsync *bool `mandatory:"false" json:"isAsync"`

	UpdateRunDeploymentConcurrencyConflictPolicyDetails *UpdateRunDeploymentConcurrencyConflictPolicyDetails `mandatory:"false" json:"updateRunDeploymentConcurrencyConflictPolicyDetails"`
}

//GetDisplayName returns DisplayName
func (m UpdateRunDeploymentStageDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetTimeoutInSeconds returns TimeoutInSeconds
func (m UpdateRunDeploymentStageDetails) GetTimeoutInSeconds() *int {
	return m.TimeoutInSeconds
}

//GetStagePredecessorCollection returns StagePredecessorCollection
func (m UpdateRunDeploymentStageDetails) GetStagePredecessorCollection() *StagePredecessorCollection {
	return m.StagePredecessorCollection
}

//GetFreeformTags returns FreeformTags
func (m UpdateRunDeploymentStageDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m UpdateRunDeploymentStageDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m UpdateRunDeploymentStageDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m UpdateRunDeploymentStageDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateRunDeploymentStageDetails UpdateRunDeploymentStageDetails
	s := struct {
		DiscriminatorParam string `json:"stageType"`
		MarshalTypeUpdateRunDeploymentStageDetails
	}{
		"RUN_DEPLOYMENT_PIPELINE",
		(MarshalTypeUpdateRunDeploymentStageDetails)(m),
	}

	return json.Marshal(&s)
}
