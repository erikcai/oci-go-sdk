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

// DataAssetDetails The details for the data asset.
type DataAssetDetails interface {

	// Generated key that can be used in API calls to identify data asset.
	GetKey() *string

	// modelVersion
	GetModelVersion() *string

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value can be edited by the user and it is restricted to 1000 characters
	GetName() *string

	// Descriptive text for the object.
	GetDescription() *string

	// Status of object, can set this to value 1 for shallow references across objects, other values reserved.
	GetObjectStatus() *int

	// Value can only contain upper case letters, underscore and numbers. It should begin with upper case letter or underscore. The value can be edited by the user.
	GetIdentifier() *string

	// The external key for the object
	GetExternalKey() *string

	// assetProperties
	GetAssetProperties() map[string]string

	GetNativeTypeSystem() *TypeSystem

	// The version of the object, to track changes in the object instance
	GetObjectVersion() *int

	GetParentRef() *ParentReference
}

type dataassetdetails struct {
	JsonData         []byte
	Key              *string           `mandatory:"false" json:"key"`
	ModelVersion     *string           `mandatory:"false" json:"modelVersion"`
	Name             *string           `mandatory:"false" json:"name"`
	Description      *string           `mandatory:"false" json:"description"`
	ObjectStatus     *int              `mandatory:"false" json:"objectStatus"`
	Identifier       *string           `mandatory:"false" json:"identifier"`
	ExternalKey      *string           `mandatory:"false" json:"externalKey"`
	AssetProperties  map[string]string `mandatory:"false" json:"assetProperties"`
	NativeTypeSystem *TypeSystem       `mandatory:"false" json:"nativeTypeSystem"`
	ObjectVersion    *int              `mandatory:"false" json:"objectVersion"`
	ParentRef        *ParentReference  `mandatory:"false" json:"parentRef"`
	ModelType        string            `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *dataassetdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdataassetdetails dataassetdetails
	s := struct {
		Model Unmarshalerdataassetdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Key = s.Model.Key
	m.ModelVersion = s.Model.ModelVersion
	m.Name = s.Model.Name
	m.Description = s.Model.Description
	m.ObjectStatus = s.Model.ObjectStatus
	m.Identifier = s.Model.Identifier
	m.ExternalKey = s.Model.ExternalKey
	m.AssetProperties = s.Model.AssetProperties
	m.NativeTypeSystem = s.Model.NativeTypeSystem
	m.ObjectVersion = s.Model.ObjectVersion
	m.ParentRef = s.Model.ParentRef
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *dataassetdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "ORACLE_ATP_DATA_ASSET":
		mm := DataAssetFromOracleAtpDataAssetDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_DATA_ASSET":
		mm := DataAssetFromOracleDataAssetDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_OBJECT_STORAGE_DATA_ASSET":
		mm := DataAssetFromOracleObjectStorageDataAssetDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_ADWC_DATA_ASSET":
		mm := DataAssetFromOracleAdwcDataAssetDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetKey returns Key
func (m dataassetdetails) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m dataassetdetails) GetModelVersion() *string {
	return m.ModelVersion
}

//GetName returns Name
func (m dataassetdetails) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m dataassetdetails) GetDescription() *string {
	return m.Description
}

//GetObjectStatus returns ObjectStatus
func (m dataassetdetails) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetIdentifier returns Identifier
func (m dataassetdetails) GetIdentifier() *string {
	return m.Identifier
}

//GetExternalKey returns ExternalKey
func (m dataassetdetails) GetExternalKey() *string {
	return m.ExternalKey
}

//GetAssetProperties returns AssetProperties
func (m dataassetdetails) GetAssetProperties() map[string]string {
	return m.AssetProperties
}

//GetNativeTypeSystem returns NativeTypeSystem
func (m dataassetdetails) GetNativeTypeSystem() *TypeSystem {
	return m.NativeTypeSystem
}

//GetObjectVersion returns ObjectVersion
func (m dataassetdetails) GetObjectVersion() *int {
	return m.ObjectVersion
}

//GetParentRef returns ParentRef
func (m dataassetdetails) GetParentRef() *ParentReference {
	return m.ParentRef
}

func (m dataassetdetails) String() string {
	return common.PointerString(m)
}

// DataAssetDetailsModelTypeEnum Enum with underlying type: string
type DataAssetDetailsModelTypeEnum string

// Set of constants representing the allowable values for DataAssetDetailsModelTypeEnum
const (
	DataAssetDetailsModelTypeOracleDataAsset              DataAssetDetailsModelTypeEnum = "ORACLE_DATA_ASSET"
	DataAssetDetailsModelTypeMysqlStore                   DataAssetDetailsModelTypeEnum = "MYSQL_STORE"
	DataAssetDetailsModelTypeOracleObjectStorageDataAsset DataAssetDetailsModelTypeEnum = "ORACLE_OBJECT_STORAGE_DATA_ASSET"
	DataAssetDetailsModelTypeOracleAtpDataAsset           DataAssetDetailsModelTypeEnum = "ORACLE_ATP_DATA_ASSET"
	DataAssetDetailsModelTypeSqlserverStore               DataAssetDetailsModelTypeEnum = "SQLSERVER_STORE"
	DataAssetDetailsModelTypeOracleAdwcDataAsset          DataAssetDetailsModelTypeEnum = "ORACLE_ADWC_DATA_ASSET"
	DataAssetDetailsModelTypeDerbyStore                   DataAssetDetailsModelTypeEnum = "DERBY_STORE"
)

var mappingDataAssetDetailsModelType = map[string]DataAssetDetailsModelTypeEnum{
	"ORACLE_DATA_ASSET":                DataAssetDetailsModelTypeOracleDataAsset,
	"MYSQL_STORE":                      DataAssetDetailsModelTypeMysqlStore,
	"ORACLE_OBJECT_STORAGE_DATA_ASSET": DataAssetDetailsModelTypeOracleObjectStorageDataAsset,
	"ORACLE_ATP_DATA_ASSET":            DataAssetDetailsModelTypeOracleAtpDataAsset,
	"SQLSERVER_STORE":                  DataAssetDetailsModelTypeSqlserverStore,
	"ORACLE_ADWC_DATA_ASSET":           DataAssetDetailsModelTypeOracleAdwcDataAsset,
	"DERBY_STORE":                      DataAssetDetailsModelTypeDerbyStore,
}

// GetDataAssetDetailsModelTypeEnumValues Enumerates the set of values for DataAssetDetailsModelTypeEnum
func GetDataAssetDetailsModelTypeEnumValues() []DataAssetDetailsModelTypeEnum {
	values := make([]DataAssetDetailsModelTypeEnum, 0)
	for _, v := range mappingDataAssetDetailsModelType {
		values = append(values, v)
	}
	return values
}
