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

// CreateKeyDetails The representation of CreateKeyDetails
type CreateKeyDetails struct {

	// The OCID of the compartment that contains this master encryption key.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name for the key. It does not have to be unique, and it is changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	KeyShape *KeyShape `mandatory:"true" json:"keyShape"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Protection mode indicates how the key is persisted and how cryptographic operations are performed using the same.
	// Using 'HSM' indicates that the key is persisted in the 'HSM' and all cryptographic operations are performed inside
	// the 'HSM'. Using 'SOFTWARE' indicates that the key is securely persisted in the service layer with the root key persisted
	// in the 'HSM' and all operations using these keys are performed at the service layer. Default protection mode is 'HSM'.
	// Once a protection mode has been selected, it cannot be modified.
	ProtectionMode CreateKeyDetailsProtectionModeEnum `mandatory:"false" json:"protectionMode,omitempty"`
}

func (m CreateKeyDetails) String() string {
	return common.PointerString(m)
}

// CreateKeyDetailsProtectionModeEnum Enum with underlying type: string
type CreateKeyDetailsProtectionModeEnum string

// Set of constants representing the allowable values for CreateKeyDetailsProtectionModeEnum
const (
	CreateKeyDetailsProtectionModeHsm      CreateKeyDetailsProtectionModeEnum = "HSM"
	CreateKeyDetailsProtectionModeSoftware CreateKeyDetailsProtectionModeEnum = "SOFTWARE"
)

var mappingCreateKeyDetailsProtectionMode = map[string]CreateKeyDetailsProtectionModeEnum{
	"HSM":      CreateKeyDetailsProtectionModeHsm,
	"SOFTWARE": CreateKeyDetailsProtectionModeSoftware,
}

// GetCreateKeyDetailsProtectionModeEnumValues Enumerates the set of values for CreateKeyDetailsProtectionModeEnum
func GetCreateKeyDetailsProtectionModeEnumValues() []CreateKeyDetailsProtectionModeEnum {
	values := make([]CreateKeyDetailsProtectionModeEnum, 0)
	for _, v := range mappingCreateKeyDetailsProtectionMode {
		values = append(values, v)
	}
	return values
}
