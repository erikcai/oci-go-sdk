// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management Service APIs.
//
// This file contains the customer facing APIs for Database Management service.
//

package databasemanagement

// DatabaseSubTypeEnum Enum with underlying type: string
type DatabaseSubTypeEnum string

// Set of constants representing the allowable values for DatabaseSubTypeEnum
const (
	DatabaseSubTypeCdb    DatabaseSubTypeEnum = "CDB"
	DatabaseSubTypePdb    DatabaseSubTypeEnum = "PDB"
	DatabaseSubTypeNonCdb DatabaseSubTypeEnum = "NON_CDB"
)

var mappingDatabaseSubType = map[string]DatabaseSubTypeEnum{
	"CDB":     DatabaseSubTypeCdb,
	"PDB":     DatabaseSubTypePdb,
	"NON_CDB": DatabaseSubTypeNonCdb,
}

// GetDatabaseSubTypeEnumValues Enumerates the set of values for DatabaseSubTypeEnum
func GetDatabaseSubTypeEnumValues() []DatabaseSubTypeEnum {
	values := make([]DatabaseSubTypeEnum, 0)
	for _, v := range mappingDatabaseSubType {
		values = append(values, v)
	}
	return values
}