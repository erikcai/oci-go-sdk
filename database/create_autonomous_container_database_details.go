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

// CreateAutonomousContainerDatabaseDetails Describes the required parameters for the creation of an Autonomous Container Database.
type CreateAutonomousContainerDatabaseDetails struct {

	// The display name for the Autonomous Container Database.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the Autonomous DB System.
	AutonomousDbSystemId *string `mandatory:"true" json:"autonomousDbSystemId"`

	// Database Patch model preference.
	PatchModel CreateAutonomousContainerDatabaseDetailsPatchModelEnum `mandatory:"true" json:"patchModel"`

	// Maintenance Window preference, if set to WEEK_TWO, database patch will be applied in the second week during each patching cycle.
	MaintenanceWindow CreateAutonomousContainerDatabaseDetailsMaintenanceWindowEnum `mandatory:"true" json:"maintenanceWindow"`

	// The service level agreement type of the Autonomous Container Database. The default is STANDARD. For a Mission Critical Container Database, the specified Autonomous Exadata Infrastructure must be associated with a remote Autonomous Exadata Infrastructure.
	ServiceLevelAgreementType CreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeEnum `mandatory:"false" json:"serviceLevelAgreementType,omitempty"`

	// The OCID of the Peer Autonomous DB System for Mission Critical Association
	PeerAutonomousDbSystemId *string `mandatory:"false" json:"peerAutonomousDbSystemId"`

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the compartment containing the Autonomous Container Database.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

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

func (m CreateAutonomousContainerDatabaseDetails) String() string {
	return common.PointerString(m)
}

// CreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeEnum Enum with underlying type: string
type CreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeEnum string

// Set of constants representing the allowable values for CreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeEnum
const (
	CreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeStandard        CreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeEnum = "STANDARD"
	CreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeMissionCritical CreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeEnum = "MISSION_CRITICAL"
)

var mappingCreateAutonomousContainerDatabaseDetailsServiceLevelAgreementType = map[string]CreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeEnum{
	"STANDARD":         CreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeStandard,
	"MISSION_CRITICAL": CreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeMissionCritical,
}

// GetCreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeEnumValues Enumerates the set of values for CreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeEnum
func GetCreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeEnumValues() []CreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeEnum {
	values := make([]CreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeEnum, 0)
	for _, v := range mappingCreateAutonomousContainerDatabaseDetailsServiceLevelAgreementType {
		values = append(values, v)
	}
	return values
}

// CreateAutonomousContainerDatabaseDetailsPatchModelEnum Enum with underlying type: string
type CreateAutonomousContainerDatabaseDetailsPatchModelEnum string

// Set of constants representing the allowable values for CreateAutonomousContainerDatabaseDetailsPatchModelEnum
const (
	CreateAutonomousContainerDatabaseDetailsPatchModelUpdates         CreateAutonomousContainerDatabaseDetailsPatchModelEnum = "RELEASE_UPDATES"
	CreateAutonomousContainerDatabaseDetailsPatchModelUpdateRevisions CreateAutonomousContainerDatabaseDetailsPatchModelEnum = "RELEASE_UPDATE_REVISIONS"
)

var mappingCreateAutonomousContainerDatabaseDetailsPatchModel = map[string]CreateAutonomousContainerDatabaseDetailsPatchModelEnum{
	"RELEASE_UPDATES":          CreateAutonomousContainerDatabaseDetailsPatchModelUpdates,
	"RELEASE_UPDATE_REVISIONS": CreateAutonomousContainerDatabaseDetailsPatchModelUpdateRevisions,
}

// GetCreateAutonomousContainerDatabaseDetailsPatchModelEnumValues Enumerates the set of values for CreateAutonomousContainerDatabaseDetailsPatchModelEnum
func GetCreateAutonomousContainerDatabaseDetailsPatchModelEnumValues() []CreateAutonomousContainerDatabaseDetailsPatchModelEnum {
	values := make([]CreateAutonomousContainerDatabaseDetailsPatchModelEnum, 0)
	for _, v := range mappingCreateAutonomousContainerDatabaseDetailsPatchModel {
		values = append(values, v)
	}
	return values
}

// CreateAutonomousContainerDatabaseDetailsMaintenanceWindowEnum Enum with underlying type: string
type CreateAutonomousContainerDatabaseDetailsMaintenanceWindowEnum string

// Set of constants representing the allowable values for CreateAutonomousContainerDatabaseDetailsMaintenanceWindowEnum
const (
	CreateAutonomousContainerDatabaseDetailsMaintenanceWindowAny       CreateAutonomousContainerDatabaseDetailsMaintenanceWindowEnum = "ANY"
	CreateAutonomousContainerDatabaseDetailsMaintenanceWindowWeekTwo   CreateAutonomousContainerDatabaseDetailsMaintenanceWindowEnum = "WEEK_TWO"
	CreateAutonomousContainerDatabaseDetailsMaintenanceWindowWeekThree CreateAutonomousContainerDatabaseDetailsMaintenanceWindowEnum = "WEEK_THREE"
	CreateAutonomousContainerDatabaseDetailsMaintenanceWindowWeekFour  CreateAutonomousContainerDatabaseDetailsMaintenanceWindowEnum = "WEEK_FOUR"
	CreateAutonomousContainerDatabaseDetailsMaintenanceWindowWeekFive  CreateAutonomousContainerDatabaseDetailsMaintenanceWindowEnum = "WEEK_FIVE"
	CreateAutonomousContainerDatabaseDetailsMaintenanceWindowWeekSix   CreateAutonomousContainerDatabaseDetailsMaintenanceWindowEnum = "WEEK_SIX"
)

var mappingCreateAutonomousContainerDatabaseDetailsMaintenanceWindow = map[string]CreateAutonomousContainerDatabaseDetailsMaintenanceWindowEnum{
	"ANY":        CreateAutonomousContainerDatabaseDetailsMaintenanceWindowAny,
	"WEEK_TWO":   CreateAutonomousContainerDatabaseDetailsMaintenanceWindowWeekTwo,
	"WEEK_THREE": CreateAutonomousContainerDatabaseDetailsMaintenanceWindowWeekThree,
	"WEEK_FOUR":  CreateAutonomousContainerDatabaseDetailsMaintenanceWindowWeekFour,
	"WEEK_FIVE":  CreateAutonomousContainerDatabaseDetailsMaintenanceWindowWeekFive,
	"WEEK_SIX":   CreateAutonomousContainerDatabaseDetailsMaintenanceWindowWeekSix,
}

// GetCreateAutonomousContainerDatabaseDetailsMaintenanceWindowEnumValues Enumerates the set of values for CreateAutonomousContainerDatabaseDetailsMaintenanceWindowEnum
func GetCreateAutonomousContainerDatabaseDetailsMaintenanceWindowEnumValues() []CreateAutonomousContainerDatabaseDetailsMaintenanceWindowEnum {
	values := make([]CreateAutonomousContainerDatabaseDetailsMaintenanceWindowEnum, 0)
	for _, v := range mappingCreateAutonomousContainerDatabaseDetailsMaintenanceWindow {
		values = append(values, v)
	}
	return values
}
