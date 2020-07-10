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

// SummarizeDatabaseInsightResourceStatisticsAggregationCollection Returns list of the Databases with resource statistics like usage,capacity,utilization and usage change percent.
type SummarizeDatabaseInsightResourceStatisticsAggregationCollection struct {

	// Defines the type of resource metric (CPU, STORAGE, IO, MEMORY)
	ResourceMetric SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnum `mandatory:"true" json:"resourceMetric"`

	// Displays usage unit ( CORES, GB)
	UsageUnit UsageUnitEnum `mandatory:"true" json:"usageUnit"`

	// Collection of Resource Statistics items
	Items []ResourceStatisticsAggregation `mandatory:"true" json:"items"`
}

func (m SummarizeDatabaseInsightResourceStatisticsAggregationCollection) String() string {
	return common.PointerString(m)
}

// SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnum Enum with underlying type: string
type SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnum string

// Set of constants representing the allowable values for SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnum
const (
	SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricCpu     SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnum = "CPU"
	SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricStorage SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnum = "STORAGE"
	SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricIo      SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnum = "IO"
	SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricMemory  SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnum = "MEMORY"
)

var mappingSummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetric = map[string]SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnum{
	"CPU":     SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricCpu,
	"STORAGE": SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricStorage,
	"IO":      SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricIo,
	"MEMORY":  SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricMemory,
}

// GetSummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnumValues Enumerates the set of values for SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnum
func GetSummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnumValues() []SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnum {
	values := make([]SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnum, 0)
	for _, v := range mappingSummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetric {
		values = append(values, v)
	}
	return values
}