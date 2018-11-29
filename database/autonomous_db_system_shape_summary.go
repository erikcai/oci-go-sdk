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

// AutonomousDbSystemShapeSummary The shape of the Autonomous DB System. The shape determines resources to allocate to the Autonomous DB system -  CPU cores, memory and storage for shapes.
// For a description of shapes, see Auotonomous DB System Launch Options (https://docs.us-phoenix-1.oraclecloud.com/Content/Database/References/launchoptions.htm).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an administrator.
// If you're an administrator who needs to write policies to give users access,
// see Getting Started with Policies (https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/policygetstarted.htm).
type AutonomousDbSystemShapeSummary struct {

	// The name of the shape used for the Autonomous DB System.
	Name *string `mandatory:"true" json:"name"`

	// The maximum number of CPU cores that can be enabled on the Autonomoous DB System.
	AvailableCoreCount *int `mandatory:"true" json:"availableCoreCount"`

	// The minimum number of CPU cores that can be enabled on the Autonomous DB System.
	MinimumCoreCount *int `mandatory:"false" json:"minimumCoreCount"`

	// The increment in which core count can be increased or decreased.
	CoreCountIncrement *int `mandatory:"false" json:"coreCountIncrement"`

	// The minimum number of nodes available for the shape.
	MinimumNodeCount *int `mandatory:"false" json:"minimumNodeCount"`

	// The maximum number of nodes available for the shape.
	MaximumNodeCount *int `mandatory:"false" json:"maximumNodeCount"`
}

func (m AutonomousDbSystemShapeSummary) String() string {
	return common.PointerString(m)
}