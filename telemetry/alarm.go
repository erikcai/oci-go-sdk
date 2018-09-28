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

// Alarm The properties that define an alarm.
// For more information, see Alarms Overview (https://docs.us-phoenix-1.oraclecloud.com/iaas/Content/Alarms/Concepts/alarmsoverview.htm).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.us-phoenix-1.oraclecloud.com/iaas/Content/Identity/Concepts/policygetstarted.htm).
// For information about endpoints and signing API requests, see
// About the API (https://docs.us-phoenix-1.oraclecloud.com/iaas/Content/API/Concepts/usingapi.htm). For information about available SDKs and tools, see
// SDKS and Other Tools (https://docs.us-phoenix-1.oraclecloud.com/iaas/Content/API/Concepts/sdks.htm).
type Alarm struct {

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

	// The perceived type of response required when the alarm is in the "FIRING" state.
	// Example: `CRITICAL`
	Severity AlarmSeverityEnum `mandatory:"true" json:"severity"`

	// An array of OCIDs (https://docs.us-phoenix-1.oraclecloud.com/iaas/Content/General/Concepts/identifiers.htm) to which the notifications for
	// this alarm will be delivered. An example destination is an OCID for a topic managed by the
	// Oracle Cloud Infrastructure Notification service.
	Destinations []string `mandatory:"true" json:"destinations"`

	// Whether the alarm is enabled.
	// Example: `true`
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// The date and time the alarm was created. Format defined by RFC3339.
	// Example: `2018-02-01T01:02:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the alarm was last updated. Format defined by RFC3339.
	// Example: `2018-02-03T01:02:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The source service or application emitting the metric that is evaluated by the alarm. If not specified, then all available sources are used.
	// Example: `oci_computeagent`
	Namespace *string `mandatory:"false" json:"namespace"`

	// The time between calculated aggregation windows. Use with the query interval to vary the
	// frequency at which aggregated data points are returned. For example, use a query interval of
	// 5 minutes with a resolution of 1 minute to retrieve five-minute aggregations at a one-minute
	// frequency. The resolution must be equal or less than the interval in the query. The default
	// resolution is 1m (one minute). Supported values: `1m`, `5m`, `1h`.
	// Example: `5m`
	Resolution *string `mandatory:"false" json:"resolution"`

	// The period of time that the condition defined in the alarm must persist before the alarm
	// state changes from "OK" to "FIRING" or vice versa. For example, a value of 5 minutes means
	// that five minutes of the alarm firing is required before the alarm can update its state to
	// "FIRING" and five minutes of the alarm not firing is required before the alarm can update
	// its state to "OK."
	// The duration is specified as a string in ISO 8601 format (`PT10M` for ten minutes or `PT1H`
	// for one hour). Minimum: PT1M. Maximum: PT1H. Default: PT1M.
	// Under the default value of PT1M, the first evaluation that breaches the alarm updates the
	// state to "FIRING" and the first evaluation that does not breach the alarm updates the state
	// to "OK".
	// Example: `PT5M`
	PendingDuration *string `mandatory:"false" json:"pendingDuration"`

	// The human-readable content of the notification delivered. Oracle recommends providing guidance
	// to operators for resolving the alarm condition. Consider adding links to standard runbook
	// practices.
	// Example: `High CPU usage alert. Follow runbook instructions for resolution.`
	Body *string `mandatory:"false" json:"body"`

	// The frequency at which notifications are re-submitted, if the alarm keeps firing without
	// interruption. Format defined by ISO 8601. For example, `PT4H` indicates four hours.
	// Minimum: PT1M. Maximum: P30D.
	// Default value: null (notifications are not re-submitted).
	// Example: `PT2H`
	RepeatNotificationDuration *string `mandatory:"false" json:"repeatNotificationDuration"`

	Suppression *Suppression `mandatory:"false" json:"suppression"`
}

func (m Alarm) String() string {
	return common.PointerString(m)
}

// AlarmSeverityEnum Enum with underlying type: string
type AlarmSeverityEnum string

// Set of constants representing the allowable values for AlarmSeverity
const (
	AlarmSeverityCritical AlarmSeverityEnum = "CRITICAL"
	AlarmSeverityError    AlarmSeverityEnum = "ERROR"
	AlarmSeverityWarning  AlarmSeverityEnum = "WARNING"
	AlarmSeverityInfo     AlarmSeverityEnum = "INFO"
)

var mappingAlarmSeverity = map[string]AlarmSeverityEnum{
	"CRITICAL": AlarmSeverityCritical,
	"ERROR":    AlarmSeverityError,
	"WARNING":  AlarmSeverityWarning,
	"INFO":     AlarmSeverityInfo,
}

// GetAlarmSeverityEnumValues Enumerates the set of values for AlarmSeverity
func GetAlarmSeverityEnumValues() []AlarmSeverityEnum {
	values := make([]AlarmSeverityEnum, 0)
	for _, v := range mappingAlarmSeverity {
		values = append(values, v)
	}
	return values
}
