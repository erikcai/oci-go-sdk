// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// StorageGateway API
//
// API for interfacing with StorageGateway
//

package storagegateway

import (
	"github.com/oracle/oci-go-sdk/common"
)

// MetricsIssues Metrics for different types of issues: errors, warnings, and rejecting IO.
// Get the total number of errors & warnings and collect the existence of rejecting IO.
// Non-zero rejecting IO count means rejecting IO starts for all the file systems in the
// storage gateway.
type MetricsIssues struct {

	// error
	Error *MetricsIssuesError `mandatory:"false" json:"error"`

	// warn
	Warn *MetricsIssuesWarn `mandatory:"false" json:"warn"`

	// rejectio
	Rejectio *MetricsIssuesRejectio `mandatory:"false" json:"rejectio"`
}

func (m MetricsIssues) String() string {
	return common.PointerString(m)
}
