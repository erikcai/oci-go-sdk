// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management Service APIs.
//
// This file contains the customer facing APIs for Database Management service.
//

package databasemanagement

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v30/common"
)

// CreateJobDetails Details to Create a job.
type CreateJobDetails interface {

	// Name of the job. Valid characters are uppercase or lowercase letters,
	// numbers, and underscores. It is unmodifiable and must be unique in a
	// compartment. It must begin with an alphabetic character.
	GetName() *string

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment of the job.
	GetCompartmentId() *string

	// Schedule type for the job.
	GetScheduleType() CreateJobDetails

	// Description of the job.
	GetDescription() *string

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database Group.
	GetManagedDatabaseGroupId() *string

	// Managed database OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) where job has to be executed.
	GetManagedDatabaseId() *string

	// Subtype of databases (CDB/PDB/NON_CDB) where job needs to be executed. Applicable only when managedDatabaseGroupId is provided.
	GetDatabaseSubType() DatabaseSubTypeEnum

	// Job timeout duration expressed like "1h 10m 15s".
	GetTimeout() *string

	GetResultLocation() JobExecutionResultLocation
}

type createjobdetails struct {
	JsonData               []byte
	Name                   *string                    `mandatory:"true" json:"name"`
	CompartmentId          *string                    `mandatory:"true" json:"compartmentId"`
	ScheduleType           CreateJobDetails           `mandatory:"true" json:"scheduleType"`
	Description            *string                    `mandatory:"false" json:"description"`
	ManagedDatabaseGroupId *string                    `mandatory:"false" json:"managedDatabaseGroupId"`
	ManagedDatabaseId      *string                    `mandatory:"false" json:"managedDatabaseId"`
	DatabaseSubType        DatabaseSubTypeEnum        `mandatory:"false" json:"databaseSubType,omitempty"`
	Timeout                *string                    `mandatory:"false" json:"timeout"`
	ResultLocation         JobExecutionResultLocation `mandatory:"false" json:"resultLocation"`
	JobType                string                     `json:"jobType"`
}

// UnmarshalJSON unmarshals json
func (m *createjobdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatejobdetails createjobdetails
	s := struct {
		Model Unmarshalercreatejobdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Name = s.Model.Name
	m.CompartmentId = s.Model.CompartmentId
	m.ScheduleType = s.Model.ScheduleType
	m.Description = s.Model.Description
	m.ManagedDatabaseGroupId = s.Model.ManagedDatabaseGroupId
	m.ManagedDatabaseId = s.Model.ManagedDatabaseId
	m.DatabaseSubType = s.Model.DatabaseSubType
	m.Timeout = s.Model.Timeout
	m.ResultLocation = s.Model.ResultLocation
	m.JobType = s.Model.JobType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createjobdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.JobType {
	case "SQL":
		mm := CreateSqlJobDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetName returns Name
func (m createjobdetails) GetName() *string {
	return m.Name
}

//GetCompartmentId returns CompartmentId
func (m createjobdetails) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetScheduleType returns ScheduleType
func (m createjobdetails) GetScheduleType() CreateJobDetails {
	return m.ScheduleType
}

//GetDescription returns Description
func (m createjobdetails) GetDescription() *string {
	return m.Description
}

//GetManagedDatabaseGroupId returns ManagedDatabaseGroupId
func (m createjobdetails) GetManagedDatabaseGroupId() *string {
	return m.ManagedDatabaseGroupId
}

//GetManagedDatabaseId returns ManagedDatabaseId
func (m createjobdetails) GetManagedDatabaseId() *string {
	return m.ManagedDatabaseId
}

//GetDatabaseSubType returns DatabaseSubType
func (m createjobdetails) GetDatabaseSubType() DatabaseSubTypeEnum {
	return m.DatabaseSubType
}

//GetTimeout returns Timeout
func (m createjobdetails) GetTimeout() *string {
	return m.Timeout
}

//GetResultLocation returns ResultLocation
func (m createjobdetails) GetResultLocation() JobExecutionResultLocation {
	return m.ResultLocation
}

func (m createjobdetails) String() string {
	return common.PointerString(m)
}
