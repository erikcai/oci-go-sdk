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

// ImportKeyDetails The representation of ImportKeyDetails
type ImportKeyDetails struct {

	// The OCID of the compartment that contains this key.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name for the key. It does not have to be unique, and it is changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	KeyShape *KeyShape `mandatory:"true" json:"keyShape"`

	WrappedImportKey *WrappedImportKey `mandatory:"true" json:"wrappedImportKey"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "foo-value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Simple key-value pair that is applied without any predefined name, type, or scope.
	// Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Protection mode indicates how the key is persisted and how cryptographic operations are performed using the same.
	// Using 'HSM' indicates that the key is persisted in the 'HSM' and all cryptographic operations are performed inside
	// the 'HSM'. Using 'SOFTWARE' indicates that the key is securely persisted in the service layer with the root key persisted
	// in the 'HSM' and all operations using these keys are performed at the service layer. Default protection mode is 'HSM'.
	// Once a protection mode has been selected, it cannot be modified.
	ProtectionMode ImportKeyDetailsProtectionModeEnum `mandatory:"false" json:"protectionMode,omitempty"`
}

func (m ImportKeyDetails) String() string {
	return common.PointerString(m)
}

// ImportKeyDetailsProtectionModeEnum Enum with underlying type: string
type ImportKeyDetailsProtectionModeEnum string

// Set of constants representing the allowable values for ImportKeyDetailsProtectionModeEnum
const (
	ImportKeyDetailsProtectionModeHsm      ImportKeyDetailsProtectionModeEnum = "HSM"
	ImportKeyDetailsProtectionModeSoftware ImportKeyDetailsProtectionModeEnum = "SOFTWARE"
)

var mappingImportKeyDetailsProtectionMode = map[string]ImportKeyDetailsProtectionModeEnum{
	"HSM":      ImportKeyDetailsProtectionModeHsm,
	"SOFTWARE": ImportKeyDetailsProtectionModeSoftware,
}

// GetImportKeyDetailsProtectionModeEnumValues Enumerates the set of values for ImportKeyDetailsProtectionModeEnum
func GetImportKeyDetailsProtectionModeEnumValues() []ImportKeyDetailsProtectionModeEnum {
	values := make([]ImportKeyDetailsProtectionModeEnum, 0)
	for _, v := range mappingImportKeyDetailsProtectionMode {
		values = append(values, v)
	}
	return values
}
