// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Vault Service Key Management API
//
// API for managing and performing operations with keys and vaults. (For the API for managing secrets, see the Vault Service
// Secret Management API. For the API for retrieving secrets, see the Vault Service Secret Retrieval API.)
//

package keymanagement

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ExportedKeyData Response of export Key query
type ExportedKeyData struct {

	// The OCID of the key version.
	KeyVersionId *string `mandatory:"true" json:"keyVersionId"`

	// The OCID of the key associated with this key version.
	KeyId *string `mandatory:"true" json:"keyId"`

	// The date and time this key version was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The OCID of the vault that contains this key version.
	VaultId *string `mandatory:"true" json:"vaultId"`

	// The wrapped/encrypted key in Base64 format. It is encrypted using RSA wrapped key based on algorithm provided in request.
	EncryptedKey *string `mandatory:"true" json:"encryptedKey"`

	// Encryption algorithm to be used to encrypt the software key. RSA_OAEP_AES_SHA256 will generate a temporary AES key
	// which is wrapped by RSA key provided in input. And the AES key would be used to wrap the actual software key.
	// In RSA_OAEP_SHA256, the software key will be wrapped using RSA key provided in input.
	Algorithm ExportedKeyDataAlgorithmEnum `mandatory:"true" json:"algorithm"`
}

func (m ExportedKeyData) String() string {
	return common.PointerString(m)
}

// ExportedKeyDataAlgorithmEnum Enum with underlying type: string
type ExportedKeyDataAlgorithmEnum string

// Set of constants representing the allowable values for ExportedKeyDataAlgorithmEnum
const (
	ExportedKeyDataAlgorithmAesSha256 ExportedKeyDataAlgorithmEnum = "RSA_OAEP_AES_SHA256"
	ExportedKeyDataAlgorithmSha256    ExportedKeyDataAlgorithmEnum = "RSA_OAEP_SHA256"
)

var mappingExportedKeyDataAlgorithm = map[string]ExportedKeyDataAlgorithmEnum{
	"RSA_OAEP_AES_SHA256": ExportedKeyDataAlgorithmAesSha256,
	"RSA_OAEP_SHA256":     ExportedKeyDataAlgorithmSha256,
}

// GetExportedKeyDataAlgorithmEnumValues Enumerates the set of values for ExportedKeyDataAlgorithmEnum
func GetExportedKeyDataAlgorithmEnumValues() []ExportedKeyDataAlgorithmEnum {
	values := make([]ExportedKeyDataAlgorithmEnum, 0)
	for _, v := range mappingExportedKeyDataAlgorithm {
		values = append(values, v)
	}
	return values
}
