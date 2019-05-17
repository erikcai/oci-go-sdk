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

// MetricsStatsLog Log storage stats.
type MetricsStatsLog struct {

	// Log storage utilization (percent).
	UtilPercent *float32 `mandatory:"false" json:"utilPercent"`
}

func (m MetricsStatsLog) String() string {
	return common.PointerString(m)
}
