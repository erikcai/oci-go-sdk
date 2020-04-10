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

// ConnectionFromOracleObjectStorageConnectionDetails The Object Storage connection details.
type ConnectionFromOracleObjectStorageConnectionDetails struct {

	// Generated key that can be used in API calls to identify connection. On scenarios where reference to the connection is needed, a value can be passed in create.
	Key *string `mandatory:"false" json:"key"`

	// modelVersion
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value can be edited by the user and it is restricted to 1000 characters
	Name *string `mandatory:"false" json:"name"`

	// Descriptive text for the object.
	Description *string `mandatory:"false" json:"description"`

	// The version of the object, to track changes in the object instance
	ObjectVersion *int `mandatory:"false" json:"objectVersion"`

	// Status of object, can set this to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Value can only contain upper case letters, underscore and numbers. It should begin with upper case letter or underscore. The value can be edited by the user.
	Identifier *string `mandatory:"false" json:"identifier"`

	PrimarySchema *Schema `mandatory:"false" json:"primarySchema"`

	// connectionProperties
	ConnectionProperties []ConnectionProperty `mandatory:"false" json:"connectionProperties"`

	// credentialFileContent
	CredentialFileContent *string `mandatory:"false" json:"credentialFileContent"`

	// The OCI user OCID for the user to connect to.
	UserId *string `mandatory:"false" json:"userId"`

	// fingerPrint
	FingerPrint *string `mandatory:"false" json:"fingerPrint"`

	// passPhrase
	PassPhrase *string `mandatory:"false" json:"passPhrase"`
}

//GetKey returns Key
func (m ConnectionFromOracleObjectStorageConnectionDetails) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m ConnectionFromOracleObjectStorageConnectionDetails) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m ConnectionFromOracleObjectStorageConnectionDetails) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetName returns Name
func (m ConnectionFromOracleObjectStorageConnectionDetails) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m ConnectionFromOracleObjectStorageConnectionDetails) GetDescription() *string {
	return m.Description
}

//GetObjectVersion returns ObjectVersion
func (m ConnectionFromOracleObjectStorageConnectionDetails) GetObjectVersion() *int {
	return m.ObjectVersion
}

//GetObjectStatus returns ObjectStatus
func (m ConnectionFromOracleObjectStorageConnectionDetails) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetIdentifier returns Identifier
func (m ConnectionFromOracleObjectStorageConnectionDetails) GetIdentifier() *string {
	return m.Identifier
}

//GetPrimarySchema returns PrimarySchema
func (m ConnectionFromOracleObjectStorageConnectionDetails) GetPrimarySchema() *Schema {
	return m.PrimarySchema
}

//GetConnectionProperties returns ConnectionProperties
func (m ConnectionFromOracleObjectStorageConnectionDetails) GetConnectionProperties() []ConnectionProperty {
	return m.ConnectionProperties
}

func (m ConnectionFromOracleObjectStorageConnectionDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ConnectionFromOracleObjectStorageConnectionDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeConnectionFromOracleObjectStorageConnectionDetails ConnectionFromOracleObjectStorageConnectionDetails
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeConnectionFromOracleObjectStorageConnectionDetails
	}{
		"ORACLE_OBJECT_STORAGE_CONNECTION",
		(MarshalTypeConnectionFromOracleObjectStorageConnectionDetails)(m),
	}

	return json.Marshal(&s)
}
