// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Storage Gateway API
//
// API for the Storage Gateway service. Use this API to manage storage gateways and related items. For more
// information, see Overview of Storage Gateway (https://docs.cloud.oracle.com/iaas/Content/StorageGateway/Concepts/storagegatewayoverview.htm).
//

package storagegateway

import (
	"github.com/oracle/oci-go-sdk/common"
)

// MetricsStats Storage Gateway statistics.
type MetricsStats struct {

	// CPU statistics.
	Cpu *MetricsStatsCpu `mandatory:"false" json:"cpu"`

	// Memory statistics.
	Memory *MetricsStatsMem `mandatory:"false" json:"memory"`

	// File system cache statistics.
	Cache *MetricsStatsCache `mandatory:"false" json:"cache"`

	// Metadata storage statistics.
	Metadata *MetricsStatsMetadata `mandatory:"false" json:"metadata"`

	// Log storage statistics.
	Log *MetricsStatsLog `mandatory:"false" json:"log"`

	// File systems statistics.
	Filesystems *MetricsStatsFilesystems `mandatory:"false" json:"filesystems"`

	// Cloud syncs statistics.
	Cloudsyncs *MetricsStatsCloudsyncs `mandatory:"false" json:"cloudsyncs"`

	// Data activity statistics.
	Data *MetricsStatsData `mandatory:"false" json:"data"`
}

func (m MetricsStats) String() string {
	return common.PointerString(m)
}
