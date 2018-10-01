// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// EventsControlService API
//
// This service exposes APIs to create, update and delete Rules. Rules are used to tap into the Events stream.
//

package cloudevents

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Rule Detailed representation of a rule.
type Rule struct {

	// A string that describes the rule
	DisplayName *string `mandatory:"true" json:"displayName"`

	// specifies current state of the rule
	LifecycleState RuleLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// specifies what the rule matches
	Condition *string `mandatory:"true" json:"condition"`

	// OCID of the compartment the rule belongs to
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// whether or not this rule is currently enabled
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// List of action object.
	Actions *ActionList `mandatory:"true" json:"actions"`

	// rule ocid
	Id *string `mandatory:"true" json:"id"`

	// time rule was created
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. Exists for cross-compatibility only.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}'
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Service-generated message about the current state of this rule
	LifecycleMessage *string `mandatory:"false" json:"lifecycleMessage"`
}

func (m Rule) String() string {
	return common.PointerString(m)
}

// RuleLifecycleStateEnum Enum with underlying type: string
type RuleLifecycleStateEnum string

// Set of constants representing the allowable values for RuleLifecycleStateEnum
const (
	RuleLifecycleStateCreating RuleLifecycleStateEnum = "CREATING"
	RuleLifecycleStateActive   RuleLifecycleStateEnum = "ACTIVE"
	RuleLifecycleStateInactive RuleLifecycleStateEnum = "INACTIVE"
	RuleLifecycleStateUpdating RuleLifecycleStateEnum = "UPDATING"
	RuleLifecycleStateDeleting RuleLifecycleStateEnum = "DELETING"
	RuleLifecycleStateDeleted  RuleLifecycleStateEnum = "DELETED"
	RuleLifecycleStateFailed   RuleLifecycleStateEnum = "FAILED"
)

var mappingRuleLifecycleState = map[string]RuleLifecycleStateEnum{
	"CREATING": RuleLifecycleStateCreating,
	"ACTIVE":   RuleLifecycleStateActive,
	"INACTIVE": RuleLifecycleStateInactive,
	"UPDATING": RuleLifecycleStateUpdating,
	"DELETING": RuleLifecycleStateDeleting,
	"DELETED":  RuleLifecycleStateDeleted,
	"FAILED":   RuleLifecycleStateFailed,
}

// GetRuleLifecycleStateEnumValues Enumerates the set of values for RuleLifecycleStateEnum
func GetRuleLifecycleStateEnumValues() []RuleLifecycleStateEnum {
	values := make([]RuleLifecycleStateEnum, 0)
	for _, v := range mappingRuleLifecycleState {
		values = append(values, v)
	}
	return values
}
