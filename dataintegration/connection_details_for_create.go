// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration Service API Specification
//
// Data Integration Service API Specification
//

package dataintegration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// ConnectionDetailsForCreate The connection details object.
type ConnectionDetailsForCreate interface {

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value can be edited by the user and it is restricted to 1000 characters
	GetName() *string

	// Value can only contain upper case letters, underscore and numbers. It should begin with upper case letter or underscore. The value can be edited by the user.
	GetIdentifier() *string

	// Generated key that can be used in API calls to identify connection. On scenarios where reference to the connection is needed, a value can be passed in create.
	GetKey() *string

	// modelVersion
	GetModelVersion() *string

	GetParentRef() *ParentReference

	// Descriptive text for the object.
	GetDescription() *string

	// Status of object, can set this to value 1 for shallow references across objects, other values reserved.
	GetObjectStatus() *int

	GetPrimarySchema() *Schema

	// connectionProperties
	GetConnectionProperties() []ConnectionProperty
}

type connectiondetailsforcreate struct {
	JsonData             []byte
	Name                 *string              `mandatory:"true" json:"name"`
	Identifier           *string              `mandatory:"true" json:"identifier"`
	Key                  *string              `mandatory:"false" json:"key"`
	ModelVersion         *string              `mandatory:"false" json:"modelVersion"`
	ParentRef            *ParentReference     `mandatory:"false" json:"parentRef"`
	Description          *string              `mandatory:"false" json:"description"`
	ObjectStatus         *int                 `mandatory:"false" json:"objectStatus"`
	PrimarySchema        *Schema              `mandatory:"false" json:"primarySchema"`
	ConnectionProperties []ConnectionProperty `mandatory:"false" json:"connectionProperties"`
	ModelType            string               `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *connectiondetailsforcreate) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerconnectiondetailsforcreate connectiondetailsforcreate
	s := struct {
		Model Unmarshalerconnectiondetailsforcreate
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Name = s.Model.Name
	m.Identifier = s.Model.Identifier
	m.Key = s.Model.Key
	m.ModelVersion = s.Model.ModelVersion
	m.ParentRef = s.Model.ParentRef
	m.Description = s.Model.Description
	m.ObjectStatus = s.Model.ObjectStatus
	m.PrimarySchema = s.Model.PrimarySchema
	m.ConnectionProperties = s.Model.ConnectionProperties
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *connectiondetailsforcreate) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "ORACLE_ADWC_CONNECTION":
		mm := ConnectionFromOracleAdwcCreateDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_ATP_CONNECTION":
		mm := ConnectionFromOracleAtpCreateDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLEDB_CONNECTION":
		mm := ConnectionFromOracleCreateDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_OBJECT_STORAGE_CONNECTION":
		mm := ConnectionFromOracleObjectStorageCreateDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetName returns Name
func (m connectiondetailsforcreate) GetName() *string {
	return m.Name
}

//GetIdentifier returns Identifier
func (m connectiondetailsforcreate) GetIdentifier() *string {
	return m.Identifier
}

//GetKey returns Key
func (m connectiondetailsforcreate) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m connectiondetailsforcreate) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m connectiondetailsforcreate) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetDescription returns Description
func (m connectiondetailsforcreate) GetDescription() *string {
	return m.Description
}

//GetObjectStatus returns ObjectStatus
func (m connectiondetailsforcreate) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetPrimarySchema returns PrimarySchema
func (m connectiondetailsforcreate) GetPrimarySchema() *Schema {
	return m.PrimarySchema
}

//GetConnectionProperties returns ConnectionProperties
func (m connectiondetailsforcreate) GetConnectionProperties() []ConnectionProperty {
	return m.ConnectionProperties
}

func (m connectiondetailsforcreate) String() string {
	return common.PointerString(m)
}

// ConnectionDetailsForCreateModelTypeEnum Enum with underlying type: string
type ConnectionDetailsForCreateModelTypeEnum string

// Set of constants representing the allowable values for ConnectionDetailsForCreateModelTypeEnum
const (
	ConnectionDetailsForCreateModelTypeOracleAdwcConnection          ConnectionDetailsForCreateModelTypeEnum = "ORACLE_ADWC_CONNECTION"
	ConnectionDetailsForCreateModelTypeOracleAtpConnection           ConnectionDetailsForCreateModelTypeEnum = "ORACLE_ATP_CONNECTION"
	ConnectionDetailsForCreateModelTypeOracleObjectStorageConnection ConnectionDetailsForCreateModelTypeEnum = "ORACLE_OBJECT_STORAGE_CONNECTION"
	ConnectionDetailsForCreateModelTypeOracledbConnection            ConnectionDetailsForCreateModelTypeEnum = "ORACLEDB_CONNECTION"
)

var mappingConnectionDetailsForCreateModelType = map[string]ConnectionDetailsForCreateModelTypeEnum{
	"ORACLE_ADWC_CONNECTION":           ConnectionDetailsForCreateModelTypeOracleAdwcConnection,
	"ORACLE_ATP_CONNECTION":            ConnectionDetailsForCreateModelTypeOracleAtpConnection,
	"ORACLE_OBJECT_STORAGE_CONNECTION": ConnectionDetailsForCreateModelTypeOracleObjectStorageConnection,
	"ORACLEDB_CONNECTION":              ConnectionDetailsForCreateModelTypeOracledbConnection,
}

// GetConnectionDetailsForCreateModelTypeEnumValues Enumerates the set of values for ConnectionDetailsForCreateModelTypeEnum
func GetConnectionDetailsForCreateModelTypeEnumValues() []ConnectionDetailsForCreateModelTypeEnum {
	values := make([]ConnectionDetailsForCreateModelTypeEnum, 0)
	for _, v := range mappingConnectionDetailsForCreateModelType {
		values = append(values, v)
	}
	return values
}
