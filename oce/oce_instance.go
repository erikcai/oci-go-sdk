// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// OceInstance API
//
// Oracle Content and Experience is a cloud-based content hub to drive omni-channel content management and accelerate experience delivery
//

package oce

import (
	"github.com/oracle/oci-go-sdk/common"
)

// OceInstance Details of OceInstance.
type OceInstance struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// Unique GUID identifier that is immutable on creation
	Guid *string `mandatory:"true" json:"guid"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// OceInstance Name
	Name *string `mandatory:"true" json:"name"`

	// Tenancy Identifier
	TenancyId *string `mandatory:"true" json:"tenancyId"`

	// IDCS Tenancy Identifier
	IdcsTenancy *string `mandatory:"true" json:"idcsTenancy"`

	// Tenancy Name
	TenancyName *string `mandatory:"true" json:"tenancyName"`

	// Object Storage Namespace of tenancy
	ObjectStorageNamespace *string `mandatory:"true" json:"objectStorageNamespace"`

	// Admin Email for Notification
	AdminEmail *string `mandatory:"true" json:"adminEmail"`

	// OceInstance description, can be updated
	Description *string `mandatory:"false" json:"description"`

	IdentityStripe *IdentityStripeDetails `mandatory:"false" json:"identityStripe"`

	// The time the the OceInstance was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the OceInstance was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the file system.
	LifecycleState OceInstanceLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// An message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	StateMessage *string `mandatory:"false" json:"stateMessage"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// SERVICE data.
	// Example: `{"service": {"IDCS": "value"}}`
	Service map[string]interface{} `mandatory:"false" json:"service"`
}

func (m OceInstance) String() string {
	return common.PointerString(m)
}

// OceInstanceLifecycleStateEnum Enum with underlying type: string
type OceInstanceLifecycleStateEnum string

// Set of constants representing the allowable values for OceInstanceLifecycleStateEnum
const (
	OceInstanceLifecycleStateCreating OceInstanceLifecycleStateEnum = "CREATING"
	OceInstanceLifecycleStateUpdating OceInstanceLifecycleStateEnum = "UPDATING"
	OceInstanceLifecycleStateActive   OceInstanceLifecycleStateEnum = "ACTIVE"
	OceInstanceLifecycleStateDeleting OceInstanceLifecycleStateEnum = "DELETING"
	OceInstanceLifecycleStateDeleted  OceInstanceLifecycleStateEnum = "DELETED"
	OceInstanceLifecycleStateFailed   OceInstanceLifecycleStateEnum = "FAILED"
)

var mappingOceInstanceLifecycleState = map[string]OceInstanceLifecycleStateEnum{
	"CREATING": OceInstanceLifecycleStateCreating,
	"UPDATING": OceInstanceLifecycleStateUpdating,
	"ACTIVE":   OceInstanceLifecycleStateActive,
	"DELETING": OceInstanceLifecycleStateDeleting,
	"DELETED":  OceInstanceLifecycleStateDeleted,
	"FAILED":   OceInstanceLifecycleStateFailed,
}

// GetOceInstanceLifecycleStateEnumValues Enumerates the set of values for OceInstanceLifecycleStateEnum
func GetOceInstanceLifecycleStateEnumValues() []OceInstanceLifecycleStateEnum {
	values := make([]OceInstanceLifecycleStateEnum, 0)
	for _, v := range mappingOceInstanceLifecycleState {
		values = append(values, v)
	}
	return values
}
