// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Data Science API
//
// The Data Science service enables data science teams to organize their work, easily access data and computing resources, and build, train, deploy, and manage ML/AI models on the Oracle Cloud.
//

package datascience

// ProjectLifecycleStateEnum Enum with underlying type: string
type ProjectLifecycleStateEnum string

// Set of constants representing the allowable values for ProjectLifecycleStateEnum
const (
	ProjectLifecycleStateActive   ProjectLifecycleStateEnum = "ACTIVE"
	ProjectLifecycleStateDeleting ProjectLifecycleStateEnum = "DELETING"
	ProjectLifecycleStateDeleted  ProjectLifecycleStateEnum = "DELETED"
)

var mappingProjectLifecycleState = map[string]ProjectLifecycleStateEnum{
	"ACTIVE":   ProjectLifecycleStateActive,
	"DELETING": ProjectLifecycleStateDeleting,
	"DELETED":  ProjectLifecycleStateDeleted,
}

// GetProjectLifecycleStateEnumValues Enumerates the set of values for ProjectLifecycleStateEnum
func GetProjectLifecycleStateEnumValues() []ProjectLifecycleStateEnum {
	values := make([]ProjectLifecycleStateEnum, 0)
	for _, v := range mappingProjectLifecycleState {
		values = append(values, v)
	}
	return values
}
