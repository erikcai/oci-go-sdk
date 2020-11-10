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
	"github.com/oracle/oci-go-sdk/v28/common"
)

// MetricsResourceFilesystems File systems resource information.
type MetricsResourceFilesystems struct {

	// The maximum number of file systems that can be created per storage gateway.
	// Example: `10`
	MaxCount *float32 `mandatory:"false" json:"maxCount"`
}

func (m MetricsResourceFilesystems) String() string {
	return common.PointerString(m)
}
