// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
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

// ConnectionFromOracleAdwcCreateDetails The ADWC connection details object.
type ConnectionFromOracleAdwcCreateDetails struct {

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
func (m ConnectionFromOracleAdwcCreateDetails) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m ConnectionFromOracleAdwcCreateDetails) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m ConnectionFromOracleAdwcCreateDetails) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetName returns Name
func (m ConnectionFromOracleAdwcCreateDetails) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m ConnectionFromOracleAdwcCreateDetails) GetDescription() *string {
	return m.Description
}

//GetObjectStatus returns ObjectStatus
func (m ConnectionFromOracleAdwcCreateDetails) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetIdentifier returns Identifier
func (m ConnectionFromOracleAdwcCreateDetails) GetIdentifier() *string {
	return m.Identifier
}

//GetPrimarySchema returns PrimarySchema
func (m ConnectionFromOracleAdwcCreateDetails) GetPrimarySchema() *Schema {
	return m.PrimarySchema
}

//GetConnectionProperties returns ConnectionProperties
func (m ConnectionFromOracleAdwcCreateDetails) GetConnectionProperties() []ConnectionProperty {
	return m.ConnectionProperties
}

func (m ConnectionFromOracleAdwcCreateDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ConnectionFromOracleAdwcCreateDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeConnectionFromOracleAdwcCreateDetails ConnectionFromOracleAdwcCreateDetails
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeConnectionFromOracleAdwcCreateDetails
	}{
		"ORACLE_ADWC_CONNECTION",
		(MarshalTypeConnectionFromOracleAdwcCreateDetails)(m),
	}

	return json.Marshal(&s)
}
