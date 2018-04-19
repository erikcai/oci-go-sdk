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

// Vault The representation of Vault
type Vault struct {

	// The OCID of the Compartment containing this resource.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The endpoint to perform cryptographic operations against.  TODO description.
	CryptoEndpoint *string `mandatory:"true" json:"cryptoEndpoint"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the resource.
	Id *string `mandatory:"true" json:"id"`

	// The Vault's current state.
	LifecycleState VaultLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The endpoint to perform management operations against.  TODO description.
	ManagementEndpoint *string `mandatory:"true" json:"managementEndpoint"`

	// The date and time this was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// TODO
	VaultType VaultVaultTypeEnum `mandatory:"true" json:"vaultType"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "foo-value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Simple key-value pair that is applied without any predefined name, type or scope.
	// Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`
}

func (m Vault) String() string {
	return common.PointerString(m)
}

// VaultLifecycleStateEnum Enum with underlying type: string
type VaultLifecycleStateEnum string

// Set of constants representing the allowable values for VaultLifecycleState
const (
	VaultLifecycleStateCreating VaultLifecycleStateEnum = "CREATING"
	VaultLifecycleStateActive   VaultLifecycleStateEnum = "ACTIVE"
	VaultLifecycleStateDeleting VaultLifecycleStateEnum = "DELETING"
	VaultLifecycleStateDeleted  VaultLifecycleStateEnum = "DELETED"
)

var mappingVaultLifecycleState = map[string]VaultLifecycleStateEnum{
	"CREATING": VaultLifecycleStateCreating,
	"ACTIVE":   VaultLifecycleStateActive,
	"DELETING": VaultLifecycleStateDeleting,
	"DELETED":  VaultLifecycleStateDeleted,
}

// GetVaultLifecycleStateEnumValues Enumerates the set of values for VaultLifecycleState
func GetVaultLifecycleStateEnumValues() []VaultLifecycleStateEnum {
	values := make([]VaultLifecycleStateEnum, 0)
	for _, v := range mappingVaultLifecycleState {
		values = append(values, v)
	}
	return values
}

// VaultVaultTypeEnum Enum with underlying type: string
type VaultVaultTypeEnum string

// Set of constants representing the allowable values for VaultVaultType
const (
	VaultVaultTypePrivate VaultVaultTypeEnum = "VIRTUAL_PRIVATE"
)

var mappingVaultVaultType = map[string]VaultVaultTypeEnum{
	"VIRTUAL_PRIVATE": VaultVaultTypePrivate,
}

// GetVaultVaultTypeEnumValues Enumerates the set of values for VaultVaultType
func GetVaultVaultTypeEnumValues() []VaultVaultTypeEnum {
	values := make([]VaultVaultTypeEnum, 0)
	for _, v := range mappingVaultVaultType {
		values = append(values, v)
	}
	return values
}
