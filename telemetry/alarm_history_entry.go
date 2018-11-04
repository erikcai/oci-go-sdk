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

// AlarmHistoryEntry An alarm history entry indicating a description of the entry and the time that the entry occurred.
// If the entry corresponds to a state transition, such as OK to Firing, then the entry also includes a transition timestamp.
type AlarmHistoryEntry struct {

	// Description for this alarm history entry.
	// Example 1 - alarm state history entry: `The alarm state is FIRING`
	// Example 2 - alarm state transition history entry: `State transitioned from OK to Firing`
	Summary *string `mandatory:"true" json:"summary"`

	// Timestamp for this alarm history entry. Format defined by RFC3339.
	// Example: `2018-02-01T01:02:29.600Z`
	Timestamp *common.SDKTime `mandatory:"true" json:"timestamp"`

	// Timestamp for the transition of the alarm state. For example, the time when the alarm transitioned from OK to Firing.
	// Available for state transition entries only. Note: A three-minute lag for this value accounts for any late-arriving metrics.
	// Example: `2018-02-01T0:59:00.000Z`
	TimestampTriggered *common.SDKTime `mandatory:"false" json:"timestampTriggered"`
}

func (m AlarmHistoryEntry) String() string {
	return common.PointerString(m)
}
