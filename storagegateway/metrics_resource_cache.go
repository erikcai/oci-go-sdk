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

// MetricsResourceCache File system cache resource info.
type MetricsResourceCache struct {

	// Total configured file system cache size (bytes).
	SizeInBytes *float32 `mandatory:"false" json:"sizeInBytes"`
}

func (m MetricsResourceCache) String() string {
	return common.PointerString(m)
}
