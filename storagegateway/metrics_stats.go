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

// MetricsStats Storage Gateway statistics.
type MetricsStats struct {

	// cpu
	Cpu *MetricsStatsCpu `mandatory:"false" json:"cpu"`

	// memory
	Memory *MetricsStatsMem `mandatory:"false" json:"memory"`

	// cache
	Cache *MetricsStatsCache `mandatory:"false" json:"cache"`

	// metadata
	Metadata *MetricsStatsMetadata `mandatory:"false" json:"metadata"`

	// log
	Log *MetricsStatsLog `mandatory:"false" json:"log"`

	// filesystems
	Filesystems *MetricsStatsFilesystems `mandatory:"false" json:"filesystems"`

	// cloudsyncs
	Cloudsyncs *MetricsStatsCloudsyncs `mandatory:"false" json:"cloudsyncs"`

	// data
	Data *MetricsStatsData `mandatory:"false" json:"data"`
}

func (m MetricsStats) String() string {
	return common.PointerString(m)
}
