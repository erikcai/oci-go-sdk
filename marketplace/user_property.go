// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Marketplace REST API endpoint
//
// Marketplace REST API for loom plugin
//

package marketplace

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UserProperty The model of user property
type UserProperty struct {

	// Name of the user property
	Name *string `mandatory:"false" json:"name"`

	// Resource tag of the user property
	ResourceTag *string `mandatory:"false" json:"resourceTag"`

	// Value of the user property
	Value *string `mandatory:"false" json:"value"`
}

func (m UserProperty) String() string {
	return common.PointerString(m)
}
