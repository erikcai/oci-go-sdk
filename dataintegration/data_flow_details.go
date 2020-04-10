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

// DataFlowDetails auto generated description
type DataFlowDetails struct {

	// Generated key that can be used in API calls to identify data flow. On scenarios where reference to the data flow is needed, a value can be passed in create.
	Key *string `mandatory:"false" json:"key"`

	// The type of the object.
	ModelType *string `mandatory:"false" json:"modelType"`

	// modelVersion
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value can be edited by the user and it is restricted to 1000 characters
	Name *string `mandatory:"false" json:"name"`

	// Value can only contain upper case letters, underscore and numbers. It should begin with upper case letter or underscore. The value can be edited by the user.
	Identifier *string `mandatory:"false" json:"identifier"`

	// The version of the object, to track changes in the object instance
	ObjectVersion *int `mandatory:"false" json:"objectVersion"`

	// nodes
	Nodes []FlowNode `mandatory:"false" json:"nodes"`

	// parameters
	Parameters []Parameter `mandatory:"false" json:"parameters"`

	// Descriptive text for the object.
	Description *string `mandatory:"false" json:"description"`

	FlowConfigValues *ConfigValues `mandatory:"false" json:"flowConfigValues"`

	// Status of object, can set this to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// External reference objects that are used within the data flow.
	ExternalRefObjects []interface{} `mandatory:"false" json:"externalRefObjects"`

	// wrappedExternalRefObjects
	WrappedExternalRefObjects []WrappedFco `mandatory:"false" json:"wrappedExternalRefObjects"`
}

func (m DataFlowDetails) String() string {
	return common.PointerString(m)
}
