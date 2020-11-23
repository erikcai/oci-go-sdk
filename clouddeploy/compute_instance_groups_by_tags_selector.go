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

// ComputeInstanceGroupsByTagsSelector Specifies the Compute Instance Group environment by tags of the compute instances
type ComputeInstanceGroupsByTagsSelector struct {

	// The compartment id of the instances in the group.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The region of the instances in the group.
	Region *string `mandatory:"true" json:"region"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m ComputeInstanceGroupsByTagsSelector) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ComputeInstanceGroupsByTagsSelector) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeComputeInstanceGroupsByTagsSelector ComputeInstanceGroupsByTagsSelector
	s := struct {
		DiscriminatorParam string `json:"selectorType"`
		MarshalTypeComputeInstanceGroupsByTagsSelector
	}{
		"INSTANCE_TAGS",
		(MarshalTypeComputeInstanceGroupsByTagsSelector)(m),
	}

	return json.Marshal(&s)
}