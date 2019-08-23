// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// CreateDatabaseFromDatabase Details for creating a database by restoring from a database backup.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type CreateDatabaseFromDatabase struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the database home.
	DbHomeId *string `mandatory:"true" json:"dbHomeId"`

	Database *CreateDatabaseFromAnotherDatabaseDetails `mandatory:"true" json:"database"`

	// A valid Oracle Database version. To get a list of supported versions, use the ListDbVersions operation.
	DbVersion *string `mandatory:"false" json:"dbVersion"`
}

//GetDbHomeId returns DbHomeId
func (m CreateDatabaseFromDatabase) GetDbHomeId() *string {
	return m.DbHomeId
}

//GetDbVersion returns DbVersion
func (m CreateDatabaseFromDatabase) GetDbVersion() *string {
	return m.DbVersion
}

func (m CreateDatabaseFromDatabase) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateDatabaseFromDatabase) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateDatabaseFromDatabase CreateDatabaseFromDatabase
	s := struct {
		DiscriminatorParam string `json:"source"`
		MarshalTypeCreateDatabaseFromDatabase
	}{
		"DATABASE",
		(MarshalTypeCreateDatabaseFromDatabase)(m),
	}

	return json.Marshal(&s)
}
