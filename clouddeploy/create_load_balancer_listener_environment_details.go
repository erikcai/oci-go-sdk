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
	"github.com/oracle/oci-go-sdk/v33/common"
)

// CreateLoadBalancerListenerEnvironmentDetails Specifies Load Balancer Listener environment.
type CreateLoadBalancerListenerEnvironmentDetails struct {

	// Environment Identifier
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Application Identifier
	ApplicationId *string `mandatory:"true" json:"applicationId"`

	// The OCID of Load Balancer
	LoadBalancerId *string `mandatory:"true" json:"loadBalancerId"`

	// Name of the Load Balancer Listener
	ListenerName *string `mandatory:"true" json:"listenerName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

//GetDisplayName returns DisplayName
func (m CreateLoadBalancerListenerEnvironmentDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetApplicationId returns ApplicationId
func (m CreateLoadBalancerListenerEnvironmentDetails) GetApplicationId() *string {
	return m.ApplicationId
}

//GetFreeformTags returns FreeformTags
func (m CreateLoadBalancerListenerEnvironmentDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m CreateLoadBalancerListenerEnvironmentDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateLoadBalancerListenerEnvironmentDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateLoadBalancerListenerEnvironmentDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateLoadBalancerListenerEnvironmentDetails CreateLoadBalancerListenerEnvironmentDetails
	s := struct {
		DiscriminatorParam string `json:"environmentType"`
		MarshalTypeCreateLoadBalancerListenerEnvironmentDetails
	}{
		"LOAD_BALANCE_LISTENER",
		(MarshalTypeCreateLoadBalancerListenerEnvironmentDetails)(m),
	}

	return json.Marshal(&s)
}
