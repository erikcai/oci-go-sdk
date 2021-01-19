// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DeploymentService API
//
// A description of the DeploymentService API
//

package clouddeploy

import (
	"github.com/oracle/oci-go-sdk/v33/common"
)

// ComputeInstanceGroupRollingDeploymentRollOutPolicy Specifies rollout policy for Compute Instance group rolling deployment stage.
type ComputeInstanceGroupRollingDeploymentRollOutPolicy struct {

	// Indicates maximum number of instances deploying at a time.
	MaxConcurrentInstances *int `mandatory:"true" json:"maxConcurrentInstances"`

	RollOutStrategy *ComputeInstanceGroupLinearRollOutStrategy `mandatory:"true" json:"rollOutStrategy"`

	// The number of seconds in delay between batch rollout. The default delay is 1 minute.
	BatchDelayInSeconds *int `mandatory:"false" json:"batchDelayInSeconds"`
}

func (m ComputeInstanceGroupRollingDeploymentRollOutPolicy) String() string {
	return common.PointerString(m)
}
