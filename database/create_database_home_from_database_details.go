// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
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

// CreateDatabaseHomeFromDatabaseDetails CreateDatabaseHomeFromDatabaseDetails
type CreateDatabaseHomeFromDatabaseDetails struct {
	Database *CreateDatabaseFromAnotherDatabaseDetails `mandatory:"true" json:"database"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the DB system.
	DbSystemId *string `mandatory:"false" json:"dbSystemId"`

	// The user-provided name of the database home.
	DisplayName *string `mandatory:"false" json:"displayName"`
}

//GetDbSystemId returns DbSystemId
func (m CreateDatabaseHomeFromDatabaseDetails) GetDbSystemId() *string {
	return m.DbSystemId
}

//GetDisplayName returns DisplayName
func (m CreateDatabaseHomeFromDatabaseDetails) GetDisplayName() *string {
	return m.DisplayName
}

func (m CreateDatabaseHomeFromDatabaseDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateDatabaseHomeFromDatabaseDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateDatabaseHomeFromDatabaseDetails CreateDatabaseHomeFromDatabaseDetails
	s := struct {
		DiscriminatorParam string `json:"source"`
		MarshalTypeCreateDatabaseHomeFromDatabaseDetails
	}{
		"DATABASE_V2",
		(MarshalTypeCreateDatabaseHomeFromDatabaseDetails)(m),
	}

	return json.Marshal(&s)
}
