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

// MetricsIssuesWarn Warnings info.
type MetricsIssuesWarn struct {

	// Total number of warnings found from the logs.
	Count *float32 `mandatory:"false" json:"count"`
}

func (m MetricsIssuesWarn) String() string {
	return common.PointerString(m)
}
