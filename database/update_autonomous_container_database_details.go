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

// UpdateAutonomousContainerDatabaseDetails Describes the modification parameters for the Autonomous Container Database.
type UpdateAutonomousContainerDatabaseDetails struct {

	// The display name for the Autonomous Container Database.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Database Patch model preference.
	PatchModel UpdateAutonomousContainerDatabaseDetailsPatchModelEnum `mandatory:"false" json:"patchModel,omitempty"`

	// Maintenance Window preference, if set to WEEK_TWO, database patch will be applied in the second week during each patching cycle.
	MaintenanceWindow UpdateAutonomousContainerDatabaseDetailsMaintenanceWindowEnum `mandatory:"false" json:"maintenanceWindow,omitempty"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	BackupConfig *AutonomousContainerDatabaseBackupConfig `mandatory:"false" json:"backupConfig"`
}

func (m UpdateAutonomousContainerDatabaseDetails) String() string {
	return common.PointerString(m)
}

// UpdateAutonomousContainerDatabaseDetailsPatchModelEnum Enum with underlying type: string
type UpdateAutonomousContainerDatabaseDetailsPatchModelEnum string

// Set of constants representing the allowable values for UpdateAutonomousContainerDatabaseDetailsPatchModelEnum
const (
	UpdateAutonomousContainerDatabaseDetailsPatchModelUpdates         UpdateAutonomousContainerDatabaseDetailsPatchModelEnum = "RELEASE_UPDATES"
	UpdateAutonomousContainerDatabaseDetailsPatchModelUpdateRevisions UpdateAutonomousContainerDatabaseDetailsPatchModelEnum = "RELEASE_UPDATE_REVISIONS"
)

var mappingUpdateAutonomousContainerDatabaseDetailsPatchModel = map[string]UpdateAutonomousContainerDatabaseDetailsPatchModelEnum{
	"RELEASE_UPDATES":          UpdateAutonomousContainerDatabaseDetailsPatchModelUpdates,
	"RELEASE_UPDATE_REVISIONS": UpdateAutonomousContainerDatabaseDetailsPatchModelUpdateRevisions,
}

// GetUpdateAutonomousContainerDatabaseDetailsPatchModelEnumValues Enumerates the set of values for UpdateAutonomousContainerDatabaseDetailsPatchModelEnum
func GetUpdateAutonomousContainerDatabaseDetailsPatchModelEnumValues() []UpdateAutonomousContainerDatabaseDetailsPatchModelEnum {
	values := make([]UpdateAutonomousContainerDatabaseDetailsPatchModelEnum, 0)
	for _, v := range mappingUpdateAutonomousContainerDatabaseDetailsPatchModel {
		values = append(values, v)
	}
	return values
}

// UpdateAutonomousContainerDatabaseDetailsMaintenanceWindowEnum Enum with underlying type: string
type UpdateAutonomousContainerDatabaseDetailsMaintenanceWindowEnum string

// Set of constants representing the allowable values for UpdateAutonomousContainerDatabaseDetailsMaintenanceWindowEnum
const (
	UpdateAutonomousContainerDatabaseDetailsMaintenanceWindowAny       UpdateAutonomousContainerDatabaseDetailsMaintenanceWindowEnum = "ANY"
	UpdateAutonomousContainerDatabaseDetailsMaintenanceWindowWeekTwo   UpdateAutonomousContainerDatabaseDetailsMaintenanceWindowEnum = "WEEK_TWO"
	UpdateAutonomousContainerDatabaseDetailsMaintenanceWindowWeekThree UpdateAutonomousContainerDatabaseDetailsMaintenanceWindowEnum = "WEEK_THREE"
	UpdateAutonomousContainerDatabaseDetailsMaintenanceWindowWeekFour  UpdateAutonomousContainerDatabaseDetailsMaintenanceWindowEnum = "WEEK_FOUR"
	UpdateAutonomousContainerDatabaseDetailsMaintenanceWindowWeekFive  UpdateAutonomousContainerDatabaseDetailsMaintenanceWindowEnum = "WEEK_FIVE"
	UpdateAutonomousContainerDatabaseDetailsMaintenanceWindowWeekSix   UpdateAutonomousContainerDatabaseDetailsMaintenanceWindowEnum = "WEEK_SIX"
)

var mappingUpdateAutonomousContainerDatabaseDetailsMaintenanceWindow = map[string]UpdateAutonomousContainerDatabaseDetailsMaintenanceWindowEnum{
	"ANY":        UpdateAutonomousContainerDatabaseDetailsMaintenanceWindowAny,
	"WEEK_TWO":   UpdateAutonomousContainerDatabaseDetailsMaintenanceWindowWeekTwo,
	"WEEK_THREE": UpdateAutonomousContainerDatabaseDetailsMaintenanceWindowWeekThree,
	"WEEK_FOUR":  UpdateAutonomousContainerDatabaseDetailsMaintenanceWindowWeekFour,
	"WEEK_FIVE":  UpdateAutonomousContainerDatabaseDetailsMaintenanceWindowWeekFive,
	"WEEK_SIX":   UpdateAutonomousContainerDatabaseDetailsMaintenanceWindowWeekSix,
}

// GetUpdateAutonomousContainerDatabaseDetailsMaintenanceWindowEnumValues Enumerates the set of values for UpdateAutonomousContainerDatabaseDetailsMaintenanceWindowEnum
func GetUpdateAutonomousContainerDatabaseDetailsMaintenanceWindowEnumValues() []UpdateAutonomousContainerDatabaseDetailsMaintenanceWindowEnum {
	values := make([]UpdateAutonomousContainerDatabaseDetailsMaintenanceWindowEnum, 0)
	for _, v := range mappingUpdateAutonomousContainerDatabaseDetailsMaintenanceWindow {
		values = append(values, v)
	}
	return values
}
