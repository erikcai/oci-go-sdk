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

// UpdateInstanceDetails The representation of UpdateInstanceDetails
type UpdateInstanceDetails struct {

	// DEPRECATED -- The Oracle Cloud ID (OCID) of the DbSystem the MySQLaaS instance belongs in.
	DbSystemId *string `mandatory:"false" json:"dbSystemId"`

	// DEPRECATED -- Role played by this Instance in the DbSystem.
	Role InstanceRoleEnum `mandatory:"false" json:"role,omitempty"`

	// DEPRECATED -- The OCID of the compartment to move the MySQLaas Instance to.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// User-provided data about the MySQLaaS Instance.
	Description *string `mandatory:"false" json:"description"`

	// The user-friendly name for the MySQLaaS Instance. It does not have to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// DEPRECATED -- this field will be only available on DbSystem once implementation catches up.
	// The username for the administrative user for the MySQLaaS Instance.
	AdminUsername *string `mandatory:"false" json:"adminUsername"`

	// DEPRECATED -- this field will be only available on DbSystem once implementation catches up.
	// The password for the administrative user. The password must be
	// between 8 and 32 characters long, and must contain at least 1
	// numeric character, 1 lowercase character, 1 uppercase character, and
	// 1 special (nonalphanumeric) character.
	AdminPassword *string `mandatory:"false" json:"adminPassword"`

	// DEPRECATED -- this field will be only available on DbSystem once implementation catches up.
	// The Availability Domain where the MySQLaaS instance is located.
	// Changes in Availability Domain will result in
	// downtime as the MySQLaaS instance and
	// storage are migrated to the new Compute instance.
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// The name of the Fault Domain the MySQLaaS instance is located in.
	// Changes in Fault Domain result in downtime as the MySQLaaS
	// instance is migrated to the new Fault Domain.
	FaultDomain *string `mandatory:"false" json:"faultDomain"`

	// DEPRECATED -- this field will be only available on DbSystem once implementation catches up.
	// The OCID of the subnet the MySQLaaS instance is associated with.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// DEPRECATED -- this field will be only available on DbSystem once implementation catches up.
	// The shape of the MySQLaaS instance. The shape determines resources
	// allocated to the MySQLaaS instance - CPU cores and memory for VM
	// shapes; CPU cores, memory and storage for non-VM (or bare metal)
	// shapes. To get a list of shapes, use (FIXME: correct link for
	// MySQLaaS shapes) the
	// ListShapes
	// operation.
	// Changes in Shape result in downtime as the MySQLaaS instance
	// is migrated to the new Compute instance.
	ShapeName *string `mandatory:"false" json:"shapeName"`

	// DEPRECATED -- this field will be only available on DbSystem once implementation catches up.
	// OCID of the MySQL Configuration associated with this instance.
	ConfigurationId *string `mandatory:"false" json:"configurationId"`

	// DEPRECATED -- this field will be only available on DbSystem once implementation catches up.
	// New size of the data volume in GBs that will be created and attached.
	// Increases in data storage size happen asynchronously and on-line.
	// Decreases in data storage size are not supported.
	DataStorageSizeInGBs *int `mandatory:"false" json:"dataStorageSizeInGBs"`

	// The port MySQLaaS instance is configured to listen on.
	// Changes to the port result in downtime as the service is restarted to bind to the new port.
	Port *int `mandatory:"false" json:"port"`

	// The network port on which X Plugin listens for TCP/IP connections. This is the X Plugin equivalent of port.
	// Changes to this port result in downtime as the service is restarted to bind to the new port.
	MysqlxPort *int `mandatory:"false" json:"mysqlxPort"`

	// The host name for the MySQLaaS Instance. The host name must begin
	// with an alphabetic character and can contain a maximum of 30
	// alphanumeric characters, including hyphens (-).
	// The maximum length of the combined hostname and domain is 63 characters.
	// **Note:** The hostname must be unique within the subnet. If it is not unique,
	// the MySQLaaS instance will fail to provision.
	// Changes in hostname will be done in real time. (FIXME: is this true?)
	Hostname *string `mandatory:"false" json:"hostname"`

	AnalyticsCluster *UpdateAnalyticsClusterDetails `mandatory:"false" json:"analyticsCluster"`

	// Simple key-value pair applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateInstanceDetails) String() string {
	return common.PointerString(m)
}
