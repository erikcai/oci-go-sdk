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

// ChangeAutonomousVmClusterCompartmentDetails The configuration details for moving the Autonomous VM cluster.
type ChangeAutonomousVmClusterCompartmentDetails struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to move the Autonomous VM cluster to.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`
}

func (m ChangeAutonomousVmClusterCompartmentDetails) String() string {
	return common.PointerString(m)
}
