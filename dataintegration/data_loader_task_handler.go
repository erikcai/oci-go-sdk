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

// DataLoaderTaskHandler auto generated description
type DataLoaderTaskHandler struct {

	// auto generated description
	ModelType DataLoaderTaskHandlerModelTypeEnum `mandatory:"true" json:"modelType"`

	// Object key
	Key *string `mandatory:"false" json:"key"`

	// modelVersion
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Status of object, can set this to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`
}

func (m DataLoaderTaskHandler) String() string {
	return common.PointerString(m)
}

// DataLoaderTaskHandlerModelTypeEnum Enum with underlying type: string
type DataLoaderTaskHandlerModelTypeEnum string

// Set of constants representing the allowable values for DataLoaderTaskHandlerModelTypeEnum
const (
	DataLoaderTaskHandlerModelTypeIntegrationTaskHandler DataLoaderTaskHandlerModelTypeEnum = "INTEGRATION_TASK_HANDLER"
	DataLoaderTaskHandlerModelTypeDataLoaderTaskHandler  DataLoaderTaskHandlerModelTypeEnum = "DATA_LOADER_TASK_HANDLER"
)

var mappingDataLoaderTaskHandlerModelType = map[string]DataLoaderTaskHandlerModelTypeEnum{
	"INTEGRATION_TASK_HANDLER": DataLoaderTaskHandlerModelTypeIntegrationTaskHandler,
	"DATA_LOADER_TASK_HANDLER": DataLoaderTaskHandlerModelTypeDataLoaderTaskHandler,
}

// GetDataLoaderTaskHandlerModelTypeEnumValues Enumerates the set of values for DataLoaderTaskHandlerModelTypeEnum
func GetDataLoaderTaskHandlerModelTypeEnumValues() []DataLoaderTaskHandlerModelTypeEnum {
	values := make([]DataLoaderTaskHandlerModelTypeEnum, 0)
	for _, v := range mappingDataLoaderTaskHandlerModelType {
		values = append(values, v)
	}
	return values
}
