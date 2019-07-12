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

// CreateOdaInstanceDetails Properties required to create a new ODA instance
type CreateOdaInstanceDetails struct {

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Shape or size of the instance
	ShapeName CreateOdaInstanceDetailsShapeNameEnum `mandatory:"true" json:"shapeName"`

	// ODA instance Identifier
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description of the ODA instance
	Description *string `mandatory:"false" json:"description"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateOdaInstanceDetails) String() string {
	return common.PointerString(m)
}

// CreateOdaInstanceDetailsShapeNameEnum Enum with underlying type: string
type CreateOdaInstanceDetailsShapeNameEnum string

// Set of constants representing the allowable values for CreateOdaInstanceDetailsShapeNameEnum
const (
	CreateOdaInstanceDetailsShapeNameDevelopment CreateOdaInstanceDetailsShapeNameEnum = "DEVELOPMENT"
	CreateOdaInstanceDetailsShapeNameProduction  CreateOdaInstanceDetailsShapeNameEnum = "PRODUCTION"
)

var mappingCreateOdaInstanceDetailsShapeName = map[string]CreateOdaInstanceDetailsShapeNameEnum{
	"DEVELOPMENT": CreateOdaInstanceDetailsShapeNameDevelopment,
	"PRODUCTION":  CreateOdaInstanceDetailsShapeNameProduction,
}

// GetCreateOdaInstanceDetailsShapeNameEnumValues Enumerates the set of values for CreateOdaInstanceDetailsShapeNameEnum
func GetCreateOdaInstanceDetailsShapeNameEnumValues() []CreateOdaInstanceDetailsShapeNameEnum {
	values := make([]CreateOdaInstanceDetailsShapeNameEnum, 0)
	for _, v := range mappingCreateOdaInstanceDetailsShapeName {
		values = append(values, v)
	}
	return values
}
