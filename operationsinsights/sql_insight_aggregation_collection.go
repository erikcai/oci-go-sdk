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

// SqlInsightAggregationCollection SQL Insights response.
type SqlInsightAggregationCollection struct {

	// The start timestamp that was passed into the request.
	TimeIntervalStart *common.SDKTime `mandatory:"true" json:"timeIntervalStart"`

	// The end timestamp that was passed into the request.
	TimeIntervalEnd *common.SDKTime `mandatory:"true" json:"timeIntervalEnd"`

	Inventory *SqlInventory `mandatory:"true" json:"inventory"`

	// List of insights.
	Items []SqlInsightAggregation `mandatory:"true" json:"items"`

	Thresholds *SqlInsightThresholds `mandatory:"true" json:"thresholds"`
}

func (m SqlInsightAggregationCollection) String() string {
	return common.PointerString(m)
}
