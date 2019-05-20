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

// MetricsStatsCache File system cache stats.
type MetricsStatsCache struct {

	// File system cache utilization (percent).
	UtilPercent *float32 `mandatory:"false" json:"utilPercent"`

	// File system hit utilization (percent).
	HitPercent *float32 `mandatory:"false" json:"hitPercent"`
}

func (m MetricsStatsCache) String() string {
	return common.PointerString(m)
}