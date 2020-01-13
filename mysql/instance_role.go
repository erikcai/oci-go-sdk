// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// MySQL as a Service API
//
// The API for the MySQL Service
//

package mysql

// InstanceRoleEnum Enum with underlying type: string
type InstanceRoleEnum string

// Set of constants representing the allowable values for InstanceRoleEnum
const (
	InstanceRoleStandalone InstanceRoleEnum = "STANDALONE"
	InstanceRolePrimary    InstanceRoleEnum = "PRIMARY"
	InstanceRoleSecondary  InstanceRoleEnum = "SECONDARY"
)

var mappingInstanceRole = map[string]InstanceRoleEnum{
	"STANDALONE": InstanceRoleStandalone,
	"PRIMARY":    InstanceRolePrimary,
	"SECONDARY":  InstanceRoleSecondary,
}

// GetInstanceRoleEnumValues Enumerates the set of values for InstanceRoleEnum
func GetInstanceRoleEnumValues() []InstanceRoleEnum {
	values := make([]InstanceRoleEnum, 0)
	for _, v := range mappingInstanceRole {
		values = append(values, v)
	}
	return values
}
