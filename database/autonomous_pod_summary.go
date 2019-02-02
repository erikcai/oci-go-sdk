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

// AutonomousPodSummary An Autonomous Pod is a container database service that enables the customer to host one or more databases within the container database. There are two types of Autonomous Pods, the basic container database, and the Mission Critical container database. A basic container database runs on a single Autonomous Exadata Infrastructure from an availability domain without the Extreme Availability features enabled. A Mission Critical container database runs on two different Autonomous Exadata Infrastructures from two different availability domains, with Extreme Availability enabled.
type AutonomousPodSummary struct {

	// The OCID of the Autonomous Pod.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The user-provided name for the Autonomous Pod.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The service level agreement type of the container database. The default is STANDARD.
	ServiceLevelAgreementType AutonomousPodSummaryServiceLevelAgreementTypeEnum `mandatory:"true" json:"serviceLevelAgreementType"`

	// The OCID of the Autonomous DB System.
	AutonomousDbSystemId *string `mandatory:"true" json:"autonomousDbSystemId"`

	// The current state of the Autonomous Pod.
	LifecycleState AutonomousPodSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Database Patch model preference.
	PatchModel AutonomousPodSummaryPatchModelEnum `mandatory:"true" json:"patchModel"`

	// Maintenance Week preference, if set to WEEK_TWO, database patch will be applied in the second week during each patching cycle.
	MaintenanceWeek AutonomousPodSummaryMaintenanceWeekEnum `mandatory:"true" json:"maintenanceWeek"`

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

	// The role of the Autonomous Pod in this Mission Critical Association.
	Role AutonomousPodSummaryRoleEnum `mandatory:"false" json:"role,omitempty"`

	// The availability domain of the Autonomous Pod.
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	BackupConfig *AutonomousPodBackupConfig `mandatory:"false" json:"backupConfig"`
}

func (m AutonomousPodSummary) String() string {
	return common.PointerString(m)
}

// AutonomousPodSummaryServiceLevelAgreementTypeEnum Enum with underlying type: string
type AutonomousPodSummaryServiceLevelAgreementTypeEnum string

// Set of constants representing the allowable values for AutonomousPodSummaryServiceLevelAgreementTypeEnum
const (
	AutonomousPodSummaryServiceLevelAgreementTypeStandard        AutonomousPodSummaryServiceLevelAgreementTypeEnum = "STANDARD"
	AutonomousPodSummaryServiceLevelAgreementTypeMissionCritical AutonomousPodSummaryServiceLevelAgreementTypeEnum = "MISSION_CRITICAL"
)

var mappingAutonomousPodSummaryServiceLevelAgreementType = map[string]AutonomousPodSummaryServiceLevelAgreementTypeEnum{
	"STANDARD":         AutonomousPodSummaryServiceLevelAgreementTypeStandard,
	"MISSION_CRITICAL": AutonomousPodSummaryServiceLevelAgreementTypeMissionCritical,
}

// GetAutonomousPodSummaryServiceLevelAgreementTypeEnumValues Enumerates the set of values for AutonomousPodSummaryServiceLevelAgreementTypeEnum
func GetAutonomousPodSummaryServiceLevelAgreementTypeEnumValues() []AutonomousPodSummaryServiceLevelAgreementTypeEnum {
	values := make([]AutonomousPodSummaryServiceLevelAgreementTypeEnum, 0)
	for _, v := range mappingAutonomousPodSummaryServiceLevelAgreementType {
		values = append(values, v)
	}
	return values
}

// AutonomousPodSummaryLifecycleStateEnum Enum with underlying type: string
type AutonomousPodSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for AutonomousPodSummaryLifecycleStateEnum
const (
	AutonomousPodSummaryLifecycleStateProvisioning     AutonomousPodSummaryLifecycleStateEnum = "PROVISIONING"
	AutonomousPodSummaryLifecycleStateAvailable        AutonomousPodSummaryLifecycleStateEnum = "AVAILABLE"
	AutonomousPodSummaryLifecycleStateUpdating         AutonomousPodSummaryLifecycleStateEnum = "UPDATING"
	AutonomousPodSummaryLifecycleStateTerminating      AutonomousPodSummaryLifecycleStateEnum = "TERMINATING"
	AutonomousPodSummaryLifecycleStateTerminated       AutonomousPodSummaryLifecycleStateEnum = "TERMINATED"
	AutonomousPodSummaryLifecycleStateFailed           AutonomousPodSummaryLifecycleStateEnum = "FAILED"
	AutonomousPodSummaryLifecycleStateBackupInProgress AutonomousPodSummaryLifecycleStateEnum = "BACKUP_IN_PROGRESS"
	AutonomousPodSummaryLifecycleStateRestoring        AutonomousPodSummaryLifecycleStateEnum = "RESTORING"
	AutonomousPodSummaryLifecycleStateRestoreFailed    AutonomousPodSummaryLifecycleStateEnum = "RESTORE_FAILED"
)

var mappingAutonomousPodSummaryLifecycleState = map[string]AutonomousPodSummaryLifecycleStateEnum{
	"PROVISIONING":       AutonomousPodSummaryLifecycleStateProvisioning,
	"AVAILABLE":          AutonomousPodSummaryLifecycleStateAvailable,
	"UPDATING":           AutonomousPodSummaryLifecycleStateUpdating,
	"TERMINATING":        AutonomousPodSummaryLifecycleStateTerminating,
	"TERMINATED":         AutonomousPodSummaryLifecycleStateTerminated,
	"FAILED":             AutonomousPodSummaryLifecycleStateFailed,
	"BACKUP_IN_PROGRESS": AutonomousPodSummaryLifecycleStateBackupInProgress,
	"RESTORING":          AutonomousPodSummaryLifecycleStateRestoring,
	"RESTORE_FAILED":     AutonomousPodSummaryLifecycleStateRestoreFailed,
}

