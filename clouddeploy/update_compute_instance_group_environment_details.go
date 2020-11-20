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

// UpdateComputeInstanceGroupEnvironmentDetails Specifies the Compute Instance Group environment.
type UpdateComputeInstanceGroupEnvironmentDetails struct {

	// Environment Identifier
	DisplayName *string `mandatory:"true" json:"displayName"`

	ComputeInstanceGroupSelectors *ComputeInstanceGroupSelectorCollection `mandatory:"true" json:"computeInstanceGroupSelectors"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

//GetDisplayName returns DisplayName
func (m UpdateComputeInstanceGroupEnvironmentDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetFreeformTags returns FreeformTags
func (m UpdateComputeInstanceGroupEnvironmentDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m UpdateComputeInstanceGroupEnvironmentDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m UpdateComputeInstanceGroupEnvironmentDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m UpdateComputeInstanceGroupEnvironmentDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateComputeInstanceGroupEnvironmentDetails UpdateComputeInstanceGroupEnvironmentDetails
	s := struct {
		DiscriminatorParam string `json:"environmentType"`
		MarshalTypeUpdateComputeInstanceGroupEnvironmentDetails
	}{
		"COMPUTE_INSTANCE_GROUP",
		(MarshalTypeUpdateComputeInstanceGroupEnvironmentDetails)(m),
	}

	return json.Marshal(&s)
}
