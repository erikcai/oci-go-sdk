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

// AlarmHistoryEntry An alarm history entry, consisting of a description and a timestamp.
type AlarmHistoryEntry struct {

	// Description for this alarm history entry.
	// Example: `Alarm transitioned from FIRING to OK.`
	Summary *string `mandatory:"true" json:"summary"`

	// Timestamp for this alarm history entry. Format defined by RFC3339.
	// Example: `2018-02-01T01:02:29.600Z`
	Timestamp *common.SDKTime `mandatory:"true" json:"timestamp"`
}

func (m AlarmHistoryEntry) String() string {
	return common.PointerString(m)
}
