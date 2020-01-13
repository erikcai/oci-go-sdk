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

// UpdateAnalyticsClusterDetails Details of the MySQLaaS Analytics Cluster.
type UpdateAnalyticsClusterDetails struct {

	// User-provided data about the MySQLaaS Instance.
	Description *string `mandatory:"false" json:"description"`

	// A change to the shape of the nodes in the Analytics Cluster will
	// result in the entire cluster being torn down and re-created with
	// Compute instances of the new Shape. This may result in significant
	// downtime for the analytics capability while the Analytics Cluster is
	// re-provisioned.
	ShapeName *string `mandatory:"false" json:"shapeName"`

	// A change to the number of nodes in the Analytics Cluster will result
	// in the entire cluster being torn down and re-created with the new
	// cluster of nodes. This may result in a significant downtime for the
	// analytics capability while the Analytics Cluster is
	// re-provisioned.
	ClusterSize *int `mandatory:"false" json:"clusterSize"`

	// Simple key-value pair applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateAnalyticsClusterDetails) String() string {
	return common.PointerString(m)
}
