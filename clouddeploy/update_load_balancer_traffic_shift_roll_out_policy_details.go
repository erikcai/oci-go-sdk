// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DeploymentService API
//
// A description of the DeploymentService API
//

package clouddeploy

import (
	"github.com/oracle/oci-go-sdk/v31/common"
)

// UpdateLoadBalancerTrafficShiftRollOutPolicyDetails Description of rollout policy for Load Balancer Traffic Shift stage.
type UpdateLoadBalancerTrafficShiftRollOutPolicyDetails struct {

	// Specifies number of batches for this stage.
	BatchCount *int `mandatory:"true" json:"batchCount"`

	// Specifies delay in seconds between batches.
	BatchDelayInSeconds *int `mandatory:"false" json:"batchDelayInSeconds"`

	// Indicates the stoping criteria.
	RampLimitPercent *float32 `mandatory:"false" json:"rampLimitPercent"`
}

func (m UpdateLoadBalancerTrafficShiftRollOutPolicyDetails) String() string {
	return common.PointerString(m)
}
