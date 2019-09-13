// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Key Management Service API
//
// API for managing and performing operations with keys and vaults.
//

package keymanagement

import (
	"github.com/oracle/oci-go-sdk/common"
)

// WrappedImportKey The representation of WrappedImportKey
type WrappedImportKey struct {

	// The wrapped/encrypted key material to import. It is encrypted using RSA wrapped key and Base64 encoded.
	KeyMaterial *string `mandatory:"true" json:"keyMaterial"`

	// The wrapping mechanism to be used during key import
	WrappingAlgorithm WrappedImportKeyWrappingAlgorithmEnum `mandatory:"true" json:"wrappingAlgorithm"`
}

func (m WrappedImportKey) String() string {
	return common.PointerString(m)
}

// WrappedImportKeyWrappingAlgorithmEnum Enum with underlying type: string
type WrappedImportKeyWrappingAlgorithmEnum string

// Set of constants representing the allowable values for WrappedImportKeyWrappingAlgorithmEnum
const (
	WrappedImportKeyWrappingAlgorithmSha1   WrappedImportKeyWrappingAlgorithmEnum = "RSA_OEAP_SHA1"
	WrappedImportKeyWrappingAlgorithmSha256 WrappedImportKeyWrappingAlgorithmEnum = "RSA_OEAP_SHA256"
	WrappedImportKeyWrappingAlgorithmSha512 WrappedImportKeyWrappingAlgorithmEnum = "RSA_OEAP_SHA512"
)

var mappingWrappedImportKeyWrappingAlgorithm = map[string]WrappedImportKeyWrappingAlgorithmEnum{
	"RSA_OEAP_SHA1":   WrappedImportKeyWrappingAlgorithmSha1,
	"RSA_OEAP_SHA256": WrappedImportKeyWrappingAlgorithmSha256,
	"RSA_OEAP_SHA512": WrappedImportKeyWrappingAlgorithmSha512,
}

// GetWrappedImportKeyWrappingAlgorithmEnumValues Enumerates the set of values for WrappedImportKeyWrappingAlgorithmEnum
func GetWrappedImportKeyWrappingAlgorithmEnumValues() []WrappedImportKeyWrappingAlgorithmEnum {
	values := make([]WrappedImportKeyWrappingAlgorithmEnum, 0)
	for _, v := range mappingWrappedImportKeyWrappingAlgorithm {
		values = append(values, v)
	}
	return values
}
