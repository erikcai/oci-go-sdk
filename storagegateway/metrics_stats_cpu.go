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

// MetricsStatsCpu CPU stats.
type MetricsStatsCpu struct {

	// CPU utilization (percent).
	UtilPercent *float32 `mandatory:"false" json:"utilPercent"`
}

func (m MetricsStatsCpu) String() string {
	return common.PointerString(m)
}
