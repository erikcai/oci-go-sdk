// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// CIMS API
// 
 // Cloud Incident Management Service
//

package cims


    // ProblemTypeEnum Enum with underlying type: string
type ProblemTypeEnum string

// Set of constants representing the allowable values for ProblemTypeEnum
const (
    ProblemTypeLimit ProblemTypeEnum = "LIMIT"
    ProblemTypeLegacyLimit ProblemTypeEnum = "LEGACY_LIMIT"
    ProblemTypeTech ProblemTypeEnum = "TECH"
    ProblemTypeAccount ProblemTypeEnum = "ACCOUNT"
)

var mappingProblemType = map[string]ProblemTypeEnum { 
    "LIMIT": ProblemTypeLimit,
    "LEGACY_LIMIT": ProblemTypeLegacyLimit,
    "TECH": ProblemTypeTech,
    "ACCOUNT": ProblemTypeAccount,
}

// GetProblemTypeEnumValues Enumerates the set of values for ProblemTypeEnum
func GetProblemTypeEnumValues() []ProblemTypeEnum {
   values := make([]ProblemTypeEnum, 0)
   for _, v := range mappingProblemType {
       values = append(values, v)
   }
   return values
}


