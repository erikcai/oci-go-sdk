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

// DataFlowValidationDetails The information about DataFlow object validation
type DataFlowValidationDetails struct {

	// Total number of validation messages
	TotalMessageCount *int `mandatory:"false" json:"totalMessageCount"`

	// Total number of validation error messages
	ErrorMessageCount *int `mandatory:"false" json:"errorMessageCount"`

	// Total number of validation warning messages
	WarnMessageCount *int `mandatory:"false" json:"warnMessageCount"`

	// Total number of validation information messages
	InfoMessageCount *int `mandatory:"false" json:"infoMessageCount"`

	// Detailed information of the DataFlow object validation.
	ValidationMessages map[string]ValidationMessage `mandatory:"false" json:"validationMessages"`

	// Objects will use a 36 character key as unique ID. It is system generated and cannot be edited by user
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

	// Status of object, can set this to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Value can only contain upper case letters, underscore and numbers. It should begin with upper case letter or underscore. The value can be edited by the user.
	Identifier *string `mandatory:"false" json:"identifier"`
}

func (m DataFlowValidationDetails) String() string {
	return common.PointerString(m)
}
