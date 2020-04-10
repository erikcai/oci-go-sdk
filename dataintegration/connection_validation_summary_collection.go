// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Data Integration Service API Specification
//
// Data Integration Service API Specification
//

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ConnectionValidationSummaryCollection List of connection validation summaries
type ConnectionValidationSummaryCollection struct {

	// The array of validation summaries
	Items []ConnectionValidationSummary `mandatory:"true" json:"items"`
}

func (m ConnectionValidationSummaryCollection) String() string {
	return common.PointerString(m)
}
