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

// TaskRunSummaryCollection List of taskRun summaries
type TaskRunSummaryCollection struct {

	// The array of taskRun summaries
	Items []TaskRunSummary `mandatory:"true" json:"items"`
}

func (m TaskRunSummaryCollection) String() string {
	return common.PointerString(m)
}
