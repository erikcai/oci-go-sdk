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
	"github.com/oracle/oci-go-sdk/v32/common"
)

// UpdateStageDetails The information to be updated.
type UpdateStageDetails interface {

	// Stage Identifier
	GetDisplayName() *string

	// A timeout in seconds, this stage may take to return.
	GetTimeoutInSeconds() *int

	GetStagePredecessorCollection() *StagePredecessorCollection

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}
}

type updatestagedetails struct {
	JsonData                   []byte
	DisplayName                *string                           `mandatory:"true" json:"displayName"`
	TimeoutInSeconds           *int                              `mandatory:"false" json:"timeoutInSeconds"`
	StagePredecessorCollection *StagePredecessorCollection       `mandatory:"false" json:"stagePredecessorCollection"`
	FreeformTags               map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags                map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	StageType                  string                            `json:"stageType"`
}

// UnmarshalJSON unmarshals json
func (m *updatestagedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdatestagedetails updatestagedetails
	s := struct {
		Model Unmarshalerupdatestagedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DisplayName = s.Model.DisplayName
	m.TimeoutInSeconds = s.Model.TimeoutInSeconds
	m.StagePredecessorCollection = s.Model.StagePredecessorCollection
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.StageType = s.Model.StageType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updatestagedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.StageType {
	case "WAIT":
		mm := UpdateWaitStageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "LOAD_BALANCER_TRAFFIC_SHIFT":
		mm := UpdateLoadBalancerTrafficShiftStageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RUN_JOB_ON_COMPUTE_INSTANCE_GROUP":
		mm := UpdateRunJobOnComputeInstanceGroupStageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RUN_VALIDATION_TEST_ON_COMPUTE_INSTANCE":
		mm := UpdateRunValidationTestOnComputeInstanceStageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "UPDATE_FUNCTION_APPLICATION":
		mm := UpdateFunctionApplicationStageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RUN_VALIDATION_TEST_ON_FUNCTION":
		mm := UpdateRunValidationTestOnFunctionStageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RUN_JOB_ON_COMPUTE_INSTANCE":
		mm := UpdateRunJobOnComputeInstanceStageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RUN_DEPLOYMENT_PIPELINE":
		mm := UpdateRunDeploymentStageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OKE_DEPLOYMENT":
		mm := UpdateOkeDeploymentStageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RUN_OKE_JOB":
		mm := UpdateRunOkeJobStageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MANUAL_APPROVAL":
		mm := UpdateManualApprovalStageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT":
		mm := UpdateComputeInstanceGroupRollingDeploymentStageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DEPLOY_FUNCTION":
		mm := UpdateDeployFunctionStageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RUN_VALIDATION_TEST_ON_OKE":
		mm := UpdateRunValidationTestOnOkeStageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "INVOKE_FUNCTION":
		mm := UpdateInvokeFunctionStageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetDisplayName returns DisplayName
func (m updatestagedetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetTimeoutInSeconds returns TimeoutInSeconds
func (m updatestagedetails) GetTimeoutInSeconds() *int {
	return m.TimeoutInSeconds
}

//GetStagePredecessorCollection returns StagePredecessorCollection
func (m updatestagedetails) GetStagePredecessorCollection() *StagePredecessorCollection {
	return m.StagePredecessorCollection
}

//GetFreeformTags returns FreeformTags
func (m updatestagedetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m updatestagedetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m updatestagedetails) String() string {
	return common.PointerString(m)
}
