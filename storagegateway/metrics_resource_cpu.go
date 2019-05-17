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

// MetricsResourceCpu CPU resource info.
type MetricsResourceCpu struct {

	// Total number of vCPUs available in storage gateway instance.
	Count *float32 `mandatory:"false" json:"count"`
}

func (m MetricsResourceCpu) String() string {
	return common.PointerString(m)
}
