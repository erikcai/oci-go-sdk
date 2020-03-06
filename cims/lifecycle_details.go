// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// CIMS API
// 
 // Cloud Incident Management Service
//

package cims


    // LifecycleDetailsEnum Enum with underlying type: string
type LifecycleDetailsEnum string

// Set of constants representing the allowable values for LifecycleDetailsEnum
const (
    LifecycleDetailsPendingWithOracle LifecycleDetailsEnum = "PENDING_WITH_ORACLE"
    LifecycleDetailsPendingWithCustomer LifecycleDetailsEnum = "PENDING_WITH_CUSTOMER"
    LifecycleDetailsCloseRequested LifecycleDetailsEnum = "CLOSE_REQUESTED"
    LifecycleDetailsClosed LifecycleDetailsEnum = "CLOSED"
)

var mappingLifecycleDetails = map[string]LifecycleDetailsEnum { 
    "PENDING_WITH_ORACLE": LifecycleDetailsPendingWithOracle,
    "PENDING_WITH_CUSTOMER": LifecycleDetailsPendingWithCustomer,
    "CLOSE_REQUESTED": LifecycleDetailsCloseRequested,
    "CLOSED": LifecycleDetailsClosed,
}

// GetLifecycleDetailsEnumValues Enumerates the set of values for LifecycleDetailsEnum
func GetLifecycleDetailsEnumValues() []LifecycleDetailsEnum {
   values := make([]LifecycleDetailsEnum, 0)
   for _, v := range mappingLifecycleDetails {
       values = append(values, v)
   }
   return values
}


