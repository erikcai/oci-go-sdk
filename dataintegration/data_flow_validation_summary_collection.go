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

// DataFlowValidationSummaryCollection List of dataflow validation summaries
type DataFlowValidationSummaryCollection struct {

	// The array of validation summaries
	Items []DataFlowValidationSummary `mandatory:"true" json:"items"`
}

func (m DataFlowValidationSummaryCollection) String() string {
	return common.PointerString(m)
}
