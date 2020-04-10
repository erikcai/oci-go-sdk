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

// PatchSummaryCollection This is the collection of patch summaries, it may be a collection of lightweight details or full definitions.
type PatchSummaryCollection struct {

	// The array of patch summaries
	Items []PatchSummary `mandatory:"true" json:"items"`
}

func (m PatchSummaryCollection) String() string {
	return common.PointerString(m)
}
