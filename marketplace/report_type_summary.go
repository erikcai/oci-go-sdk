// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Manage applications in Oracle Cloud Infrastructure Marketplace.
//

package marketplace

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ReportTypeSummary The model of the description of a report.
type ReportTypeSummary struct {

	// The type of the report.
	ReportType *string `mandatory:"false" json:"reportType"`

	// The name of the report.
	Name *string `mandatory:"false" json:"name"`

	// A description of the report.
	Description *string `mandatory:"false" json:"description"`

	// The columns in the report.
	Columns []string `mandatory:"false" json:"columns"`
}

func (m ReportTypeSummary) String() string {
	return common.PointerString(m)
}
