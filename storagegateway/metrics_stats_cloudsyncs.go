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

// MetricsStatsCloudsyncs Cloud syncs stats.
type MetricsStatsCloudsyncs struct {

	// Total number of cloud syncs.
	Count *float32 `mandatory:"false" json:"count"`

	// Number of running cloud syncs.
	RunningCount *float32 `mandatory:"false" json:"runningCount"`

	// Number of created cloud syncs.
	CreatedCount *float32 `mandatory:"false" json:"createdCount"`

	// Number of completed cloud syncs.
	CompletedCount *float32 `mandatory:"false" json:"completedCount"`

	// Number of canceled cloud syncs.
	CanceledCount *float32 `mandatory:"false" json:"canceledCount"`

	// Number of failed cloud syncs.
	FailedCount *float32 `mandatory:"false" json:"failedCount"`

	// Number of upload cloud syncs.
	UploadCount *float32 `mandatory:"false" json:"uploadCount"`

	// Number of download cloud syncs.
	DownloadCount *float32 `mandatory:"false" json:"downloadCount"`
}

func (m MetricsStatsCloudsyncs) String() string {
	return common.PointerString(m)
}
