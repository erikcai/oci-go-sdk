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

// CreateAutonomousPodDetails Describes the required parameters for the creation of an Autonomous Pod.
type CreateAutonomousPodDetails struct {

	// The display name for the Autonomous Pod.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the Autonomous DB System.
	AutonomousDbSystemId *string `mandatory:"true" json:"autonomousDbSystemId"`

	// Database Patch model preference.
	PatchModel CreateAutonomousPodDetailsPatchModelEnum `mandatory:"true" json:"patchModel"`

	// Maintenance Window preference, if set to WEEK_TWO, database patch will be applied in the second week during each patching cycle.
	MaintenanceWindow CreateAutonomousPodDetailsMaintenanceWindowEnum `mandatory:"true" json:"maintenanceWindow"`

	// The service level agreement type of the Autonomous Pod. The default is STANDARD. For a Mission Critical Pod, the specified Autonomous Exadata Infrastructure must be associated with a remote Autonomous Exadata Infrastructure.
	ServiceLevelAgreementType CreateAutonomousPodDetailsServiceLevelAgreementTypeEnum `mandatory:"false" json:"serviceLevelAgreementType,omitempty"`

	// The OCID of the Peer Autonomous DB System for Mission Critical Association
	PeerAutonomousDbSystemId *string `mandatory:"false" json:"peerAutonomousDbSystemId"`

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the compartment containing the Autonomous Pod.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

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

func (m CreateAutonomousPodDetails) String() string {
	return common.PointerString(m)
}

// CreateAutonomousPodDetailsServiceLevelAgreementTypeEnum Enum with underlying type: string
type CreateAutonomousPodDetailsServiceLevelAgreementTypeEnum string

// Set of constants representing the allowable values for CreateAutonomousPodDetailsServiceLevelAgreementTypeEnum
const (
	CreateAutonomousPodDetailsServiceLevelAgreementTypeStandard        CreateAutonomousPodDetailsServiceLevelAgreementTypeEnum = "STANDARD"
	CreateAutonomousPodDetailsServiceLevelAgreementTypeMissionCritical CreateAutonomousPodDetailsServiceLevelAgreementTypeEnum = "MISSION_CRITICAL"
)

var mappingCreateAutonomousPodDetailsServiceLevelAgreementType = map[string]CreateAutonomousPodDetailsServiceLevelAgreementTypeEnum{
	"STANDARD":         CreateAutonomousPodDetailsServiceLevelAgreementTypeStandard,
	"MISSION_CRITICAL": CreateAutonomousPodDetailsServiceLevelAgreementTypeMissionCritical,
}

// GetCreateAutonomousPodDetailsServiceLevelAgreementTypeEnumValues Enumerates the set of values for CreateAutonomousPodDetailsServiceLevelAgreementTypeEnum
func GetCreateAutonomousPodDetailsServiceLevelAgreementTypeEnumValues() []CreateAutonomousPodDetailsServiceLevelAgreementTypeEnum {
	values := make([]CreateAutonomousPodDetailsServiceLevelAgreementTypeEnum, 0)
	for _, v := range mappingCreateAutonomousPodDetailsServiceLevelAgreementType {
		values = append(values, v)
	}
	return values
}

// CreateAutonomousPodDetailsPatchModelEnum Enum with underlying type: string
type CreateAutonomousPodDetailsPatchModelEnum string

// Set of constants representing the allowable values for CreateAutonomousPodDetailsPatchModelEnum
const (
	CreateAutonomousPodDetailsPatchModelUpdates         CreateAutonomousPodDetailsPatchModelEnum = "RELEASE_UPDATES"
	CreateAutonomousPodDetailsPatchModelUpdateRevisions CreateAutonomousPodDetailsPatchModelEnum = "RELEASE_UPDATE_REVISIONS"
)

var mappingCreateAutonomousPodDetailsPatchModel = map[string]CreateAutonomousPodDetailsPatchModelEnum{
	"RELEASE_UPDATES":          CreateAutonomousPodDetailsPatchModelUpdates,
	"RELEASE_UPDATE_REVISIONS": CreateAutonomousPodDetailsPatchModelUpdateRevisions,
}

// GetCreateAutonomousPodDetailsPatchModelEnumValues Enumerates the set of values for CreateAutonomousPodDetailsPatchModelEnum
func GetCreateAutonomousPodDetailsPatchModelEnumValues() []CreateAutonomousPodDetailsPatchModelEnum {
	values := make([]CreateAutonomousPodDetailsPatchModelEnum, 0)
	for _, v := range mappingCreateAutonomousPodDetailsPatchModel {
		values = append(values, v)
	}
	return values
}

// CreateAutonomousPodDetailsMaintenanceWindowEnum Enum with underlying type: string
type CreateAutonomousPodDetailsMaintenanceWindowEnum string

// Set of constants representing the allowable values for CreateAutonomousPodDetailsMaintenanceWindowEnum
const (
	CreateAutonomousPodDetailsMaintenanceWindowAny       CreateAutonomousPodDetailsMaintenanceWindowEnum = "ANY"
	CreateAutonomousPodDetailsMaintenanceWindowWeekTwo   CreateAutonomousPodDetailsMaintenanceWindowEnum = "WEEK_TWO"
	CreateAutonomousPodDetailsMaintenanceWindowWeekThree CreateAutonomousPodDetailsMaintenanceWindowEnum = "WEEK_THREE"
	CreateAutonomousPodDetailsMaintenanceWindowWeekFour  CreateAutonomousPodDetailsMaintenanceWindowEnum = "WEEK_FOUR"
	CreateAutonomousPodDetailsMaintenanceWindowWeekFive  CreateAutonomousPodDetailsMaintenanceWindowEnum = "WEEK_FIVE"
	CreateAutonomousPodDetailsMaintenanceWindowWeekSix   CreateAutonomousPodDetailsMaintenanceWindowEnum = "WEEK_SIX"
)

var mappingCreateAutonomousPodDetailsMaintenanceWindow = map[string]CreateAutonomousPodDetailsMaintenanceWindowEnum{
	"ANY":        CreateAutonomousPodDetailsMaintenanceWindowAny,
	"WEEK_TWO":   CreateAutonomousPodDetailsMaintenanceWindowWeekTwo,
	"WEEK_THREE": CreateAutonomousPodDetailsMaintenanceWindowWeekThree,
	"WEEK_FOUR":  CreateAutonomousPodDetailsMaintenanceWindowWeekFour,
	"WEEK_FIVE":  CreateAutonomousPodDetailsMaintenanceWindowWeekFive,
	"WEEK_SIX":   CreateAutonomousPodDetailsMaintenanceWindowWeekSix,
}

// GetCreateAutonomousPodDetailsMaintenanceWindowEnumValues Enumerates the set of values for CreateAutonomousPodDetailsMaintenanceWindowEnum
func GetCreateAutonomousPodDetailsMaintenanceWindowEnumValues() []CreateAutonomousPodDetailsMaintenanceWindowEnum {
	values := make([]CreateAutonomousPodDetailsMaintenanceWindowEnum, 0)
	for _, v := range mappingCreateAutonomousPodDetailsMaintenanceWindow {
		values = append(values, v)
	}
	return values
}
