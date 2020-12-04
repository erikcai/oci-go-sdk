// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management Service APIs.
//
// This file contains the customer facing APIs for Database Management service.
//

package databasemanagement

import (
	"github.com/oracle/oci-go-sdk/v30/common"
)

// ManagedDatabase An Oracle database that is being managed.
type ManagedDatabase struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name of the managed database.
	Name *string `mandatory:"true" json:"name"`

	// The type of Oracle database installation.
	DatabaseType DatabaseTypeEnum `mandatory:"true" json:"databaseType"`

	// Indicates whether an Oracle database is a Container/Pluggable/Non-Container database.
	DatabaseSubType DatabaseSubTypeEnum `mandatory:"true" json:"databaseSubType"`

	// Indicates whether the database is part of a cluster.
	IsCluster *bool `mandatory:"true" json:"isCluster"`

	// The date and time the Managed Database was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the parent Container Database (CDB)
	// if this database is a Pluggable Database (PDB).
	ParentContainerId *string `mandatory:"false" json:"parentContainerId"`

	// List of Managed Database Groups that this database belongs to.
	ManagedDatabaseGroups []ParentGroup `mandatory:"false" json:"managedDatabaseGroups"`

	// Indicates whether database status is Up/Down/Unknown.
	DatabaseStatus DatabaseStatusEnum `mandatory:"false" json:"databaseStatus,omitempty"`

	// Additional database details specific to a kind of database defined by database properties in `{"key": "value"}` format.
	// Example: `{"bar-key": "value"}`
	AdditionalDetails map[string]string `mandatory:"false" json:"additionalDetails"`
}

func (m ManagedDatabase) String() string {
	return common.PointerString(m)
}
