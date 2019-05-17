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

// MetricsResourceLog Log storage resource info.
type MetricsResourceLog struct {

	// Total configured log storage size (bytes).
	SizeInBytes *float32 `mandatory:"false" json:"sizeInBytes"`
}

func (m MetricsResourceLog) String() string {
	return common.PointerString(m)
}
