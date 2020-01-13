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

// CreateInstanceDetails Details required to create a new instance in an existing DbSystem.
type CreateInstanceDetails struct {

	// The Availability Domain where the MySQLaaS instance is located.
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The Oracle Cloud ID (OCID) of the DbSystem the MySQLaaS instance belongs in.
	DbSystemId *string `mandatory:"false" json:"dbSystemId"`

	// Role played by this Instance in the DbSystem.
	Role InstanceRoleEnum `mandatory:"false" json:"role,omitempty"`

	// DEPRECATED -- The Oracle Cloud ID (OCID) of the compartment the MySQLaaS instance belongs in.
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

	// The name of the Fault Domain the MySQLaaS instance is located in.
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
	// ListShapes operation.
	ShapeName *string `mandatory:"false" json:"shapeName"`

	// DEPRECATED -- this field will be only available on DbSystem once implementation catches up.
	// OCID of the MySQL Configuration associated with this instance.
	ConfigurationId *string `mandatory:"false" json:"configurationId"`

	// DEPRECATED -- this field will be only available on DbSystem once implementation catches up.
	// Initial size (GBs) of the data volume created and attached.
	// This only specifies the size of the database data volume, the log volume
	// for the database is scaled according to the specified shape.
	// (TODO: Include more information about db and log BVs and BVG for MySQLaaS.)
	DataStorageSizeInGBs *int `mandatory:"false" json:"dataStorageSizeInGBs"`

	// The port MySQLaaS instance is configured to listen on.
	Port *int `mandatory:"false" json:"port"`

	// The network port on which X Plugin listens for TCP/IP connections. This is the X Plugin equivalent of port.
	MysqlxPort *int `mandatory:"false" json:"mysqlxPort"`

	// The host name for the MySQLaaS Instance. The host name must begin with an alphabetic character and
	// can contain a maximum of 30 alphanumeric characters, including hyphens (-).
	// The maximum length of the combined hostname and domain is 63 characters.
	// **Note:** The hostname must be unique within the subnet.
	//  It is not possible to provision a MySQLaaS instance if the hostname is not unique.
	Hostname *string `mandatory:"false" json:"hostname"`

	AnalyticsCluster *CreateAnalyticsClusterDetails `mandatory:"false" json:"analyticsCluster"`

	// [PRIVATE API] OCID of the host image to use for this Instance. If not set, the control plane default will be used.
	HostImageId *string `mandatory:"false" json:"hostImageId"`

	// Simple key-value pair applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The Backup Id from which the instance should be restored.
	BackupId *string `mandatory:"false" json:"backupId"`
}

func (m CreateInstanceDetails) String() string {
	return common.PointerString(m)
}
