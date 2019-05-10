// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ChangeCompartmentDetails Details regarding target compartment id
type ChangeCompartmentDetails struct {

	// The Oracle Cloud ID (OCID) of the compartment the resource belongs in.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`
}

func (m ChangeCompartmentDetails) String() string {
	return common.PointerString(m)
}
