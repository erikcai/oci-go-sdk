// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Digital Assistant Control Plane API
//
// API to create and maintain Digital Assistant (ODA) service instances.
//

package oda

import (
	"github.com/oracle/oci-go-sdk/common"
)

// OdaInstance Description of OdaServiceInstance.
type OdaInstance struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Shape or size of the instance
	ShapeName OdaInstanceShapeNameEnum `mandatory:"true" json:"shapeName"`

	// OdaServiceInstance Identifier, can be renamed
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description of the ODA instance
	Description *string `mandatory:"false" json:"description"`

	// URL for the ODA web application associated with the instance
	WebAppUrl *string `mandatory:"false" json:"webAppUrl"`

	// URL for the Connectors endpoint
	ConnectorUrl *string `mandatory:"false" json:"connectorUrl"`

	// The time the the OdaServiceInstance was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the OdaServiceInstance was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the ODA instance
	LifecycleState OdaInstanceLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// An message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	StateMessage *string `mandatory:"false" json:"stateMessage"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m OdaInstance) String() string {
	return common.PointerString(m)
}

// OdaInstanceShapeNameEnum Enum with underlying type: string
type OdaInstanceShapeNameEnum string

// Set of constants representing the allowable values for OdaInstanceShapeNameEnum
const (
	OdaInstanceShapeNameDevelopment OdaInstanceShapeNameEnum = "DEVELOPMENT"
	OdaInstanceShapeNameProduction  OdaInstanceShapeNameEnum = "PRODUCTION"
)

var mappingOdaInstanceShapeName = map[string]OdaInstanceShapeNameEnum{
	"DEVELOPMENT": OdaInstanceShapeNameDevelopment,
	"PRODUCTION":  OdaInstanceShapeNameProduction,
}

// GetOdaInstanceShapeNameEnumValues Enumerates the set of values for OdaInstanceShapeNameEnum
func GetOdaInstanceShapeNameEnumValues() []OdaInstanceShapeNameEnum {
	values := make([]OdaInstanceShapeNameEnum, 0)
	for _, v := range mappingOdaInstanceShapeName {
		values = append(values, v)
	}
	return values
}

// OdaInstanceLifecycleStateEnum Enum with underlying type: string
type OdaInstanceLifecycleStateEnum string

// Set of constants representing the allowable values for OdaInstanceLifecycleStateEnum
const (
	OdaInstanceLifecycleStateCreating OdaInstanceLifecycleStateEnum = "CREATING"
	OdaInstanceLifecycleStateUpdating OdaInstanceLifecycleStateEnum = "UPDATING"
	OdaInstanceLifecycleStateActive   OdaInstanceLifecycleStateEnum = "ACTIVE"
	OdaInstanceLifecycleStateInactive OdaInstanceLifecycleStateEnum = "INACTIVE"
	OdaInstanceLifecycleStateDeleting OdaInstanceLifecycleStateEnum = "DELETING"
	OdaInstanceLifecycleStateDeleted  OdaInstanceLifecycleStateEnum = "DELETED"
	OdaInstanceLifecycleStateFailed   OdaInstanceLifecycleStateEnum = "FAILED"
)

var mappingOdaInstanceLifecycleState = map[string]OdaInstanceLifecycleStateEnum{
	"CREATING": OdaInstanceLifecycleStateCreating,
	"UPDATING": OdaInstanceLifecycleStateUpdating,
	"ACTIVE":   OdaInstanceLifecycleStateActive,
	"INACTIVE": OdaInstanceLifecycleStateInactive,
	"DELETING": OdaInstanceLifecycleStateDeleting,
	"DELETED":  OdaInstanceLifecycleStateDeleted,
	"FAILED":   OdaInstanceLifecycleStateFailed,
}

// GetOdaInstanceLifecycleStateEnumValues Enumerates the set of values for OdaInstanceLifecycleStateEnum
func GetOdaInstanceLifecycleStateEnumValues() []OdaInstanceLifecycleStateEnum {
	values := make([]OdaInstanceLifecycleStateEnum, 0)
	for _, v := range mappingOdaInstanceLifecycleState {
		values = append(values, v)
	}
	return values
}
