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


    
 // CreateKeyDetails The representation of CreateKeyDetails
type CreateKeyDetails struct {
    
 // The OCID of the Compartment containing this resource.
    CompartmentId *string `mandatory:"true" json:"compartmentId"`
    
 // A user-friendly name. Does not have to be unique, and it's changeable.
 // Avoid entering confidential information.
    DisplayName *string `mandatory:"true" json:"displayName"`
    
    KeyShape *KeyShape `mandatory:"true" json:"keyShape"`
}

func (m CreateKeyDetails) String() string {
    return common.PointerString(m)
}





