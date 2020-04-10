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

// ReportTypeCollection A collection of report types.
type ReportTypeCollection struct {

	// Array of report types.
	Items []ReportTypeSummary `mandatory:"true" json:"items"`
}

func (m ReportTypeCollection) String() string {
	return common.PointerString(m)
}
