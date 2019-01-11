// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Data Security Control Plane API
//
// The API to manage data security instance creation and deletion
//

package datasecurity

import (
	"github.com/oracle/oci-go-sdk/common"
)

// DataSecurityInstance Object of DataSecurityInstance
type DataSecurityInstance struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Data security instance name, can be renamed
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Service URL of the data security instance
	Url *string `mandatory:"false" json:"url"`

	// Description of the data security instance
	Description *string `mandatory:"false" json:"description"`

	// The time the the data security instance was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The current state of the data security instance.
	LifecycleState DataSecurityInstanceLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m DataSecurityInstance) String() string {
	return common.PointerString(m)
}

// DataSecurityInstanceLifecycleStateEnum Enum with underlying type: string
type DataSecurityInstanceLifecycleStateEnum string

// Set of constants representing the allowable values for DataSecurityInstanceLifecycleStateEnum
const (
	DataSecurityInstanceLifecycleStateCreating DataSecurityInstanceLifecycleStateEnum = "CREATING"
	DataSecurityInstanceLifecycleStateUpdating DataSecurityInstanceLifecycleStateEnum = "UPDATING"
	DataSecurityInstanceLifecycleStateActive   DataSecurityInstanceLifecycleStateEnum = "ACTIVE"
	DataSecurityInstanceLifecycleStateDeleting DataSecurityInstanceLifecycleStateEnum = "DELETING"
	DataSecurityInstanceLifecycleStateDeleted  DataSecurityInstanceLifecycleStateEnum = "DELETED"
	DataSecurityInstanceLifecycleStateFailed   DataSecurityInstanceLifecycleStateEnum = "FAILED"
)

var mappingDataSecurityInstanceLifecycleState = map[string]DataSecurityInstanceLifecycleStateEnum{
	"CREATING": DataSecurityInstanceLifecycleStateCreating,
	"UPDATING": DataSecurityInstanceLifecycleStateUpdating,
	"ACTIVE":   DataSecurityInstanceLifecycleStateActive,
	"DELETING": DataSecurityInstanceLifecycleStateDeleting,
	"DELETED":  DataSecurityInstanceLifecycleStateDeleted,
	"FAILED":   DataSecurityInstanceLifecycleStateFailed,
}

// GetDataSecurityInstanceLifecycleStateEnumValues Enumerates the set of values for DataSecurityInstanceLifecycleStateEnum
func GetDataSecurityInstanceLifecycleStateEnumValues() []DataSecurityInstanceLifecycleStateEnum {
	values := make([]DataSecurityInstanceLifecycleStateEnum, 0)
	for _, v := range mappingDataSecurityInstanceLifecycleState {
		values = append(values, v)
	}
	return values
}
