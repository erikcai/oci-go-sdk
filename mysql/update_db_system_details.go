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

// UpdateDbSystemDetails Details required to create a DbSystem.
type UpdateDbSystemDetails struct {

	// The user-friendly name for the DbSystem. It does not have to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// User-provided data about the DbSystem.
	Description *string `mandatory:"false" json:"description"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The OCID of the subnet the MySQLaaS DbSystem is associated with.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	AvailabilityPolicy *DbSystemAvailabilityPolicy `mandatory:"false" json:"availabilityPolicy"`

	// The shape of the MySQLaaS instance. The shape determines resources
	// allocated to the MySQLaaS instance - CPU cores and memory for VM
	// shapes; CPU cores, memory and storage for non-VM (or bare metal)
	// shapes. To get a list of shapes, use (FIXME: correct link for
	// MySQLaaS shapes) the
	// ListShapes
	// operation.
	// Changes in Shape will result in a downtime as the MySQLaaS instance
	// is migrated to the new Compute instance.
	ShapeName *string `mandatory:"false" json:"shapeName"`

	// The specific MySQL version identifier.
	MysqlVersion *string `mandatory:"false" json:"mysqlVersion"`

	// The OCID of the Configuration to be used for Instances in this DbSystem.
	ConfigurationId *string `mandatory:"false" json:"configurationId"`

	// The username for the administrative user for the MySQLaaS Instance.
	AdminUsername *string `mandatory:"false" json:"adminUsername"`

	// The password for the administrative user. The password must be
	// between 8 and 32 characters long, and must contain at least 1
	// numeric character, 1 lowercase character, 1 uppercase character, and
	// 1 special (nonalphanumeric) character.
	AdminPassword *string `mandatory:"false" json:"adminPassword"`

	// New size of the data volume in GBs that will be created and attached.
	// Increases in data storage size will happen asynchronously and on-line.
	// Decreases in data storage size are not supported.
	DataStorageSizeInGBs *int `mandatory:"false" json:"dataStorageSizeInGBs"`

	BackupPolicy *CreateUpdateBackupPolicy `mandatory:"false" json:"backupPolicy"`

	// Simple key-value pair applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateDbSystemDetails) String() string {
	return common.PointerString(m)
}
