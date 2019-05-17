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

// MetricsCloudsync cloud sync-specific metrics.
type MetricsCloudsync struct {

	// Error count.
	ErrorCount *float32 `mandatory:"false" json:"errorCount"`

	// Warning count.
	WarnCount *float32 `mandatory:"false" json:"warnCount"`

	// Cloud sync cache utilization (percent).
	CacheUtilPercent *float32 `mandatory:"false" json:"cacheUtilPercent"`

	// Cloud sync cache hit (percent).
	CacheHitPercent *float32 `mandatory:"false" json:"cacheHitPercent"`

	// Cloud sync uploaded data in bytes.
	UploadedDataInBytes *float32 `mandatory:"false" json:"uploadedDataInBytes"`

	// Cloud sync downloaded data in bytes.
	DownloadedDataInBytes *float32 `mandatory:"false" json:"downloadedDataInBytes"`

	// Cloud sync write (ingestion) data in bytes.
	WriteDataInBytes *float32 `mandatory:"false" json:"writeDataInBytes"`

	// Cloud sync read data in bytes.
	ReadDataInBytes *float32 `mandatory:"false" json:"readDataInBytes"`

	// Cloud sync pending data in bytes.
	PendingDataInBytes *float32 `mandatory:"false" json:"pendingDataInBytes"`

	// Source data size in bytes.
	SourceDataInBytes *float32 `mandatory:"false" json:"sourceDataInBytes"`
}

func (m MetricsCloudsync) String() string {
	return common.PointerString(m)
}
