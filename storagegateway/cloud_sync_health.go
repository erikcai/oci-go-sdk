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
	"github.com/oracle/oci-go-sdk/v31/common"
)

// CloudSyncHealth The current health of the specified cloud sync.
type CloudSyncHealth struct {

	// Metrics data about the specified cloud sync.
	Metrics *MetricsCloudsync `mandatory:"true" json:"metrics"`

	// Additional information about WARNING and CRITICAL health statuses.
	Reasons *StatusReasons `mandatory:"false" json:"reasons"`
}

func (m CloudSyncHealth) String() string {
	return common.PointerString(m)
}
