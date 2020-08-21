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

// SummarizeDatabaseInsightResourceUsageAggregation Resource usage summation for the current time period
type SummarizeDatabaseInsightResourceUsageAggregation struct {

	// The start timestamp that was passed into the request.
	TimeIntervalStart *common.SDKTime `mandatory:"true" json:"timeIntervalStart"`

	// The end timestamp that was passed into the request.
	TimeIntervalEnd *common.SDKTime `mandatory:"true" json:"timeIntervalEnd"`

	// Defines the type of resource metric (CPU, STORAGE)
	ResourceMetric SummarizeDatabaseInsightResourceUsageAggregationResourceMetricEnum `mandatory:"true" json:"resourceMetric"`

	// Displays usage unit (CORES, GB)
	UsageUnit UsageUnitEnum `mandatory:"true" json:"usageUnit"`

	// Total amount used of the resource metric type (CPU, STORAGE).
	Usage *float64 `mandatory:"true" json:"usage"`

	// The maximum allocated amount of the resource metric type  (CPU, STORAGE).
	Capacity *float64 `mandatory:"true" json:"capacity"`

	// Percentage change in resource usage during the current period calculated using linear regression functions
	UsageChangePercent *float64 `mandatory:"true" json:"usageChangePercent"`
}

func (m SummarizeDatabaseInsightResourceUsageAggregation) String() string {
	return common.PointerString(m)
}

// SummarizeDatabaseInsightResourceUsageAggregationResourceMetricEnum Enum with underlying type: string
type SummarizeDatabaseInsightResourceUsageAggregationResourceMetricEnum string

// Set of constants representing the allowable values for SummarizeDatabaseInsightResourceUsageAggregationResourceMetricEnum
const (
	SummarizeDatabaseInsightResourceUsageAggregationResourceMetricCpu     SummarizeDatabaseInsightResourceUsageAggregationResourceMetricEnum = "CPU"
	SummarizeDatabaseInsightResourceUsageAggregationResourceMetricStorage SummarizeDatabaseInsightResourceUsageAggregationResourceMetricEnum = "STORAGE"
)

var mappingSummarizeDatabaseInsightResourceUsageAggregationResourceMetric = map[string]SummarizeDatabaseInsightResourceUsageAggregationResourceMetricEnum{
	"CPU":     SummarizeDatabaseInsightResourceUsageAggregationResourceMetricCpu,
	"STORAGE": SummarizeDatabaseInsightResourceUsageAggregationResourceMetricStorage,
}

// GetSummarizeDatabaseInsightResourceUsageAggregationResourceMetricEnumValues Enumerates the set of values for SummarizeDatabaseInsightResourceUsageAggregationResourceMetricEnum
func GetSummarizeDatabaseInsightResourceUsageAggregationResourceMetricEnumValues() []SummarizeDatabaseInsightResourceUsageAggregationResourceMetricEnum {
	values := make([]SummarizeDatabaseInsightResourceUsageAggregationResourceMetricEnum, 0)
	for _, v := range mappingSummarizeDatabaseInsightResourceUsageAggregationResourceMetric {
		values = append(values, v)
	}
	return values
}
