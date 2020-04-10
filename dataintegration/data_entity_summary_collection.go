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

// DataEntitySummaryCollection This is the collection of data entity summaries, it may be a collection of lightweight details or full definitions.
type DataEntitySummaryCollection struct {

	// The array of DataEntity summaries
	Items []DataEntitySummary `mandatory:"true" json:"items"`
}

func (m DataEntitySummaryCollection) String() string {
	return common.PointerString(m)
}
