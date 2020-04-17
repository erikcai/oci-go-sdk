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

// DataAssetFromOracleAtpDataAssetDetails The ATP data asset details.
type DataAssetFromOracleAtpDataAssetDetails struct {

	// Generated key that can be used in API calls to identify data asset.
	Key *string `mandatory:"false" json:"key"`

	// modelVersion
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value can be edited by the user and it is restricted to 1000 characters
	Name *string `mandatory:"false" json:"name"`

	// Descriptive text for the object.
	Description *string `mandatory:"false" json:"description"`

	// Status of object, can set this to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Value can only contain upper case letters, underscore and numbers. It should begin with upper case letter or underscore. The value can be edited by the user.
	Identifier *string `mandatory:"false" json:"identifier"`

	// The external key for the object
	ExternalKey *string `mandatory:"false" json:"externalKey"`

	// assetProperties
	AssetProperties map[string]string `mandatory:"false" json:"assetProperties"`

	NativeTypeSystem *TypeSystem `mandatory:"false" json:"nativeTypeSystem"`

	// The version of the object, to track changes in the object instance
	ObjectVersion *int `mandatory:"false" json:"objectVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

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

	DefaultConnection *ConnectionFromOracleAtpConnectionDetails `mandatory:"false" json:"defaultConnection"`
}

//GetKey returns Key
func (m DataAssetFromOracleAtpDataAssetDetails) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m DataAssetFromOracleAtpDataAssetDetails) GetModelVersion() *string {
	return m.ModelVersion
}

//GetName returns Name
func (m DataAssetFromOracleAtpDataAssetDetails) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m DataAssetFromOracleAtpDataAssetDetails) GetDescription() *string {
	return m.Description
}

//GetObjectStatus returns ObjectStatus
func (m DataAssetFromOracleAtpDataAssetDetails) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetIdentifier returns Identifier
func (m DataAssetFromOracleAtpDataAssetDetails) GetIdentifier() *string {
	return m.Identifier
}

//GetExternalKey returns ExternalKey
func (m DataAssetFromOracleAtpDataAssetDetails) GetExternalKey() *string {
	return m.ExternalKey
}

//GetAssetProperties returns AssetProperties
func (m DataAssetFromOracleAtpDataAssetDetails) GetAssetProperties() map[string]string {
	return m.AssetProperties
}

//GetNativeTypeSystem returns NativeTypeSystem
func (m DataAssetFromOracleAtpDataAssetDetails) GetNativeTypeSystem() *TypeSystem {
	return m.NativeTypeSystem
}

//GetObjectVersion returns ObjectVersion
func (m DataAssetFromOracleAtpDataAssetDetails) GetObjectVersion() *int {
	return m.ObjectVersion
}

//GetParentRef returns ParentRef
func (m DataAssetFromOracleAtpDataAssetDetails) GetParentRef() *ParentReference {
	return m.ParentRef
}

func (m DataAssetFromOracleAtpDataAssetDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m DataAssetFromOracleAtpDataAssetDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDataAssetFromOracleAtpDataAssetDetails DataAssetFromOracleAtpDataAssetDetails
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeDataAssetFromOracleAtpDataAssetDetails
	}{
		"ORACLE_ATP_DATA_ASSET",
		(MarshalTypeDataAssetFromOracleAtpDataAssetDetails)(m),
	}

	return json.Marshal(&s)
}
