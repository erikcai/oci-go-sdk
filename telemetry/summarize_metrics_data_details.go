// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Telemetry API
//
// Use the Telemetry API to manage metric queries and alarms for assessing the health, capacity, and performance of your cloud resources.
// For information about metrics, see Telemetry Overview (https://docs.us-phoenix-1.oraclecloud.com/iaas/Content/Telemetry/Concepts/telemetryoverview.htm).
// For information about alarms, see Alarms Overview (https://docs.us-phoenix-1.oraclecloud.com/iaas/Content/Alarms/Concepts/alarmsoverview.htm).
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
	// Example: `oci_computeagent`
	Namespace *string `mandatory:"true" json:"namespace"`

	// The Telemetry Query Language (TQL) expression to use when searching for metric data points to
	// aggregate. The query must specify a metric, statistic, and interval. Supported values for
	// interval: `1m`-`60m` (also `1h`). You can optionally specify dimensions and grouping functions.
	// Supported grouping functions: `grouping()`, `groupBy()`.
	// For available dimensions, review the metric definition.
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

	// The time between calculated aggregation windows. Use with the query interval to vary the
	// frequency at which aggregated data points are returned. For example, use a query interval of
	// 5 minutes with a resolution of 1 minute to retrieve five-minute aggregations at a one-minute
	// frequency. The resolution must be equal or less than the interval in the query. The default
	// resolution is 1m (one minute). Supported values: `1m`-`60m` (also `1h`).
	// Example: `5m`
	Resolution *string `mandatory:"false" json:"resolution"`
}

func (m SummarizeMetricsDataDetails) String() string {
	return common.PointerString(m)
}
