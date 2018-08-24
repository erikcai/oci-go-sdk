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

// UpdateAlarmDetails The configuration details for updating an alarm.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type UpdateAlarmDetails struct {

	// A user-friendly name for the alarm. It does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// This name is sent as the title for notifications related to this alarm.
	// Example: `High CPU Utilization`
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the compartment containing the alarm.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the compartment containing the metric
	// being evaluated by the alarm.
	MetricCompartmentId *string `mandatory:"false" json:"metricCompartmentId"`

	// The source service or application emitting the metric that is evaluated by the alarm.
	// Example: `oci/compute`
	Namespace *string `mandatory:"false" json:"namespace"`

	// The Telemetry Query Language (TQL) expression to evaluate for the alarm. For each time series,
	// the query must evaluate to a Boolean value, where zero represents false and non-zero values
	// represent true. The query must specify a metric, statistic, interval, and  comparison operator
	// (unless defining a query for an absence alarm). Supported values for
	// interval: `1m`, `5m`, `1h`. You can optionally specify dimensions and grouping functions.
	// Supported grouping functions: `grouping()`, `groupBy()`.
	// For available dimensions, review the metric definition.
	// Example of threshold alarm:
	// `CpuUtilization[1m]{availabilityDomain="cumS:PHX-AD-1"}.groupBy(availabilityDomain).percentile(0.9) > 85`
	// Example of absence alarm:
	// `CpuUtilization[1m]{availabilityDomain="cumS:PHX-AD-1"}.absent()`
	Query *string `mandatory:"false" json:"query"`

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

	// The perceived severity of the alarm with regard to the affected system.
	// Example: `CRITICAL`
	Severity AlarmSeverityEnum `mandatory:"false" json:"severity,omitempty"`

	// The human-readable content of the notification delivered. Oracle recommends providing guidance
	// to operators for resolving the alarm condition. Consider adding links to standard runbook
	// practices.
	// Example: `High CPU usage alert. Follow runbook instructions for resolution.`
	Body *string `mandatory:"false" json:"body"`

	// An array of OCIDs (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) to which the notifications for
	// this alarm will be delivered. An example destination is an OCID for a topic managed by the
	// Oracle Cloud Infrastructure Notification service.
	Destinations []string `mandatory:"false" json:"destinations"`

	// The frequency at which notifications are re-submitted, if the alarm keeps firing without
	// interruption. Format defined by ISO 8601. For example, `PT4H` indicates four hours.
	// Default value: null (notifications are not re-submitted).
	// Example: `PT2H`
	RepeatNotificationDuration *string `mandatory:"false" json:"repeatNotificationDuration"`

	Suppression *Suppression `mandatory:"false" json:"suppression"`

	// Whether the alarm is enabled.
	// Example: `true`
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`
}

func (m UpdateAlarmDetails) String() string {
	return common.PointerString(m)
}
