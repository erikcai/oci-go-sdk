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
	"github.com/oracle/oci-go-sdk/v27/common"
)

// StatusReasons The reasons for the overall health status. The object can include an array of reason strings for the 'CRITICAL'
// health status and another array of reason strings for the 'WARNING' health status. The 'CRITICAL' health status
// includes both arrays if a warning also exists.
type StatusReasons struct {

	// An array of reasons for the critical status.
	// Example: `Rejecting IO due to low cache space`
	Critical []string `mandatory:"false" json:"critical"`

	// An array of reasons for the warning status.
	// Example: `A newer version is available`
	Warning []string `mandatory:"false" json:"warning"`
}

func (m StatusReasons) String() string {
	return common.PointerString(m)
}
