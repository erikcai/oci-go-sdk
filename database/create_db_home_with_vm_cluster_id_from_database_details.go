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

// CreateDbHomeWithVmClusterIdFromDatabaseDetails Note that a valid `vmClusterId` value must be supplied for the `CreateDbHomeWithVmClusterIdFromDatabase` API operation to successfully complete.
type CreateDbHomeWithVmClusterIdFromDatabaseDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the DB system.
	VmClusterId *string `mandatory:"true" json:"vmClusterId"`

	Database *CreateDatabaseFromAnotherDatabaseDetails `mandatory:"true" json:"database"`

	// The user-provided name of the database home.
	DisplayName *string `mandatory:"false" json:"displayName"`
}

//GetDisplayName returns DisplayName
func (m CreateDbHomeWithVmClusterIdFromDatabaseDetails) GetDisplayName() *string {
	return m.DisplayName
}

func (m CreateDbHomeWithVmClusterIdFromDatabaseDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateDbHomeWithVmClusterIdFromDatabaseDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateDbHomeWithVmClusterIdFromDatabaseDetails CreateDbHomeWithVmClusterIdFromDatabaseDetails
	s := struct {
		DiscriminatorParam string `json:"source"`
		MarshalTypeCreateDbHomeWithVmClusterIdFromDatabaseDetails
	}{
		"VM_CLUSTER_DATABASE",
		(MarshalTypeCreateDbHomeWithVmClusterIdFromDatabaseDetails)(m),
	}

	return json.Marshal(&s)
}
