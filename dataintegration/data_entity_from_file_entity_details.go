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

// DataEntityFromFileEntityDetails The file data entity details.
type DataEntityFromFileEntityDetails struct {

	// Object key
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

	// The external key for the object.
	ExternalKey *string `mandatory:"false" json:"externalKey"`

	Shape *Shape `mandatory:"false" json:"shape"`

	// shapeId
	ShapeId *string `mandatory:"false" json:"shapeId"`

	Types *TypeLibrary `mandatory:"false" json:"types"`

	// otherTypeLabel
	OtherTypeLabel *string `mandatory:"false" json:"otherTypeLabel"`

	// uniqueKeys
	UniqueKeys []UniqueKey `mandatory:"false" json:"uniqueKeys"`

	// foreignKeys
	ForeignKeys []ForeignKey `mandatory:"false" json:"foreignKeys"`

	// resourceName
	ResourceName *string `mandatory:"false" json:"resourceName"`

	// header
	IsHeader *bool `mandatory:"false" json:"isHeader"`

	// firstDataRow
	FirstDataRow *float32 `mandatory:"false" json:"firstDataRow"`

	DataFormat *DataFormat `mandatory:"false" json:"dataFormat"`

	// Status of object, can set this to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Value can only contain upper case letters, underscore and numbers. It should begin with upper case letter or underscore. The value can be edited by the user.
	Identifier *string `mandatory:"false" json:"identifier"`

	// entityType
	EntityType DataEntityFromFileEntityDetailsEntityTypeEnum `mandatory:"false" json:"entityType,omitempty"`
}

func (m DataEntityFromFileEntityDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m DataEntityFromFileEntityDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDataEntityFromFileEntityDetails DataEntityFromFileEntityDetails
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeDataEntityFromFileEntityDetails
	}{
		"FILE_ENTITY",
		(MarshalTypeDataEntityFromFileEntityDetails)(m),
	}

	return json.Marshal(&s)
}

// DataEntityFromFileEntityDetailsEntityTypeEnum Enum with underlying type: string
type DataEntityFromFileEntityDetailsEntityTypeEnum string

// Set of constants representing the allowable values for DataEntityFromFileEntityDetailsEntityTypeEnum
const (
	DataEntityFromFileEntityDetailsEntityTypeTable  DataEntityFromFileEntityDetailsEntityTypeEnum = "TABLE"
	DataEntityFromFileEntityDetailsEntityTypeView   DataEntityFromFileEntityDetailsEntityTypeEnum = "VIEW"
	DataEntityFromFileEntityDetailsEntityTypeFile   DataEntityFromFileEntityDetailsEntityTypeEnum = "FILE"
	DataEntityFromFileEntityDetailsEntityTypeQueue  DataEntityFromFileEntityDetailsEntityTypeEnum = "QUEUE"
	DataEntityFromFileEntityDetailsEntityTypeStream DataEntityFromFileEntityDetailsEntityTypeEnum = "STREAM"
	DataEntityFromFileEntityDetailsEntityTypeOther  DataEntityFromFileEntityDetailsEntityTypeEnum = "OTHER"
)

var mappingDataEntityFromFileEntityDetailsEntityType = map[string]DataEntityFromFileEntityDetailsEntityTypeEnum{
	"TABLE":  DataEntityFromFileEntityDetailsEntityTypeTable,
	"VIEW":   DataEntityFromFileEntityDetailsEntityTypeView,
	"FILE":   DataEntityFromFileEntityDetailsEntityTypeFile,
	"QUEUE":  DataEntityFromFileEntityDetailsEntityTypeQueue,
	"STREAM": DataEntityFromFileEntityDetailsEntityTypeStream,
	"OTHER":  DataEntityFromFileEntityDetailsEntityTypeOther,
}

// GetDataEntityFromFileEntityDetailsEntityTypeEnumValues Enumerates the set of values for DataEntityFromFileEntityDetailsEntityTypeEnum
func GetDataEntityFromFileEntityDetailsEntityTypeEnumValues() []DataEntityFromFileEntityDetailsEntityTypeEnum {
	values := make([]DataEntityFromFileEntityDetailsEntityTypeEnum, 0)
	for _, v := range mappingDataEntityFromFileEntityDetailsEntityType {
		values = append(values, v)
	}
	return values
}
