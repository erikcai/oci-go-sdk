// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Key Management Service API
// 
 // APIs for managing and performing operations with keys and vaults.
//

package kms

import (
    "github.com/oracle/oci-go-sdk/common"
)


    
 // CreateVaultDetails The representation of CreateVaultDetails
type CreateVaultDetails struct {
    
 // The OCID of the Compartment containing this resource.
    CompartmentId *string `mandatory:"true" json:"compartmentId"`
    
 // A user-friendly name. Does not have to be unique, and it's changeable.
 // Avoid entering confidential information.
    DisplayName *string `mandatory:"true" json:"displayName"`
    
 // TODO
    VaultType CreateVaultDetailsVaultTypeEnum `mandatory:"true" json:"vaultType"`
}

func (m CreateVaultDetails) String() string {
    return common.PointerString(m)
}


// CreateVaultDetailsVaultTypeEnum Enum with underlying type: string
type CreateVaultDetailsVaultTypeEnum string

// Set of constants representing the allowable values for CreateVaultDetailsVaultType
const (
    CreateVaultDetailsVaultTypePrivate CreateVaultDetailsVaultTypeEnum = "VIRTUAL_PRIVATE"
)

var mappingCreateVaultDetailsVaultType = map[string]CreateVaultDetailsVaultTypeEnum { 
    "VIRTUAL_PRIVATE": CreateVaultDetailsVaultTypePrivate,
}

// GetCreateVaultDetailsVaultTypeEnumValues Enumerates the set of values for CreateVaultDetailsVaultType
func GetCreateVaultDetailsVaultTypeEnumValues() []CreateVaultDetailsVaultTypeEnum {
   values := make([]CreateVaultDetailsVaultTypeEnum, 0)
   for _, v := range mappingCreateVaultDetailsVaultType {
       values = append(values, v)
   }
   return values
}



