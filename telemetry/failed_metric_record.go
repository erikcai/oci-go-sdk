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

// FailedMetricRecord The record of a single metric object that failed input validation and the reason for the failure.
type FailedMetricRecord struct {

	// An error message indicating the reason that the indicated metric object failed input validation.
	Message *string `mandatory:"true" json:"message"`

	// Identifier of a metric object that failed input validation.
	MetricData *MetricDataDetails `mandatory:"true" json:"metricData"`
}

func (m FailedMetricRecord) String() string {
	return common.PointerString(m)
}
