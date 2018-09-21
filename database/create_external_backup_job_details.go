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

// CreateExternalBackupJobDetails The representation of CreateExternalBackupJobDetails
type CreateExternalBackupJobDetails struct {

	// The targeted availability domain for the backup.
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The character set for the database.
	CharacterSet *string `mandatory:"true" json:"characterSet"`

	// The OCID of the compartment where this backup should be created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The Oracle database edition to use for restoring this backup.
	// Note that 2 node RAC DB Systems require ENTERPRISE_EDITION_EXTREME_PERFORMANCE.
	DatabaseEdition CreateExternalBackupJobDetailsDatabaseEditionEnum `mandatory:"true" json:"databaseEdition"`

	// The mode of the database.
	DatabaseMode CreateExternalBackupJobDetailsDatabaseModeEnum `mandatory:"true" json:"databaseMode"`

	// The name of the database from which the backup is being taken.
	DbName *string `mandatory:"true" json:"dbName"`

	// A valid Oracle database version.
	DbVersion *string `mandatory:"true" json:"dbVersion"`

	// A user-friendly name for the backup. The displayName value does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The dbid of the Oracle database being backed up.
	ExternalDatabaseIdentifier *int64 `mandatory:"true" json:"externalDatabaseIdentifier"`

	// The national character set for the database.
	NcharacterSet *string `mandatory:"true" json:"ncharacterSet"`

	// The dbunique name of the database being backed up.
	DbUniqueName *string `mandatory:"false" json:"dbUniqueName"`

	// The pluggable database name.
	PdbName *string `mandatory:"false" json:"pdbName"`
}

func (m CreateExternalBackupJobDetails) String() string {
	return common.PointerString(m)
}

// CreateExternalBackupJobDetailsDatabaseEditionEnum Enum with underlying type: string
type CreateExternalBackupJobDetailsDatabaseEditionEnum string

// Set of constants representing the allowable values for CreateExternalBackupJobDetailsDatabaseEdition
const (
	CreateExternalBackupJobDetailsDatabaseEditionStandardEdition                     CreateExternalBackupJobDetailsDatabaseEditionEnum = "STANDARD_EDITION"
	CreateExternalBackupJobDetailsDatabaseEditionEnterpriseEdition                   CreateExternalBackupJobDetailsDatabaseEditionEnum = "ENTERPRISE_EDITION"
	CreateExternalBackupJobDetailsDatabaseEditionEnterpriseEditionHighPerformance    CreateExternalBackupJobDetailsDatabaseEditionEnum = "ENTERPRISE_EDITION_HIGH_PERFORMANCE"
	CreateExternalBackupJobDetailsDatabaseEditionEnterpriseEditionExtremePerformance CreateExternalBackupJobDetailsDatabaseEditionEnum = "ENTERPRISE_EDITION_EXTREME_PERFORMANCE"
)

var mappingCreateExternalBackupJobDetailsDatabaseEdition = map[string]CreateExternalBackupJobDetailsDatabaseEditionEnum{
	"STANDARD_EDITION":                       CreateExternalBackupJobDetailsDatabaseEditionStandardEdition,
	"ENTERPRISE_EDITION":                     CreateExternalBackupJobDetailsDatabaseEditionEnterpriseEdition,
	"ENTERPRISE_EDITION_HIGH_PERFORMANCE":    CreateExternalBackupJobDetailsDatabaseEditionEnterpriseEditionHighPerformance,
	"ENTERPRISE_EDITION_EXTREME_PERFORMANCE": CreateExternalBackupJobDetailsDatabaseEditionEnterpriseEditionExtremePerformance,
}

// GetCreateExternalBackupJobDetailsDatabaseEditionEnumValues Enumerates the set of values for CreateExternalBackupJobDetailsDatabaseEdition
func GetCreateExternalBackupJobDetailsDatabaseEditionEnumValues() []CreateExternalBackupJobDetailsDatabaseEditionEnum {
	values := make([]CreateExternalBackupJobDetailsDatabaseEditionEnum, 0)
	for _, v := range mappingCreateExternalBackupJobDetailsDatabaseEdition {
		values = append(values, v)
	}
	return values
}

// CreateExternalBackupJobDetailsDatabaseModeEnum Enum with underlying type: string
type CreateExternalBackupJobDetailsDatabaseModeEnum string

// Set of constants representing the allowable values for CreateExternalBackupJobDetailsDatabaseMode
const (
	CreateExternalBackupJobDetailsDatabaseModeSi  CreateExternalBackupJobDetailsDatabaseModeEnum = "SI"
	CreateExternalBackupJobDetailsDatabaseModeRac CreateExternalBackupJobDetailsDatabaseModeEnum = "RAC"
)

var mappingCreateExternalBackupJobDetailsDatabaseMode = map[string]CreateExternalBackupJobDetailsDatabaseModeEnum{
	"SI":  CreateExternalBackupJobDetailsDatabaseModeSi,
	"RAC": CreateExternalBackupJobDetailsDatabaseModeRac,
}

// GetCreateExternalBackupJobDetailsDatabaseModeEnumValues Enumerates the set of values for CreateExternalBackupJobDetailsDatabaseMode
func GetCreateExternalBackupJobDetailsDatabaseModeEnumValues() []CreateExternalBackupJobDetailsDatabaseModeEnum {
	values := make([]CreateExternalBackupJobDetailsDatabaseModeEnum, 0)
	for _, v := range mappingCreateExternalBackupJobDetailsDatabaseMode {
		values = append(values, v)
	}
	return values
}
