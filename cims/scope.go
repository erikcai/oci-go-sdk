// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// CIMS API
// 
 // Cloud Incident Management Service
//

package cims


    // ScopeEnum Enum with underlying type: string
type ScopeEnum string

// Set of constants representing the allowable values for ScopeEnum
const (
    ScopeAd ScopeEnum = "AD"
    ScopeRegion ScopeEnum = "REGION"
    ScopeTenancy ScopeEnum = "TENANCY"
    ScopeNone ScopeEnum = "NONE"
)

var mappingScope = map[string]ScopeEnum { 
    "AD": ScopeAd,
    "REGION": ScopeRegion,
    "TENANCY": ScopeTenancy,
    "NONE": ScopeNone,
}

// GetScopeEnumValues Enumerates the set of values for ScopeEnum
func GetScopeEnumValues() []ScopeEnum {
   values := make([]ScopeEnum, 0)
   for _, v := range mappingScope {
       values = append(values, v)
   }
   return values
}


