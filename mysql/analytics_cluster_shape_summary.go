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

// AnalyticsClusterShapeSummary The shape of the MySQLaaS Analytics Cluster. The shape determines
// resources to allocate to each node in the MySQLaaS Analytics Cluster.
// For a description of shapes, see MySQLaaS Analytics Cluster Create
// Options (https://docs.cloud.oracle.com/Content/MySQL/References/MySqlAnalyticslaunchoptions.htm).  To
// use any of the API operations, you must be authorized in an IAM
// policy. If you are not authorized, talk to an administrator.  If you are
// an administrator who needs to write policies to give users access, see
// Getting Started with
// Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
type AnalyticsClusterShapeSummary struct {

	// The name of the shape used for each compute node in MySQLaaS Analytics Cluster.
	Name *string `mandatory:"true" json:"name"`

	// The number of CPU Cores the Instance provides. These are "OCPU"s.
	CpuCoreCount *int `mandatory:"true" json:"cpuCoreCount"`

	// The amount of RAM the Instance provides. This is an IEC base-2 number.
	MemorySizeInGBs *int `mandatory:"true" json:"memorySizeInGBs"`
}

func (m AnalyticsClusterShapeSummary) String() string {
	return common.PointerString(m)
}
