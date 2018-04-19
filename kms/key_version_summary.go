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

// KeyVersionSummary The representation of KeyVersionSummary
type KeyVersionSummary struct {

	// The OCID of the Compartment containing this resource.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the resource.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the Key containing this resource.
	KeyId *string `mandatory:"true" json:"keyId"`

	// The date and time this was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The OCID of the Vault containing this resource.
	VaultId *string `mandatory:"true" json:"vaultId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`
}

func (m KeyVersionSummary) String() string {
	return common.PointerString(m)
}
