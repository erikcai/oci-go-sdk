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

// Metrics The latest telemetry data posted from the storage gateway instance. This telemetry
// data include five areas: 'resource', 'stats', 'issues', 'filesystems', and 'cloudsyncs'.
// The 'resource' contains telemetry data on storage gateway resource capacity, such
// as the number of vCPUs, memory size, etc. The 'stats' contains telemetry data on
// the resource usages, such as CPU/memory utilization, created/connected file systems
// count, etc. The 'issues' contains telemetry data on errors, warnings, and rejecting IO
// on the storage gateway. The 'filesystems' and 'cloudsyncs' respectively contain
// telemetry data on each of the file systems and cloud syncs on the storage gateway.
type Metrics struct {

	// resource
	Resource *MetricsResource `mandatory:"false" json:"resource"`

	// stats
	Stats *MetricsStats `mandatory:"false" json:"stats"`

	// issues
	Issues *MetricsIssues `mandatory:"false" json:"issues"`

	// filesystems, where each key represents a file system name.
	Filesystems map[string]MetricsFilesystem `mandatory:"false" json:"filesystems"`

	// cloudsyncs, where each key represents a cloud sync name.
	Cloudsyncs map[string]MetricsCloudsync `mandatory:"false" json:"cloudsyncs"`
}

func (m Metrics) String() string {
	return common.PointerString(m)
}
