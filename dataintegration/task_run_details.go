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

// TaskRunDetails The task run object provides information on the execution of a task.
type TaskRunDetails struct {

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

	TaskRef *Task `mandatory:"false" json:"taskRef"`

	// taskType
	TaskType *string `mandatory:"false" json:"taskType"`

	// status
	Status TaskRunDetailsStatusEnum `mandatory:"false" json:"status,omitempty"`

	// startTimeMillis
	StartTimeMillis *int `mandatory:"false" json:"startTimeMillis"`

	// endTimeMillis
	EndTimeMillis *int `mandatory:"false" json:"endTimeMillis"`

	// lastUpdated
	LastUpdated *int `mandatory:"false" json:"lastUpdated"`

	// recordsWritten
	RecordsWritten *int `mandatory:"false" json:"recordsWritten"`

	// logs
	Logs *string `mandatory:"false" json:"logs"`

	// Status of object, can set this to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Value can only contain upper case letters, underscore and numbers. It should begin with upper case letter or underscore. The value can be edited by the user.
	Identifier *string `mandatory:"false" json:"identifier"`
}

func (m TaskRunDetails) String() string {
	return common.PointerString(m)
}

// TaskRunDetailsStatusEnum Enum with underlying type: string
type TaskRunDetailsStatusEnum string

// Set of constants representing the allowable values for TaskRunDetailsStatusEnum
const (
	TaskRunDetailsStatusNotStarted  TaskRunDetailsStatusEnum = "NOT_STARTED"
	TaskRunDetailsStatusQueued      TaskRunDetailsStatusEnum = "QUEUED"
	TaskRunDetailsStatusRunning     TaskRunDetailsStatusEnum = "RUNNING"
	TaskRunDetailsStatusTerminating TaskRunDetailsStatusEnum = "TERMINATING"
	TaskRunDetailsStatusTerminated  TaskRunDetailsStatusEnum = "TERMINATED"
	TaskRunDetailsStatusSuccess     TaskRunDetailsStatusEnum = "SUCCESS"
	TaskRunDetailsStatusError       TaskRunDetailsStatusEnum = "ERROR"
)

var mappingTaskRunDetailsStatus = map[string]TaskRunDetailsStatusEnum{
	"NOT_STARTED": TaskRunDetailsStatusNotStarted,
	"QUEUED":      TaskRunDetailsStatusQueued,
	"RUNNING":     TaskRunDetailsStatusRunning,
	"TERMINATING": TaskRunDetailsStatusTerminating,
	"TERMINATED":  TaskRunDetailsStatusTerminated,
	"SUCCESS":     TaskRunDetailsStatusSuccess,
	"ERROR":       TaskRunDetailsStatusError,
}

// GetTaskRunDetailsStatusEnumValues Enumerates the set of values for TaskRunDetailsStatusEnum
func GetTaskRunDetailsStatusEnumValues() []TaskRunDetailsStatusEnum {
	values := make([]TaskRunDetailsStatusEnum, 0)
	for _, v := range mappingTaskRunDetailsStatus {
		values = append(values, v)
	}
	return values
}
