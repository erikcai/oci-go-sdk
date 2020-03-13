// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// AutonomousDatabaseStandbySummary Autonomous Database Standby Details.
type AutonomousDatabaseStandbySummary struct {

	// The lag time set between data on the source database and data on the cloned database. From 5 mins to 7 days.
	LagTimeInSeconds *int `mandatory:"false" json:"lagTimeInSeconds"`

	// The current state of the Autonomous Database.
	LifecycleState AutonomousDatabaseStandbySummaryLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m AutonomousDatabaseStandbySummary) String() string {
	return common.PointerString(m)
}

// AutonomousDatabaseStandbySummaryLifecycleStateEnum Enum with underlying type: string
type AutonomousDatabaseStandbySummaryLifecycleStateEnum string

// Set of constants representing the allowable values for AutonomousDatabaseStandbySummaryLifecycleStateEnum
const (
	AutonomousDatabaseStandbySummaryLifecycleStateProvisioning            AutonomousDatabaseStandbySummaryLifecycleStateEnum = "PROVISIONING"
	AutonomousDatabaseStandbySummaryLifecycleStateAvailable               AutonomousDatabaseStandbySummaryLifecycleStateEnum = "AVAILABLE"
	AutonomousDatabaseStandbySummaryLifecycleStateStopping                AutonomousDatabaseStandbySummaryLifecycleStateEnum = "STOPPING"
	AutonomousDatabaseStandbySummaryLifecycleStateStopped                 AutonomousDatabaseStandbySummaryLifecycleStateEnum = "STOPPED"
	AutonomousDatabaseStandbySummaryLifecycleStateStarting                AutonomousDatabaseStandbySummaryLifecycleStateEnum = "STARTING"
	AutonomousDatabaseStandbySummaryLifecycleStateTerminating             AutonomousDatabaseStandbySummaryLifecycleStateEnum = "TERMINATING"
	AutonomousDatabaseStandbySummaryLifecycleStateTerminated              AutonomousDatabaseStandbySummaryLifecycleStateEnum = "TERMINATED"
	AutonomousDatabaseStandbySummaryLifecycleStateUnavailable             AutonomousDatabaseStandbySummaryLifecycleStateEnum = "UNAVAILABLE"
	AutonomousDatabaseStandbySummaryLifecycleStateRestoreInProgress       AutonomousDatabaseStandbySummaryLifecycleStateEnum = "RESTORE_IN_PROGRESS"
	AutonomousDatabaseStandbySummaryLifecycleStateRestoreFailed           AutonomousDatabaseStandbySummaryLifecycleStateEnum = "RESTORE_FAILED"
	AutonomousDatabaseStandbySummaryLifecycleStateBackupInProgress        AutonomousDatabaseStandbySummaryLifecycleStateEnum = "BACKUP_IN_PROGRESS"
	AutonomousDatabaseStandbySummaryLifecycleStateScaleInProgress         AutonomousDatabaseStandbySummaryLifecycleStateEnum = "SCALE_IN_PROGRESS"
	AutonomousDatabaseStandbySummaryLifecycleStateAvailableNeedsAttention AutonomousDatabaseStandbySummaryLifecycleStateEnum = "AVAILABLE_NEEDS_ATTENTION"
	AutonomousDatabaseStandbySummaryLifecycleStateUpdating                AutonomousDatabaseStandbySummaryLifecycleStateEnum = "UPDATING"
	AutonomousDatabaseStandbySummaryLifecycleStateMaintenanceInProgress   AutonomousDatabaseStandbySummaryLifecycleStateEnum = "MAINTENANCE_IN_PROGRESS"
	AutonomousDatabaseStandbySummaryLifecycleStateRestarting              AutonomousDatabaseStandbySummaryLifecycleStateEnum = "RESTARTING"
	AutonomousDatabaseStandbySummaryLifecycleStateRecreating              AutonomousDatabaseStandbySummaryLifecycleStateEnum = "RECREATING"
)

var mappingAutonomousDatabaseStandbySummaryLifecycleState = map[string]AutonomousDatabaseStandbySummaryLifecycleStateEnum{
	"PROVISIONING":              AutonomousDatabaseStandbySummaryLifecycleStateProvisioning,
	"AVAILABLE":                 AutonomousDatabaseStandbySummaryLifecycleStateAvailable,
	"STOPPING":                  AutonomousDatabaseStandbySummaryLifecycleStateStopping,
	"STOPPED":                   AutonomousDatabaseStandbySummaryLifecycleStateStopped,
	"STARTING":                  AutonomousDatabaseStandbySummaryLifecycleStateStarting,
	"TERMINATING":               AutonomousDatabaseStandbySummaryLifecycleStateTerminating,
	"TERMINATED":                AutonomousDatabaseStandbySummaryLifecycleStateTerminated,
	"UNAVAILABLE":               AutonomousDatabaseStandbySummaryLifecycleStateUnavailable,
	"RESTORE_IN_PROGRESS":       AutonomousDatabaseStandbySummaryLifecycleStateRestoreInProgress,
	"RESTORE_FAILED":            AutonomousDatabaseStandbySummaryLifecycleStateRestoreFailed,
	"BACKUP_IN_PROGRESS":        AutonomousDatabaseStandbySummaryLifecycleStateBackupInProgress,
	"SCALE_IN_PROGRESS":         AutonomousDatabaseStandbySummaryLifecycleStateScaleInProgress,
	"AVAILABLE_NEEDS_ATTENTION": AutonomousDatabaseStandbySummaryLifecycleStateAvailableNeedsAttention,
	"UPDATING":                  AutonomousDatabaseStandbySummaryLifecycleStateUpdating,
	"MAINTENANCE_IN_PROGRESS":   AutonomousDatabaseStandbySummaryLifecycleStateMaintenanceInProgress,
	"RESTARTING":                AutonomousDatabaseStandbySummaryLifecycleStateRestarting,
	"RECREATING":                AutonomousDatabaseStandbySummaryLifecycleStateRecreating,
}

// GetAutonomousDatabaseStandbySummaryLifecycleStateEnumValues Enumerates the set of values for AutonomousDatabaseStandbySummaryLifecycleStateEnum
func GetAutonomousDatabaseStandbySummaryLifecycleStateEnumValues() []AutonomousDatabaseStandbySummaryLifecycleStateEnum {
	values := make([]AutonomousDatabaseStandbySummaryLifecycleStateEnum, 0)
	for _, v := range mappingAutonomousDatabaseStandbySummaryLifecycleState {
		values = append(values, v)
	}
	return values
}
