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

// UpdateImageDetails The representation of UpdateImageDetails
type UpdateImageDetails struct {

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "foo-value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The non-unique, changeable name of the image.
	// Example: `My custom Oracle Linux image`
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Image operating system
	OperatingSystem *string `mandatory:"false" json:"operatingSystem"`

	// Image operating system version
	OperatingSystemVersion *string `mandatory:"false" json:"operatingSystemVersion"`
}

func (m UpdateImageDetails) String() string {
	return common.PointerString(m)
}
