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

// JobSummary Details of a Job. Jobs are scheduled instances of a JobDefinition.
type JobSummary struct {

	// Unique key of the Job.
	Key *string `mandatory:"true" json:"key"`

	// URI to the Job instance in the API.
	Uri *string `mandatory:"false" json:"uri"`

	// The display name of a user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The Catalog's Oracle ID (OCID).
	CatalogId *string `mandatory:"false" json:"catalogId"`

	// The unique key of the job definition resource which defined the scope of this job.
	JobDefinitionKey *string `mandatory:"false" json:"jobDefinitionKey"`

	// Lifecycle state of the job. For eg: Running, Paused, Completed etc.
	LifecycleState JobLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Type of the Job.
	JobType JobTypeEnum `mandatory:"false" json:"jobType,omitempty"`

	// Type of Job Schedule which is inferred from the scheduling properties.
	ScheduleType *string `mandatory:"false" json:"scheduleType"`

	// Detailed description of the Job.
	Description *string `mandatory:"false" json:"description"`

	// The date and time the Job was created, in the format defined by RFC3339.
	// Example: `2019-03-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m JobSummary) String() string {
	return common.PointerString(m)
}
