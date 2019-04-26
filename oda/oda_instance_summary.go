// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Digital Assistant Control Plane API
//
// API to create and maintain Digital Assistant (ODA) service instances.
//

package oda

import (
	"github.com/oracle/oci-go-sdk/common"
)

// OdaInstanceSummary Summary of the ODA instance
type OdaInstanceSummary struct {

	// Unique identifier of the ODA instance
	Id *string `mandatory:"true" json:"id"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The current state of the ODA instance
	LifecycleState OdaInstanceSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Name of the ODA instance (can be modified)
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description of the ODA instance
	Description *string `mandatory:"false" json:"description"`

	// Type of the instance, corresponds to the SKU and metering method
	InstanceType OdaInstanceSummaryInstanceTypeEnum `mandatory:"false" json:"instanceType,omitempty"`

	// Shape or size of the instance
	ShapeName OdaInstanceSummaryShapeNameEnum `mandatory:"false" json:"shapeName,omitempty"`

	// The time the the ODA instance was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the ODA instance was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// An message describing the current state in more detail. For example, can be used to provide actionable
	// information for a resource in Failed state.
	StateMessage *string `mandatory:"false" json:"stateMessage"`
}

func (m OdaInstanceSummary) String() string {
	return common.PointerString(m)
}

// OdaInstanceSummaryInstanceTypeEnum Enum with underlying type: string
type OdaInstanceSummaryInstanceTypeEnum string

// Set of constants representing the allowable values for OdaInstanceSummaryInstanceTypeEnum
const (
	OdaInstanceSummaryInstanceTypeUcm              OdaInstanceSummaryInstanceTypeEnum = "UCM"
	OdaInstanceSummaryInstanceTypeSaasUserSessions OdaInstanceSummaryInstanceTypeEnum = "SAAS_USER_SESSIONS"
	OdaInstanceSummaryInstanceTypeSaasEmployees    OdaInstanceSummaryInstanceTypeEnum = "SAAS_EMPLOYEES"
)

var mappingOdaInstanceSummaryInstanceType = map[string]OdaInstanceSummaryInstanceTypeEnum{
	"UCM":                OdaInstanceSummaryInstanceTypeUcm,
	"SAAS_USER_SESSIONS": OdaInstanceSummaryInstanceTypeSaasUserSessions,
	"SAAS_EMPLOYEES":     OdaInstanceSummaryInstanceTypeSaasEmployees,
}

// GetOdaInstanceSummaryInstanceTypeEnumValues Enumerates the set of values for OdaInstanceSummaryInstanceTypeEnum
func GetOdaInstanceSummaryInstanceTypeEnumValues() []OdaInstanceSummaryInstanceTypeEnum {
	values := make([]OdaInstanceSummaryInstanceTypeEnum, 0)
	for _, v := range mappingOdaInstanceSummaryInstanceType {
		values = append(values, v)
	}
	return values
}

// OdaInstanceSummaryShapeNameEnum Enum with underlying type: string
type OdaInstanceSummaryShapeNameEnum string

// Set of constants representing the allowable values for OdaInstanceSummaryShapeNameEnum
const (
	OdaInstanceSummaryShapeNameDevelopment OdaInstanceSummaryShapeNameEnum = "DEVELOPMENT"
	OdaInstanceSummaryShapeNameTrial       OdaInstanceSummaryShapeNameEnum = "TRIAL"
	OdaInstanceSummaryShapeNameTest        OdaInstanceSummaryShapeNameEnum = "TEST"
	OdaInstanceSummaryShapeNameProduction  OdaInstanceSummaryShapeNameEnum = "PRODUCTION"
)

var mappingOdaInstanceSummaryShapeName = map[string]OdaInstanceSummaryShapeNameEnum{
	"DEVELOPMENT": OdaInstanceSummaryShapeNameDevelopment,
	"TRIAL":       OdaInstanceSummaryShapeNameTrial,
	"TEST":        OdaInstanceSummaryShapeNameTest,
	"PRODUCTION":  OdaInstanceSummaryShapeNameProduction,
}

// GetOdaInstanceSummaryShapeNameEnumValues Enumerates the set of values for OdaInstanceSummaryShapeNameEnum
func GetOdaInstanceSummaryShapeNameEnumValues() []OdaInstanceSummaryShapeNameEnum {
	values := make([]OdaInstanceSummaryShapeNameEnum, 0)
	for _, v := range mappingOdaInstanceSummaryShapeName {
		values = append(values, v)
	}
	return values
}

// OdaInstanceSummaryLifecycleStateEnum Enum with underlying type: string
type OdaInstanceSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for OdaInstanceSummaryLifecycleStateEnum
const (
	OdaInstanceSummaryLifecycleStateCreating OdaInstanceSummaryLifecycleStateEnum = "CREATING"
	OdaInstanceSummaryLifecycleStateUpdating OdaInstanceSummaryLifecycleStateEnum = "UPDATING"
	OdaInstanceSummaryLifecycleStateActive   OdaInstanceSummaryLifecycleStateEnum = "ACTIVE"
	OdaInstanceSummaryLifecycleStateInactive OdaInstanceSummaryLifecycleStateEnum = "INACTIVE"
	OdaInstanceSummaryLifecycleStateDeleting OdaInstanceSummaryLifecycleStateEnum = "DELETING"
	OdaInstanceSummaryLifecycleStateDeleted  OdaInstanceSummaryLifecycleStateEnum = "DELETED"
	OdaInstanceSummaryLifecycleStateFailed   OdaInstanceSummaryLifecycleStateEnum = "FAILED"
)

var mappingOdaInstanceSummaryLifecycleState = map[string]OdaInstanceSummaryLifecycleStateEnum{
	"CREATING": OdaInstanceSummaryLifecycleStateCreating,
	"UPDATING": OdaInstanceSummaryLifecycleStateUpdating,
	"ACTIVE":   OdaInstanceSummaryLifecycleStateActive,
	"INACTIVE": OdaInstanceSummaryLifecycleStateInactive,
	"DELETING": OdaInstanceSummaryLifecycleStateDeleting,
	"DELETED":  OdaInstanceSummaryLifecycleStateDeleted,
	"FAILED":   OdaInstanceSummaryLifecycleStateFailed,
}

// GetOdaInstanceSummaryLifecycleStateEnumValues Enumerates the set of values for OdaInstanceSummaryLifecycleStateEnum
func GetOdaInstanceSummaryLifecycleStateEnumValues() []OdaInstanceSummaryLifecycleStateEnum {
	values := make([]OdaInstanceSummaryLifecycleStateEnum, 0)
	for _, v := range mappingOdaInstanceSummaryLifecycleState {
		values = append(values, v)
	}
	return values
}
