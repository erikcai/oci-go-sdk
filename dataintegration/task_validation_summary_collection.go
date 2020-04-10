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

// TaskValidationSummaryCollection List of task validation summaries
type TaskValidationSummaryCollection struct {

	// The array of validation summaries
	Items []TaskValidationSummary `mandatory:"true" json:"items"`
}

func (m TaskValidationSummaryCollection) String() string {
	return common.PointerString(m)
}
