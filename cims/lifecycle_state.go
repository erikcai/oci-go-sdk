// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// CIMS API
// 
 // Cloud Incident Management Service
//

package cims


    // LifecycleStateEnum Enum with underlying type: string
type LifecycleStateEnum string

// Set of constants representing the allowable values for LifecycleStateEnum
const (
    LifecycleStateActive LifecycleStateEnum = "ACTIVE"
    LifecycleStateClosed LifecycleStateEnum = "CLOSED"
)

var mappingLifecycleState = map[string]LifecycleStateEnum { 
    "ACTIVE": LifecycleStateActive,
    "CLOSED": LifecycleStateClosed,
}

// GetLifecycleStateEnumValues Enumerates the set of values for LifecycleStateEnum
func GetLifecycleStateEnumValues() []LifecycleStateEnum {
   values := make([]LifecycleStateEnum, 0)
   for _, v := range mappingLifecycleState {
       values = append(values, v)
   }
   return values
}


