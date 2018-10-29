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

// CategorySummary The summary of available categories
type CategorySummary struct {

	// The name of the category
	Name *string `mandatory:"false" json:"name"`

	// The code of the category
	Code *string `mandatory:"false" json:"code"`
}

func (m CategorySummary) String() string {
	return common.PointerString(m)
}
