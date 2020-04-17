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

// ExportKeyDetails Parameters to fetch encrypted Key.
type ExportKeyDetails struct {

	// The OCID of the key associated with key version.
	KeyId *string `mandatory:"true" json:"keyId"`

	// Encryption algorithm to be used to encrypt the software key. RSA_OAEP_AES_SHA256 will generate a temporary AES key
	// which is wrapped by RSA key provided in input. And the AES key would be used to wrap the actual software key.
	// In RSA_OAEP_SHA256, the software key will be wrapped using RSA key provided in input.
	Algorithm ExportKeyDetailsAlgorithmEnum `mandatory:"true" json:"algorithm"`

	// PEM format of RSA2048/RSA3072/RSA4096 Public Key, to be used to encrypt the key.
	PublicKey *string `mandatory:"true" json:"publicKey"`

	// The OCID of the key version that has to be exported. If not supplied, the current key version
	// would be considered for export.
	KeyVersionId *string `mandatory:"false" json:"keyVersionId"`

	// Information that provides context for audit logging. You can provide this additional
	// data as key-value pairs to include in the audit logs when audit logging is enabled.
	LoggingContext map[string]string `mandatory:"false" json:"loggingContext"`
}

func (m ExportKeyDetails) String() string {
	return common.PointerString(m)
}

// ExportKeyDetailsAlgorithmEnum Enum with underlying type: string
type ExportKeyDetailsAlgorithmEnum string

// Set of constants representing the allowable values for ExportKeyDetailsAlgorithmEnum
const (
	ExportKeyDetailsAlgorithmAesSha256 ExportKeyDetailsAlgorithmEnum = "RSA_OAEP_AES_SHA256"
	ExportKeyDetailsAlgorithmSha256    ExportKeyDetailsAlgorithmEnum = "RSA_OAEP_SHA256"
)

var mappingExportKeyDetailsAlgorithm = map[string]ExportKeyDetailsAlgorithmEnum{
	"RSA_OAEP_AES_SHA256": ExportKeyDetailsAlgorithmAesSha256,
	"RSA_OAEP_SHA256":     ExportKeyDetailsAlgorithmSha256,
}

// GetExportKeyDetailsAlgorithmEnumValues Enumerates the set of values for ExportKeyDetailsAlgorithmEnum
func GetExportKeyDetailsAlgorithmEnumValues() []ExportKeyDetailsAlgorithmEnum {
	values := make([]ExportKeyDetailsAlgorithmEnum, 0)
	for _, v := range mappingExportKeyDetailsAlgorithm {
		values = append(values, v)
	}
	return values
}
