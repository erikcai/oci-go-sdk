// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Telemetry Service API
//
// Use the Telemetry Service API to access metric data about the health, capacity, and performance of your cloud resources.
// For more information on monitoring metrics, see Monitoring Overview (https://docs.us-phoenix-1.oraclecloud.com/Content/Monitoring/Concepts/overview.htm).
//

package telemetry

import (
	"github.com/oracle/oci-go-sdk/common"
)

// SummarizeMetricsDataDetails The request details for retrieving aggregated data.
// Use the query and optional properties to filter the returned results.
// Maximum data points returned per call: 10,080
type SummarizeMetricsDataDetails struct {

	// The source service or application to use when searching for metric data points to aggregate.
	// Example: `oci/compute`
	Namespace *string `mandatory:"true" json:"namespace"`

	// The query expression to use when searching for metric datapoints to aggregate.
	// The query must specify a metric, an aggregation function, and an interval. You can optionally specify
	// dimensions and metadata that exist in the metric definition. Responses include the metric name, its source compartment and namespace.
	// Example: `CpuUtilization[1m].sum()`
	Query *string `mandatory:"true" json:"query"`

	// The beginning of the time range to use when searching for metric data points.
	// Format is defined by RFC3339. The response includes metric data points for the startTime.
	// Default value: the timestamp 3 hours before the call was sent.
	// Example: `2018-02-01T01:02:29.600Z`
	StartTime *common.SDKTime `mandatory:"false" json:"startTime"`

	// The end of the time range to use when searching for metric data points.
	// Format is defined by RFC3339. The response excludes metric data points for the endTime.
	// Default value: the timestamp representing when the call was sent.
	// Example: `2018-02-01T02:02:29.600Z`
	EndTime *common.SDKTime `mandatory:"false" json:"endTime"`

	// The time between calculated aggregation windows. Use with the query interval to vary the frequency
	// at which aggregated data points are returned. For example, use a query interval of 5 minutes with
	// a resolution of 1 minute to retrieve five-minute aggregations at a one-minute frequency.
	// Example: `1m`
	Resolution *string `mandatory:"false" json:"resolution"`
}

func (m SummarizeMetricsDataDetails) String() string {
	return common.PointerString(m)
}
