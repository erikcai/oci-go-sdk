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

// InstanceDbSystemRole Describes the Role the Instance fills in relation to the DbSystem
type InstanceDbSystemRole struct {

	// Role played by this Instance in the DbSystem.
	Role InstanceRoleEnum `mandatory:"true" json:"role"`

	// The OCID of the DbSystem.
	Id *string `mandatory:"true" json:"id"`
}

func (m InstanceDbSystemRole) String() string {
	return common.PointerString(m)
}
