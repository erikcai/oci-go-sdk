// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// PublicLoggingControlplane API
//
// PublicLoggingControlplane API specification
//

package logging

// LogRuleLifecycleStateEnum Enum with underlying type: string
type LogRuleLifecycleStateEnum string

// Set of constants representing the allowable values for LogRuleLifecycleStateEnum
const (
	LogRuleLifecycleStateCreating LogRuleLifecycleStateEnum = "CREATING"
	LogRuleLifecycleStateActive   LogRuleLifecycleStateEnum = "ACTIVE"
	LogRuleLifecycleStateUpdating LogRuleLifecycleStateEnum = "UPDATING"
	LogRuleLifecycleStateInactive LogRuleLifecycleStateEnum = "INACTIVE"
	LogRuleLifecycleStateDeleting LogRuleLifecycleStateEnum = "DELETING"
)

var mappingLogRuleLifecycleState = map[string]LogRuleLifecycleStateEnum{
	"CREATING": LogRuleLifecycleStateCreating,
	"ACTIVE":   LogRuleLifecycleStateActive,
	"UPDATING": LogRuleLifecycleStateUpdating,
	"INACTIVE": LogRuleLifecycleStateInactive,
	"DELETING": LogRuleLifecycleStateDeleting,
}

// GetLogRuleLifecycleStateEnumValues Enumerates the set of values for LogRuleLifecycleStateEnum
func GetLogRuleLifecycleStateEnumValues() []LogRuleLifecycleStateEnum {
	values := make([]LogRuleLifecycleStateEnum, 0)
	for _, v := range mappingLogRuleLifecycleState {
		values = append(values, v)
	}
	return values
}
