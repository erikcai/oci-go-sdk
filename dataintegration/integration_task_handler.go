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

// IntegrationTaskHandler auto generated description
type IntegrationTaskHandler struct {

	// auto generated description
	ModelType IntegrationTaskHandlerModelTypeEnum `mandatory:"true" json:"modelType"`

	// Object key
	Key *string `mandatory:"false" json:"key"`

	// modelVersion
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Status of object, can set this to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`
}

func (m IntegrationTaskHandler) String() string {
	return common.PointerString(m)
}

// IntegrationTaskHandlerModelTypeEnum Enum with underlying type: string
type IntegrationTaskHandlerModelTypeEnum string

// Set of constants representing the allowable values for IntegrationTaskHandlerModelTypeEnum
const (
	IntegrationTaskHandlerModelTypeIntegrationTaskHandler IntegrationTaskHandlerModelTypeEnum = "INTEGRATION_TASK_HANDLER"
	IntegrationTaskHandlerModelTypeDataLoaderTaskHandler  IntegrationTaskHandlerModelTypeEnum = "DATA_LOADER_TASK_HANDLER"
)

var mappingIntegrationTaskHandlerModelType = map[string]IntegrationTaskHandlerModelTypeEnum{
	"INTEGRATION_TASK_HANDLER": IntegrationTaskHandlerModelTypeIntegrationTaskHandler,
	"DATA_LOADER_TASK_HANDLER": IntegrationTaskHandlerModelTypeDataLoaderTaskHandler,
}

// GetIntegrationTaskHandlerModelTypeEnumValues Enumerates the set of values for IntegrationTaskHandlerModelTypeEnum
func GetIntegrationTaskHandlerModelTypeEnumValues() []IntegrationTaskHandlerModelTypeEnum {
	values := make([]IntegrationTaskHandlerModelTypeEnum, 0)
	for _, v := range mappingIntegrationTaskHandlerModelType {
		values = append(values, v)
	}
	return values
}
