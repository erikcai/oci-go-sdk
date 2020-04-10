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

// FolderDetailsForCreate The details incuding name, description for the folder, which is a container of other folders, tasks and dataflows.
type FolderDetailsForCreate struct {

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value can be edited by the user and it is restricted to 1000 characters
	Name *string `mandatory:"true" json:"name"`

	// Value can only contain upper case letters, underscore and numbers. It should begin with upper case letter or underscore. The value can be edited by the user.
	Identifier *string `mandatory:"true" json:"identifier"`

	// Currently not used on folder creation. Reserved for future.
	Key *string `mandatory:"false" json:"key"`

	// modelVersion
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	// Descriptive text for the object.
	Description *string `mandatory:"false" json:"description"`

	// categoryName
	CategoryName *string `mandatory:"false" json:"categoryName"`

	// Status of object, can set this to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`
}

func (m FolderDetailsForCreate) String() string {
	return common.PointerString(m)
}
