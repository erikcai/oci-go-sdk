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

// DataAssetDetailsForCreate The details for the data asset.
type DataAssetDetailsForCreate interface {

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value can be edited by the user and it is restricted to 1000 characters
	GetName() *string

	// Value can only contain upper case letters, underscore and numbers. It should begin with upper case letter or underscore. The value can be edited by the user.
	GetIdentifier() *string

	// Currently not used on data asset creation. Reserved for future.
	GetKey() *string

	// modelVersion
	GetModelVersion() *string

	// Descriptive text for the object.
	GetDescription() *string

	// Status of object, can set this to value 1 for shallow references across objects, other values reserved.
	GetObjectStatus() *int

	// The external key for the object
	GetExternalKey() *string

	// assetProperties
	GetAssetProperties() map[string]string
}

type dataassetdetailsforcreate struct {
	JsonData        []byte
	Name            *string           `mandatory:"true" json:"name"`
	Identifier      *string           `mandatory:"true" json:"identifier"`
	Key             *string           `mandatory:"false" json:"key"`
	ModelVersion    *string           `mandatory:"false" json:"modelVersion"`
	Description     *string           `mandatory:"false" json:"description"`
	ObjectStatus    *int              `mandatory:"false" json:"objectStatus"`
	ExternalKey     *string           `mandatory:"false" json:"externalKey"`
	AssetProperties map[string]string `mandatory:"false" json:"assetProperties"`
	ModelType       string            `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *dataassetdetailsforcreate) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdataassetdetailsforcreate dataassetdetailsforcreate
	s := struct {
		Model Unmarshalerdataassetdetailsforcreate
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Name = s.Model.Name
	m.Identifier = s.Model.Identifier
	m.Key = s.Model.Key
	m.ModelVersion = s.Model.ModelVersion
	m.Description = s.Model.Description
	m.ObjectStatus = s.Model.ObjectStatus
	m.ExternalKey = s.Model.ExternalKey
	m.AssetProperties = s.Model.AssetProperties
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *dataassetdetailsforcreate) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "ORACLE_DATA_ASSET":
		mm := DataAssetFromOracleDataAssetCreate{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_ATP_DATA_ASSET":
		mm := DataAssetFromOracleAtpDataAssetCreate{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_ADWC_DATA_ASSET":
		mm := DataAssetFromOracleAdwcDataAssetCreate{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_OBJECT_STORAGE_DATA_ASSET":
		mm := DataAssetFromOracleObjectStorageDataAssetCreate{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetName returns Name
func (m dataassetdetailsforcreate) GetName() *string {
	return m.Name
}

//GetIdentifier returns Identifier
func (m dataassetdetailsforcreate) GetIdentifier() *string {
	return m.Identifier
}

//GetKey returns Key
func (m dataassetdetailsforcreate) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m dataassetdetailsforcreate) GetModelVersion() *string {
	return m.ModelVersion
}

//GetDescription returns Description
func (m dataassetdetailsforcreate) GetDescription() *string {
	return m.Description
}

//GetObjectStatus returns ObjectStatus
func (m dataassetdetailsforcreate) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetExternalKey returns ExternalKey
func (m dataassetdetailsforcreate) GetExternalKey() *string {
	return m.ExternalKey
}

//GetAssetProperties returns AssetProperties
func (m dataassetdetailsforcreate) GetAssetProperties() map[string]string {
	return m.AssetProperties
}

func (m dataassetdetailsforcreate) String() string {
	return common.PointerString(m)
}

// DataAssetDetailsForCreateModelTypeEnum Enum with underlying type: string
type DataAssetDetailsForCreateModelTypeEnum string

// Set of constants representing the allowable values for DataAssetDetailsForCreateModelTypeEnum
const (
	DataAssetDetailsForCreateModelTypeOracleDataAsset              DataAssetDetailsForCreateModelTypeEnum = "ORACLE_DATA_ASSET"
	DataAssetDetailsForCreateModelTypeMysqlStore                   DataAssetDetailsForCreateModelTypeEnum = "MYSQL_STORE"
	DataAssetDetailsForCreateModelTypeOracleObjectStorageDataAsset DataAssetDetailsForCreateModelTypeEnum = "ORACLE_OBJECT_STORAGE_DATA_ASSET"
	DataAssetDetailsForCreateModelTypeOracleAtpDataAsset           DataAssetDetailsForCreateModelTypeEnum = "ORACLE_ATP_DATA_ASSET"
	DataAssetDetailsForCreateModelTypeSqlserverStore               DataAssetDetailsForCreateModelTypeEnum = "SQLSERVER_STORE"
	DataAssetDetailsForCreateModelTypeOracleAdwcDataAsset          DataAssetDetailsForCreateModelTypeEnum = "ORACLE_ADWC_DATA_ASSET"
	DataAssetDetailsForCreateModelTypeDerbyStore                   DataAssetDetailsForCreateModelTypeEnum = "DERBY_STORE"
)

var mappingDataAssetDetailsForCreateModelType = map[string]DataAssetDetailsForCreateModelTypeEnum{
	"ORACLE_DATA_ASSET":                DataAssetDetailsForCreateModelTypeOracleDataAsset,
	"MYSQL_STORE":                      DataAssetDetailsForCreateModelTypeMysqlStore,
	"ORACLE_OBJECT_STORAGE_DATA_ASSET": DataAssetDetailsForCreateModelTypeOracleObjectStorageDataAsset,
	"ORACLE_ATP_DATA_ASSET":            DataAssetDetailsForCreateModelTypeOracleAtpDataAsset,
	"SQLSERVER_STORE":                  DataAssetDetailsForCreateModelTypeSqlserverStore,
	"ORACLE_ADWC_DATA_ASSET":           DataAssetDetailsForCreateModelTypeOracleAdwcDataAsset,
	"DERBY_STORE":                      DataAssetDetailsForCreateModelTypeDerbyStore,
}

// GetDataAssetDetailsForCreateModelTypeEnumValues Enumerates the set of values for DataAssetDetailsForCreateModelTypeEnum
func GetDataAssetDetailsForCreateModelTypeEnumValues() []DataAssetDetailsForCreateModelTypeEnum {
	values := make([]DataAssetDetailsForCreateModelTypeEnum, 0)
	for _, v := range mappingDataAssetDetailsForCreateModelType {
		values = append(values, v)
	}
	return values
}
