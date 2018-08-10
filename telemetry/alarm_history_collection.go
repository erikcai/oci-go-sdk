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

// AlarmHistoryCollection The configuration details for retrieving alarm history.
type AlarmHistoryCollection struct {

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the alarm for which to retrieve history.
	AlarmId *string `mandatory:"true" json:"alarmId"`

	// Whether the alarm is enabled.
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// The set of history entries retrieved for the alarm.
	Entries []AlarmHistoryEntry `mandatory:"true" json:"entries"`
}

func (m AlarmHistoryCollection) String() string {
	return common.PointerString(m)
}
