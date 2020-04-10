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

// DataAssetFromOracleObjectStorageDataAssetDetails The Object Storage data asset details.
type DataAssetFromOracleObjectStorageDataAssetDetails struct {

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

	// url
	Url *string `mandatory:"false" json:"url"`

	// The OCI tenancy OCID.
	TenancyId *string `mandatory:"false" json:"tenancyId"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value can be edited by the user and it is restricted to 1000 characters
	Namespace *string `mandatory:"false" json:"namespace"`

	DefaultConnection *ConnectionFromOracleObjectStorageConnectionDetails `mandatory:"false" json:"defaultConnection"`
}

//GetKey returns Key
func (m DataAssetFromOracleObjectStorageDataAssetDetails) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m DataAssetFromOracleObjectStorageDataAssetDetails) GetModelVersion() *string {
	return m.ModelVersion
}

//GetName returns Name
func (m DataAssetFromOracleObjectStorageDataAssetDetails) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m DataAssetFromOracleObjectStorageDataAssetDetails) GetDescription() *string {
	return m.Description
}

//GetObjectStatus returns ObjectStatus
func (m DataAssetFromOracleObjectStorageDataAssetDetails) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetIdentifier returns Identifier
func (m DataAssetFromOracleObjectStorageDataAssetDetails) GetIdentifier() *string {
	return m.Identifier
}

//GetExternalKey returns ExternalKey
func (m DataAssetFromOracleObjectStorageDataAssetDetails) GetExternalKey() *string {
	return m.ExternalKey
}

//GetAssetProperties returns AssetProperties
func (m DataAssetFromOracleObjectStorageDataAssetDetails) GetAssetProperties() map[string]string {
	return m.AssetProperties
}

//GetNativeTypeSystem returns NativeTypeSystem
func (m DataAssetFromOracleObjectStorageDataAssetDetails) GetNativeTypeSystem() *TypeSystem {
	return m.NativeTypeSystem
}

//GetObjectVersion returns ObjectVersion
func (m DataAssetFromOracleObjectStorageDataAssetDetails) GetObjectVersion() *int {
	return m.ObjectVersion
}

//GetParentRef returns ParentRef
func (m DataAssetFromOracleObjectStorageDataAssetDetails) GetParentRef() *ParentReference {
	return m.ParentRef
}

func (m DataAssetFromOracleObjectStorageDataAssetDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m DataAssetFromOracleObjectStorageDataAssetDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDataAssetFromOracleObjectStorageDataAssetDetails DataAssetFromOracleObjectStorageDataAssetDetails
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeDataAssetFromOracleObjectStorageDataAssetDetails
	}{
		"ORACLE_OBJECT_STORAGE_DATA_ASSET",
		(MarshalTypeDataAssetFromOracleObjectStorageDataAssetDetails)(m),
	}

	return json.Marshal(&s)
}
