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

// DbSystemSummary A summary of a DbSystem.
type DbSystemSummary struct {

	// The OCID of the DbSystem.
	Id *string `mandatory:"true" json:"id"`

	// The user-friendly name for the DbSystem. It does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	AvailabilityPolicy *DbSystemAvailabilityPolicy `mandatory:"true" json:"availabilityPolicy"`

	// The current state of the DbSystem.
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Name of the MySQL Version in use for the DbSystem.
	MysqlVersion *string `mandatory:"true" json:"mysqlVersion"`

	// The date and time the DbSystem was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the DbSystem was last updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// User-provided data about the DbSystem.
	Description *string `mandatory:"false" json:"description"`

	// The Oracle Cloud ID (OCID) of the compartment the DbSystem belongs in.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	Instances []DbSystemInstance `mandatory:"false" json:"instances"`

	// Simple key-value pair applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m DbSystemSummary) String() string {
	return common.PointerString(m)
}
