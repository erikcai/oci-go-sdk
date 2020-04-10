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

// DataFlowSummaryCollection This is the collection of data flow summaries, it may be a collection of lightweight details or full definitions.
type DataFlowSummaryCollection struct {

	// The array of DataFlow summaries
	Items []DataFlowSummary `mandatory:"true" json:"items"`
}

func (m DataFlowSummaryCollection) String() string {
	return common.PointerString(m)
}
