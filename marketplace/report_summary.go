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

// ReportSummary The model of a single report.
type ReportSummary struct {

	// The type of the report.
	ReportType *string `mandatory:"true" json:"reportType"`

	// Date in YYYYMMDD format.
	Date *string `mandatory:"true" json:"date"`

	// The columns in the report content.
	Columns []string `mandatory:"true" json:"columns"`

	// The content of report in CSV string format.
	Content *string `mandatory:"true" json:"content"`
}

func (m ReportSummary) String() string {
	return common.PointerString(m)
}
