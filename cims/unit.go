// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// CIMS API
// 
 // Cloud Incident Management Service
//

package cims


    // UnitEnum Enum with underlying type: string
type UnitEnum string

// Set of constants representing the allowable values for UnitEnum
const (
    UnitCount UnitEnum = "COUNT"
    UnitGb UnitEnum = "GB"
    UnitNone UnitEnum = "NONE"
)

var mappingUnit = map[string]UnitEnum { 
    "COUNT": UnitCount,
    "GB": UnitGb,
    "NONE": UnitNone,
}

// GetUnitEnumValues Enumerates the set of values for UnitEnum
func GetUnitEnumValues() []UnitEnum {
   values := make([]UnitEnum, 0)
   for _, v := range mappingUnit {
       values = append(values, v)
   }
   return values
}


