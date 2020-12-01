// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Autoscaling API
//
// APIs for dynamically scaling Compute resources to meet application requirements. For more information about
// autoscaling, see Autoscaling (https://docs.cloud.oracle.com/Content/Compute/Tasks/autoscalinginstancepools.htm). For information about the
// Compute service, see Overview of the Compute Service (https://docs.cloud.oracle.com/Content/Compute/Concepts/computeoverview.htm).
// **Note:** Autoscaling is not available in US Government Cloud tenancies. For more information, see
// Oracle Cloud Infrastructure US Government Cloud (https://docs.cloud.oracle.com/Content/General/Concepts/govoverview.htm).
//

package autoscaling

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v30/common"
)

// InstanceResource A Compute instance.
type InstanceResource struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the resource that is managed by the autoscaling configuration.
	Id *string `mandatory:"true" json:"id"`
}

//GetId returns Id
func (m InstanceResource) GetId() *string {
	return m.Id
}

func (m InstanceResource) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m InstanceResource) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeInstanceResource InstanceResource
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeInstanceResource
	}{
		"instance",
		(MarshalTypeInstanceResource)(m),
	}

	return json.Marshal(&s)
}
