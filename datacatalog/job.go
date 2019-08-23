// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// DataCatalog API
//
// A description of the DataCatalog API
//

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Job Details of a Job. Jobs are scheduled instances of a JobDefinition.
type Job struct {

	// Unique key of the Job resource
	Key *string `mandatory:"true" json:"key"`

	// The display name of a user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Detailed description of the Job.
	Description *string `mandatory:"false" json:"description"`

	// The Catalog's Oracle ID (OCID).
	CatalogId *string `mandatory:"false" json:"catalogId"`

	// Lifecycle State for Job.
	LifecycleState JobLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The date and time the Job was created, in the format defined by RFC3339.
	// Example: `2019-03-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Time that this Job was last updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Type of the Job.
	JobType JobTypeEnum `mandatory:"false" json:"jobType,omitempty"`

	// Schedule specified in the cron expression format that has seven fields for second , minute , hour , day-of-month , month , day-of-week , year .
	// It can also include special characters like * for all and ? for any . There are also pre-defined schedules that can be specified using
	// special strings. For example, @hourly will run the job every hour.
	ScheduleCronExpression *int `mandatory:"false" json:"scheduleCronExpression"`

	// Date that the schedule should be operational. An RFC3339 formatted datetime string.
	TimeScheduleBegin *common.SDKTime `mandatory:"false" json:"timeScheduleBegin"`

	// Date that the schedule should end from being operational. An RFC3339 formatted datetime string.
	TimeScheduleEnd *common.SDKTime `mandatory:"false" json:"timeScheduleEnd"`

	// Type of Job Schedule which is inferred from the scheduling properties.
	ScheduleType JobScheduleTypeEnum `mandatory:"false" json:"scheduleType,omitempty"`

	// The key of the connection used by the Job. This connection will override the default connection specified in
	// the associated Job Definition. All executions will use this connection.
	ConnectionKey *string `mandatory:"false" json:"connectionKey"`

	// The unique key of the job definition resource which defined the scope of this job.
	JobDefinitionKey *string `mandatory:"false" json:"jobDefinitionKey"`

	// Internal Version of the Job resource
	InternalVersion *string `mandatory:"false" json:"internalVersion"`

	// The total number of executions for this job schedule.
	ExecutionCount *int `mandatory:"false" json:"executionCount"`

	// The date and time of the most recent execution for this Job, in the format defined by RFC3339.
	// Example: `2019-03-25T21:10:29.600Z`
	TimeOfLatestExecution *common.SDKTime `mandatory:"false" json:"timeOfLatestExecution"`

	// Id (OCID) of the user who created this job.
	CreatedById *string `mandatory:"false" json:"createdById"`

	// Id (OCID) of the user who updated this job.
	UpdatedById *string `mandatory:"false" json:"updatedById"`

	// URI to the Job instance in the API.
	Uri *string `mandatory:"false" json:"uri"`
}

func (m Job) String() string {
	return common.PointerString(m)
}
