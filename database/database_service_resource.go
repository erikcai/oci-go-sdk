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

// DatabaseServiceResource Details of each database service resource and its corresponding operations.
type DatabaseServiceResource struct {

	// Resource Name.
	Name DatabaseServiceResourceNameEnum `mandatory:"true" json:"name"`

	// The DB system operations.
	ApiOperations []ApiOperation `mandatory:"true" json:"apiOperations"`
}

func (m DatabaseServiceResource) String() string {
	return common.PointerString(m)
}

// DatabaseServiceResourceNameEnum Enum with underlying type: string
type DatabaseServiceResourceNameEnum string

// Set of constants representing the allowable values for DatabaseServiceResourceNameEnum
const (
	DatabaseServiceResourceNameDbsystems          DatabaseServiceResourceNameEnum = "DBSYSTEMS"
	DatabaseServiceResourceNameDbnodes            DatabaseServiceResourceNameEnum = "DBNODES"
	DatabaseServiceResourceNameDbhomes            DatabaseServiceResourceNameEnum = "DBHOMES"
	DatabaseServiceResourceNameDatabases          DatabaseServiceResourceNameEnum = "DATABASES"
	DatabaseServiceResourceNameBackups            DatabaseServiceResourceNameEnum = "BACKUPS"
	DatabaseServiceResourceNameExternalbackupjobs DatabaseServiceResourceNameEnum = "EXTERNALBACKUPJOBS"
)

var mappingDatabaseServiceResourceName = map[string]DatabaseServiceResourceNameEnum{
	"DBSYSTEMS":          DatabaseServiceResourceNameDbsystems,
	"DBNODES":            DatabaseServiceResourceNameDbnodes,
	"DBHOMES":            DatabaseServiceResourceNameDbhomes,
	"DATABASES":          DatabaseServiceResourceNameDatabases,
	"BACKUPS":            DatabaseServiceResourceNameBackups,
	"EXTERNALBACKUPJOBS": DatabaseServiceResourceNameExternalbackupjobs,
}

// GetDatabaseServiceResourceNameEnumValues Enumerates the set of values for DatabaseServiceResourceNameEnum
func GetDatabaseServiceResourceNameEnumValues() []DatabaseServiceResourceNameEnum {
	values := make([]DatabaseServiceResourceNameEnum, 0)
	for _, v := range mappingDatabaseServiceResourceName {
		values = append(values, v)
	}
	return values
}
