// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// MySQL as a Service API
//
// The API for the MySQL Service
//

package mysql

// ConfigurationLifecycleStateEnum Enum with underlying type: string
type ConfigurationLifecycleStateEnum string

// Set of constants representing the allowable values for ConfigurationLifecycleStateEnum
const (
	ConfigurationLifecycleStateActive  ConfigurationLifecycleStateEnum = "ACTIVE"
	ConfigurationLifecycleStateDeleted ConfigurationLifecycleStateEnum = "DELETED"
)

var mappingConfigurationLifecycleState = map[string]ConfigurationLifecycleStateEnum{
	"ACTIVE":  ConfigurationLifecycleStateActive,
	"DELETED": ConfigurationLifecycleStateDeleted,
}

// GetConfigurationLifecycleStateEnumValues Enumerates the set of values for ConfigurationLifecycleStateEnum
func GetConfigurationLifecycleStateEnumValues() []ConfigurationLifecycleStateEnum {
	values := make([]ConfigurationLifecycleStateEnum, 0)
	for _, v := range mappingConfigurationLifecycleState {
		values = append(values, v)
	}
	return values
}
