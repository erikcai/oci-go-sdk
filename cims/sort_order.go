// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Support Management API
// 
 // Use the Support Management API to manage support requests. For more information, see Getting Help and Contacting Support (https://docs.cloud.oracle.com/iaas/Content/GSG/Tasks/contactingsupport.htm). **Note**: Before you can create service requests with this API, you need to have an Oracle Single Sign On (SSO) account, and you need to register your Customer Support Identifier (CSI) with My Oracle Support.
//

package cims


    // SortOrderEnum Enum with underlying type: string
type SortOrderEnum string

// Set of constants representing the allowable values for SortOrderEnum
const (
    SortOrderAsc SortOrderEnum = "ASC"
    SortOrderDesc SortOrderEnum = "DESC"
)

var mappingSortOrder = map[string]SortOrderEnum { 
    "ASC": SortOrderAsc,
    "DESC": SortOrderDesc,
}

// GetSortOrderEnumValues Enumerates the set of values for SortOrderEnum
func GetSortOrderEnumValues() []SortOrderEnum {
   values := make([]SortOrderEnum, 0)
   for _, v := range mappingSortOrder {
       values = append(values, v)
   }
   return values
}


