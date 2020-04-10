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

// ConnectionSummaryCollection This is the collection of connection summaries, it may be a collection of lightweight details or full definitions.
type ConnectionSummaryCollection struct {

	// The array of Connection summaries
	Items []ConnectionSummary `mandatory:"true" json:"items"`
}

func (m ConnectionSummaryCollection) String() string {
	return common.PointerString(m)
}
