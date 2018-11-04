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

// MetricDataDetails A metric object containing raw metric data points to be posted to the Telemetry service.
type MetricDataDetails struct {

	// The source service or application emitting the metric.
	// Example: `oci_computeagent`
	Namespace *string `mandatory:"true" json:"namespace"`

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to use for metrics.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name of the metric.
	// Example: `CpuUtilization`
	Name *string `mandatory:"true" json:"name"`

	// A list of metric values with timestamps. At least one data point is required per call.
	Datapoints []Datapoint `mandatory:"true" json:"datapoints"`

	// Qualifiers provided in a metric definition. Available dimensions vary by metric namespace.
	// Each dimension takes the form of a key-value pair.
	// Example: `resourceId`
	Dimensions map[string]string `mandatory:"false" json:"dimensions"`

	// Properties describing metrics. These are not part of the unique fields identifying the metric.
	// Example: `"unit": "bytes"`
	Metadata map[string]string `mandatory:"false" json:"metadata"`
}

func (m MetricDataDetails) String() string {
	return common.PointerString(m)
}
