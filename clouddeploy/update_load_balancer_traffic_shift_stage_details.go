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

// UpdateLoadBalancerTrafficShiftStageDetails Specifies Load Balncer traffic shift stage.
type UpdateLoadBalancerTrafficShiftStageDetails struct {

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

	// Existing Load Balance Listener environment ocid.
	// Listener of this environment should have Backends or BackendSet set up.
	LoadBalancerListenerEnvironmentId *string `mandatory:"false" json:"loadBalancerListenerEnvironmentId"`

	BlueBackendIps *BackendSetIpCollection `mandatory:"false" json:"blueBackendIps"`

	GreenBackendIps *BackendSetIpCollection `mandatory:"false" json:"greenBackendIps"`

	RolloutPolicy *UpdateLoadBalancerTrafficShiftRollOutPolicyDetails `mandatory:"false" json:"rolloutPolicy"`

	// Specifies the target or destination Backend set.
	TrafficShiftTarget LoadBalancerTrafficShiftStageTrafficShiftTargetEnum `mandatory:"false" json:"trafficShiftTarget,omitempty"`
}

//GetDisplayName returns DisplayName
func (m UpdateLoadBalancerTrafficShiftStageDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetTimeoutInSeconds returns TimeoutInSeconds
func (m UpdateLoadBalancerTrafficShiftStageDetails) GetTimeoutInSeconds() *int {
	return m.TimeoutInSeconds
}

//GetStagePredecessorCollection returns StagePredecessorCollection
func (m UpdateLoadBalancerTrafficShiftStageDetails) GetStagePredecessorCollection() *StagePredecessorCollection {
	return m.StagePredecessorCollection
}

//GetFreeformTags returns FreeformTags
func (m UpdateLoadBalancerTrafficShiftStageDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m UpdateLoadBalancerTrafficShiftStageDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m UpdateLoadBalancerTrafficShiftStageDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m UpdateLoadBalancerTrafficShiftStageDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateLoadBalancerTrafficShiftStageDetails UpdateLoadBalancerTrafficShiftStageDetails
	s := struct {
		DiscriminatorParam string `json:"stageType"`
		MarshalTypeUpdateLoadBalancerTrafficShiftStageDetails
	}{
		"LOAD_BALANCER_TRAFFIC_SHIFT",
		(MarshalTypeUpdateLoadBalancerTrafficShiftStageDetails)(m),
	}

	return json.Marshal(&s)
}
