// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Application Migration Service API
//
// API for the Application Migration service. Use this API to migrate applications from Oracle Cloud Infrastructure - Classic to Oracle Cloud Infrastructure.
//

package applicationmigration

// SourceTypesEnum Enum with underlying type: string
type SourceTypesEnum string

// Set of constants representing the allowable values for SourceTypesEnum
const (
	SourceTypesOcic SourceTypesEnum = "OCIC"
)

var mappingSourceTypes = map[string]SourceTypesEnum{
	"OCIC": SourceTypesOcic,
}

// GetSourceTypesEnumValues Enumerates the set of values for SourceTypesEnum
func GetSourceTypesEnumValues() []SourceTypesEnum {
	values := make([]SourceTypesEnum, 0)
	for _, v := range mappingSourceTypes {
		values = append(values, v)
	}
	return values
}
