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

// SchemaSummaryCollection This is the collection of schema summaries, it may be a collection of lightweight details or full definitions.
type SchemaSummaryCollection struct {

	// The array of Schema summaries
	Items []SchemaSummary `mandatory:"true" json:"items"`
}

func (m SchemaSummaryCollection) String() string {
	return common.PointerString(m)
}
