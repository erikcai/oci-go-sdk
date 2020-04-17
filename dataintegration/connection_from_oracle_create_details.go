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

// ConnectionFromOracleCreateDetails The Oracle connection details object.
type ConnectionFromOracleCreateDetails struct {

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value can be edited by the user and it is restricted to 1000 characters
	Name *string `mandatory:"true" json:"name"`

	// Value can only contain upper case letters, underscore and numbers. It should begin with upper case letter or underscore. The value can be edited by the user.
	Identifier *string `mandatory:"true" json:"identifier"`

	// Generated key that can be used in API calls to identify connection. On scenarios where reference to the connection is needed, a value can be passed in create.
	Key *string `mandatory:"false" json:"key"`

	// modelVersion
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Descriptive text for the object.
	Description *string `mandatory:"false" json:"description"`

	// Status of object, can set this to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	PrimarySchema *Schema `mandatory:"false" json:"primarySchema"`

	// connectionProperties
	ConnectionProperties []ConnectionProperty `mandatory:"false" json:"connectionProperties"`

	// Connection Username
	Username *string `mandatory:"false" json:"username"`

	// password
	Password *string `mandatory:"false" json:"password"`
}

//GetKey returns Key
func (m ConnectionFromOracleCreateDetails) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m ConnectionFromOracleCreateDetails) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m ConnectionFromOracleCreateDetails) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetName returns Name
func (m ConnectionFromOracleCreateDetails) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m ConnectionFromOracleCreateDetails) GetDescription() *string {
	return m.Description
}

//GetObjectStatus returns ObjectStatus
func (m ConnectionFromOracleCreateDetails) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetIdentifier returns Identifier
func (m ConnectionFromOracleCreateDetails) GetIdentifier() *string {
	return m.Identifier
}

//GetPrimarySchema returns PrimarySchema
func (m ConnectionFromOracleCreateDetails) GetPrimarySchema() *Schema {
	return m.PrimarySchema
}

//GetConnectionProperties returns ConnectionProperties
func (m ConnectionFromOracleCreateDetails) GetConnectionProperties() []ConnectionProperty {
	return m.ConnectionProperties
}

func (m ConnectionFromOracleCreateDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ConnectionFromOracleCreateDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeConnectionFromOracleCreateDetails ConnectionFromOracleCreateDetails
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeConnectionFromOracleCreateDetails
	}{
		"ORACLEDB_CONNECTION",
		(MarshalTypeConnectionFromOracleCreateDetails)(m),
	}

	return json.Marshal(&s)
}
