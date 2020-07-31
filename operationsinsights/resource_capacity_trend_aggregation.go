// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OperationsInsights API
//
// The OperationsInsights API for OPSI service.
//

package operationsinsights

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ResourceCapacityTrendAggregation Resource Capacity samples
type ResourceCapacityTrendAggregation struct {

	// The timestamp in which the current sampling period ends in RFC 3339 format.
	EndTimestamp *common.SDKTime `mandatory:"true" json:"endTimestamp"`

	// The maximum allocated amount of the resource metric type  (CPU, STORAGE).
	Capacity *float64 `mandatory:"true" json:"capacity"`

	// The base allocated amount of the resource metric type  (CPU, STORAGE).
	BaseCapacity *float64 `mandatory:"true" json:"baseCapacity"`
}

func (m ResourceCapacityTrendAggregation) String() string {
	return common.PointerString(m)
}
