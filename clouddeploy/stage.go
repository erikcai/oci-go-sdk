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

// Stage Description of Stage.
type Stage interface {

	// Unique identifier that is immutable on creation
	GetId() *string

	// Application Identifier
	GetApplicationId() *string

	// Pipeline Identifier
	GetPipelineId() *string

	// Compartment Identifier
	GetCompartmentId() *string

	// Stage Identifier, can be renamed
	GetDisplayName() *string

	// The time the the Stage was created. An RFC3339 formatted datetime string
	GetTimeCreated() *common.SDKTime

	// The time the Stage was updated. An RFC3339 formatted datetime string
	GetTimeUpdated() *common.SDKTime

	// The current state of the Stage.
	GetLifecycleState() StageLifecycleStateEnum

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	GetLifecyleDetails() *string

	// A timeout in seconds, this stage may take to return.
	GetTimeoutInSeconds() *int

	GetStagePredecessorCollection() *StagePredecessorCollection

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	GetSystemTags() map[string]map[string]interface{}
}

type stage struct {
	JsonData                   []byte
	Id                         *string                           `mandatory:"true" json:"id"`
	ApplicationId              *string                           `mandatory:"true" json:"applicationId"`
	PipelineId                 *string                           `mandatory:"true" json:"pipelineId"`
	CompartmentId              *string                           `mandatory:"true" json:"compartmentId"`
	DisplayName                *string                           `mandatory:"false" json:"displayName"`
	TimeCreated                *common.SDKTime                   `mandatory:"false" json:"timeCreated"`
	TimeUpdated                *common.SDKTime                   `mandatory:"false" json:"timeUpdated"`
	LifecycleState             StageLifecycleStateEnum           `mandatory:"false" json:"lifecycleState,omitempty"`
	LifecyleDetails            *string                           `mandatory:"false" json:"lifecyleDetails"`
	TimeoutInSeconds           *int                              `mandatory:"false" json:"timeoutInSeconds"`
	StagePredecessorCollection *StagePredecessorCollection       `mandatory:"false" json:"stagePredecessorCollection"`
	FreeformTags               map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags                map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	SystemTags                 map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	StageType                  string                            `json:"stageType"`
}

// UnmarshalJSON unmarshals json
func (m *stage) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerstage stage
	s := struct {
		Model Unmarshalerstage
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.ApplicationId = s.Model.ApplicationId
	m.PipelineId = s.Model.PipelineId
	m.CompartmentId = s.Model.CompartmentId
	m.DisplayName = s.Model.DisplayName
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
	m.LifecycleState = s.Model.LifecycleState
	m.LifecyleDetails = s.Model.LifecyleDetails
	m.TimeoutInSeconds = s.Model.TimeoutInSeconds
	m.StagePredecessorCollection = s.Model.StagePredecessorCollection
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.StageType = s.Model.StageType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *stage) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.StageType {
	case "RUN_VALIDATION_TEST_ON_OKE":
		mm := RunValidationTestOnOkeStage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "LOAD_BALANCER_TRAFFIC_SHIFT":
		mm := LoadBalancerTrafficShiftStage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "INVOKE_FUNCTION":
		mm := InvokeFunctionStage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RUN_VALIDATION_TEST_ON_COMPUTE_INSTANCE":
		mm := RunValidationTestOnComputeInstanceStage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RUN_OKE_JOB":
		mm := RunOkeJobStage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "UPDATE_FUNCTION_APPLICATION":
		mm := UpdateFunctionApplicationStage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RUN_JOB_ON_COMPUTE_INSTANCE_GROUP":
		mm := RunJobOnComputeInstanceGroupStage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OKE_DEPLOYMENT":
		mm := OkeDeploymentStage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT":
		mm := ComputeInstanceGroupRollingDeploymentStage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RUN_JOB_ON_COMPUTE_INSTANCE":
		mm := RunJobOnComputeInstanceStage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "WAIT":
		mm := WaitStage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RUN_DEPLOYMENT_PIPELINE":
		mm := RunDeploymentStage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RUN_VALIDATION_TEST_ON_FUNCTION":
		mm := RunValidationTestOnFunctionStage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MANUAL_APPROVAL":
		mm := ManualApprovalStage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DEPLOY_FUNCTION":
		mm := DeployFunctionStage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetId returns Id
func (m stage) GetId() *string {
	return m.Id
}

//GetApplicationId returns ApplicationId
func (m stage) GetApplicationId() *string {
	return m.ApplicationId
}

//GetPipelineId returns PipelineId
func (m stage) GetPipelineId() *string {
	return m.PipelineId
}

//GetCompartmentId returns CompartmentId
func (m stage) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetDisplayName returns DisplayName
func (m stage) GetDisplayName() *string {
	return m.DisplayName
}

//GetTimeCreated returns TimeCreated
func (m stage) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m stage) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetLifecycleState returns LifecycleState
func (m stage) GetLifecycleState() StageLifecycleStateEnum {
	return m.LifecycleState
}

//GetLifecyleDetails returns LifecyleDetails
func (m stage) GetLifecyleDetails() *string {
	return m.LifecyleDetails
}

//GetTimeoutInSeconds returns TimeoutInSeconds
func (m stage) GetTimeoutInSeconds() *int {
	return m.TimeoutInSeconds
}

//GetStagePredecessorCollection returns StagePredecessorCollection
func (m stage) GetStagePredecessorCollection() *StagePredecessorCollection {
	return m.StagePredecessorCollection
}

