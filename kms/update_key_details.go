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


    
 // UpdateKeyDetails The representation of UpdateKeyDetails
type UpdateKeyDetails struct {
    
 // A user-friendly name. Does not have to be unique, and it's changeable.
 // Avoid entering confidential information.
    DisplayName *string `mandatory:"false" json:"displayName"`
}

func (m UpdateKeyDetails) String() string {
    return common.PointerString(m)
}





