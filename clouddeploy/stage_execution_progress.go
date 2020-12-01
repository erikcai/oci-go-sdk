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

// StageExecutionProgress The details about the execution progress of a Stage in a Deployment.
type StageExecutionProgress interface {

	// The time the Stage was started executing. An RFC3339 formatted datetime string
	GetTimeStarted() *common.SDKTime

	// The time the Stage was finished executing. An RFC3339 formatted datetime string
	GetTimeFinished() *common.SDKTime

	// The current state of the Stage.
	GetStatus() StageExecutionProgressStatusEnum

	// A message describing the current state of the stage in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	GetStatusMessages() []string
}

type stageexecutionprogress struct {
	JsonData       []byte
	TimeStarted    *common.SDKTime                  `mandatory:"false" json:"timeStarted"`
	TimeFinished   *common.SDKTime                  `mandatory:"false" json:"timeFinished"`
	Status         StageExecutionProgressStatusEnum `mandatory:"false" json:"status,omitempty"`
	StatusMessages []string                         `mandatory:"false" json:"statusMessages"`
	StageType      string                           `json:"stageType"`
}

// UnmarshalJSON unmarshals json
func (m *stageexecutionprogress) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerstageexecutionprogress stageexecutionprogress
	s := struct {
		Model Unmarshalerstageexecutionprogress
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.TimeStarted = s.Model.TimeStarted
	m.TimeFinished = s.Model.TimeFinished
	m.Status = s.Model.Status
	m.StatusMessages = s.Model.StatusMessages
	m.StageType = s.Model.StageType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *stageexecutionprogress) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.StageType {
	case "COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT":
		mm := ComputeInstanceGroupRollingDeploymentStageExecutionProgress{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RUN_OKE_JOB":
		mm := RunOkeJobStageExecutionProgress{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OKE_DEPLOYMENT":
		mm := OkeDeploymentStageExecutionProgress{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RUN_VALIDATION_TEST_ON_COMPUTE_INSTANCE":
		mm := RunValidationTestOnComputeInstanceStageExecutionProgress{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "WAIT":
		mm := WaitStageExecutionProgress{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RUN_DEPLOYMENT_PIPELINE":
		mm := RunDeploymentPipelineStageExecutionProgress{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "INVOKE_FUNCTION":
		mm := InvokeFunctionStageExecutionProgress{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DEPLOY_FUNCTION":
		mm := DeployFunctionStageExecutionProgress{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RUN_JOB_ON_COMPUTE_INSTANCE_GROUP":
		mm := RunJobOnComputeInstanceGroupStageExecutionProgress{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "UPDATE_FUNCTION_APPLICATION":
		mm := UpdateFunctionApplicationStageExecutionProgress{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RUN_VALIDATION_TEST_ON_FUNCTION":
		mm := RunValidationTestOnFunctionStageExecutionProgress{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RUN_VALIDATION_TEST_ON_OKE":
		mm := RunValidationTestOnOkeStageExecutionProgress{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "LOAD_BALANCER_TRAFFIC_SHIFT":
		mm := LoadBalancerTrafficShiftStageExecutionProgress{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RUN_JOB_ON_COMPUTE_INSTANCE":
		mm := RunJobOnComputeInstanceStageExecutionProgress{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MANUAL_APPROVAL":
		mm := ManualApprovalStageExecutionProgress{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetTimeStarted returns TimeStarted
func (m stageexecutionprogress) GetTimeStarted() *common.SDKTime {
	return m.TimeStarted
}

//GetTimeFinished returns TimeFinished
func (m stageexecutionprogress) GetTimeFinished() *common.SDKTime {
	return m.TimeFinished
}

//GetStatus returns Status
func (m stageexecutionprogress) GetStatus() StageExecutionProgressStatusEnum {
	return m.Status
}

//GetStatusMessages returns StatusMessages
func (m stageexecutionprogress) GetStatusMessages() []string {
	return m.StatusMessages
}

func (m stageexecutionprogress) String() string {
	return common.PointerString(m)
}

// StageExecutionProgressStatusEnum Enum with underlying type: string
type StageExecutionProgressStatusEnum string

// Set of constants representing the allowable values for StageExecutionProgressStatusEnum
const (
	StageExecutionProgressStatusAccepted   StageExecutionProgressStatusEnum = "ACCEPTED"
	StageExecutionProgressStatusInProgress StageExecutionProgressStatusEnum = "IN_PROGRESS"
	StageExecutionProgressStatusFailed     StageExecutionProgressStatusEnum = "FAILED"
	StageExecutionProgressStatusSucceeded  StageExecutionProgressStatusEnum = "SUCCEEDED"
	StageExecutionProgressStatusCanceling  StageExecutionProgressStatusEnum = "CANCELING"
	StageExecutionProgressStatusCanceled   StageExecutionProgressStatusEnum = "CANCELED"
)

var mappingStageExecutionProgressStatus = map[string]StageExecutionProgressStatusEnum{
	"ACCEPTED":    StageExecutionProgressStatusAccepted,
	"IN_PROGRESS": StageExecutionProgressStatusInProgress,
	"FAILED":      StageExecutionProgressStatusFailed,
	"SUCCEEDED":   StageExecutionProgressStatusSucceeded,
	"CANCELING":   StageExecutionProgressStatusCanceling,
	"CANCELED":    StageExecutionProgressStatusCanceled,
}

// GetStageExecutionProgressStatusEnumValues Enumerates the set of values for StageExecutionProgressStatusEnum
func GetStageExecutionProgressStatusEnumValues() []StageExecutionProgressStatusEnum {
	values := make([]StageExecutionProgressStatusEnum, 0)
	for _, v := range mappingStageExecutionProgressStatus {
		values = append(values, v)
	}
	return values
}
