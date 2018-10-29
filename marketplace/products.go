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

// Products The model of products
type Products struct {

	// Code of the product
	Code *string `mandatory:"false" json:"code"`

	// Name of the product
	Name *string `mandatory:"false" json:"name"`

	// Categories of the product
	Categories []CategorySummary `mandatory:"false" json:"categories"`

	// Filters of the product
	Filters []ProductFilter `mandatory:"false" json:"filters"`

	// Group the product belongs to
	Group *string `mandatory:"false" json:"group"`
}

func (m Products) String() string {
	return common.PointerString(m)
}
