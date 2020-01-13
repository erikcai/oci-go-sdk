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

// AnalyticsClusterNode The representation of AnalyticsClusterNode
type AnalyticsClusterNode struct {

	// The ID of the node within MySQLaaS Analytics Cluster.
	NodeId *string `mandatory:"true" json:"nodeId"`

	// The current state of the MySQLaaS Analytics Cluster node.
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the MySQLaaS Analytics Cluster node was created, as described by RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m AnalyticsClusterNode) String() string {
	return common.PointerString(m)
}
