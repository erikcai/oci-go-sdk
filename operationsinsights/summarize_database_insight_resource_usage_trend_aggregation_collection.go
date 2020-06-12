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

// SummarizeDatabaseInsightResourceUsageTrendAggregationCollection Top level response object.
type SummarizeDatabaseInsightResourceUsageTrendAggregationCollection struct {

	// Defines the type of resource metric (CPU, STORAGE, IO, MEMORY)
	ResourceMetric SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricEnum `mandatory:"true" json:"resourceMetric"`

	// Displays usage unit ( CORES, GB)
	UsageUnit UsageUnitEnum `mandatory:"true" json:"usageUnit"`

	// Time duration in milliseconds between data points (one hour or one day).
	ItemDurationInMs *int64 `mandatory:"true" json:"itemDurationInMs"`

	// Usage Data with time stamps
	UsageData []ResourceUsageTrendAggregation `mandatory:"true" json:"usageData"`

	// The start timestamp that was passed into the request.
	TimeIntervalStart *common.SDKTime `mandatory:"false" json:"timeIntervalStart"`

	// The end timestamp that was passed into the request.
	TimeIntervalEnd *common.SDKTime `mandatory:"false" json:"timeIntervalEnd"`
}

func (m SummarizeDatabaseInsightResourceUsageTrendAggregationCollection) String() string {
	return common.PointerString(m)
}

// SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricEnum Enum with underlying type: string
type SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricEnum string

// Set of constants representing the allowable values for SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricEnum
const (
	SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricCpu     SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricEnum = "CPU"
	SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricStorage SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricEnum = "STORAGE"
	SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricIo      SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricEnum = "IO"
	SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricMemory  SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricEnum = "MEMORY"
)

var mappingSummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetric = map[string]SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricEnum{
	"CPU":     SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricCpu,
	"STORAGE": SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricStorage,
	"IO":      SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricIo,
	"MEMORY":  SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricMemory,
}

// GetSummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricEnumValues Enumerates the set of values for SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricEnum
func GetSummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricEnumValues() []SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricEnum {
	values := make([]SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricEnum, 0)
	for _, v := range mappingSummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetric {
		values = append(values, v)
	}
	return values
}
