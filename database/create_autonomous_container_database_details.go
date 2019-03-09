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

	// The OCID of the Autonomous Exadata Infrastructure.
	AutonomousExadataInfrastructureId *string `mandatory:"true" json:"autonomousExadataInfrastructureId"`

	// Database Patch model preference.
	PatchModel CreateAutonomousContainerDatabaseDetailsPatchModelEnum `mandatory:"true" json:"patchModel"`

	// The service level agreement type of the Autonomous Container Database. The default is STANDARD. For a Mission Critical Container Database, the specified Autonomous Exadata Infrastructure must be associated with a remote Autonomous Exadata Infrastructure.
	ServiceLevelAgreementType CreateAutonomousContainerDatabaseDetailsServiceLevelAgreementTypeEnum `mandatory:"false" json:"serviceLevelAgreementType,omitempty"`

	// The OCID of the peer Autonomous Exadata Infrastructure for a Mission Critical Association
	PeerAutonomousExadataInfrastructureId *string `mandatory:"false" json:"peerAutonomousExadataInfrastructureId"`

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the compartment containing the Autonomous Container Database.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the last maintenance run.
	LastMaintenanceRunId *string `mandatory:"false" json:"lastMaintenanceRunId"`

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the next maintenance run.
	NextMaintenanceRunId *string `mandatory:"false" json:"nextMaintenanceRunId"`

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
