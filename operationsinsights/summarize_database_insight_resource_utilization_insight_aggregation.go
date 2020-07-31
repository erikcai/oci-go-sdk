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

// SummarizeDatabaseInsightResourceUtilizationInsightAggregation Insights response containing current/projected groups for storage or CPU.
type SummarizeDatabaseInsightResourceUtilizationInsightAggregation struct {

	// Defines the type of resource metric (CPU, STORAGE)
	ResourceMetric SummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricEnum `mandatory:"true" json:"resourceMetric"`

	ProjectedUtilization *ResourceInsightProjectedUtilization `mandatory:"true" json:"projectedUtilization"`

	CurrentUtilization *ResourceInsightCurrentUtilization `mandatory:"true" json:"currentUtilization"`
}

func (m SummarizeDatabaseInsightResourceUtilizationInsightAggregation) String() string {
	return common.PointerString(m)
}

// SummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricEnum Enum with underlying type: string
type SummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricEnum string

// Set of constants representing the allowable values for SummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricEnum
const (
	SummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricCpu     SummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricEnum = "CPU"
	SummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricStorage SummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricEnum = "STORAGE"
)

var mappingSummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetric = map[string]SummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricEnum{
	"CPU":     SummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricCpu,
	"STORAGE": SummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricStorage,
}

// GetSummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricEnumValues Enumerates the set of values for SummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricEnum
func GetSummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricEnumValues() []SummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricEnum {
	values := make([]SummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricEnum, 0)
	for _, v := range mappingSummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetric {
		values = append(values, v)
	}
	return values
}
