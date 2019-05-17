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

// MetricsResource Storage Gateway resources.
type MetricsResource struct {

	// cpu
	Cpu *MetricsResourceCpu `mandatory:"false" json:"cpu"`

	// memory
	Memory *MetricsResourceMem `mandatory:"false" json:"memory"`

	// cache
	Cache *MetricsResourceCache `mandatory:"false" json:"cache"`

	// metadata
	Metadata *MetricsResourceMetadata `mandatory:"false" json:"metadata"`

	// log
	Log *MetricsResourceLog `mandatory:"false" json:"log"`

	// filesystems
	Filesystems *MetricsResourceFilesystems `mandatory:"false" json:"filesystems"`

	// cloudsyncs
	Cloudsyncs *MetricsResourceCloudsyncs `mandatory:"false" json:"cloudsyncs"`
}

func (m MetricsResource) String() string {
	return common.PointerString(m)
}
