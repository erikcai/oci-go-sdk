// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateInstanceConsoleConnectionDetails The details for creating a instance console connection.
// The instance console connection is created in the same compartment as the instance.
type CreateInstanceConsoleConnectionDetails struct {

	// The OCID of the instance to create the console connection to.
	InstanceId *string `mandatory:"true" json:"instanceId"`

	// The SSH public key used to authenticate the console connection.
	PublicKey *string `mandatory:"true" json:"publicKey"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "foo-value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`
}

func (m CreateInstanceConsoleConnectionDetails) String() string {
	return common.PointerString(m)
}
