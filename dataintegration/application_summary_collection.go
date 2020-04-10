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

// ApplicationSummaryCollection This is the collection of application summaries, it may be a collection of lightweight details or full definitions.
type ApplicationSummaryCollection struct {

	// The array of Application summaries
	Items []ApplicationSummary `mandatory:"true" json:"items"`
}

func (m ApplicationSummaryCollection) String() string {
	return common.PointerString(m)
}
