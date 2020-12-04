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

// CreateSqlJobDetails Details specific to SQL job request.
type CreateSqlJobDetails struct {

	// Name of the job. Valid characters are uppercase or lowercase letters,
	// numbers, and underscores. It is unmodifiable and must be unique in a
	// compartment. It must begin with an alphabetic character.
	Name *string `mandatory:"true" json:"name"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment of the job.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Description of the job.
	Description *string `mandatory:"false" json:"description"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database Group.
	ManagedDatabaseGroupId *string `mandatory:"false" json:"managedDatabaseGroupId"`

	// Managed database OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) where job has to be executed.
	ManagedDatabaseId *string `mandatory:"false" json:"managedDatabaseId"`

	// Job timeout duration expressed like "1h 10m 15s".
	Timeout *string `mandatory:"false" json:"timeout"`

	ResultLocation JobExecutionResultLocation `mandatory:"false" json:"resultLocation"`

	// SQL text to be executed as the job.
	SqlText *string `mandatory:"false" json:"sqlText"`

	// Database user name used for executing the SQL job. If managedDatabaseGroupId is provided, this user name should exist on all the databases in the group with the same password.
	UserName *string `mandatory:"false" json:"userName"`

	// Password for the database user used for executing the SQL job.
	Password *string `mandatory:"false" json:"password"`

	// Schedule type for the job.
	ScheduleType JobScheduleTypeEnum `mandatory:"true" json:"scheduleType"`

	// Subtype of databases (CDB/PDB/NON_CDB) where job needs to be executed. Applicable only when managedDatabaseGroupId is provided.
	DatabaseSubType DatabaseSubTypeEnum `mandatory:"false" json:"databaseSubType,omitempty"`

	SqlType SqlJobSqlTypeEnum `mandatory:"false" json:"sqlType,omitempty"`

	// SQL operation type.
	OperationType SqlJobOperationTypeEnum `mandatory:"true" json:"operationType"`

	// Role of database user
	Role SqlJobRoleEnum `mandatory:"false" json:"role,omitempty"`
}

//GetName returns Name
func (m CreateSqlJobDetails) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m CreateSqlJobDetails) GetDescription() *string {
	return m.Description
}

//GetCompartmentId returns CompartmentId
func (m CreateSqlJobDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetManagedDatabaseGroupId returns ManagedDatabaseGroupId
func (m CreateSqlJobDetails) GetManagedDatabaseGroupId() *string {
	return m.ManagedDatabaseGroupId
}

//GetManagedDatabaseId returns ManagedDatabaseId
func (m CreateSqlJobDetails) GetManagedDatabaseId() *string {
	return m.ManagedDatabaseId
}

//GetDatabaseSubType returns DatabaseSubType
func (m CreateSqlJobDetails) GetDatabaseSubType() DatabaseSubTypeEnum {
	return m.DatabaseSubType
}

//GetScheduleType returns ScheduleType
func (m CreateSqlJobDetails) GetScheduleType() JobScheduleTypeEnum {
	return m.ScheduleType
}

//GetTimeout returns Timeout
func (m CreateSqlJobDetails) GetTimeout() *string {
	return m.Timeout
}

//GetResultLocation returns ResultLocation
func (m CreateSqlJobDetails) GetResultLocation() JobExecutionResultLocation {
	return m.ResultLocation
}

func (m CreateSqlJobDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateSqlJobDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateSqlJobDetails CreateSqlJobDetails
	s := struct {
		DiscriminatorParam string `json:"jobType"`
		MarshalTypeCreateSqlJobDetails
	}{
		"SQL",
		(MarshalTypeCreateSqlJobDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CreateSqlJobDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description            *string                    `json:"description"`
		ManagedDatabaseGroupId *string                    `json:"managedDatabaseGroupId"`
		ManagedDatabaseId      *string                    `json:"managedDatabaseId"`
		DatabaseSubType        DatabaseSubTypeEnum        `json:"databaseSubType"`
		Timeout                *string                    `json:"timeout"`
		ResultLocation         jobexecutionresultlocation `json:"resultLocation"`
		SqlText                *string                    `json:"sqlText"`
		SqlType                SqlJobSqlTypeEnum          `json:"sqlType"`
		UserName               *string                    `json:"userName"`
		Password               *string                    `json:"password"`
		Role                   SqlJobRoleEnum             `json:"role"`
		Name                   *string                    `json:"name"`
		CompartmentId          *string                    `json:"compartmentId"`
		ScheduleType           JobScheduleTypeEnum        `json:"scheduleType"`
		OperationType          SqlJobOperationTypeEnum    `json:"operationType"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.ManagedDatabaseGroupId = model.ManagedDatabaseGroupId

	m.ManagedDatabaseId = model.ManagedDatabaseId

	m.DatabaseSubType = model.DatabaseSubType

	m.Timeout = model.Timeout

	nn, e = model.ResultLocation.UnmarshalPolymorphicJSON(model.ResultLocation.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ResultLocation = nn.(JobExecutionResultLocation)
	} else {
		m.ResultLocation = nil
	}

	m.SqlText = model.SqlText

	m.SqlType = model.SqlType

	m.UserName = model.UserName

	m.Password = model.Password

	m.Role = model.Role

	m.Name = model.Name

	m.CompartmentId = model.CompartmentId

	m.ScheduleType = model.ScheduleType

	m.OperationType = model.OperationType

	return
}
