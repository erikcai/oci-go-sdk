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

// TaskSummaryCollection This is the collection of task summaries, it may be a collection of lightweight details or full definitions.
type TaskSummaryCollection struct {

	// The array of Task summaries
	Items []TaskSummary `mandatory:"true" json:"items"`
}

func (m TaskSummaryCollection) String() string {
	return common.PointerString(m)
}
