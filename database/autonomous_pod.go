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

// AutonomousPod The representation of AutonomousPod
type AutonomousPod struct {

	// The OCID of the Autonomous Pod.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The user-provided name for the Autonomous Pod.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Service level agreement type of the Pod. Default is STANDARD.
	ServiceLevelAgreementType AutonomousPodServiceLevelAgreementTypeEnum `mandatory:"true" json:"serviceLevelAgreementType"`

	// The OCID of the Autonomous DB System.
	AutonomousDbSystemId *string `mandatory:"true" json:"autonomousDbSystemId"`

	// The current state of the Autonomous Pod.
	LifecycleState AutonomousPodLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Database Patch model preference.
	PatchModel AutonomousPodPatchModelEnum `mandatory:"true" json:"patchModel"`

	// Maintenance Week preference, if set to WEEK_TWO, database patch will be applied in the second week during each patching cycle.
	MaintenanceWeek AutonomousPodMaintenanceWeekEnum `mandatory:"true" json:"maintenanceWeek"`

	// Additional information about the current lifecycleState.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The date and time the Autonomous was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	LastMaintenanceRun *MaintenanceRun `mandatory:"false" json:"lastMaintenanceRun"`

	NextMaintenanceRun *MaintenanceRun `mandatory:"false" json:"nextMaintenanceRun"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The role of the Autonomous Pod in this Mission Critical association.
	Role AutonomousPodRoleEnum `mandatory:"false" json:"role,omitempty"`

	// The Availability Domain where the Autonomous Pod located.
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	BackupConfig *AutonomousPodBackupConfig `mandatory:"false" json:"backupConfig"`
}

func (m AutonomousPod) String() string {
	return common.PointerString(m)
}

// AutonomousPodServiceLevelAgreementTypeEnum Enum with underlying type: string
type AutonomousPodServiceLevelAgreementTypeEnum string

// Set of constants representing the allowable values for AutonomousPodServiceLevelAgreementTypeEnum
const (
	AutonomousPodServiceLevelAgreementTypeStandard        AutonomousPodServiceLevelAgreementTypeEnum = "STANDARD"
	AutonomousPodServiceLevelAgreementTypeMissionCritical AutonomousPodServiceLevelAgreementTypeEnum = "MISSION_CRITICAL"
)

var mappingAutonomousPodServiceLevelAgreementType = map[string]AutonomousPodServiceLevelAgreementTypeEnum{
	"STANDARD":         AutonomousPodServiceLevelAgreementTypeStandard,
	"MISSION_CRITICAL": AutonomousPodServiceLevelAgreementTypeMissionCritical,
}

// GetAutonomousPodServiceLevelAgreementTypeEnumValues Enumerates the set of values for AutonomousPodServiceLevelAgreementTypeEnum
func GetAutonomousPodServiceLevelAgreementTypeEnumValues() []AutonomousPodServiceLevelAgreementTypeEnum {
	values := make([]AutonomousPodServiceLevelAgreementTypeEnum, 0)
	for _, v := range mappingAutonomousPodServiceLevelAgreementType {
		values = append(values, v)
	}
	return values
}

// AutonomousPodLifecycleStateEnum Enum with underlying type: string
type AutonomousPodLifecycleStateEnum string

// Set of constants representing the allowable values for AutonomousPodLifecycleStateEnum
const (
	AutonomousPodLifecycleStateProvisioning     AutonomousPodLifecycleStateEnum = "PROVISIONING"
	AutonomousPodLifecycleStateAvailable        AutonomousPodLifecycleStateEnum = "AVAILABLE"
	AutonomousPodLifecycleStateUpdating         AutonomousPodLifecycleStateEnum = "UPDATING"
	AutonomousPodLifecycleStateTerminating      AutonomousPodLifecycleStateEnum = "TERMINATING"
	AutonomousPodLifecycleStateTerminated       AutonomousPodLifecycleStateEnum = "TERMINATED"
	AutonomousPodLifecycleStateFailed           AutonomousPodLifecycleStateEnum = "FAILED"
	AutonomousPodLifecycleStateBackupInProgress AutonomousPodLifecycleStateEnum = "BACKUP_IN_PROGRESS"
	AutonomousPodLifecycleStateRestoring        AutonomousPodLifecycleStateEnum = "RESTORING"
	AutonomousPodLifecycleStateRestoreFailed    AutonomousPodLifecycleStateEnum = "RESTORE_FAILED"
)

var mappingAutonomousPodLifecycleState = map[string]AutonomousPodLifecycleStateEnum{
	"PROVISIONING":       AutonomousPodLifecycleStateProvisioning,
	"AVAILABLE":          AutonomousPodLifecycleStateAvailable,
	"UPDATING":           AutonomousPodLifecycleStateUpdating,
	"TERMINATING":        AutonomousPodLifecycleStateTerminating,
	"TERMINATED":         AutonomousPodLifecycleStateTerminated,
	"FAILED":             AutonomousPodLifecycleStateFailed,
	"BACKUP_IN_PROGRESS": AutonomousPodLifecycleStateBackupInProgress,
	"RESTORING":          AutonomousPodLifecycleStateRestoring,
	"RESTORE_FAILED":     AutonomousPodLifecycleStateRestoreFailed,
}

// GetAutonomousPodLifecycleStateEnumValues Enumerates the set of values for AutonomousPodLifecycleStateEnum
func GetAutonomousPodLifecycleStateEnumValues() []AutonomousPodLifecycleStateEnum {
	values := make([]AutonomousPodLifecycleStateEnum, 0)
	for _, v := range mappingAutonomousPodLifecycleState {
		values = append(values, v)
	}
	return values
}

// AutonomousPodPatchModelEnum Enum with underlying type: string
type AutonomousPodPatchModelEnum string

// Set of constants representing the allowable values for AutonomousPodPatchModelEnum
const (
	AutonomousPodPatchModelUpdates         AutonomousPodPatchModelEnum = "RELEASE_UPDATES"
	AutonomousPodPatchModelUpdateRevisions AutonomousPodPatchModelEnum = "RELEASE_UPDATE_REVISIONS"
)

var mappingAutonomousPodPatchModel = map[string]AutonomousPodPatchModelEnum{
	"RELEASE_UPDATES":          AutonomousPodPatchModelUpdates,
	"RELEASE_UPDATE_REVISIONS": AutonomousPodPatchModelUpdateRevisions,
}

// GetAutonomousPodPatchModelEnumValues Enumerates the set of values for AutonomousPodPatchModelEnum
func GetAutonomousPodPatchModelEnumValues() []AutonomousPodPatchModelEnum {
	values := make([]AutonomousPodPatchModelEnum, 0)
	for _, v := range mappingAutonomousPodPatchModel {
		values = append(values, v)
	}
	return values
}

// AutonomousPodMaintenanceWeekEnum Enum with underlying type: string
type AutonomousPodMaintenanceWeekEnum string

// Set of constants representing the allowable values for AutonomousPodMaintenanceWeekEnum
const (
	AutonomousPodMaintenanceWeekAny       AutonomousPodMaintenanceWeekEnum = "ANY"
	AutonomousPodMaintenanceWeekWeekTwo   AutonomousPodMaintenanceWeekEnum = "WEEK_TWO"
	AutonomousPodMaintenanceWeekWeekThree AutonomousPodMaintenanceWeekEnum = "WEEK_THREE"
	AutonomousPodMaintenanceWeekWeekFour  AutonomousPodMaintenanceWeekEnum = "WEEK_FOUR"
	AutonomousPodMaintenanceWeekWeekFive  AutonomousPodMaintenanceWeekEnum = "WEEK_FIVE"
	AutonomousPodMaintenanceWeekWeekSix   AutonomousPodMaintenanceWeekEnum = "WEEK_SIX"
)

var mappingAutonomousPodMaintenanceWeek = map[string]AutonomousPodMaintenanceWeekEnum{
	"ANY":        AutonomousPodMaintenanceWeekAny,
	"WEEK_TWO":   AutonomousPodMaintenanceWeekWeekTwo,
	"WEEK_THREE": AutonomousPodMaintenanceWeekWeekThree,
	"WEEK_FOUR":  AutonomousPodMaintenanceWeekWeekFour,
	"WEEK_FIVE":  AutonomousPodMaintenanceWeekWeekFive,
	"WEEK_SIX":   AutonomousPodMaintenanceWeekWeekSix,
}

// GetAutonomousPodMaintenanceWeekEnumValues Enumerates the set of values for AutonomousPodMaintenanceWeekEnum
func GetAutonomousPodMaintenanceWeekEnumValues() []AutonomousPodMaintenanceWeekEnum {
	values := make([]AutonomousPodMaintenanceWeekEnum, 0)
	for _, v := range mappingAutonomousPodMaintenanceWeek {
		values = append(values, v)
	}
	return values
}

// AutonomousPodRoleEnum Enum with underlying type: string
type AutonomousPodRoleEnum string

// Set of constants representing the allowable values for AutonomousPodRoleEnum
const (
	AutonomousPodRolePrimary         AutonomousPodRoleEnum = "PRIMARY"
	AutonomousPodRoleStandby         AutonomousPodRoleEnum = "STANDBY"
	AutonomousPodRoleDisabledStandby AutonomousPodRoleEnum = "DISABLED_STANDBY"
)

var mappingAutonomousPodRole = map[string]AutonomousPodRoleEnum{
	"PRIMARY":          AutonomousPodRolePrimary,
	"STANDBY":          AutonomousPodRoleStandby,
	"DISABLED_STANDBY": AutonomousPodRoleDisabledStandby,
}

// GetAutonomousPodRoleEnumValues Enumerates the set of values for AutonomousPodRoleEnum
func GetAutonomousPodRoleEnumValues() []AutonomousPodRoleEnum {
	values := make([]AutonomousPodRoleEnum, 0)
	for _, v := range mappingAutonomousPodRole {
		values = append(values, v)
	}
	return values
}
