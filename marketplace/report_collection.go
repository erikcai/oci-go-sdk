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

// ReportCollection A collection of reports.
type ReportCollection struct {

	// Array of reports.
	Items []ReportSummary `mandatory:"true" json:"items"`
}

func (m ReportCollection) String() string {
	return common.PointerString(m)
}
