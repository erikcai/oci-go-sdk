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

// ApiOperation Details of each API operation for an Oracle Database release.
type ApiOperation struct {

	// Type of DB system operation.
	Name *string `mandatory:"true" json:"name"`

	// True if API operation is allowed for an Oracle Database release.
	IsAllowed *bool `mandatory:"true" json:"isAllowed"`

	// The DB system's source at the time of provisioning. The source can be a backup, a running database, or the new DB system itself (when the DB system is not provisioned from an existing database).
	DataSource []OperationFieldSummary `mandatory:"false" json:"dataSource"`

	// Lists the DB system properties that can be updated.
	Properties []OperationFieldSummary `mandatory:"false" json:"properties"`
}

func (m ApiOperation) String() string {
	return common.PointerString(m)
}
