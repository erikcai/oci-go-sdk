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


    
 // GenerateKeyDetails The representation of GenerateKeyDetails
type GenerateKeyDetails struct {
    
 // If true, the generated key is also returned unencrypted.
    IncludePlaintextKey *bool `mandatory:"true" json:"includePlaintextKey"`
    
 // The OCID of the Key to encrypt the generated key with.
    KeyId *string `mandatory:"true" json:"keyId"`
    
    KeyShape *KeyShape `mandatory:"true" json:"keyShape"`
    
 // Any associated data.  The string representation of the associatedData
 // must be less than 4096 characters.  Description TODO
    AssociatedData map[string]string `mandatory:"false" json:"associatedData"`
}

func (m GenerateKeyDetails) String() string {
    return common.PointerString(m)
}





