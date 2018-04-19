// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Orchestration
//
// Orchestration Maestro API
//

package orchestration

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Job The representation of Job
type Job struct {
	Id *string `mandatory:"false" json:"id"`

	StackId *string `mandatory:"false" json:"stackId"`

	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	DisplayName *string `mandatory:"false" json:"displayName"`

	// The type of job executing
	Operation JobOperationEnum `mandatory:"false" json:"operation,omitempty"`

	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the job finished running independent on whether or not the job was successful.
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

	LifecycleState JobLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	Variables map[string]string `mandatory:"false" json:"variables"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m Job) String() string {
	return common.PointerString(m)
}

// JobOperationEnum Enum with underlying type: string
type JobOperationEnum string

// Set of constants representing the allowable values for JobOperation
const (
	JobOperationPlan    JobOperationEnum = "PLAN"
	JobOperationApply   JobOperationEnum = "APPLY"
	JobOperationDestroy JobOperationEnum = "DESTROY"
)

var mappingJobOperation = map[string]JobOperationEnum{
	"PLAN":    JobOperationPlan,
	"APPLY":   JobOperationApply,
	"DESTROY": JobOperationDestroy,
}

// GetJobOperationEnumValues Enumerates the set of values for JobOperation
func GetJobOperationEnumValues() []JobOperationEnum {
	values := make([]JobOperationEnum, 0)
	for _, v := range mappingJobOperation {
		values = append(values, v)
	}
	return values
}

// JobLifecycleStateEnum Enum with underlying type: string
type JobLifecycleStateEnum string

// Set of constants representing the allowable values for JobLifecycleState
const (
	JobLifecycleStateAccepted   JobLifecycleStateEnum = "ACCEPTED"
	JobLifecycleStateInProgress JobLifecycleStateEnum = "IN_PROGRESS"
	JobLifecycleStateFailed     JobLifecycleStateEnum = "FAILED"
	JobLifecycleStateSucceeded  JobLifecycleStateEnum = "SUCCEEDED"
	JobLifecycleStateCanceling  JobLifecycleStateEnum = "CANCELING"
	JobLifecycleStateCanceled   JobLifecycleStateEnum = "CANCELED"
)

var mappingJobLifecycleState = map[string]JobLifecycleStateEnum{
	"ACCEPTED":    JobLifecycleStateAccepted,
	"IN_PROGRESS": JobLifecycleStateInProgress,
	"FAILED":      JobLifecycleStateFailed,
	"SUCCEEDED":   JobLifecycleStateSucceeded,
	"CANCELING":   JobLifecycleStateCanceling,
	"CANCELED":    JobLifecycleStateCanceled,
}

// GetJobLifecycleStateEnumValues Enumerates the set of values for JobLifecycleState
func GetJobLifecycleStateEnumValues() []JobLifecycleStateEnum {
	values := make([]JobLifecycleStateEnum, 0)
	for _, v := range mappingJobLifecycleState {
		values = append(values, v)
	}
	return values
}
