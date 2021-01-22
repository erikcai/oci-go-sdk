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

// LoadBalancerTrafficShiftStage Specifies Load Balncer traffic shift stage.
type LoadBalancerTrafficShiftStage struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// Application Identifier
	ApplicationId *string `mandatory:"true" json:"applicationId"`

	// Pipeline Identifier
	PipelineId *string `mandatory:"true" json:"pipelineId"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Existing Load Balance Listener environment ocid.
	// Listener of this environment should have Backends or BackendSet set up.
	LoadBalancerListenerEnvironmentId *string `mandatory:"true" json:"loadBalancerListenerEnvironmentId"`

	BlueBackendIps *BackendSetIpCollection `mandatory:"true" json:"blueBackendIps"`

	GreenBackendIps *BackendSetIpCollection `mandatory:"true" json:"greenBackendIps"`

	RolloutPolicy *LoadBalancerTrafficShiftRollOutPolicy `mandatory:"true" json:"rolloutPolicy"`

	// Stage Identifier, can be renamed
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The time the the Stage was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the Stage was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecyleDetails *string `mandatory:"false" json:"lifecyleDetails"`

	// A timeout in seconds, this stage may take to return.
	TimeoutInSeconds *int `mandatory:"false" json:"timeoutInSeconds"`

	StagePredecessorCollection *StagePredecessorCollection `mandatory:"false" json:"stagePredecessorCollection"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// Specifies the target or destination Backend set.
	TrafficShiftTarget LoadBalancerTrafficShiftStageTrafficShiftTargetEnum `mandatory:"true" json:"trafficShiftTarget"`

	// The current state of the Stage.
	LifecycleState StageLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

//GetId returns Id
func (m LoadBalancerTrafficShiftStage) GetId() *string {
	return m.Id
}

//GetDisplayName returns DisplayName
func (m LoadBalancerTrafficShiftStage) GetDisplayName() *string {
	return m.DisplayName
}

//GetApplicationId returns ApplicationId
func (m LoadBalancerTrafficShiftStage) GetApplicationId() *string {
	return m.ApplicationId
}

//GetPipelineId returns PipelineId
func (m LoadBalancerTrafficShiftStage) GetPipelineId() *string {
	return m.PipelineId
}

//GetCompartmentId returns CompartmentId
func (m LoadBalancerTrafficShiftStage) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetTimeCreated returns TimeCreated
func (m LoadBalancerTrafficShiftStage) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m LoadBalancerTrafficShiftStage) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetLifecycleState returns LifecycleState
func (m LoadBalancerTrafficShiftStage) GetLifecycleState() StageLifecycleStateEnum {
	return m.LifecycleState
}

//GetLifecyleDetails returns LifecyleDetails
func (m LoadBalancerTrafficShiftStage) GetLifecyleDetails() *string {
	return m.LifecyleDetails
}

//GetTimeoutInSeconds returns TimeoutInSeconds
func (m LoadBalancerTrafficShiftStage) GetTimeoutInSeconds() *int {
	return m.TimeoutInSeconds
}

//GetStagePredecessorCollection returns StagePredecessorCollection
func (m LoadBalancerTrafficShiftStage) GetStagePredecessorCollection() *StagePredecessorCollection {
	return m.StagePredecessorCollection
}

//GetFreeformTags returns FreeformTags
func (m LoadBalancerTrafficShiftStage) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m LoadBalancerTrafficShiftStage) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetSystemTags returns SystemTags
func (m LoadBalancerTrafficShiftStage) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m LoadBalancerTrafficShiftStage) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m LoadBalancerTrafficShiftStage) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeLoadBalancerTrafficShiftStage LoadBalancerTrafficShiftStage
	s := struct {
		DiscriminatorParam string `json:"stageType"`
		MarshalTypeLoadBalancerTrafficShiftStage
	}{
		"LOAD_BALANCER_TRAFFIC_SHIFT",
		(MarshalTypeLoadBalancerTrafficShiftStage)(m),
	}

	return json.Marshal(&s)
}

// LoadBalancerTrafficShiftStageTrafficShiftTargetEnum Enum with underlying type: string
type LoadBalancerTrafficShiftStageTrafficShiftTargetEnum string

// Set of constants representing the allowable values for LoadBalancerTrafficShiftStageTrafficShiftTargetEnum
const (
	LoadBalancerTrafficShiftStageTrafficShiftTargetBlue  LoadBalancerTrafficShiftStageTrafficShiftTargetEnum = "BLUE"
	LoadBalancerTrafficShiftStageTrafficShiftTargetGreen LoadBalancerTrafficShiftStageTrafficShiftTargetEnum = "GREEN"
)

var mappingLoadBalancerTrafficShiftStageTrafficShiftTarget = map[string]LoadBalancerTrafficShiftStageTrafficShiftTargetEnum{
	"BLUE":  LoadBalancerTrafficShiftStageTrafficShiftTargetBlue,
	"GREEN": LoadBalancerTrafficShiftStageTrafficShiftTargetGreen,
}

// GetLoadBalancerTrafficShiftStageTrafficShiftTargetEnumValues Enumerates the set of values for LoadBalancerTrafficShiftStageTrafficShiftTargetEnum
func GetLoadBalancerTrafficShiftStageTrafficShiftTargetEnumValues() []LoadBalancerTrafficShiftStageTrafficShiftTargetEnum {
	values := make([]LoadBalancerTrafficShiftStageTrafficShiftTargetEnum, 0)
	for _, v := range mappingLoadBalancerTrafficShiftStageTrafficShiftTarget {
		values = append(values, v)
	}
	return values
}
