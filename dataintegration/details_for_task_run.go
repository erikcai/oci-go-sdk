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

// DetailsForTaskRun Detail of task run object.
type DetailsForTaskRun struct {

	// Objects will use a 36 character key as unique ID. It is system generated and cannot be edited by user
	Key *string `mandatory:"false" json:"key"`

	// Version for object
	ObjectVersion *string `mandatory:"false" json:"objectVersion"`

	// Status of object, can set this to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Value can only contain upper case letters, underscore and numbers. It should begin with upper case letter or underscore. The value can be edited by the user.
	Identifier *string `mandatory:"false" json:"identifier"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value can be edited by the user and it is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// Descriptive text for the object, can be up to 2000 characters.
	Description *string `mandatory:"false" json:"description"`

	ConfigProvider *ConfigProvider `mandatory:"false" json:"configProvider"`

	// Key of object to execute; this may be a task, a dataflow or a published object in an application
	ObjectToExecute *string `mandatory:"false" json:"objectToExecute"`
}

func (m DetailsForTaskRun) String() string {
	return common.PointerString(m)
}
