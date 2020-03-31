// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Object Storage Service API
//
// Common set of Object Storage and Archive Storage APIs for managing buckets, objects, and related resources.
// For more information, see Overview of Object Storage (https://docs.cloud.oracle.com/Content/Object/Concepts/objectstorageoverview.htm) and
// Overview of Archive Storage (https://docs.cloud.oracle.com/Content/Archive/Concepts/archivestorageoverview.htm).
//

package objectstorage

import (
	"github.com/oracle/oci-go-sdk/common"
)

// SseCustomerKeyDetails Specifies the details of the customer-provided key associated with an object.
type SseCustomerKeyDetails struct {

	// Specifies the encryption algorithm. The only supported value is "AES256".
	Algorithm SseCustomerKeyDetailsAlgorithmEnum `mandatory:"true" json:"algorithm"`

	// Specifies the base64-encoded 256-bit encryption key to use to encrypt or decrypt the object data.
	Key *string `mandatory:"true" json:"key"`

	// Specifies the base64-encoded SHA256 hash of the encryption key. This value is used to check the integrity
	// of the encryption key.
	KeySha256 *string `mandatory:"true" json:"keySha256"`
}

func (m SseCustomerKeyDetails) String() string {
	return common.PointerString(m)
}

// SseCustomerKeyDetailsAlgorithmEnum Enum with underlying type: string
type SseCustomerKeyDetailsAlgorithmEnum string

// Set of constants representing the allowable values for SseCustomerKeyDetailsAlgorithmEnum
const (
	SseCustomerKeyDetailsAlgorithmAes256 SseCustomerKeyDetailsAlgorithmEnum = "AES256"
)

var mappingSseCustomerKeyDetailsAlgorithm = map[string]SseCustomerKeyDetailsAlgorithmEnum{
	"AES256": SseCustomerKeyDetailsAlgorithmAes256,
}

// GetSseCustomerKeyDetailsAlgorithmEnumValues Enumerates the set of values for SseCustomerKeyDetailsAlgorithmEnum
func GetSseCustomerKeyDetailsAlgorithmEnumValues() []SseCustomerKeyDetailsAlgorithmEnum {
	values := make([]SseCustomerKeyDetailsAlgorithmEnum, 0)
	for _, v := range mappingSseCustomerKeyDetailsAlgorithm {
		values = append(values, v)
	}
	return values
}