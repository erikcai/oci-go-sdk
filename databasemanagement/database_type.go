// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management Service APIs.
//
// This file contains the customer facing APIs for Database Management service.
//

package databasemanagement

// DatabaseTypeEnum Enum with underlying type: string
type DatabaseTypeEnum string

// Set of constants representing the allowable values for DatabaseTypeEnum
const (
	DatabaseTypeExternalSidb DatabaseTypeEnum = "EXTERNAL_SIDB"
	DatabaseTypeExternalRac  DatabaseTypeEnum = "EXTERNAL_RAC"
)

var mappingDatabaseType = map[string]DatabaseTypeEnum{
	"EXTERNAL_SIDB": DatabaseTypeExternalSidb,
	"EXTERNAL_RAC":  DatabaseTypeExternalRac,
}

// GetDatabaseTypeEnumValues Enumerates the set of values for DatabaseTypeEnum
func GetDatabaseTypeEnumValues() []DatabaseTypeEnum {
	values := make([]DatabaseTypeEnum, 0)
	for _, v := range mappingDatabaseType {
		values = append(values, v)
	}
	return values
}
