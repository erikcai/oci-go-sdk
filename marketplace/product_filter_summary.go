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

// ProductFilterSummary The summary of the product filters for the search bar which is in the response
type ProductFilterSummary struct {

	// Name of the filter
	Name *string `mandatory:"false" json:"name"`

	// Code of the filter
	Code *string `mandatory:"false" json:"code"`

	// Instruction of the filter
	UsageInstructions *string `mandatory:"false" json:"usageInstructions"`

	// Whether the filter is multiple selectable
	IsMultiSelect *bool `mandatory:"false" json:"isMultiSelect"`

	// Default value of the filter
	DefaultValue *Item `mandatory:"false" json:"defaultValue"`

	// Values of the filter
	Values []Item `mandatory:"false" json:"values"`
}

func (m ProductFilterSummary) String() string {
	return common.PointerString(m)
}