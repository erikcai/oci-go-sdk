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

// CreateAnalyticsClusterDetails Details required to create a MySQLaaS Analytics Cluster.
type CreateAnalyticsClusterDetails struct {

	// The shape determines resources to allocate to the MySQLaaS Analyics
	// Cluster nodes - CPU cores, memory.
	ShapeName *string `mandatory:"true" json:"shapeName"`

	// The number of analytics-processing nodes provisioned for the
	// MySQLaaS Analytics Cluster.
	ClusterSize *int `mandatory:"true" json:"clusterSize"`

	// DEPRECATED -- this field will be only available on DbSystem once implementation catches up.
	// The OCID of the MySQLaaS Instance this Analytics Cluster will be associated with.
	InstanceId *string `mandatory:"false" json:"instanceId"`

	// User-provided data about the MySQLaaS Instance.
	Description *string `mandatory:"false" json:"description"`

	// [PRIVATE API] OCID of the host image used for this Analytics Cluster. If not set, the control plane default will be used.
	HostImageId *string `mandatory:"false" json:"hostImageId"`

	// Simple key-value pair applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateAnalyticsClusterDetails) String() string {
	return common.PointerString(m)
}
