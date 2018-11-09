// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// AutonomousDbSystemMaintenanceWindow Autonomous DB System maintenance window details for quarterly patching.
type AutonomousDbSystemMaintenanceWindow struct {

	// Day of the first week, the patch should be applied on the Autonomous DbSystem.
	DayOfWeek AutonomousDbSystemMaintenanceWindowDayOfWeekEnum `mandatory:"true" json:"dayOfWeek"`

	// Hour of the day, the patch should be applied.
	HourOfDay *int `mandatory:"false" json:"hourOfDay"`
}

func (m AutonomousDbSystemMaintenanceWindow) String() string {
	return common.PointerString(m)
}

// AutonomousDbSystemMaintenanceWindowDayOfWeekEnum Enum with underlying type: string
type AutonomousDbSystemMaintenanceWindowDayOfWeekEnum string

// Set of constants representing the allowable values for AutonomousDbSystemMaintenanceWindowDayOfWeekEnum
const (
	AutonomousDbSystemMaintenanceWindowDayOfWeekAny       AutonomousDbSystemMaintenanceWindowDayOfWeekEnum = "ANY"
	AutonomousDbSystemMaintenanceWindowDayOfWeekSunday    AutonomousDbSystemMaintenanceWindowDayOfWeekEnum = "SUNDAY"
	AutonomousDbSystemMaintenanceWindowDayOfWeekMonday    AutonomousDbSystemMaintenanceWindowDayOfWeekEnum = "MONDAY"
	AutonomousDbSystemMaintenanceWindowDayOfWeekTuesday   AutonomousDbSystemMaintenanceWindowDayOfWeekEnum = "TUESDAY"
	AutonomousDbSystemMaintenanceWindowDayOfWeekWednesday AutonomousDbSystemMaintenanceWindowDayOfWeekEnum = "WEDNESDAY"
	AutonomousDbSystemMaintenanceWindowDayOfWeekThursday  AutonomousDbSystemMaintenanceWindowDayOfWeekEnum = "THURSDAY"
	AutonomousDbSystemMaintenanceWindowDayOfWeekFriday    AutonomousDbSystemMaintenanceWindowDayOfWeekEnum = "FRIDAY"
	AutonomousDbSystemMaintenanceWindowDayOfWeekSaturday  AutonomousDbSystemMaintenanceWindowDayOfWeekEnum = "SATURDAY"
)

var mappingAutonomousDbSystemMaintenanceWindowDayOfWeek = map[string]AutonomousDbSystemMaintenanceWindowDayOfWeekEnum{
	"ANY":       AutonomousDbSystemMaintenanceWindowDayOfWeekAny,
	"SUNDAY":    AutonomousDbSystemMaintenanceWindowDayOfWeekSunday,
	"MONDAY":    AutonomousDbSystemMaintenanceWindowDayOfWeekMonday,
	"TUESDAY":   AutonomousDbSystemMaintenanceWindowDayOfWeekTuesday,
	"WEDNESDAY": AutonomousDbSystemMaintenanceWindowDayOfWeekWednesday,
	"THURSDAY":  AutonomousDbSystemMaintenanceWindowDayOfWeekThursday,
	"FRIDAY":    AutonomousDbSystemMaintenanceWindowDayOfWeekFriday,
	"SATURDAY":  AutonomousDbSystemMaintenanceWindowDayOfWeekSaturday,
}

// GetAutonomousDbSystemMaintenanceWindowDayOfWeekEnumValues Enumerates the set of values for AutonomousDbSystemMaintenanceWindowDayOfWeekEnum
func GetAutonomousDbSystemMaintenanceWindowDayOfWeekEnumValues() []AutonomousDbSystemMaintenanceWindowDayOfWeekEnum {
	values := make([]AutonomousDbSystemMaintenanceWindowDayOfWeekEnum, 0)
	for _, v := range mappingAutonomousDbSystemMaintenanceWindowDayOfWeek {
		values = append(values, v)
	}
	return values
}
