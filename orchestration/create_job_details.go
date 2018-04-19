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

// CreateJobDetails The representation of CreateJobDetails
type CreateJobDetails struct {
	StackId *string `mandatory:"true" json:"stackId"`

	// Terraform specific operation to execute
	Operation CreateJobDetailsOperationEnum `mandatory:"true" json:"operation"`

	DisplayName *string `mandatory:"false" json:"displayName"`

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

func (m CreateJobDetails) String() string {
	return common.PointerString(m)
}

// CreateJobDetailsOperationEnum Enum with underlying type: string
type CreateJobDetailsOperationEnum string

// Set of constants representing the allowable values for CreateJobDetailsOperation
const (
	CreateJobDetailsOperationPlan    CreateJobDetailsOperationEnum = "PLAN"
	CreateJobDetailsOperationApply   CreateJobDetailsOperationEnum = "APPLY"
	CreateJobDetailsOperationDestroy CreateJobDetailsOperationEnum = "DESTROY"
)

var mappingCreateJobDetailsOperation = map[string]CreateJobDetailsOperationEnum{
	"PLAN":    CreateJobDetailsOperationPlan,
	"APPLY":   CreateJobDetailsOperationApply,
	"DESTROY": CreateJobDetailsOperationDestroy,
}

// GetCreateJobDetailsOperationEnumValues Enumerates the set of values for CreateJobDetailsOperation
func GetCreateJobDetailsOperationEnumValues() []CreateJobDetailsOperationEnum {
	values := make([]CreateJobDetailsOperationEnum, 0)
	for _, v := range mappingCreateJobDetailsOperation {
		values = append(values, v)
	}
	return values
}