// GetAutonomousPodSummaryLifecycleStateEnumValues Enumerates the set of values for AutonomousPodSummaryLifecycleStateEnum
func GetAutonomousPodSummaryLifecycleStateEnumValues() []AutonomousPodSummaryLifecycleStateEnum {
	values := make([]AutonomousPodSummaryLifecycleStateEnum, 0)
	for _, v := range mappingAutonomousPodSummaryLifecycleState {
		values = append(values, v)
	}
	return values
}

// AutonomousPodSummaryPatchModelEnum Enum with underlying type: string
type AutonomousPodSummaryPatchModelEnum string

// Set of constants representing the allowable values for AutonomousPodSummaryPatchModelEnum
const (
	AutonomousPodSummaryPatchModelUpdates         AutonomousPodSummaryPatchModelEnum = "RELEASE_UPDATES"
	AutonomousPodSummaryPatchModelUpdateRevisions AutonomousPodSummaryPatchModelEnum = "RELEASE_UPDATE_REVISIONS"
)

var mappingAutonomousPodSummaryPatchModel = map[string]AutonomousPodSummaryPatchModelEnum{
	"RELEASE_UPDATES":          AutonomousPodSummaryPatchModelUpdates,
	"RELEASE_UPDATE_REVISIONS": AutonomousPodSummaryPatchModelUpdateRevisions,
}

// GetAutonomousPodSummaryPatchModelEnumValues Enumerates the set of values for AutonomousPodSummaryPatchModelEnum
func GetAutonomousPodSummaryPatchModelEnumValues() []AutonomousPodSummaryPatchModelEnum {
	values := make([]AutonomousPodSummaryPatchModelEnum, 0)
	for _, v := range mappingAutonomousPodSummaryPatchModel {
		values = append(values, v)
	}
	return values
}

// AutonomousPodSummaryMaintenanceWeekEnum Enum with underlying type: string
type AutonomousPodSummaryMaintenanceWeekEnum string

// Set of constants representing the allowable values for AutonomousPodSummaryMaintenanceWeekEnum
const (
	AutonomousPodSummaryMaintenanceWeekAny       AutonomousPodSummaryMaintenanceWeekEnum = "ANY"
	AutonomousPodSummaryMaintenanceWeekWeekTwo   AutonomousPodSummaryMaintenanceWeekEnum = "WEEK_TWO"
	AutonomousPodSummaryMaintenanceWeekWeekThree AutonomousPodSummaryMaintenanceWeekEnum = "WEEK_THREE"
	AutonomousPodSummaryMaintenanceWeekWeekFour  AutonomousPodSummaryMaintenanceWeekEnum = "WEEK_FOUR"
	AutonomousPodSummaryMaintenanceWeekWeekFive  AutonomousPodSummaryMaintenanceWeekEnum = "WEEK_FIVE"
	AutonomousPodSummaryMaintenanceWeekWeekSix   AutonomousPodSummaryMaintenanceWeekEnum = "WEEK_SIX"
)

var mappingAutonomousPodSummaryMaintenanceWeek = map[string]AutonomousPodSummaryMaintenanceWeekEnum{
	"ANY":        AutonomousPodSummaryMaintenanceWeekAny,
	"WEEK_TWO":   AutonomousPodSummaryMaintenanceWeekWeekTwo,
	"WEEK_THREE": AutonomousPodSummaryMaintenanceWeekWeekThree,
	"WEEK_FOUR":  AutonomousPodSummaryMaintenanceWeekWeekFour,
	"WEEK_FIVE":  AutonomousPodSummaryMaintenanceWeekWeekFive,
	"WEEK_SIX":   AutonomousPodSummaryMaintenanceWeekWeekSix,
}

// GetAutonomousPodSummaryMaintenanceWeekEnumValues Enumerates the set of values for AutonomousPodSummaryMaintenanceWeekEnum
func GetAutonomousPodSummaryMaintenanceWeekEnumValues() []AutonomousPodSummaryMaintenanceWeekEnum {
	values := make([]AutonomousPodSummaryMaintenanceWeekEnum, 0)
	for _, v := range mappingAutonomousPodSummaryMaintenanceWeek {
		values = append(values, v)
	}
	return values
}

// AutonomousPodSummaryRoleEnum Enum with underlying type: string
type AutonomousPodSummaryRoleEnum string

// Set of constants representing the allowable values for AutonomousPodSummaryRoleEnum
const (
	AutonomousPodSummaryRolePrimary         AutonomousPodSummaryRoleEnum = "PRIMARY"
	AutonomousPodSummaryRoleStandby         AutonomousPodSummaryRoleEnum = "STANDBY"
	AutonomousPodSummaryRoleDisabledStandby AutonomousPodSummaryRoleEnum = "DISABLED_STANDBY"
)

var mappingAutonomousPodSummaryRole = map[string]AutonomousPodSummaryRoleEnum{
	"PRIMARY":          AutonomousPodSummaryRolePrimary,
	"STANDBY":          AutonomousPodSummaryRoleStandby,
	"DISABLED_STANDBY": AutonomousPodSummaryRoleDisabledStandby,
}

// GetAutonomousPodSummaryRoleEnumValues Enumerates the set of values for AutonomousPodSummaryRoleEnum
func GetAutonomousPodSummaryRoleEnumValues() []AutonomousPodSummaryRoleEnum {
	values := make([]AutonomousPodSummaryRoleEnum, 0)
	for _, v := range mappingAutonomousPodSummaryRole {
		values = append(values, v)
	}
	return values
}
