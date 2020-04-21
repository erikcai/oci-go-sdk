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

// DataFlowDetailsForCreate auto generated description
type DataFlowDetailsForCreate struct {

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value can be edited by the user and it is restricted to 1000 characters
	Name *string `mandatory:"true" json:"name"`

	// Value can only contain upper case letters, underscore and numbers. It should begin with upper case letter or underscore. The value can be edited by the user.
	Identifier *string `mandatory:"true" json:"identifier"`

	// Generated key that can be used in API calls to identify data flow. On scenarios where reference to the data flow is needed, a value can be passed in create.
	Key *string `mandatory:"false" json:"key"`

	// modelVersion
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

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

func (m DataFlowDetailsForCreate) String() string {
	return common.PointerString(m)
}