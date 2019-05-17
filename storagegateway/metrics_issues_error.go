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

// MetricsIssuesError Errors info.
type MetricsIssuesError struct {

	// Total number of errors found from the logs.
	Count *float32 `mandatory:"false" json:"count"`
}

func (m MetricsIssuesError) String() string {
	return common.PointerString(m)
}
