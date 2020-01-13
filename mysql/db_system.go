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

// DbSystem A DbSystem is a logical set of data and collection of Instances to provide various levels of MySQL as a Service.
type DbSystem struct {

	// The OCID of the DbSystem.
	Id *string `mandatory:"true" json:"id"`

	// The user-friendly name for the DbSystem. It does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The Oracle Cloud ID (OCID) of the compartment the DbSystem belongs in.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the subnet the MySQLaaS DbSystem is associated with.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	AvailabilityPolicy *DbSystemAvailabilityPolicy `mandatory:"true" json:"availabilityPolicy"`

	// Name of the MySQL Version in use for the DbSystem.
	MysqlVersion *string `mandatory:"true" json:"mysqlVersion"`

	// Initial size of the data volume in GiBs that will be created and attached.
	DataStorageSizeInGBs *int `mandatory:"true" json:"dataStorageSizeInGBs"`

	// The current state of the DbSystem.
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the DbSystem was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the DbSystem was last updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// User-provided data about the DbSystem.
	Description *string `mandatory:"false" json:"description"`

	// The shape of the primary instances of the DbSystem. The shape
	// determines resources allocated to a MySQLaaS instance - CPU cores
	// and memory for VM shapes; CPU cores, memory and storage for non-VM
	// (or bare metal) shapes. To get a list of shapes, use (FIXME: correct
	// link for MySQLaaS shapes) the
	// ListShapes operation.
	ShapeName *string `mandatory:"false" json:"shapeName"`

	Instances []DbSystemInstance `mandatory:"false" json:"instances"`

	BackupPolicy *BackupPolicy `mandatory:"false" json:"backupPolicy"`

	// The OCID of the Configuration to be used for Instances in this DbSystem.
	ConfigurationId *string `mandatory:"false" json:"configurationId"`

	InBoundChannels []InBoundChannel `mandatory:"false" json:"inBoundChannels"`

	Endpoints []DbSystemEndpoint `mandatory:"false" json:"endpoints"`

	// Additional information about the current lifecycleState.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Simple key-value pair applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m DbSystem) String() string {
	return common.PointerString(m)
}
