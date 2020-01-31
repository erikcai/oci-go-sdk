// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// MySQL as a Service API
//
// The API for the MySQL Service
//

package mysql

// InnoDbShutdownModeEnum Enum with underlying type: string
type InnoDbShutdownModeEnum string

// Set of constants representing the allowable values for InnoDbShutdownModeEnum
const (
	InnoDbShutdownModeImmediate InnoDbShutdownModeEnum = "IMMEDIATE"
	InnoDbShutdownModeFast      InnoDbShutdownModeEnum = "FAST"
	InnoDbShutdownModeSlow      InnoDbShutdownModeEnum = "SLOW"
)

var mappingInnoDbShutdownMode = map[string]InnoDbShutdownModeEnum{
	"IMMEDIATE": InnoDbShutdownModeImmediate,
	"FAST":      InnoDbShutdownModeFast,
	"SLOW":      InnoDbShutdownModeSlow,
}

// GetInnoDbShutdownModeEnumValues Enumerates the set of values for InnoDbShutdownModeEnum
func GetInnoDbShutdownModeEnumValues() []InnoDbShutdownModeEnum {
	values := make([]InnoDbShutdownModeEnum, 0)
	for _, v := range mappingInnoDbShutdownMode {
		values = append(values, v)
	}
	return values
}