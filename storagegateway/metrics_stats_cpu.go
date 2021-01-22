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

// MetricsStatsCpu CPU statistics.
type MetricsStatsCpu struct {

	// CPU utilization (percent).
	// Example: `60`
	UtilPercent *float32 `mandatory:"false" json:"utilPercent"`
}

func (m MetricsStatsCpu) String() string {
	return common.PointerString(m)
}
