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

// UpdateAlarmDetails The document with the alarm definition that specifies among other things, the query to evaluate alarms, the
// severity, and the notification destinations.
type UpdateAlarmDetails struct {

	// A user-friendly name for the alarm. It does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// The name is used as the title for the alarm notifications.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The OCID of the compartment for the alarm.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The OCID of the compartment for the metric in the alarm.
	MetricCompartmentId *string `mandatory:"false" json:"metricCompartmentId"`

	// The namespace for the metric in the alarm.
	Namespace *string `mandatory:"false" json:"namespace"`

	// The Telemetry Query Language (TQL) expression that evaluates into a "boolean" time series, where zero
	// represents false and one represents true.
	Query *string `mandatory:"false" json:"query"`

	// Desired resolution between datapoints in query response. This determines the
	// time window which will be used to query for data. For example, if cpu[5m].sum() is
	// provided as the query with a resolution of 1m, then the sum of cpu metric for 5m intervals
	// will be returned, every 1m, between the start and end time ranges.
	// The resolution must be equal or less than the interval in the query expression.
	// The default resolution is 1m.
	Resolution *string `mandatory:"false" json:"resolution"`

	// The time duration in which the query is breaching, prior to FIRING the alarm. During this time, the
	// breaching alarm is considered PENDING.
	// This also determines the duration in which the query must be not breaching, prior to transitioning the
	// alarm from FIRING into OK.
	// The duration is specified as a string in ISO 8601 format, e.g. PT10M, ten minutes, or one hour, PT1H.
	// The minimum duration is PT1M, and the maximum duration is PT1H.
	// The default value is PT1M, so the first evaluation breaching will be FIRING the alarm, and the first
	// evaluation not breaching will set to OK the alarm.
	PendingDuration *string `mandatory:"false" json:"pendingDuration"`

	// The severity of the alarm.
	Severity UpdateAlarmDetailsSeverityEnum `mandatory:"false" json:"severity,omitempty"`

	// The body of the notification delivered.
	Body *string `mandatory:"false" json:"body"`

	// A list of OCID that identify the destinations for the notification. For example, the destinations
	// could contain the OCID of an ONS topic, for delivering the notification.
	Destinations []string `mandatory:"false" json:"destinations"`

	// The frequency in which notifications are re-submitted, if the alarm keeps firing. The frequency
	// is specified as string in ISO 8601 format, e.g. PT4H, for four hours.
	// Default does not repeat notifications.
	RepeatNotificationDuration *string `mandatory:"false" json:"repeatNotificationDuration"`

	Suppression *Suppression `mandatory:"false" json:"suppression"`

	// Whether the alarm is enabled.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`
}

func (m UpdateAlarmDetails) String() string {
	return common.PointerString(m)
}

// UpdateAlarmDetailsSeverityEnum Enum with underlying type: string
type UpdateAlarmDetailsSeverityEnum string

// Set of constants representing the allowable values for UpdateAlarmDetailsSeverity
const (
	UpdateAlarmDetailsSeverityCritical UpdateAlarmDetailsSeverityEnum = "CRITICAL"
	UpdateAlarmDetailsSeverityError    UpdateAlarmDetailsSeverityEnum = "ERROR"
	UpdateAlarmDetailsSeverityWarning  UpdateAlarmDetailsSeverityEnum = "WARNING"
	UpdateAlarmDetailsSeverityInfo     UpdateAlarmDetailsSeverityEnum = "INFO"
)

var mappingUpdateAlarmDetailsSeverity = map[string]UpdateAlarmDetailsSeverityEnum{
	"CRITICAL": UpdateAlarmDetailsSeverityCritical,
	"ERROR":    UpdateAlarmDetailsSeverityError,
	"WARNING":  UpdateAlarmDetailsSeverityWarning,
	"INFO":     UpdateAlarmDetailsSeverityInfo,
}

// GetUpdateAlarmDetailsSeverityEnumValues Enumerates the set of values for UpdateAlarmDetailsSeverity
func GetUpdateAlarmDetailsSeverityEnumValues() []UpdateAlarmDetailsSeverityEnum {
	values := make([]UpdateAlarmDetailsSeverityEnum, 0)
	for _, v := range mappingUpdateAlarmDetailsSeverity {
		values = append(values, v)
	}
	return values
}
