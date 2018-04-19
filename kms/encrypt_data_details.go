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

// EncryptDataDetails The representation of EncryptDataDetails
type EncryptDataDetails struct {

	// The OCID of the Key to encrypt with.
	KeyId *string `mandatory:"true" json:"keyId"`

	// The plaintext data to encrypt.
	Plaintext *string `mandatory:"true" json:"plaintext"`

	// Any associated data.  The string representation of the associatedData
	// must be less than 4096 characters.  Description TODO
	AssociatedData map[string]string `mandatory:"false" json:"associatedData"`
}

func (m EncryptDataDetails) String() string {
	return common.PointerString(m)
}
