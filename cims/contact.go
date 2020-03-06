// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// CIMS API
// 
 // Cloud Incident Management Service
//

package cims

import (
    "github.com/oracle/oci-go-sdk/common"
)


    
 // Contact Contact Details of the Customer
type Contact struct {
    
 // Contact person name
    ContactName *string `mandatory:"false" json:"contactName"`
    
 // Contact person email
    ContactEmail *string `mandatory:"false" json:"contactEmail"`
    
 // Contact person phone number
    ContactPhone *string `mandatory:"false" json:"contactPhone"`
    
 // ContactType enum. eg: MANAGER, PRIMARY
    ContactType ContactContactTypeEnum `mandatory:"false" json:"contactType,omitempty"`
}

func (m Contact) String() string {
    return common.PointerString(m)
}




// ContactContactTypeEnum Enum with underlying type: string
type ContactContactTypeEnum string

// Set of constants representing the allowable values for ContactContactTypeEnum
const (
    ContactContactTypePrimary ContactContactTypeEnum = "PRIMARY"
    ContactContactTypeAlternate ContactContactTypeEnum = "ALTERNATE"
    ContactContactTypeSecondary ContactContactTypeEnum = "SECONDARY"
    ContactContactTypeAdmin ContactContactTypeEnum = "ADMIN"
    ContactContactTypeManager ContactContactTypeEnum = "MANAGER"
)

var mappingContactContactType = map[string]ContactContactTypeEnum { 
    "PRIMARY": ContactContactTypePrimary,
    "ALTERNATE": ContactContactTypeAlternate,
    "SECONDARY": ContactContactTypeSecondary,
    "ADMIN": ContactContactTypeAdmin,
    "MANAGER": ContactContactTypeManager,
}

// GetContactContactTypeEnumValues Enumerates the set of values for ContactContactTypeEnum
func GetContactContactTypeEnumValues() []ContactContactTypeEnum {
   values := make([]ContactContactTypeEnum, 0)
   for _, v := range mappingContactContactType {
       values = append(values, v)
   }
   return values
}



