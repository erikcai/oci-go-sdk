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

// SqlJob SQL job
type SqlJob struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the job.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment of the job.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Display name of the job.
	Name *string `mandatory:"true" json:"name"`

	// Time when the job was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Time when the job was last updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Description of the job.
	Description *string `mandatory:"false" json:"description"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database Group.
	ManagedDatabaseGroupId *string `mandatory:"false" json:"managedDatabaseGroupId"`

	// Managed database OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) where job has to be executed.
	ManagedDatabaseId *string `mandatory:"false" json:"managedDatabaseId"`

	// Details of managed databases where job needs to be executed
	ManagedDatabasesDetails []JobDatabase `mandatory:"false" json:"managedDatabasesDetails"`

	// Job timeout duration expressed like "1h 10m 15s".
	Timeout *string `mandatory:"false" json:"timeout"`

	ResultLocation JobExecutionResultLocation `mandatory:"false" json:"resultLocation"`

	// Error message if job submission failed else null.
	SubmissionErrorMessage *string `mandatory:"false" json:"submissionErrorMessage"`

	// SQL text to be executed as the job. This is a mandatory field for EXECUTE_SQL operationType.
	SqlText *string `mandatory:"false" json:"sqlText"`

	// Database user name used for executing the SQL job. If managedDatabaseGroupId is provided, this user name should exist on all the databases in the group with the same password.
	UserName *string `mandatory:"false" json:"userName"`

	// Type of the SQL. This is a mandatory field for EXECUTE_SQL operationType.
	SqlType SqlJobSqlTypeEnum `mandatory:"false" json:"sqlType,omitempty"`

	// SQL operation type.
	OperationType SqlJobOperationTypeEnum `mandatory:"true" json:"operationType"`

	// Role of database user
	Role SqlJobRoleEnum `mandatory:"false" json:"role,omitempty"`

	// Subtype of databases (CDB/PDB/NON_CDB) where job needs to be executed. Applicable only when managedDatabaseGroupId is provided.
	DatabaseSubType DatabaseSubTypeEnum `mandatory:"false" json:"databaseSubType,omitempty"`

	// Schedule type for the job.
	ScheduleType JobScheduleTypeEnum `mandatory:"true" json:"scheduleType"`

	// Lifecycle state of the job.
	LifecycleState JobLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

//GetId returns Id
func (m SqlJob) GetId() *string {
	return m.Id
}

//GetCompartmentId returns CompartmentId
func (m SqlJob) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetName returns Name
func (m SqlJob) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m SqlJob) GetDescription() *string {
	return m.Description
}

//GetManagedDatabaseGroupId returns ManagedDatabaseGroupId
func (m SqlJob) GetManagedDatabaseGroupId() *string {
	return m.ManagedDatabaseGroupId
}

//GetManagedDatabaseId returns ManagedDatabaseId
func (m SqlJob) GetManagedDatabaseId() *string {
	return m.ManagedDatabaseId
}

//GetManagedDatabasesDetails returns ManagedDatabasesDetails
func (m SqlJob) GetManagedDatabasesDetails() []JobDatabase {
	return m.ManagedDatabasesDetails
}

//GetDatabaseSubType returns DatabaseSubType
func (m SqlJob) GetDatabaseSubType() DatabaseSubTypeEnum {
	return m.DatabaseSubType
}

//GetScheduleType returns ScheduleType
func (m SqlJob) GetScheduleType() JobScheduleTypeEnum {
	return m.ScheduleType
}

//GetLifecycleState returns LifecycleState
func (m SqlJob) GetLifecycleState() JobLifecycleStateEnum {
	return m.LifecycleState
}

//GetTimeout returns Timeout
func (m SqlJob) GetTimeout() *string {
	return m.Timeout
}

//GetResultLocation returns ResultLocation
func (m SqlJob) GetResultLocation() JobExecutionResultLocation {
	return m.ResultLocation
}

//GetSubmissionErrorMessage returns SubmissionErrorMessage
func (m SqlJob) GetSubmissionErrorMessage() *string {
	return m.SubmissionErrorMessage
}

//GetTimeCreated returns TimeCreated
func (m SqlJob) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m SqlJob) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m SqlJob) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m SqlJob) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSqlJob SqlJob
	s := struct {
		DiscriminatorParam string `json:"jobType"`
		MarshalTypeSqlJob
	}{
		"SQL",
		(MarshalTypeSqlJob)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *SqlJob) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description             *string                    `json:"description"`
		ManagedDatabaseGroupId  *string                    `json:"managedDatabaseGroupId"`
		ManagedDatabaseId       *string                    `json:"managedDatabaseId"`
		ManagedDatabasesDetails []JobDatabase              `json:"managedDatabasesDetails"`
		DatabaseSubType         DatabaseSubTypeEnum        `json:"databaseSubType"`
		Timeout                 *string                    `json:"timeout"`
		ResultLocation          jobexecutionresultlocation `json:"resultLocation"`
		SubmissionErrorMessage  *string                    `json:"submissionErrorMessage"`
		SqlType                 SqlJobSqlTypeEnum          `json:"sqlType"`
		SqlText                 *string                    `json:"sqlText"`
		UserName                *string                    `json:"userName"`
		Role                    SqlJobRoleEnum             `json:"role"`
		Id                      *string                    `json:"id"`
		CompartmentId           *string                    `json:"compartmentId"`
		Name                    *string                    `json:"name"`
		ScheduleType            JobScheduleTypeEnum        `json:"scheduleType"`
		LifecycleState          JobLifecycleStateEnum      `json:"lifecycleState"`
		TimeCreated             *common.SDKTime            `json:"timeCreated"`
		TimeUpdated             *common.SDKTime            `json:"timeUpdated"`
		OperationType           SqlJobOperationTypeEnum    `json:"operationType"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.ManagedDatabaseGroupId = model.ManagedDatabaseGroupId

	m.ManagedDatabaseId = model.ManagedDatabaseId

	m.ManagedDatabasesDetails = make([]JobDatabase, len(model.ManagedDatabasesDetails))
	for i, n := range model.ManagedDatabasesDetails {
		m.ManagedDatabasesDetails[i] = n
	}

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

	m.SubmissionErrorMessage = model.SubmissionErrorMessage

	m.SqlType = model.SqlType

	m.SqlText = model.SqlText

	m.UserName = model.UserName

	m.Role = model.Role

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.Name = model.Name

	m.ScheduleType = model.ScheduleType

	m.LifecycleState = model.LifecycleState

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.OperationType = model.OperationType

	return
}

