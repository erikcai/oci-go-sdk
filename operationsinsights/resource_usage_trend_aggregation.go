// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OperationsInsights API
//
// The OperationsInsights API for OPSI service.
//

package operationsinsights

import (
	"github.com/oracle/oci-go-sdk/v25/common"
)

// ResourceUsageTrendAggregation Aggregate usage samples
type ResourceUsageTrendAggregation struct {

	// The timestamp in which the current sampling period ends in RFC 3339 format.
	EndTimestamp *common.SDKTime `mandatory:"true" json:"endTimestamp"`

	// Total amount used of the resource metric type (CPU, STORAGE).
	Usage *float64 `mandatory:"true" json:"usage"`

	// The maximum allocated amount of the resource metric type  (CPU, STORAGE).
	Capacity *float64 `mandatory:"true" json:"capacity"`
}

func (m ResourceUsageTrendAggregation) String() string {
	return common.PointerString(m)
}
