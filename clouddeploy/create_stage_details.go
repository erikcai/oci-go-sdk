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

// CreateStageDetails The information about new Stage.
type CreateStageDetails interface {

	// Stage Identifier
	GetDisplayName() *string

	// Pipeline Identifier
	GetPipelineId() *string

	GetStagePredecessorCollection() *StagePredecessorCollection

	// A timeout in seconds, this stage may take to return.
	GetTimeoutInSeconds() *int

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}
}

type createstagedetails struct {
	JsonData                   []byte
	DisplayName                *string                           `mandatory:"true" json:"displayName"`
	PipelineId                 *string                           `mandatory:"true" json:"pipelineId"`
	StagePredecessorCollection *StagePredecessorCollection       `mandatory:"true" json:"stagePredecessorCollection"`
	TimeoutInSeconds           *int                              `mandatory:"false" json:"timeoutInSeconds"`
	FreeformTags               map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags                map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	StageType                  string                            `json:"stageType"`
}

// UnmarshalJSON unmarshals json
func (m *createstagedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatestagedetails createstagedetails
	s := struct {
		Model Unmarshalercreatestagedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DisplayName = s.Model.DisplayName
	m.PipelineId = s.Model.PipelineId
	m.StagePredecessorCollection = s.Model.StagePredecessorCollection
	m.TimeoutInSeconds = s.Model.TimeoutInSeconds
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.StageType = s.Model.StageType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createstagedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.StageType {
	case "RUN_VALIDATION_TEST_ON_FUNCTION":
		mm := CreateRunValidationTestOnFunctionStageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DEPLOY_FUNCTION":
		mm := CreateDeployFunctionStageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RUN_VALIDATION_TEST_ON_COMPUTE_INSTANCE":
		mm := CreateRunValidationTestOnComputeInstanceStageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OKE_DEPLOYMENT":
		mm := CreateOkeDeploymentStageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MANUAL_APPROVAL":
		mm := CreateManualApprovalStageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RUN_DEPLOYMENT_PIPELINE":
		mm := CreateRunDeploymentStageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "UPDATE_FUNCTION_APPLICATION":
		mm := CreateUpdateFunctionApplicationStageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "LOAD_BALANCER_TRAFFIC_SHIFT":
		mm := CreateLoadBalancerTrafficShiftStageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "INVOKE_FUNCTION":
		mm := CreateInvokeFunctionStageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RUN_JOB_ON_COMPUTE_INSTANCE":
		mm := CreateRunJobOnComputeInstanceStageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RUN_JOB_ON_COMPUTE_INSTANCE_GROUP":
		mm := CreateRunJobOnComputeInstanceGroupStageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "WAIT":
		mm := CreateWaitStageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT":
		mm := CreateComputeInstanceGroupRollingDeploymentStageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RUN_VALIDATION_TEST_ON_OKE":
		mm := CreateRunValidationTestOnOkeStageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RUN_OKE_JOB":
		mm := CreateRunOkeJobStageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetDisplayName returns DisplayName
func (m createstagedetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetPipelineId returns PipelineId
func (m createstagedetails) GetPipelineId() *string {
	return m.PipelineId
}

//GetStagePredecessorCollection returns StagePredecessorCollection
func (m createstagedetails) GetStagePredecessorCollection() *StagePredecessorCollection {
	return m.StagePredecessorCollection
}

//GetTimeoutInSeconds returns TimeoutInSeconds
func (m createstagedetails) GetTimeoutInSeconds() *int {
	return m.TimeoutInSeconds
}

//GetFreeformTags returns FreeformTags
func (m createstagedetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m createstagedetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m createstagedetails) String() string {
	return common.PointerString(m)
}
