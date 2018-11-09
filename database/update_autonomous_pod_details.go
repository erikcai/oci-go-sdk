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

// UpdateAutonomousPodDetails Describe the modification parameters for the Autonomous Pod.
type UpdateAutonomousPodDetails struct {

	// The display name for the Autonomous Pod.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Database Patch model preference.
	PatchModel UpdateAutonomousPodDetailsPatchModelEnum `mandatory:"false" json:"patchModel,omitempty"`

	// Maintenance Window preference, if set to WEEK_TWO, database patch will be applied in the second week during each patching cycle.
	MaintenanceWindow UpdateAutonomousPodDetailsMaintenanceWindowEnum `mandatory:"false" json:"maintenanceWindow,omitempty"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	BackupConfig *AutonomousPodBackupConfig `mandatory:"false" json:"backupConfig"`
}

func (m UpdateAutonomousPodDetails) String() string {
	return common.PointerString(m)
}

// UpdateAutonomousPodDetailsPatchModelEnum Enum with underlying type: string
type UpdateAutonomousPodDetailsPatchModelEnum string

// Set of constants representing the allowable values for UpdateAutonomousPodDetailsPatchModelEnum
const (
	UpdateAutonomousPodDetailsPatchModelUpdates         UpdateAutonomousPodDetailsPatchModelEnum = "RELEASE_UPDATES"
	UpdateAutonomousPodDetailsPatchModelUpdateRevisions UpdateAutonomousPodDetailsPatchModelEnum = "RELEASE_UPDATE_REVISIONS"
)

var mappingUpdateAutonomousPodDetailsPatchModel = map[string]UpdateAutonomousPodDetailsPatchModelEnum{
	"RELEASE_UPDATES":          UpdateAutonomousPodDetailsPatchModelUpdates,
	"RELEASE_UPDATE_REVISIONS": UpdateAutonomousPodDetailsPatchModelUpdateRevisions,
}

// GetUpdateAutonomousPodDetailsPatchModelEnumValues Enumerates the set of values for UpdateAutonomousPodDetailsPatchModelEnum
func GetUpdateAutonomousPodDetailsPatchModelEnumValues() []UpdateAutonomousPodDetailsPatchModelEnum {
	values := make([]UpdateAutonomousPodDetailsPatchModelEnum, 0)
	for _, v := range mappingUpdateAutonomousPodDetailsPatchModel {
		values = append(values, v)
	}
	return values
}

// UpdateAutonomousPodDetailsMaintenanceWindowEnum Enum with underlying type: string
type UpdateAutonomousPodDetailsMaintenanceWindowEnum string

// Set of constants representing the allowable values for UpdateAutonomousPodDetailsMaintenanceWindowEnum
const (
	UpdateAutonomousPodDetailsMaintenanceWindowAny       UpdateAutonomousPodDetailsMaintenanceWindowEnum = "ANY"
	UpdateAutonomousPodDetailsMaintenanceWindowWeekTwo   UpdateAutonomousPodDetailsMaintenanceWindowEnum = "WEEK_TWO"
	UpdateAutonomousPodDetailsMaintenanceWindowWeekThree UpdateAutonomousPodDetailsMaintenanceWindowEnum = "WEEK_THREE"
	UpdateAutonomousPodDetailsMaintenanceWindowWeekFour  UpdateAutonomousPodDetailsMaintenanceWindowEnum = "WEEK_FOUR"
	UpdateAutonomousPodDetailsMaintenanceWindowWeekFive  UpdateAutonomousPodDetailsMaintenanceWindowEnum = "WEEK_FIVE"
	UpdateAutonomousPodDetailsMaintenanceWindowWeekSix   UpdateAutonomousPodDetailsMaintenanceWindowEnum = "WEEK_SIX"
)

var mappingUpdateAutonomousPodDetailsMaintenanceWindow = map[string]UpdateAutonomousPodDetailsMaintenanceWindowEnum{
	"ANY":        UpdateAutonomousPodDetailsMaintenanceWindowAny,
	"WEEK_TWO":   UpdateAutonomousPodDetailsMaintenanceWindowWeekTwo,
	"WEEK_THREE": UpdateAutonomousPodDetailsMaintenanceWindowWeekThree,
	"WEEK_FOUR":  UpdateAutonomousPodDetailsMaintenanceWindowWeekFour,
	"WEEK_FIVE":  UpdateAutonomousPodDetailsMaintenanceWindowWeekFive,
	"WEEK_SIX":   UpdateAutonomousPodDetailsMaintenanceWindowWeekSix,
}

// GetUpdateAutonomousPodDetailsMaintenanceWindowEnumValues Enumerates the set of values for UpdateAutonomousPodDetailsMaintenanceWindowEnum
func GetUpdateAutonomousPodDetailsMaintenanceWindowEnumValues() []UpdateAutonomousPodDetailsMaintenanceWindowEnum {
	values := make([]UpdateAutonomousPodDetailsMaintenanceWindowEnum, 0)
	for _, v := range mappingUpdateAutonomousPodDetailsMaintenanceWindow {
		values = append(values, v)
	}
	return values
}
