// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Data Science API
//
// The Data Science service enables data science teams to organize their work, easily access data and computing resources, and build, train, deploy, and manage ML/AI models on the Oracle Cloud.
//

package datascience

// ModelLifecycleStateEnum Enum with underlying type: string
type ModelLifecycleStateEnum string

// Set of constants representing the allowable values for ModelLifecycleStateEnum
const (
	ModelLifecycleStateActive   ModelLifecycleStateEnum = "ACTIVE"
	ModelLifecycleStateDeleted  ModelLifecycleStateEnum = "DELETED"
	ModelLifecycleStateFailed   ModelLifecycleStateEnum = "FAILED"
	ModelLifecycleStateInactive ModelLifecycleStateEnum = "INACTIVE"
)

var mappingModelLifecycleState = map[string]ModelLifecycleStateEnum{
	"ACTIVE":   ModelLifecycleStateActive,
	"DELETED":  ModelLifecycleStateDeleted,
	"FAILED":   ModelLifecycleStateFailed,
	"INACTIVE": ModelLifecycleStateInactive,
}

// GetModelLifecycleStateEnumValues Enumerates the set of values for ModelLifecycleStateEnum
func GetModelLifecycleStateEnumValues() []ModelLifecycleStateEnum {
	values := make([]ModelLifecycleStateEnum, 0)
	for _, v := range mappingModelLifecycleState {
		values = append(values, v)
	}
	return values
}
