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

// GetTargetDbHome During move database operation, the target home ID needs to be checked for the status of the move operation
type GetTargetDbHome struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the dbHome.
	TargetDbHomeId *string `mandatory:"true" json:"targetDbHomeId"`
}

func (m GetTargetDbHome) String() string {
	return common.PointerString(m)
}
