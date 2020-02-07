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

// OperationFieldSummary The database operation field summary.
type OperationFieldSummary struct {

	// The DB system or DB system resource operation name.
	Name *string `mandatory:"true" json:"name"`

	// True if operation is allowed for an Oracle Database release.
	IsAllowed *bool `mandatory:"true" json:"isAllowed"`
}

func (m OperationFieldSummary) String() string {
	return common.PointerString(m)
}
