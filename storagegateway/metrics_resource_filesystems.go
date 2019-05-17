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

// MetricsResourceFilesystems File systems resource info.
type MetricsResourceFilesystems struct {

	// Max. number of file systems that can be created per storage gateway.
	MaxCount *float32 `mandatory:"false" json:"maxCount"`
}

func (m MetricsResourceFilesystems) String() string {
	return common.PointerString(m)
}
