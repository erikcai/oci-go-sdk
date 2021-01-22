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

// ComputeInstanceGroupByIdsSelector Specifies the Compute Instance Group environment by listing the OCIDs of the compute instances
type ComputeInstanceGroupByIdsSelector struct {

	// Compute instance OCID identifiers that are members of this group.
	ComputeInstanceIds []string `mandatory:"true" json:"computeInstanceIds"`
}

func (m ComputeInstanceGroupByIdsSelector) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ComputeInstanceGroupByIdsSelector) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeComputeInstanceGroupByIdsSelector ComputeInstanceGroupByIdsSelector
	s := struct {
		DiscriminatorParam string `json:"selectorType"`
		MarshalTypeComputeInstanceGroupByIdsSelector
	}{
		"INSTANCE_IDS",
		(MarshalTypeComputeInstanceGroupByIdsSelector)(m),
	}

	return json.Marshal(&s)
}
