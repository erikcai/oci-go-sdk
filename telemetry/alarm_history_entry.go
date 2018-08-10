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
