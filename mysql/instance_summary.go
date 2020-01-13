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

// InstanceSummary MySQL-as-a-Service Instances Summary.
type InstanceSummary struct {

	// The OCID of the MySQLaaS instance.
	Id *string `mandatory:"true" json:"id"`

	// The name of the Availability Domain the MySQLaaS instance is located in.
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// DEPRECATED -- this field will be only available on DbSystem once implementation catches up.
	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// DEPRECATED -- this field will be only available on DbSystem once implementation catches up.
	// The shape of the MySQLaaS instance. The shape determines resources
	// to allocate to the MySQLaaS instance - CPU cores and memory for VM
	// shapes; CPU cores, memory and storage for non-VM (or bare metal)
	// shapes.
	ShapeName *string `mandatory:"true" json:"shapeName"`

	// DEPRECATED -- this field will be only available on DbSystem once implementation catches up.
	// Size of the data volume in GBs, defined in base-2, IEC units (GiBs).
	DataStorageSizeInGBs *int `mandatory:"true" json:"dataStorageSizeInGBs"`

	// The port MySQLaaS instance is configured to listen on.
	Port *int `mandatory:"true" json:"port"`

	// The current state of the MySQLaaS instance.
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the MySQLaaS instance was created, as described by RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the MySQLaaS instance was last updated, as described by RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The name of the Fault Domain the MySQLaaS instance is located in.
	FaultDomain *string `mandatory:"false" json:"faultDomain"`

	// The user-friendly name for the MySQLaaS instance. It does not have to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// User-provided data about the MySQLaaS Instance.
	Description *string `mandatory:"false" json:"description"`

	// The network host name for the MySQLaaS instance.
	Hostname *string `mandatory:"false" json:"hostname"`

	// DEPRECATED -- this field will be only available on DbSystem once implementation catches up.
	// The IP address MySQLaaS instance is configured to listen on.
	IpAddress *string `mandatory:"false" json:"ipAddress"`

	// The network port on which X Plugin listens for TCP/IP connections. This is the X Plugin equivalent of port.
	MysqlxPort *int `mandatory:"false" json:"mysqlxPort"`

	// The MySQL version used by the instance.
	MysqlVersion *string `mandatory:"false" json:"mysqlVersion"`

	AnalyticsCluster *AnalyticsClusterSummary `mandatory:"false" json:"analyticsCluster"`

	DbSystem *InstanceDbSystemRole `mandatory:"false" json:"dbSystem"`

	// Simple key-value pair applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m InstanceSummary) String() string {
	return common.PointerString(m)
}
