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

// JobExecutionSummary Summary of a job execution.
type JobExecutionSummary struct {

	// Identifier of the job execution.
	Id *string `mandatory:"true" json:"id"`

	// Generated name of the job execution.
	Name *string `mandatory:"true" json:"name"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment of the parent job.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of Managed Database associated with the job execution.
	ManagedDatabaseId *string `mandatory:"true" json:"managedDatabaseId"`

	// Name of the Managed Database associated with the job execution.
	ManagedDatabaseName *string `mandatory:"true" json:"managedDatabaseName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the parent job.
	JobId *string `mandatory:"true" json:"jobId"`

	// Name of the parent job.
	JobName *string `mandatory:"true" json:"jobName"`

	// Status of the job execution.
	Status JobExecutionStatusEnum `mandatory:"true" json:"status"`

	// Time when job execution was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database Group associated with the job execution.
	ManagedDatabaseGroupId *string `mandatory:"false" json:"managedDatabaseGroupId"`

	// Time when job execution was completed.
	TimeCompleted *common.SDKTime `mandatory:"false" json:"timeCompleted"`
}

func (m JobExecutionSummary) String() string {
	return common.PointerString(m)
}
