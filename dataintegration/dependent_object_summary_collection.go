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

// DependentObjectSummaryCollection List of DependentObject summaries
type DependentObjectSummaryCollection struct {

	// The array of DependentObject summaries
	Items []DependentObjectSummary `mandatory:"true" json:"items"`
}

func (m DependentObjectSummaryCollection) String() string {
	return common.PointerString(m)
}