//GetFreeformTags returns FreeformTags
func (m stage) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m stage) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetSystemTags returns SystemTags
func (m stage) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m stage) String() string {
	return common.PointerString(m)
}

// StageLifecycleStateEnum Enum with underlying type: string
type StageLifecycleStateEnum string

// Set of constants representing the allowable values for StageLifecycleStateEnum
const (
	StageLifecycleStateCreating StageLifecycleStateEnum = "CREATING"
	StageLifecycleStateUpdating StageLifecycleStateEnum = "UPDATING"
	StageLifecycleStateActive   StageLifecycleStateEnum = "ACTIVE"
	StageLifecycleStateDeleting StageLifecycleStateEnum = "DELETING"
	StageLifecycleStateDeleted  StageLifecycleStateEnum = "DELETED"
	StageLifecycleStateFailed   StageLifecycleStateEnum = "FAILED"
)

var mappingStageLifecycleState = map[string]StageLifecycleStateEnum{
	"CREATING": StageLifecycleStateCreating,
	"UPDATING": StageLifecycleStateUpdating,
	"ACTIVE":   StageLifecycleStateActive,
	"DELETING": StageLifecycleStateDeleting,
	"DELETED":  StageLifecycleStateDeleted,
	"FAILED":   StageLifecycleStateFailed,
}

// GetStageLifecycleStateEnumValues Enumerates the set of values for StageLifecycleStateEnum
func GetStageLifecycleStateEnumValues() []StageLifecycleStateEnum {
	values := make([]StageLifecycleStateEnum, 0)
	for _, v := range mappingStageLifecycleState {
		values = append(values, v)
	}
	return values
}

// StageStageTypeEnum Enum with underlying type: string
type StageStageTypeEnum string

// Set of constants representing the allowable values for StageStageTypeEnum
const (
	StageStageTypeWait                                  StageStageTypeEnum = "WAIT"
	StageStageTypeComputeInstanceGroupRollingDeployment StageStageTypeEnum = "COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT"
	StageStageTypeOkeDeployment                         StageStageTypeEnum = "OKE_DEPLOYMENT"
	StageStageTypeDeployFunction                        StageStageTypeEnum = "DEPLOY_FUNCTION"
	StageStageTypeUpdateFunctionApplication             StageStageTypeEnum = "UPDATE_FUNCTION_APPLICATION"
	StageStageTypeRunJobOnComputeInstanceGroup          StageStageTypeEnum = "RUN_JOB_ON_COMPUTE_INSTANCE_GROUP"
	StageStageTypeRunJobOnComputeInstance               StageStageTypeEnum = "RUN_JOB_ON_COMPUTE_INSTANCE"
	StageStageTypeRunOkeJob                             StageStageTypeEnum = "RUN_OKE_JOB"
	StageStageTypeInvokeFunction                        StageStageTypeEnum = "INVOKE_FUNCTION"
	StageStageTypeRunValidationTestOnComputeInstance    StageStageTypeEnum = "RUN_VALIDATION_TEST_ON_COMPUTE_INSTANCE"
	StageStageTypeRunValidationTestOnOke                StageStageTypeEnum = "RUN_VALIDATION_TEST_ON_OKE"
	StageStageTypeRunValidationTestOnFunction           StageStageTypeEnum = "RUN_VALIDATION_TEST_ON_FUNCTION"
	StageStageTypeLoadBalancerTrafficShift              StageStageTypeEnum = "LOAD_BALANCER_TRAFFIC_SHIFT"
	StageStageTypeManualApproval                        StageStageTypeEnum = "MANUAL_APPROVAL"
	StageStageTypeRunDeploymentPipeline                 StageStageTypeEnum = "RUN_DEPLOYMENT_PIPELINE"
)

var mappingStageStageType = map[string]StageStageTypeEnum{
	"WAIT": StageStageTypeWait,
	"COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT": StageStageTypeComputeInstanceGroupRollingDeployment,
	"OKE_DEPLOYMENT":                          StageStageTypeOkeDeployment,
	"DEPLOY_FUNCTION":                         StageStageTypeDeployFunction,
	"UPDATE_FUNCTION_APPLICATION":             StageStageTypeUpdateFunctionApplication,
	"RUN_JOB_ON_COMPUTE_INSTANCE_GROUP":       StageStageTypeRunJobOnComputeInstanceGroup,
	"RUN_JOB_ON_COMPUTE_INSTANCE":             StageStageTypeRunJobOnComputeInstance,
	"RUN_OKE_JOB":                             StageStageTypeRunOkeJob,
	"INVOKE_FUNCTION":                         StageStageTypeInvokeFunction,
	"RUN_VALIDATION_TEST_ON_COMPUTE_INSTANCE": StageStageTypeRunValidationTestOnComputeInstance,
	"RUN_VALIDATION_TEST_ON_OKE":              StageStageTypeRunValidationTestOnOke,
	"RUN_VALIDATION_TEST_ON_FUNCTION":         StageStageTypeRunValidationTestOnFunction,
	"LOAD_BALANCER_TRAFFIC_SHIFT":             StageStageTypeLoadBalancerTrafficShift,
	"MANUAL_APPROVAL":                         StageStageTypeManualApproval,
	"RUN_DEPLOYMENT_PIPELINE":                 StageStageTypeRunDeploymentPipeline,
}

// GetStageStageTypeEnumValues Enumerates the set of values for StageStageTypeEnum
func GetStageStageTypeEnumValues() []StageStageTypeEnum {
	values := make([]StageStageTypeEnum, 0)
	for _, v := range mappingStageStageType {
		values = append(values, v)
	}
	return values
}
