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

// CreateDatabaseFromBackupDetails The representation of CreateDatabaseFromBackupDetails
type CreateDatabaseFromBackupDetails struct {

	// A strong password for SYS, SYSTEM, PDB Admin and TDE Wallet. The password must be at least nine characters and contain at least two uppercase, two lowercase, two numbers, and two special characters. The special characters must be _, \#, or -.
	AdminPassword *string `mandatory:"true" json:"adminPassword"`

	// The backup OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).
	BackupId *string `mandatory:"true" json:"backupId"`

	// The password to open the TDE wallet.
	BackupTDEPassword *string `mandatory:"true" json:"backupTDEPassword"`

	// The database domain. In a distributed database system, DB_DOMAIN specifies the logical location of the database within the network structure.
	DbDomain *string `mandatory:"false" json:"dbDomain"`

	// New database name. It must begin with an alphabetic character and can contain a maximum of eight alphanumeric characters. Special characters are not permitted.
	DbName *string `mandatory:"false" json:"dbName"`
}

func (m CreateDatabaseFromBackupDetails) String() string {
	return common.PointerString(m)
}
