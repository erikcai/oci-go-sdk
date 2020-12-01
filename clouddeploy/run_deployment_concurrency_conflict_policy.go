// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DeploymentService API
//
// A description of the DeploymentService API
//

package clouddeploy

import (
	"github.com/oracle/oci-go-sdk/v30/common"
)

// RunDeploymentConcurrencyConflictPolicy This describes the policy in case there is a deployment already running on the target pipeline.
type RunDeploymentConcurrencyConflictPolicy struct {

	// The types of policies in case of a conflict
	ConcurrencyConflictPolicyType RunDeploymentConcurrencyConflictPolicyConcurrencyConflictPolicyTypeEnum `mandatory:"true" json:"concurrencyConflictPolicyType"`
}

func (m RunDeploymentConcurrencyConflictPolicy) String() string {
	return common.PointerString(m)
}

// RunDeploymentConcurrencyConflictPolicyConcurrencyConflictPolicyTypeEnum Enum with underlying type: string
type RunDeploymentConcurrencyConflictPolicyConcurrencyConflictPolicyTypeEnum string

// Set of constants representing the allowable values for RunDeploymentConcurrencyConflictPolicyConcurrencyConflictPolicyTypeEnum
const (
	RunDeploymentConcurrencyConflictPolicyConcurrencyConflictPolicyTypeWaitForCompletion RunDeploymentConcurrencyConflictPolicyConcurrencyConflictPolicyTypeEnum = "WAIT_FOR_COMPLETION"
	RunDeploymentConcurrencyConflictPolicyConcurrencyConflictPolicyTypeFailImmediately   RunDeploymentConcurrencyConflictPolicyConcurrencyConflictPolicyTypeEnum = "FAIL_IMMEDIATELY"
)

var mappingRunDeploymentConcurrencyConflictPolicyConcurrencyConflictPolicyType = map[string]RunDeploymentConcurrencyConflictPolicyConcurrencyConflictPolicyTypeEnum{
	"WAIT_FOR_COMPLETION": RunDeploymentConcurrencyConflictPolicyConcurrencyConflictPolicyTypeWaitForCompletion,
	"FAIL_IMMEDIATELY":    RunDeploymentConcurrencyConflictPolicyConcurrencyConflictPolicyTypeFailImmediately,
}

// GetRunDeploymentConcurrencyConflictPolicyConcurrencyConflictPolicyTypeEnumValues Enumerates the set of values for RunDeploymentConcurrencyConflictPolicyConcurrencyConflictPolicyTypeEnum
func GetRunDeploymentConcurrencyConflictPolicyConcurrencyConflictPolicyTypeEnumValues() []RunDeploymentConcurrencyConflictPolicyConcurrencyConflictPolicyTypeEnum {
	values := make([]RunDeploymentConcurrencyConflictPolicyConcurrencyConflictPolicyTypeEnum, 0)
	for _, v := range mappingRunDeploymentConcurrencyConflictPolicyConcurrencyConflictPolicyType {
		values = append(values, v)
	}
	return values
}
