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

// Datapoint Metric value for a specific timestamp.
type Datapoint struct {

	// Timestamp for this metric value. Format defined by RFC3339.
	// Example: `2018-02-01T01:02:29.600Z`
	Timestamp *common.SDKTime `mandatory:"true" json:"timestamp"`

	// Numeric value of the metric.
	// Example: `10.23`
	Value *float64 `mandatory:"true" json:"value"`

	// The number of occurrences of the associated value in the set of data.
	// Optional. Default is 1.
	Count *int `mandatory:"false" json:"count"`
}

func (m Datapoint) String() string {
	return common.PointerString(m)
}
