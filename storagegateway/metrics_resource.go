// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Storage Gateway API
//
// API for the Storage Gateway service. Use this API to manage storage gateways and related items. For more
// information, see Overview of Storage Gateway (https://docs.cloud.oracle.com/iaas/Content/StorageGateway/Concepts/storagegatewayoverview.htm).
//

package storagegateway

import (
	"github.com/erikcai/oci-go-sdk/v33/common"
)

// MetricsResource Provides metrics on storage gateway resource capacity and utilization.
type MetricsResource struct {

	// CPU resource information.
	Cpu *MetricsResourceCpu `mandatory:"false" json:"cpu"`

	// Memory resource information.
	Memory *MetricsResourceMem `mandatory:"false" json:"memory"`

	// File system cache resource information.
	Cache *MetricsResourceCache `mandatory:"false" json:"cache"`

	// Metadata storage resource information.
	Metadata *MetricsResourceMetadata `mandatory:"false" json:"metadata"`

	// Log storage resource information.
	Log *MetricsResourceLog `mandatory:"false" json:"log"`

	// File systems resource information.
	Filesystems *MetricsResourceFilesystems `mandatory:"false" json:"filesystems"`

	// Cloud syncs resource information.
	Cloudsyncs *MetricsResourceCloudsyncs `mandatory:"false" json:"cloudsyncs"`
}

func (m MetricsResource) String() string {
	return common.PointerString(m)
}
