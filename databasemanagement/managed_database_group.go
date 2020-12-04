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

// ManagedDatabaseGroup Detailed information about a Managed Database Group.
type ManagedDatabaseGroup struct {

	// The name of the Managed Database Group.
	Name *string `mandatory:"true" json:"name"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database Group.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// List of Managed Databases in the group.
	ManagedDatabases []ChildDatabase `mandatory:"true" json:"managedDatabases"`

	// The current state of the Managed Database Group.
	LifecycleState LifecycleStatesEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the Managed Database Group was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Information specified by the user about the Managed Database Group.
	Description *string `mandatory:"false" json:"description"`

	// The date and time the Managed Database Group was last updated.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m ManagedDatabaseGroup) String() string {
	return common.PointerString(m)
}
