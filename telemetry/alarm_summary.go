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


    
 // AlarmSummary The document with a summary of the alarm definition
type AlarmSummary struct {
    
 // The OCID of the alarm.
    Id *string `mandatory:"true" json:"id"`
    
 // A user-friendly name for the alarm. It does not have to be unique, and it's changeable.
 // Avoid entering confidential information.
 // The name is used as the title for the alarm notifications.
    DisplayName *string `mandatory:"true" json:"displayName"`
    
 // The OCID of the compartment for the alarm.
    CompartmentId *string `mandatory:"true" json:"compartmentId"`
    
 // The OCID of the compartment for the metric in the alarm.
    MetricCompartmentId *string `mandatory:"true" json:"metricCompartmentId"`
    
 // The namespace for the metric in the alarm.
    Namespace *string `mandatory:"true" json:"namespace"`
    
 // The severity of the alarm.
    Severity AlarmSummarySeverityEnum `mandatory:"true" json:"severity"`
    
 // A list of OCID that identify the destinations for the notification. For example, the destinations
 // could contain the OCID of an ONS topic, for delivering the notification.
    Destinations []string `mandatory:"true" json:"destinations"`
    
    Suppression *Suppression `mandatory:"false" json:"suppression"`
}

func (m AlarmSummary) String() string {
    return common.PointerString(m)
}


// AlarmSummarySeverityEnum Enum with underlying type: string
type AlarmSummarySeverityEnum string

// Set of constants representing the allowable values for AlarmSummarySeverity
const (
    AlarmSummarySeverityCritical AlarmSummarySeverityEnum = "CRITICAL"
    AlarmSummarySeverityError AlarmSummarySeverityEnum = "ERROR"
    AlarmSummarySeverityWarning AlarmSummarySeverityEnum = "WARNING"
    AlarmSummarySeverityInfo AlarmSummarySeverityEnum = "INFO"
)

var mappingAlarmSummarySeverity = map[string]AlarmSummarySeverityEnum { 
    "CRITICAL": AlarmSummarySeverityCritical,
    "ERROR": AlarmSummarySeverityError,
    "WARNING": AlarmSummarySeverityWarning,
    "INFO": AlarmSummarySeverityInfo,
}

// GetAlarmSummarySeverityEnumValues Enumerates the set of values for AlarmSummarySeverity
func GetAlarmSummarySeverityEnumValues() []AlarmSummarySeverityEnum {
   values := make([]AlarmSummarySeverityEnum, 0)
   for _, v := range mappingAlarmSummarySeverity {
       values = append(values, v)
   }
   return values
}



