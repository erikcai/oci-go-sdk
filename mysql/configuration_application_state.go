// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// MySQL as a Service API
//
// The API for the MySQL Service
//

package mysql

// ConfigurationApplicationStateEnum Enum with underlying type: string
type ConfigurationApplicationStateEnum string

// Set of constants representing the allowable values for ConfigurationApplicationStateEnum
const (
	ConfigurationApplicationStatePending  ConfigurationApplicationStateEnum = "PENDING"
	ConfigurationApplicationStateApplying ConfigurationApplicationStateEnum = "APPLYING"
	ConfigurationApplicationStateApplied  ConfigurationApplicationStateEnum = "APPLIED"
	ConfigurationApplicationStateFailed   ConfigurationApplicationStateEnum = "FAILED"
)

var mappingConfigurationApplicationState = map[string]ConfigurationApplicationStateEnum{
	"PENDING":  ConfigurationApplicationStatePending,
	"APPLYING": ConfigurationApplicationStateApplying,
	"APPLIED":  ConfigurationApplicationStateApplied,
	"FAILED":   ConfigurationApplicationStateFailed,
}

// GetConfigurationApplicationStateEnumValues Enumerates the set of values for ConfigurationApplicationStateEnum
func GetConfigurationApplicationStateEnumValues() []ConfigurationApplicationStateEnum {
	values := make([]ConfigurationApplicationStateEnum, 0)
	for _, v := range mappingConfigurationApplicationState {
		values = append(values, v)
	}
	return values
}
