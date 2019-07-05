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

// CreateDbHomeWithDbSystemIdFromDatabaseDetails Note that a valid `dbSystemId` value must be supplied for the `CreateDbHomeWithDbSystemIdFromDatabase` API operation to successfully complete.
type CreateDbHomeWithDbSystemIdFromDatabaseDetails struct {
	Database *CreateDatabaseFromAnotherDatabaseDetails `mandatory:"true" json:"database"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the DB system.
	DbSystemId *string `mandatory:"false" json:"dbSystemId"`

	// The user-provided name of the database home.
	DisplayName *string `mandatory:"false" json:"displayName"`
}

//GetDbSystemId returns DbSystemId
func (m CreateDbHomeWithDbSystemIdFromDatabaseDetails) GetDbSystemId() *string {
	return m.DbSystemId
}

//GetDisplayName returns DisplayName
func (m CreateDbHomeWithDbSystemIdFromDatabaseDetails) GetDisplayName() *string {
	return m.DisplayName
}

func (m CreateDbHomeWithDbSystemIdFromDatabaseDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateDbHomeWithDbSystemIdFromDatabaseDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateDbHomeWithDbSystemIdFromDatabaseDetails CreateDbHomeWithDbSystemIdFromDatabaseDetails
	s := struct {
		DiscriminatorParam string `json:"source"`
		MarshalTypeCreateDbHomeWithDbSystemIdFromDatabaseDetails
	}{
		"DATABASE",
		(MarshalTypeCreateDbHomeWithDbSystemIdFromDatabaseDetails)(m),
	}

	return json.Marshal(&s)
}
