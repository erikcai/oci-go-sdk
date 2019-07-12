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

// ChangeExadataInfrastructureCompartmentDetails Details regarding target compartment id
type ChangeExadataInfrastructureCompartmentDetails struct {

	// The Oracle Cloud ID (OCID) of the compartment the resource belongs in.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`
}

func (m ChangeExadataInfrastructureCompartmentDetails) String() string {
	return common.PointerString(m)
}
