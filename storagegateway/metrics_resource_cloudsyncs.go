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

// MetricsResourceCloudsyncs Cloud syncs resource info.
type MetricsResourceCloudsyncs struct {

	// Max. number of CloudSyncs that can be created per storage gateway.
	MaxCount *float32 `mandatory:"false" json:"maxCount"`
}

func (m MetricsResourceCloudsyncs) String() string {
	return common.PointerString(m)
}
