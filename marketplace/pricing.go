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

// Pricing The model of pricing for applications
type Pricing struct {

	// The type of pricing, either 'free' or 'not free'
	Type *string `mandatory:"false" json:"type"`

	// The name of pricing
	Name *string `mandatory:"false" json:"name"`

	// The description of pricing
	Description *string `mandatory:"false" json:"description"`
}

func (m Pricing) String() string {
	return common.PointerString(m)
}