// SqlJobSqlTypeEnum Enum with underlying type: string
type SqlJobSqlTypeEnum string

// Set of constants representing the allowable values for SqlJobSqlTypeEnum
const (
	SqlJobSqlTypeQuery SqlJobSqlTypeEnum = "QUERY"
	SqlJobSqlTypeDml   SqlJobSqlTypeEnum = "DML"
	SqlJobSqlTypeDdl   SqlJobSqlTypeEnum = "DDL"
	SqlJobSqlTypePlsql SqlJobSqlTypeEnum = "PLSQL"
)

var mappingSqlJobSqlType = map[string]SqlJobSqlTypeEnum{
	"QUERY": SqlJobSqlTypeQuery,
	"DML":   SqlJobSqlTypeDml,
	"DDL":   SqlJobSqlTypeDdl,
	"PLSQL": SqlJobSqlTypePlsql,
}

// GetSqlJobSqlTypeEnumValues Enumerates the set of values for SqlJobSqlTypeEnum
func GetSqlJobSqlTypeEnumValues() []SqlJobSqlTypeEnum {
	values := make([]SqlJobSqlTypeEnum, 0)
	for _, v := range mappingSqlJobSqlType {
		values = append(values, v)
	}
	return values
}

// SqlJobOperationTypeEnum Enum with underlying type: string
type SqlJobOperationTypeEnum string

// Set of constants representing the allowable values for SqlJobOperationTypeEnum
const (
	SqlJobOperationTypeExecuteSql SqlJobOperationTypeEnum = "EXECUTE_SQL"
)

var mappingSqlJobOperationType = map[string]SqlJobOperationTypeEnum{
	"EXECUTE_SQL": SqlJobOperationTypeExecuteSql,
}

// GetSqlJobOperationTypeEnumValues Enumerates the set of values for SqlJobOperationTypeEnum
func GetSqlJobOperationTypeEnumValues() []SqlJobOperationTypeEnum {
	values := make([]SqlJobOperationTypeEnum, 0)
	for _, v := range mappingSqlJobOperationType {
		values = append(values, v)
	}
	return values
}

// SqlJobRoleEnum Enum with underlying type: string
type SqlJobRoleEnum string

// Set of constants representing the allowable values for SqlJobRoleEnum
const (
	SqlJobRoleNormal SqlJobRoleEnum = "NORMAL"
	SqlJobRoleSysdba SqlJobRoleEnum = "SYSDBA"
)

var mappingSqlJobRole = map[string]SqlJobRoleEnum{
	"NORMAL": SqlJobRoleNormal,
	"SYSDBA": SqlJobRoleSysdba,
}

// GetSqlJobRoleEnumValues Enumerates the set of values for SqlJobRoleEnum
func GetSqlJobRoleEnumValues() []SqlJobRoleEnum {
	values := make([]SqlJobRoleEnum, 0)
	for _, v := range mappingSqlJobRole {
		values = append(values, v)
	}
	return values
}
