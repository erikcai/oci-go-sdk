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

// AlarmSummary A summary of properties for the specified alarm.
// For more information, see Alarms Overview (https://docs.us-phoenix-1.oraclecloud.com/iaas/Content/Alarms/Concepts/alarmsoverview.htm).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.us-phoenix-1.oraclecloud.com/iaas/Content/Identity/Concepts/policygetstarted.htm).
// For information about endpoints and signing API requests, see
// About the API (https://docs.us-phoenix-1.oraclecloud.com/iaas/Content/API/Concepts/usingapi.htm). For information about available SDKs and tools, see
// SDKS and Other Tools (https://docs.us-phoenix-1.oraclecloud.com/iaas/Content/API/Concepts/sdks.htm).
type AlarmSummary struct {

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/iaas/Content/General/Concepts/identifiers.htm) of the alarm.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name for the alarm. It does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// This name is sent as the title for notifications related to this alarm.
	// Example: `High CPU Utilization`
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the alarm.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the metric
	// being evaluated by the alarm.
	MetricCompartmentId *string `mandatory:"true" json:"metricCompartmentId"`

	// The Telemetry Query Language (TQL) expression to evaluate for the alarm. The Alarms service
	// interprets results for each returned time series as Boolean values, where zero represents false
	// and a non-zero value represents true. A true value means that the trigger rule condition has
	// been met. The query must specify a metric, statistic, interval, and trigger rule (threshold or
	// absence). Supported values for interval: `1m`, `5m`, `1h`. You can optionally specify dimensions
	// and grouping functions. Supported grouping functions: `grouping()`, `groupBy()`.
	// For available dimensions, review the metric definition.
	// Example of threshold alarm:
	//   -----
	//     `CpuUtilization[1m]{availabilityDomain="cumS:PHX-AD-1"}.groupBy(availabilityDomain).percentile(0.9) > 85`
	//   -----
	// Example of absence alarm:
	//   -----
	//     `CpuUtilization[1m]{availabilityDomain="cumS:PHX-AD-1"}.absent()`
	//   -----
	Query *string `mandatory:"true" json:"query"`

	// The perceived severity of the alarm with regard to the affected system.
	// Example: `CRITICAL`
	Severity AlarmSummarySeverityEnum `mandatory:"true" json:"severity"`

	// An array of OCIDs (https://docs.us-phoenix-1.oraclecloud.com/iaas/Content/General/Concepts/identifiers.htm) to which the notifications for
	// this alarm will be delivered. An example destination is an OCID for a topic managed by the
	// Oracle Cloud Infrastructure Notification service.
	Destinations []string `mandatory:"true" json:"destinations"`

	// Whether the alarm is enabled.
	// Example: `true`
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// The source service or application emitting the metric that is evaluated by the alarm. If not specified, then all available sources are used.
	// Example: `oci_computeagent`
	Namespace *string `mandatory:"false" json:"namespace"`

	Suppression *Suppression `mandatory:"false" json:"suppression"`
}

func (m AlarmSummary) String() string {
	return common.PointerString(m)
}

// AlarmSummarySeverityEnum Enum with underlying type: string
type AlarmSummarySeverityEnum string

// Set of constants representing the allowable values for AlarmSummarySeverityEnum
const (
	AlarmSummarySeverityCritical AlarmSummarySeverityEnum = "CRITICAL"
	AlarmSummarySeverityError    AlarmSummarySeverityEnum = "ERROR"
	AlarmSummarySeverityWarning  AlarmSummarySeverityEnum = "WARNING"
	AlarmSummarySeverityInfo     AlarmSummarySeverityEnum = "INFO"
)

var mappingAlarmSummarySeverity = map[string]AlarmSummarySeverityEnum{
	"CRITICAL": AlarmSummarySeverityCritical,
	"ERROR":    AlarmSummarySeverityError,
	"WARNING":  AlarmSummarySeverityWarning,
	"INFO":     AlarmSummarySeverityInfo,
}

// GetAlarmSummarySeverityEnumValues Enumerates the set of values for AlarmSummarySeverityEnum
func GetAlarmSummarySeverityEnumValues() []AlarmSummarySeverityEnum {
	values := make([]AlarmSummarySeverityEnum, 0)
	for _, v := range mappingAlarmSummarySeverity {
		values = append(values, v)
	}
	return values
}
