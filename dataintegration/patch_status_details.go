// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration Service API Specification
//
// Data Integration Service API Specification
//

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/common"
)

// PatchStatusDetails auto generated description
type PatchStatusDetails struct {

	// Object key
	Key *string `mandatory:"false" json:"key"`

	// The type of the object.
	ModelType *string `mandatory:"false" json:"modelType"`

	// modelVersion
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value can be edited by the user and it is restricted to 1000 characters
	Name *string `mandatory:"false" json:"name"`

	// Descriptive text for the object.
	Description *string `mandatory:"false" json:"description"`

	// The version of the object, to track changes in the object instance
	ObjectVersion *int `mandatory:"false" json:"objectVersion"`

	// version
	Version *int `mandatory:"false" json:"version"`

	// Status of object, can set this to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Value can only contain upper case letters, underscore and numbers. It should begin with upper case letter or underscore. The value can be edited by the user.
	Identifier *string `mandatory:"false" json:"identifier"`

	// timePatched
	TimePatched *common.SDKTime `mandatory:"false" json:"timePatched"`

	// Error encountered while applying the patch, if any
	ErrorMessage *string `mandatory:"false" json:"errorMessage"`

	// application version of the patch
	ApplicationVersion *string `mandatory:"false" json:"applicationVersion"`

	// patchType
	PatchType PatchStatusDetailsPatchTypeEnum `mandatory:"false" json:"patchType,omitempty"`

	// Status of the patch applied or being applied on the application
	PatchStatus PatchStatusDetailsPatchStatusEnum `mandatory:"false" json:"patchStatus,omitempty"`

	// List of dependent objects in this patch.
	DependentObjects []MetadataObjectSummary `mandatory:"false" json:"dependentObjects"`

	// List of objects that are published / unpublished in this patch.
	Objects []MetadataObjectSummary `mandatory:"false" json:"objects"`
}

func (m PatchStatusDetails) String() string {
	return common.PointerString(m)
}

// PatchStatusDetailsPatchTypeEnum Enum with underlying type: string
type PatchStatusDetailsPatchTypeEnum string

// Set of constants representing the allowable values for PatchStatusDetailsPatchTypeEnum
const (
	PatchStatusDetailsPatchTypePublish   PatchStatusDetailsPatchTypeEnum = "PUBLISH"
	PatchStatusDetailsPatchTypeUnpublish PatchStatusDetailsPatchTypeEnum = "UNPUBLISH"
)

var mappingPatchStatusDetailsPatchType = map[string]PatchStatusDetailsPatchTypeEnum{
	"PUBLISH":   PatchStatusDetailsPatchTypePublish,
	"UNPUBLISH": PatchStatusDetailsPatchTypeUnpublish,
}

// GetPatchStatusDetailsPatchTypeEnumValues Enumerates the set of values for PatchStatusDetailsPatchTypeEnum
func GetPatchStatusDetailsPatchTypeEnumValues() []PatchStatusDetailsPatchTypeEnum {
	values := make([]PatchStatusDetailsPatchTypeEnum, 0)
	for _, v := range mappingPatchStatusDetailsPatchType {
		values = append(values, v)
	}
	return values
}

// PatchStatusDetailsPatchStatusEnum Enum with underlying type: string
type PatchStatusDetailsPatchStatusEnum string

// Set of constants representing the allowable values for PatchStatusDetailsPatchStatusEnum
const (
	PatchStatusDetailsPatchStatusQueued     PatchStatusDetailsPatchStatusEnum = "QUEUED"
	PatchStatusDetailsPatchStatusSuccessful PatchStatusDetailsPatchStatusEnum = "SUCCESSFUL"
	PatchStatusDetailsPatchStatusFailed     PatchStatusDetailsPatchStatusEnum = "FAILED"
	PatchStatusDetailsPatchStatusInProgress PatchStatusDetailsPatchStatusEnum = "IN_PROGRESS"
)

var mappingPatchStatusDetailsPatchStatus = map[string]PatchStatusDetailsPatchStatusEnum{
	"QUEUED":      PatchStatusDetailsPatchStatusQueued,
	"SUCCESSFUL":  PatchStatusDetailsPatchStatusSuccessful,
	"FAILED":      PatchStatusDetailsPatchStatusFailed,
	"IN_PROGRESS": PatchStatusDetailsPatchStatusInProgress,
}

// GetPatchStatusDetailsPatchStatusEnumValues Enumerates the set of values for PatchStatusDetailsPatchStatusEnum
func GetPatchStatusDetailsPatchStatusEnumValues() []PatchStatusDetailsPatchStatusEnum {
	values := make([]PatchStatusDetailsPatchStatusEnum, 0)
	for _, v := range mappingPatchStatusDetailsPatchStatus {
		values = append(values, v)
	}
	return values
}
