// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Tenancy The representation of Tenancy
type Tenancy struct {

	// The id of the tenancy.
	Id *string `mandatory:"false" json:"id"`

	// The name of the tenancy.
	Name *string `mandatory:"false" json:"name"`

	// The description of the tenancy.
	Description *string `mandatory:"false" json:"description"`

	// The region key for home region.
	HomeRegionKey *string `mandatory:"false" json:"homeRegionKey"`
}

func (m Tenancy) String() string {
	return common.PointerString(m)
}
