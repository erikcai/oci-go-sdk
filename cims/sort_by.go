// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// CIMS API
// 
 // Cloud Incident Management Service
//

package cims


    // SortByEnum Enum with underlying type: string
type SortByEnum string

// Set of constants representing the allowable values for SortByEnum
const (
    SortByDateUpdated SortByEnum = "dateUpdated"
    SortBySeverity SortByEnum = "severity"
)

var mappingSortBy = map[string]SortByEnum { 
    "dateUpdated": SortByDateUpdated,
    "severity": SortBySeverity,
}

// GetSortByEnumValues Enumerates the set of values for SortByEnum
func GetSortByEnumValues() []SortByEnum {
   values := make([]SortByEnum, 0)
   for _, v := range mappingSortBy {
       values = append(values, v)
   }
   return values
}


