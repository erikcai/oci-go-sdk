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


    
 // DecryptedData The representation of DecryptedData
type DecryptedData struct {
    
 // The decrypted data.
    Plaintext *string `mandatory:"true" json:"plaintext"`
    
 // Checksum of the decrypted data.
    PlaintextChecksum *string `mandatory:"true" json:"plaintextChecksum"`
}

func (m DecryptedData) String() string {
    return common.PointerString(m)
}





