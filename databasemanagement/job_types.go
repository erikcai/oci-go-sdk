// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management Service APIs.
//
// This file contains the customer facing APIs for Database Management service.
//

package databasemanagement

// JobTypesEnum Enum with underlying type: string
type JobTypesEnum string

// Set of constants representing the allowable values for JobTypesEnum
const (
	JobTypesSql JobTypesEnum = "SQL"
)

var mappingJobTypes = map[string]JobTypesEnum{
	"SQL": JobTypesSql,
}

// GetJobTypesEnumValues Enumerates the set of values for JobTypesEnum
func GetJobTypesEnumValues() []JobTypesEnum {
	values := make([]JobTypesEnum, 0)
	for _, v := range mappingJobTypes {
		values = append(values, v)
	}
	return values
}
