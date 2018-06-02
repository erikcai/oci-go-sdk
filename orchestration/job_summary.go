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

// JobSummary The representation of JobSummary
type JobSummary struct {
	Id *string `mandatory:"false" json:"id"`

	StackId *string `mandatory:"false" json:"stackId"`

	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	DisplayName *string `mandatory:"false" json:"displayName"`

	// The type of job executing
	Operation JobSummaryOperationEnum `mandatory:"false" json:"operation,omitempty"`

	// The job to use as input plan for this apply. Only valid on apply. Must be the id of the most recent PLAN job. May be sentinel value LATEST, to use latest job without specifying id. May be sentinel AUTO_APPROVE to use the configuration directly without reference to any execution plan.
	PlanJob *string `mandatory:"false" json:"planJob"`

	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the job succeeded or failed
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

	LifecycleState JobSummaryLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m JobSummary) String() string {
	return common.PointerString(m)
}

// JobSummaryOperationEnum Enum with underlying type: string
type JobSummaryOperationEnum string

// Set of constants representing the allowable values for JobSummaryOperation
const (
	JobSummaryOperationPlan    JobSummaryOperationEnum = "PLAN"
	JobSummaryOperationApply   JobSummaryOperationEnum = "APPLY"
	JobSummaryOperationDestroy JobSummaryOperationEnum = "DESTROY"
)

var mappingJobSummaryOperation = map[string]JobSummaryOperationEnum{
	"PLAN":    JobSummaryOperationPlan,
	"APPLY":   JobSummaryOperationApply,
	"DESTROY": JobSummaryOperationDestroy,
}

// GetJobSummaryOperationEnumValues Enumerates the set of values for JobSummaryOperation
func GetJobSummaryOperationEnumValues() []JobSummaryOperationEnum {
	values := make([]JobSummaryOperationEnum, 0)
	for _, v := range mappingJobSummaryOperation {
		values = append(values, v)
	}
	return values
}

// JobSummaryLifecycleStateEnum Enum with underlying type: string
type JobSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for JobSummaryLifecycleState
const (
	JobSummaryLifecycleStateAccepted   JobSummaryLifecycleStateEnum = "ACCEPTED"
	JobSummaryLifecycleStateInProgress JobSummaryLifecycleStateEnum = "IN_PROGRESS"
	JobSummaryLifecycleStateFailed     JobSummaryLifecycleStateEnum = "FAILED"
	JobSummaryLifecycleStateSucceeded  JobSummaryLifecycleStateEnum = "SUCCEEDED"
	JobSummaryLifecycleStateCanceling  JobSummaryLifecycleStateEnum = "CANCELING"
	JobSummaryLifecycleStateCanceled   JobSummaryLifecycleStateEnum = "CANCELED"
)

var mappingJobSummaryLifecycleState = map[string]JobSummaryLifecycleStateEnum{
	"ACCEPTED":    JobSummaryLifecycleStateAccepted,
	"IN_PROGRESS": JobSummaryLifecycleStateInProgress,
	"FAILED":      JobSummaryLifecycleStateFailed,
	"SUCCEEDED":   JobSummaryLifecycleStateSucceeded,
	"CANCELING":   JobSummaryLifecycleStateCanceling,
	"CANCELED":    JobSummaryLifecycleStateCanceled,
}

// GetJobSummaryLifecycleStateEnumValues Enumerates the set of values for JobSummaryLifecycleState
func GetJobSummaryLifecycleStateEnumValues() []JobSummaryLifecycleStateEnum {
	values := make([]JobSummaryLifecycleStateEnum, 0)
	for _, v := range mappingJobSummaryLifecycleState {
		values = append(values, v)
	}
	return values
}
