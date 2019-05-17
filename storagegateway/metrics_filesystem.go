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

// MetricsFilesystem File system-specific metrics.
type MetricsFilesystem struct {

	// Error count.
	ErrorCount *float32 `mandatory:"false" json:"errorCount"`

	// Warning count.
	WarnCount *float32 `mandatory:"false" json:"warnCount"`

	// File system cache utilization (percent).
	CacheUtilPercent *float32 `mandatory:"false" json:"cacheUtilPercent"`

	// File system cache hit (percent).
	CacheHitPercent *float32 `mandatory:"false" json:"cacheHitPercent"`

	// File system uploaded data in bytes.
	UploadedDataInBytes *float32 `mandatory:"false" json:"uploadedDataInBytes"`

	// File system downloaded data in bytes.
	DownloadedDataInBytes *float32 `mandatory:"false" json:"downloadedDataInBytes"`

	// File system write (ingestion) data in bytes.
	WriteDataInBytes *float32 `mandatory:"false" json:"writeDataInBytes"`

	// File system read data in bytes.
	ReadDataInBytes *float32 `mandatory:"false" json:"readDataInBytes"`

	// File system pending data in bytes.
	PendingDataInBytes *float32 `mandatory:"false" json:"pendingDataInBytes"`

	// Zero unless the file system is mounted; otherwise, non-zero.
	IsConnected *float32 `mandatory:"false" json:"isConnected"`

	// Zero unless the file system is in refresh mode; otherwise, non-zero.
	IsInRefreshMode *float32 `mandatory:"false" json:"isInRefreshMode"`
}

func (m MetricsFilesystem) String() string {
	return common.PointerString(m)
}
