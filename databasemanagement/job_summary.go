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

// JobSummary Summary of the job.
type JobSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the job.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment of the job.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Display name of the job.
	Name *string `mandatory:"true" json:"name"`

	// Schedule type for the job.
	ScheduleType JobScheduleTypeEnum `mandatory:"true" json:"scheduleType"`

	// Type of the job.
	JobType JobTypesEnum `mandatory:"true" json:"jobType"`

	// Lifecycle state of the job.
	LifecycleState JobLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Time when job was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Time when the job was last updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Description of the job.
	Description *string `mandatory:"false" json:"description"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database Group.
	ManagedDatabaseGroupId *string `mandatory:"false" json:"managedDatabaseGroupId"`

	// Managed database OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) where job has to be executed.
	ManagedDatabaseId *string `mandatory:"false" json:"managedDatabaseId"`

	// Subtype of databases (CDB/PDB/NON_CDB) where job needs to be executed. Applicable only when managedDatabaseGroupId is provided.
	DatabaseSubType DatabaseSubTypeEnum `mandatory:"false" json:"databaseSubType,omitempty"`

	// Job timeout duration expressed like "1h 10m 15s".
	Timeout *string `mandatory:"false" json:"timeout"`

	// Error message if job submission failed else null.
	SubmissionErrorMessage *string `mandatory:"false" json:"submissionErrorMessage"`
}

func (m JobSummary) String() string {
	return common.PointerString(m)
}
