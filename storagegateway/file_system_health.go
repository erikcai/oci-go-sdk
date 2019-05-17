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

// FileSystemHealth The current health of the file system.
type FileSystemHealth struct {

	// metrics
	Metrics *MetricsFilesystem `mandatory:"true" json:"metrics"`

	// reasons
	Reasons *StatusReasons `mandatory:"false" json:"reasons"`
}

func (m FileSystemHealth) String() string {
	return common.PointerString(m)
}
