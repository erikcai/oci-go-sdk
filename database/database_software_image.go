// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// DatabaseSoftwareImage The Database Software Image is created using the patch set, one-off patches and patches for the database home listed by lsiventory.
type DatabaseSoftwareImage struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the database software image.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The database version with which the database software image is to be built.
	DatabaseVersion *string `mandatory:"true" json:"databaseVersion"`

	// The user-friendly name for the database software image. The name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The current state of the Database Software Image.
	LifecycleState DatabaseSoftwareImageLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the Database Software Image was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	ImageType DatabaseSoftwareImageImageTypeEnum `mandatory:"true" json:"imageType"`

	// To what shape the image is meant for.
	ImageShapeFamily DatabaseSoftwareImageImageShapeFamilyEnum `mandatory:"true" json:"imageShapeFamily"`

	// The PSU or PBP or Release Updates. To get a list of supported versions, use the ListDbVersions operation.
	PatchSet *string `mandatory:"true" json:"patchSet"`

	// Detailed message for the lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// List of one-off patches for Database Homes.
	DatabaseSoftwareImageIncludedPatches []string `mandatory:"false" json:"databaseSoftwareImageIncludedPatches"`

	// The patches included in the image and the version of the image
	IncludedPatchesSummary *string `mandatory:"false" json:"includedPatchesSummary"`
}

func (m DatabaseSoftwareImage) String() string {
	return common.PointerString(m)
}

// DatabaseSoftwareImageLifecycleStateEnum Enum with underlying type: string
type DatabaseSoftwareImageLifecycleStateEnum string

// Set of constants representing the allowable values for DatabaseSoftwareImageLifecycleStateEnum
const (
	DatabaseSoftwareImageLifecycleStateProvisioning DatabaseSoftwareImageLifecycleStateEnum = "PROVISIONING"
	DatabaseSoftwareImageLifecycleStateAvailable    DatabaseSoftwareImageLifecycleStateEnum = "AVAILABLE"
	DatabaseSoftwareImageLifecycleStateDeleting     DatabaseSoftwareImageLifecycleStateEnum = "DELETING"
	DatabaseSoftwareImageLifecycleStateDeleted      DatabaseSoftwareImageLifecycleStateEnum = "DELETED"
	DatabaseSoftwareImageLifecycleStateFailed       DatabaseSoftwareImageLifecycleStateEnum = "FAILED"
	DatabaseSoftwareImageLifecycleStateTerminating  DatabaseSoftwareImageLifecycleStateEnum = "TERMINATING"
	DatabaseSoftwareImageLifecycleStateTerminated   DatabaseSoftwareImageLifecycleStateEnum = "TERMINATED"
	DatabaseSoftwareImageLifecycleStateUpdating     DatabaseSoftwareImageLifecycleStateEnum = "UPDATING"
)

var mappingDatabaseSoftwareImageLifecycleState = map[string]DatabaseSoftwareImageLifecycleStateEnum{
	"PROVISIONING": DatabaseSoftwareImageLifecycleStateProvisioning,
	"AVAILABLE":    DatabaseSoftwareImageLifecycleStateAvailable,
	"DELETING":     DatabaseSoftwareImageLifecycleStateDeleting,
	"DELETED":      DatabaseSoftwareImageLifecycleStateDeleted,
	"FAILED":       DatabaseSoftwareImageLifecycleStateFailed,
	"TERMINATING":  DatabaseSoftwareImageLifecycleStateTerminating,
	"TERMINATED":   DatabaseSoftwareImageLifecycleStateTerminated,
	"UPDATING":     DatabaseSoftwareImageLifecycleStateUpdating,
}

// GetDatabaseSoftwareImageLifecycleStateEnumValues Enumerates the set of values for DatabaseSoftwareImageLifecycleStateEnum
func GetDatabaseSoftwareImageLifecycleStateEnumValues() []DatabaseSoftwareImageLifecycleStateEnum {
	values := make([]DatabaseSoftwareImageLifecycleStateEnum, 0)
	for _, v := range mappingDatabaseSoftwareImageLifecycleState {
		values = append(values, v)
	}
	return values
}

// DatabaseSoftwareImageImageTypeEnum Enum with underlying type: string
type DatabaseSoftwareImageImageTypeEnum string

// Set of constants representing the allowable values for DatabaseSoftwareImageImageTypeEnum
const (
	DatabaseSoftwareImageImageTypeGridImage     DatabaseSoftwareImageImageTypeEnum = "GRID_IMAGE"
	DatabaseSoftwareImageImageTypeDatabaseImage DatabaseSoftwareImageImageTypeEnum = "DATABASE_IMAGE"
)

var mappingDatabaseSoftwareImageImageType = map[string]DatabaseSoftwareImageImageTypeEnum{
	"GRID_IMAGE":     DatabaseSoftwareImageImageTypeGridImage,
	"DATABASE_IMAGE": DatabaseSoftwareImageImageTypeDatabaseImage,
}

// GetDatabaseSoftwareImageImageTypeEnumValues Enumerates the set of values for DatabaseSoftwareImageImageTypeEnum
func GetDatabaseSoftwareImageImageTypeEnumValues() []DatabaseSoftwareImageImageTypeEnum {
	values := make([]DatabaseSoftwareImageImageTypeEnum, 0)
	for _, v := range mappingDatabaseSoftwareImageImageType {
		values = append(values, v)
	}
	return values
}

// DatabaseSoftwareImageImageShapeFamilyEnum Enum with underlying type: string
type DatabaseSoftwareImageImageShapeFamilyEnum string

// Set of constants representing the allowable values for DatabaseSoftwareImageImageShapeFamilyEnum
const (
	DatabaseSoftwareImageImageShapeFamilyVmBmShape    DatabaseSoftwareImageImageShapeFamilyEnum = "VM_BM_SHAPE"
	DatabaseSoftwareImageImageShapeFamilyExadataShape DatabaseSoftwareImageImageShapeFamilyEnum = "EXADATA_SHAPE"
)

var mappingDatabaseSoftwareImageImageShapeFamily = map[string]DatabaseSoftwareImageImageShapeFamilyEnum{
	"VM_BM_SHAPE":   DatabaseSoftwareImageImageShapeFamilyVmBmShape,
	"EXADATA_SHAPE": DatabaseSoftwareImageImageShapeFamilyExadataShape,
}

// GetDatabaseSoftwareImageImageShapeFamilyEnumValues Enumerates the set of values for DatabaseSoftwareImageImageShapeFamilyEnum
func GetDatabaseSoftwareImageImageShapeFamilyEnumValues() []DatabaseSoftwareImageImageShapeFamilyEnum {
	values := make([]DatabaseSoftwareImageImageShapeFamilyEnum, 0)
	for _, v := range mappingDatabaseSoftwareImageImageShapeFamily {
		values = append(values, v)
	}
	return values
}
