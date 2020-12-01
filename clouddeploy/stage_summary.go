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

// StageSummary Summary of the Stage.
type StageSummary interface {

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
	GetLifecycleState() StageSummary

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	GetLifecycleDetails() *string

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

type stagesummary struct {
	JsonData                   []byte
	Id                         *string                           `mandatory:"true" json:"id"`
	ApplicationId              *string                           `mandatory:"true" json:"applicationId"`
	PipelineId                 *string                           `mandatory:"true" json:"pipelineId"`
	CompartmentId              *string                           `mandatory:"true" json:"compartmentId"`
	DisplayName                *string                           `mandatory:"false" json:"displayName"`
	TimeCreated                *common.SDKTime                   `mandatory:"false" json:"timeCreated"`
	TimeUpdated                *common.SDKTime                   `mandatory:"false" json:"timeUpdated"`
	LifecycleState             StageSummary                      `mandatory:"false" json:"lifecycleState,omitempty"`
	LifecycleDetails           *string                           `mandatory:"false" json:"lifecycleDetails"`
	TimeoutInSeconds           *int                              `mandatory:"false" json:"timeoutInSeconds"`
	StagePredecessorCollection *StagePredecessorCollection       `mandatory:"false" json:"stagePredecessorCollection"`
	FreeformTags               map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags                map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	SystemTags                 map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	StageType                  string                            `json:"stageType"`
}

// UnmarshalJSON unmarshals json
func (m *stagesummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerstagesummary stagesummary
	s := struct {
		Model Unmarshalerstagesummary
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
	m.LifecycleDetails = s.Model.LifecycleDetails
	m.TimeoutInSeconds = s.Model.TimeoutInSeconds
	m.StagePredecessorCollection = s.Model.StagePredecessorCollection
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.StageType = s.Model.StageType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *stagesummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.StageType {
	case "DEPLOY_FUNCTION":
		mm := DeployFunctionStageSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RUN_VALIDATION_TEST_ON_FUNCTION":
		mm := RunValidationTestOnFunctionStageSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RUN_DEPLOYMENT_PIPELINE":
		mm := RunDeploymentStageSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "WAIT":
		mm := WaitStageSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RUN_JOB_ON_COMPUTE_INSTANCE_GROUP":
		mm := RunJobOnComputeInstanceGroupStageSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OKE_DEPLOYMENT":
		mm := OkeDeploymentStageSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RUN_OKE_JOB":
		mm := RunOkeJobStageSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RUN_JOB_ON_COMPUTE_INSTANCE":
		mm := RunJobOnComputeInstanceStageSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT":
		mm := ComputeInstanceGroupRollingDeploymentStageSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "UPDATE_FUNCTION_APPLICATION":
		mm := UpdateFunctionApplicationStageSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "INVOKE_FUNCTION":
		mm := InvokeFunctionStageSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RUN_VALIDATION_TEST_ON_COMPUTE_INSTANCE":
		mm := RunValidationTestOnComputeInstanceStageSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MANUAL_APPROVAL":
		mm := ManualApprovalStageSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RUN_VALIDATION_TEST_ON_OKE":
		mm := RunValidationTestOnOkeStageSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "LOAD_BALANCER_TRAFFIC_SHIFT":
		mm := LoadBalancerTrafficShiftStageSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetId returns Id
func (m stagesummary) GetId() *string {
	return m.Id
}

//GetApplicationId returns ApplicationId
func (m stagesummary) GetApplicationId() *string {
	return m.ApplicationId
}

//GetPipelineId returns PipelineId
func (m stagesummary) GetPipelineId() *string {
	return m.PipelineId
}

//GetCompartmentId returns CompartmentId
func (m stagesummary) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetDisplayName returns DisplayName
func (m stagesummary) GetDisplayName() *string {
	return m.DisplayName
}

//GetTimeCreated returns TimeCreated
func (m stagesummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m stagesummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetLifecycleState returns LifecycleState
func (m stagesummary) GetLifecycleState() StageSummary {
	return m.LifecycleState
}

//GetLifecycleDetails returns LifecycleDetails
func (m stagesummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

//GetTimeoutInSeconds returns TimeoutInSeconds
func (m stagesummary) GetTimeoutInSeconds() *int {
	return m.TimeoutInSeconds
}

//GetStagePredecessorCollection returns StagePredecessorCollection
func (m stagesummary) GetStagePredecessorCollection() *StagePredecessorCollection {
	return m.StagePredecessorCollection
}

//GetFreeformTags returns FreeformTags
func (m stagesummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m stagesummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetSystemTags returns SystemTags
func (m stagesummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m stagesummary) String() string {
	return common.PointerString(m)
}
