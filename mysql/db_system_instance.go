// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// MySQL as a Service API
//
// The API for the MySQL Service
//

package mysql

import (
	"github.com/oracle/oci-go-sdk/common"
)

// DbSystemInstance Instance information that is accesible when regardless of AD availability.
type DbSystemInstance struct {

	// The OCID of the MySQLaaS Instance.
	Id *string `mandatory:"true" json:"id"`

	// Role played by this Instance in the DbSystem.
	Role InstanceRoleEnum `mandatory:"true" json:"role"`
}

func (m DbSystemInstance) String() string {
	return common.PointerString(m)
}
