// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DeploymentService API
//
// A description of the DeploymentService API
//

package clouddeploy

import (
	"github.com/oracle/oci-go-sdk/v29/common"
)

// UpdateRunDeploymentConcurrencyConflictPolicyDetails This describes the policy in case there is a deployment already running on the target pipeline.
type UpdateRunDeploymentConcurrencyConflictPolicyDetails struct {

	// The types of policies in case of a conflict
	ConcurrencyConflictPolicyType RunDeploymentConcurrencyConflictPolicyConcurrencyConflictPolicyTypeEnum `mandatory:"false" json:"concurrencyConflictPolicyType,omitempty"`
}

func (m UpdateRunDeploymentConcurrencyConflictPolicyDetails) String() string {
	return common.PointerString(m)
}
