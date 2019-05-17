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

// MetricsStatsData Upload/download/write/read/pending data stats.
type MetricsStatsData struct {

	// Uploaded data in bytes.
	UploadedDataInBytes *float32 `mandatory:"false" json:"uploadedDataInBytes"`

	// Downloaded data in bytes.
	DownloadedDataInBytes *float32 `mandatory:"false" json:"downloadedDataInBytes"`

	// Write (ingestion) data in bytes.
	WriteDataInBytes *float32 `mandatory:"false" json:"writeDataInBytes"`

	// Read data in bytes.
	ReadDataInBytes *float32 `mandatory:"false" json:"readDataInBytes"`

	// Pending data in bytes.
	PendingDataInBytes *float32 `mandatory:"false" json:"pendingDataInBytes"`
}

func (m MetricsStatsData) String() string {
	return common.PointerString(m)
}
