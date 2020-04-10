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

// ProjectSummaryCollection This is the collection of project summaries, it may be a collection of lightweight details or full definitions.
type ProjectSummaryCollection struct {

	// The array of Project summaries
	Items []ProjectSummary `mandatory:"true" json:"items"`
}

func (m ProjectSummaryCollection) String() string {
	return common.PointerString(m)
}
