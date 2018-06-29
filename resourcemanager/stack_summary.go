// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Oracle Resource Manager
//
// Oracle Resource Manager API
//

package resourcemanager

import (
	"github.com/oracle/oci-go-sdk/common"
)

// StackSummary The representation of StackSummary
type StackSummary struct {
	Id *string `mandatory:"false" json:"id"`

	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	DisplayName *string `mandatory:"false" json:"displayName"`

	Description *string `mandatory:"false" json:"description"`

	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	LifecycleState StackSummaryLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m StackSummary) String() string {
	return common.PointerString(m)
}

// StackSummaryLifecycleStateEnum Enum with underlying type: string
type StackSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for StackSummaryLifecycleState
const (
	StackSummaryLifecycleStateCreating StackSummaryLifecycleStateEnum = "CREATING"
	StackSummaryLifecycleStateActive   StackSummaryLifecycleStateEnum = "ACTIVE"
	StackSummaryLifecycleStateDeleting StackSummaryLifecycleStateEnum = "DELETING"
)

var mappingStackSummaryLifecycleState = map[string]StackSummaryLifecycleStateEnum{
	"CREATING": StackSummaryLifecycleStateCreating,
	"ACTIVE":   StackSummaryLifecycleStateActive,
	"DELETING": StackSummaryLifecycleStateDeleting,
}

// GetStackSummaryLifecycleStateEnumValues Enumerates the set of values for StackSummaryLifecycleState
func GetStackSummaryLifecycleStateEnumValues() []StackSummaryLifecycleStateEnum {
	values := make([]StackSummaryLifecycleStateEnum, 0)
	for _, v := range mappingStackSummaryLifecycleState {
		values = append(values, v)
	}
	return values
}
