// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Data Flow API
//
// Use the Data Flow APIs to run any Apache Spark application at any scale without deploying or managing any infrastructure.
//

package dataflow

// ApplicationTypeEnum Enum with underlying type: string
type ApplicationTypeEnum string

// Set of constants representing the allowable values for ApplicationTypeEnum
const (
	ApplicationTypeBatch   ApplicationTypeEnum = "BATCH"
	ApplicationTypeSession ApplicationTypeEnum = "SESSION"
)

var mappingApplicationType = map[string]ApplicationTypeEnum{
	"BATCH":   ApplicationTypeBatch,
	"SESSION": ApplicationTypeSession,
}

// GetApplicationTypeEnumValues Enumerates the set of values for ApplicationTypeEnum
func GetApplicationTypeEnumValues() []ApplicationTypeEnum {
	values := make([]ApplicationTypeEnum, 0)
	for _, v := range mappingApplicationType {
		values = append(values, v)
	}
	return values
}
