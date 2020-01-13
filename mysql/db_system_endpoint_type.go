// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// MySQL as a Service API
//
// The API for the MySQL Service
//

package mysql

// DbSystemEndpointTypeEnum Enum with underlying type: string
type DbSystemEndpointTypeEnum string

// Set of constants representing the allowable values for DbSystemEndpointTypeEnum
const (
	DbSystemEndpointTypeReadWrite DbSystemEndpointTypeEnum = "READ_WRITE"
	DbSystemEndpointTypeReadOnly  DbSystemEndpointTypeEnum = "READ_ONLY"
)

var mappingDbSystemEndpointType = map[string]DbSystemEndpointTypeEnum{
	"READ_WRITE": DbSystemEndpointTypeReadWrite,
	"READ_ONLY":  DbSystemEndpointTypeReadOnly,
}

// GetDbSystemEndpointTypeEnumValues Enumerates the set of values for DbSystemEndpointTypeEnum
func GetDbSystemEndpointTypeEnumValues() []DbSystemEndpointTypeEnum {
	values := make([]DbSystemEndpointTypeEnum, 0)
	for _, v := range mappingDbSystemEndpointType {
		values = append(values, v)
	}
	return values
}
