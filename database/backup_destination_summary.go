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

// BackupDestinationSummary A database backup destination summary, databases field will include an array of all the databases using this backup destination.
type BackupDestinationSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the database backup destination.
	Id *string `mandatory:"false" json:"id"`

	// The user-provided name of the database backup destination.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Type of the database backup destination.
	Type BackupDestinationSummaryTypeEnum `mandatory:"false" json:"type,omitempty"`

	// List of databases associated to this backup destination
	AssociatedDatabases []AssociatedDatabaseDetails `mandatory:"false" json:"associatedDatabases"`

	// The connection string that is used to connect to the Zero Data Loss Recovery Appliance.
	ConnectionString *string `mandatory:"false" json:"connectionString"`

	// The Virtual Private Catalog users that will be used to access the Zero Data Loss Recovery Appliance.
	VpcUsers []string `mandatory:"false" json:"vpcUsers"`

	// The network file path of the NFS device to be mounted. In the format server:/directory/folder.
	Path *string `mandatory:"false" json:"path"`

	// The current lifecycle state of the backup destination.
	LifecycleState BackupDestinationSummaryLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

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

func (m BackupDestinationSummary) String() string {
	return common.PointerString(m)
}

// BackupDestinationSummaryTypeEnum Enum with underlying type: string
type BackupDestinationSummaryTypeEnum string

// Set of constants representing the allowable values for BackupDestinationSummaryTypeEnum
const (
	BackupDestinationSummaryTypeNfs               BackupDestinationSummaryTypeEnum = "NFS"
	BackupDestinationSummaryTypeRecoveryAppliance BackupDestinationSummaryTypeEnum = "RECOVERY_APPLIANCE"
)

var mappingBackupDestinationSummaryType = map[string]BackupDestinationSummaryTypeEnum{
	"NFS":                BackupDestinationSummaryTypeNfs,
	"RECOVERY_APPLIANCE": BackupDestinationSummaryTypeRecoveryAppliance,
}

// GetBackupDestinationSummaryTypeEnumValues Enumerates the set of values for BackupDestinationSummaryTypeEnum
func GetBackupDestinationSummaryTypeEnumValues() []BackupDestinationSummaryTypeEnum {
	values := make([]BackupDestinationSummaryTypeEnum, 0)
	for _, v := range mappingBackupDestinationSummaryType {
		values = append(values, v)
	}
	return values
}

// BackupDestinationSummaryLifecycleStateEnum Enum with underlying type: string
type BackupDestinationSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for BackupDestinationSummaryLifecycleStateEnum
const (
	BackupDestinationSummaryLifecycleStateActive  BackupDestinationSummaryLifecycleStateEnum = "ACTIVE"
	BackupDestinationSummaryLifecycleStateFailed  BackupDestinationSummaryLifecycleStateEnum = "FAILED"
	BackupDestinationSummaryLifecycleStateDeleted BackupDestinationSummaryLifecycleStateEnum = "DELETED"
)

var mappingBackupDestinationSummaryLifecycleState = map[string]BackupDestinationSummaryLifecycleStateEnum{
	"ACTIVE":  BackupDestinationSummaryLifecycleStateActive,
	"FAILED":  BackupDestinationSummaryLifecycleStateFailed,
	"DELETED": BackupDestinationSummaryLifecycleStateDeleted,
}

// GetBackupDestinationSummaryLifecycleStateEnumValues Enumerates the set of values for BackupDestinationSummaryLifecycleStateEnum
func GetBackupDestinationSummaryLifecycleStateEnumValues() []BackupDestinationSummaryLifecycleStateEnum {
	values := make([]BackupDestinationSummaryLifecycleStateEnum, 0)
	for _, v := range mappingBackupDestinationSummaryLifecycleState {
		values = append(values, v)
	}
	return values
}
