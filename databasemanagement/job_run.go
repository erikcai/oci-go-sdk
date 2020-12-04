// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management Service APIs.
//
// This file contains the customer facing APIs for Database Management service.
//

package databasemanagement

import (
	"github.com/oracle/oci-go-sdk/v30/common"
)

// JobRun Details of a job run.
type JobRun struct {

	// Identifier of the job run.
	Id *string `mandatory:"true" json:"id"`

	// Generated name of the job run.
	Name *string `mandatory:"true" json:"name"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment of the parent job.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the parent job.
	JobId *string `mandatory:"true" json:"jobId"`

	// Name of the parent job.
	JobName *string `mandatory:"true" json:"jobName"`

	// Status of the job run.
	RunStatus JobRunRunStatusEnum `mandatory:"true" json:"runStatus"`

	// Time when job run was submitted.
	TimeSubmitted *common.SDKTime `mandatory:"true" json:"timeSubmitted"`

	// Time when job run was updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database Group for the job run.
	ManagedDatabaseGroupId *string `mandatory:"false" json:"managedDatabaseGroupId"`

	// Managed database OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) for the job run.
	ManagedDatabaseId *string `mandatory:"false" json:"managedDatabaseId"`
}

func (m JobRun) String() string {
	return common.PointerString(m)
}

// JobRunRunStatusEnum Enum with underlying type: string
type JobRunRunStatusEnum string

// Set of constants representing the allowable values for JobRunRunStatusEnum
const (
	JobRunRunStatusCompleted  JobRunRunStatusEnum = "COMPLETED"
	JobRunRunStatusFailed     JobRunRunStatusEnum = "FAILED"
	JobRunRunStatusInProgress JobRunRunStatusEnum = "IN_PROGRESS"
)

var mappingJobRunRunStatus = map[string]JobRunRunStatusEnum{
	"COMPLETED":   JobRunRunStatusCompleted,
	"FAILED":      JobRunRunStatusFailed,
	"IN_PROGRESS": JobRunRunStatusInProgress,
}

// GetJobRunRunStatusEnumValues Enumerates the set of values for JobRunRunStatusEnum
func GetJobRunRunStatusEnumValues() []JobRunRunStatusEnum {
	values := make([]JobRunRunStatusEnum, 0)
	for _, v := range mappingJobRunRunStatus {
		values = append(values, v)
	}
	return values
}
