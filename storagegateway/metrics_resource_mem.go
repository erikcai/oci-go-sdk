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

// MetricsResourceMem Memory resource info.
type MetricsResourceMem struct {

	// Total memory available (bytes).
	SizeInBytes *float32 `mandatory:"false" json:"sizeInBytes"`
}

func (m MetricsResourceMem) String() string {
	return common.PointerString(m)
}
