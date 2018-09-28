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

// AlarmHistoryCollection The configuration details for retrieving alarm history.
type AlarmHistoryCollection struct {

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/iaas/Content/General/Concepts/identifiers.htm) of the alarm for which to retrieve history.
	AlarmId *string `mandatory:"true" json:"alarmId"`

	// Whether the alarm is enabled.
	// Example: `true`
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// The set of history entries retrieved for the alarm.
	Entries []AlarmHistoryEntry `mandatory:"true" json:"entries"`
}

func (m AlarmHistoryCollection) String() string {
	return common.PointerString(m)
}
