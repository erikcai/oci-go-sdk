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

// DataAssetFromOracleAtpDataAssetCreate The ATP data asset details.
type DataAssetFromOracleAtpDataAssetCreate struct {

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value can be edited by the user and it is restricted to 1000 characters
	Name *string `mandatory:"true" json:"name"`

	// Value can only contain upper case letters, underscore and numbers. It should begin with upper case letter or underscore. The value can be edited by the user.
	Identifier *string `mandatory:"true" json:"identifier"`

	// Currently not used on data asset creation. Reserved for future.
	Key *string `mandatory:"false" json:"key"`

	// modelVersion
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	// Descriptive text for the object.
	Description *string `mandatory:"false" json:"description"`

	// Status of object, can set this to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// The external key for the object
	ExternalKey *string `mandatory:"false" json:"externalKey"`

	// assetProperties
	AssetProperties map[string]string `mandatory:"false" json:"assetProperties"`

	// host
	Host *string `mandatory:"false" json:"host"`

	// port
	Port *string `mandatory:"false" json:"port"`

	// serviceName
	ServiceName *string `mandatory:"false" json:"serviceName"`

	// driverClass
	DriverClass *string `mandatory:"false" json:"driverClass"`

	// credentialFileContent
	CredentialFileContent *string `mandatory:"false" json:"credentialFileContent"`

	DefaultConnection *ConnectionFromOracleAtpCreateDetails `mandatory:"false" json:"defaultConnection"`
}

//GetKey returns Key
func (m DataAssetFromOracleAtpDataAssetCreate) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m DataAssetFromOracleAtpDataAssetCreate) GetModelVersion() *string {
	return m.ModelVersion
}

//GetName returns Name
func (m DataAssetFromOracleAtpDataAssetCreate) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m DataAssetFromOracleAtpDataAssetCreate) GetDescription() *string {
	return m.Description
}

//GetObjectStatus returns ObjectStatus
func (m DataAssetFromOracleAtpDataAssetCreate) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetIdentifier returns Identifier
func (m DataAssetFromOracleAtpDataAssetCreate) GetIdentifier() *string {
	return m.Identifier
}

//GetExternalKey returns ExternalKey
func (m DataAssetFromOracleAtpDataAssetCreate) GetExternalKey() *string {
	return m.ExternalKey
}

//GetAssetProperties returns AssetProperties
func (m DataAssetFromOracleAtpDataAssetCreate) GetAssetProperties() map[string]string {
	return m.AssetProperties
}

func (m DataAssetFromOracleAtpDataAssetCreate) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m DataAssetFromOracleAtpDataAssetCreate) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDataAssetFromOracleAtpDataAssetCreate DataAssetFromOracleAtpDataAssetCreate
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeDataAssetFromOracleAtpDataAssetCreate
	}{
		"ORACLE_ATP_DATA_ASSET",
		(MarshalTypeDataAssetFromOracleAtpDataAssetCreate)(m),
	}

	return json.Marshal(&s)
}
