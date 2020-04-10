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

// PublishedObjectSummaryCollection This is the collection of published object summaries, it may be a collection of lightweight details or full definitions.
type PublishedObjectSummaryCollection struct {

	// The array of PublishedObject summaries
	Items []PublishedObjectSummary `mandatory:"true" json:"items"`
}

func (m PublishedObjectSummaryCollection) String() string {
	return common.PointerString(m)
}
