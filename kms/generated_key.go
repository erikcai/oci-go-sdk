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


    
 // GeneratedKey The representation of GeneratedKey
type GeneratedKey struct {
    
 // The encrypted generated key.
    Ciphertext *string `mandatory:"true" json:"ciphertext"`
    
 // The plaintext generated key.
    Plaintext *string `mandatory:"false" json:"plaintext"`
    
 // Checksum of the plaintext generated key.
    PlaintextChecksum *string `mandatory:"false" json:"plaintextChecksum"`
}

func (m GeneratedKey) String() string {
    return common.PointerString(m)
}





