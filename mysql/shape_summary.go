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

// ShapeSummary The shape of the MySQLaaS Instance. The shape determines resources to
// allocate to the MySQLaaS Instance - CPU cores and memory for VM shapes;
// CPU cores, memory and storage for non-VM (or bare metal) shapes.  For a
// description of shapes, see MySQLaaS Instance Create
// Options (https://docs.cloud.oracle.com/Content/MySQL/References/launchoptions.htm).  To use any of
// the API operations, you must be authorized in an IAM policy. If you're
// not authorized, talk to an administrator.  If you are an administrator
// who needs to write policies to give users access, see Getting Started
// with Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
type ShapeSummary struct {

	// The name of the shape used for the MySQLaaS Instance.
	Name *string `mandatory:"true" json:"name"`

	// The number of CPU Cores the Instance provides. These are "OCPU"s.
	CpuCoreCount *int `mandatory:"true" json:"cpuCoreCount"`

	// The amount of RAM the Instance provides. This is an IEC base-2 number.
	MemorySizeInGBs *int `mandatory:"true" json:"memorySizeInGBs"`
}

func (m ShapeSummary) String() string {
	return common.PointerString(m)
}
