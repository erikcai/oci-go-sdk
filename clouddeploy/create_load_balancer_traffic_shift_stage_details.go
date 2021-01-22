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

// CreateLoadBalancerTrafficShiftStageDetails Specifies Load Balncer traffic shift stage.
type CreateLoadBalancerTrafficShiftStageDetails struct {

	// Stage Identifier
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Pipeline Identifier
	PipelineId *string `mandatory:"true" json:"pipelineId"`

	StagePredecessorCollection *StagePredecessorCollection `mandatory:"true" json:"stagePredecessorCollection"`

	// Existing Load Balance Listener environment ocid.
	// Listener of this environment should have Backends or BackendSet set up.
	LoadBalancerListenerEnvironmentId *string `mandatory:"true" json:"loadBalancerListenerEnvironmentId"`

	BlueBackendIps *BackendSetIpCollection `mandatory:"true" json:"blueBackendIps"`

	GreenBackendIps *BackendSetIpCollection `mandatory:"true" json:"greenBackendIps"`

	RolloutPolicy *CreateLoadBalancerTrafficShiftRollOutPolicyDetails `mandatory:"true" json:"rolloutPolicy"`

	// A timeout in seconds, this stage may take to return.
	TimeoutInSeconds *int `mandatory:"false" json:"timeoutInSeconds"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Specifies the target or destination Backend set.
	TrafficShiftTarget LoadBalancerTrafficShiftStageTrafficShiftTargetEnum `mandatory:"true" json:"trafficShiftTarget"`
}

//GetDisplayName returns DisplayName
func (m CreateLoadBalancerTrafficShiftStageDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetPipelineId returns PipelineId
func (m CreateLoadBalancerTrafficShiftStageDetails) GetPipelineId() *string {
	return m.PipelineId
}

//GetTimeoutInSeconds returns TimeoutInSeconds
func (m CreateLoadBalancerTrafficShiftStageDetails) GetTimeoutInSeconds() *int {
	return m.TimeoutInSeconds
}

//GetStagePredecessorCollection returns StagePredecessorCollection
func (m CreateLoadBalancerTrafficShiftStageDetails) GetStagePredecessorCollection() *StagePredecessorCollection {
	return m.StagePredecessorCollection
}

//GetFreeformTags returns FreeformTags
func (m CreateLoadBalancerTrafficShiftStageDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m CreateLoadBalancerTrafficShiftStageDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateLoadBalancerTrafficShiftStageDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateLoadBalancerTrafficShiftStageDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateLoadBalancerTrafficShiftStageDetails CreateLoadBalancerTrafficShiftStageDetails
	s := struct {
		DiscriminatorParam string `json:"stageType"`
		MarshalTypeCreateLoadBalancerTrafficShiftStageDetails
	}{
		"LOAD_BALANCER_TRAFFIC_SHIFT",
		(MarshalTypeCreateLoadBalancerTrafficShiftStageDetails)(m),
	}

	return json.Marshal(&s)
}
