// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// BackupDestination Backup Destination Details.
type BackupDestination struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the database backup destination.
	Id *string `mandatory:"false" json:"id"`

	// The user-provided name of the database backup destination.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Type of the database backup destination.
	Type BackupDestinationTypeEnum `mandatory:"false" json:"type,omitempty"`

	// List of databases associated to this backup destination
	AssociatedDatabases []AssociatedDatabaseDetails `mandatory:"false" json:"associatedDatabases"`

	// The connection string that is used to connect to the Zero Data Loss Recovery Appliance.
	ConnectionString *string `mandatory:"false" json:"connectionString"`

	// The Virtual Private Catalog users that will be used to access the Zero Data Loss Recovery Appliance.
	VpcUsers []string `mandatory:"false" json:"vpcUsers"`

	// The network file path of the NFS device to be mounted. In the format server:/directory/folder.
	Path *string `mandatory:"false" json:"path"`

	// The current lifecycle state of the backup destination.
	LifecycleState BackupDestinationLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The date and time the backup destination was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// A descriptive text associated with the lifecycleState.
	// Typically contains additional displayable text
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m BackupDestination) String() string {
	return common.PointerString(m)
}

// BackupDestinationTypeEnum Enum with underlying type: string
type BackupDestinationTypeEnum string

// Set of constants representing the allowable values for BackupDestinationTypeEnum
const (
	BackupDestinationTypeNfs               BackupDestinationTypeEnum = "NFS"
	BackupDestinationTypeRecoveryAppliance BackupDestinationTypeEnum = "RECOVERY_APPLIANCE"
)

var mappingBackupDestinationType = map[string]BackupDestinationTypeEnum{
	"NFS":                BackupDestinationTypeNfs,
	"RECOVERY_APPLIANCE": BackupDestinationTypeRecoveryAppliance,
}

// GetBackupDestinationTypeEnumValues Enumerates the set of values for BackupDestinationTypeEnum
func GetBackupDestinationTypeEnumValues() []BackupDestinationTypeEnum {
	values := make([]BackupDestinationTypeEnum, 0)
	for _, v := range mappingBackupDestinationType {
		values = append(values, v)
	}
	return values
}

// BackupDestinationLifecycleStateEnum Enum with underlying type: string
type BackupDestinationLifecycleStateEnum string

// Set of constants representing the allowable values for BackupDestinationLifecycleStateEnum
const (
	BackupDestinationLifecycleStateActive  BackupDestinationLifecycleStateEnum = "ACTIVE"
	BackupDestinationLifecycleStateFailed  BackupDestinationLifecycleStateEnum = "FAILED"
	BackupDestinationLifecycleStateDeleted BackupDestinationLifecycleStateEnum = "DELETED"
)

var mappingBackupDestinationLifecycleState = map[string]BackupDestinationLifecycleStateEnum{
	"ACTIVE":  BackupDestinationLifecycleStateActive,
	"FAILED":  BackupDestinationLifecycleStateFailed,
	"DELETED": BackupDestinationLifecycleStateDeleted,
}

// GetBackupDestinationLifecycleStateEnumValues Enumerates the set of values for BackupDestinationLifecycleStateEnum
func GetBackupDestinationLifecycleStateEnumValues() []BackupDestinationLifecycleStateEnum {
	values := make([]BackupDestinationLifecycleStateEnum, 0)
	for _, v := range mappingBackupDestinationLifecycleState {
		values = append(values, v)
	}
	return values
}
