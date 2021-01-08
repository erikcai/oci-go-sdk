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
	"github.com/oracle/oci-go-sdk/v31/common"
)

// ComputeInstanceGroupSelectorCollection A collection of selectors. The UNION of instances matching the selectors are included in the instance group.
type ComputeInstanceGroupSelectorCollection struct {

	// An array of selectors for the instance group. UNION operator will be use for combining the instances selected by each selector.
	Items []ComputeInstanceGroupSelector `mandatory:"true" json:"items"`
}

func (m ComputeInstanceGroupSelectorCollection) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *ComputeInstanceGroupSelectorCollection) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Items []computeinstancegroupselector `json:"items"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Items = make([]ComputeInstanceGroupSelector, len(model.Items))
	for i, n := range model.Items {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Items[i] = nn.(ComputeInstanceGroupSelector)
		} else {
			m.Items[i] = nil
		}
	}

	return
}
