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

// CloudSyncHealth The current health of the cloud sync.
type CloudSyncHealth struct {

	// metrics
	Metrics *MetricsCloudsync `mandatory:"true" json:"metrics"`

	// reasons
	Reasons *StatusReasons `mandatory:"false" json:"reasons"`
}

func (m CloudSyncHealth) String() string {
	return common.PointerString(m)
}
