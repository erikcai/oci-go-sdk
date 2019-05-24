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

// CreateDatabaseFromAnotherDatabaseDetails The representation of CreateDatabaseFromAnotherDatabaseDetails
type CreateDatabaseFromAnotherDatabaseDetails struct {

	// The database OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	DatabaseId *string `mandatory:"true" json:"databaseId"`

	// The password to open the TDE wallet.
	BackupTDEPassword *string `mandatory:"true" json:"backupTDEPassword"`

	// A strong password for SYS, SYSTEM, PDB Admin and TDE Wallet. The password must be at least nine characters and contain at least two uppercase, two lowercase, two numbers, and two special characters. The special characters must be _, \#, or -.
	AdminPassword *string `mandatory:"true" json:"adminPassword"`

	// The display name of the database to be created from the backup. It must begin with an alphabetic character and can contain a maximum of eight alphanumeric characters. Special characters are not permitted.
	DbName *string `mandatory:"false" json:"dbName"`

	// The point in time of the original database from which the new database is created. If not specifed, the latest backup is used to create the database.
	TimeStampForPointInTimeRecovery *common.SDKTime `mandatory:"false" json:"timeStampForPointInTimeRecovery"`
}

func (m CreateDatabaseFromAnotherDatabaseDetails) String() string {
	return common.PointerString(m)
}
