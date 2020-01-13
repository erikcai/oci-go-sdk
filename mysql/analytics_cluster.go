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

// AnalyticsCluster The representation of AnalyticsCluster
type AnalyticsCluster struct {

	// The OCID of the MySQLaaS Analytics Cluster.
	Id *string `mandatory:"true" json:"id"`

	// OCID of the Compartment the MySQLaaS Analytics Cluster exists in.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the MySQLaaS Instance this MySQLaaS Analytics Cluster is attached to.
	InstanceId *string `mandatory:"true" json:"instanceId"`

	// The shape determines resources to allocate to the MySQLaaS Analyics
	// Cluster nodes - CPU cores, memory.
	ShapeName *string `mandatory:"true" json:"shapeName"`

	// The number of Analytics-processing compute instances, of the
	// specified shape, in the Analytics Cluster.
	ClusterSize *int `mandatory:"true" json:"clusterSize"`

	ClusterNodes []AnalyticsClusterNode `mandatory:"true" json:"clusterNodes"`

	// The current state of the MySQLaaS Analytics Cluster.
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the MySQLaaS Analytics Cluster was created, as described by RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the Analytics Cluster was last updated, as described by RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// User-provided data about the MySQLaaS Instance.
	Description *string `mandatory:"false" json:"description"`

	Instance *InstanceSummary `mandatory:"false" json:"instance"`

	// Additional information about the current lifecycleState.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// [PRIVATE API] OCID of the host image used for this Analytics Cluster. If not set, the control plane default will be used.
	HostImageId *string `mandatory:"false" json:"hostImageId"`

	// Simple key-value pair applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m AnalyticsCluster) String() string {
	return common.PointerString(m)
}
