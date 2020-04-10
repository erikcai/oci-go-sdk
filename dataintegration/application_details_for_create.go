// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Data Integration Service API Specification
//
// Data Integration Service API Specification
//

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ApplicationDetailsForCreate auto generated description
type ApplicationDetailsForCreate struct {

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value can be edited by the user and it is restricted to 1000 characters
	Name *string `mandatory:"true" json:"name"`

	// Value can only contain upper case letters, underscore and numbers. It should begin with upper case letter or underscore. The value can be edited by the user.
	Identifier *string `mandatory:"true" json:"identifier"`

	// Currently not used on application creation. Reserved for future.
	Key *string `mandatory:"false" json:"key"`

	// modelVersion
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	// The type of the application.
	ModelType ApplicationDetailsForCreateModelTypeEnum `mandatory:"false" json:"modelType,omitempty"`

	// Descriptive text for the object.
	Description *string `mandatory:"false" json:"description"`

	// version
	Version *int `mandatory:"false" json:"version"`

	// Status of object, can set this to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`
}

func (m ApplicationDetailsForCreate) String() string {
	return common.PointerString(m)
}

// ApplicationDetailsForCreateModelTypeEnum Enum with underlying type: string
type ApplicationDetailsForCreateModelTypeEnum string

// Set of constants representing the allowable values for ApplicationDetailsForCreateModelTypeEnum
const (
	ApplicationDetailsForCreateModelTypeIntegrationApplication ApplicationDetailsForCreateModelTypeEnum = "INTEGRATION_APPLICATION"
)

var mappingApplicationDetailsForCreateModelType = map[string]ApplicationDetailsForCreateModelTypeEnum{
	"INTEGRATION_APPLICATION": ApplicationDetailsForCreateModelTypeIntegrationApplication,
}

// GetApplicationDetailsForCreateModelTypeEnumValues Enumerates the set of values for ApplicationDetailsForCreateModelTypeEnum
func GetApplicationDetailsForCreateModelTypeEnumValues() []ApplicationDetailsForCreateModelTypeEnum {
	values := make([]ApplicationDetailsForCreateModelTypeEnum, 0)
	for _, v := range mappingApplicationDetailsForCreateModelType {
		values = append(values, v)
	}
	return values
}
