// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Telemetry API
//
// Use the Telemetry API to manage metric queries and alarms for assessing the health, capacity, and performance of your cloud resources.
// For information about metrics, see Telemetry Overview (https://docs.us-phoenix-1.oraclecloud.com/Content/Telemetry/Concepts/telemetryoverview.htm).
// For information about alarms, see Alarms Overview (https://docs.us-phoenix-1.oraclecloud.com/Content/Alarms/Concepts/alarmsoverview.htm).
//

package telemetry

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Metric The properties that define a metric.
// For more information on monitoring metrics, see Telemetry Overview (https://docs.us-phoenix-1.oraclecloud.com/Content/Telemetry/Concepts/telemetryoverview.htm).
type Metric struct {

	// The name of the metric.
	// Example: `CpuUtilization`
	Name *string `mandatory:"false" json:"name"`

	// The source service or application emitting the metric.
	// Example: `oci/compute`
	Namespace *string `mandatory:"false" json:"namespace"`

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the compartment containing
	// the resources monitored by the metric.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Qualifiers provided in a metric definition. Available dimensions vary by metric namespace.
	// Each dimension takes the form of a key-value pair.
	// Example: "resourceId": "<var>&lt;instance_OCID&gt;</var>"
	Dimensions map[string]string `mandatory:"false" json:"dimensions"`

	// The references provided in a metric definition to indicate extra information about the metric.
	// Example: `"unit": "milliseconds"`
	Metadata map[string]string `mandatory:"false" json:"metadata"`
}

func (m Metric) String() string {
	return common.PointerString(m)
}
