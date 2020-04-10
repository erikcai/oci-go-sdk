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

// DataEntityFromTableEntityDetails The table entity data entity.
type DataEntityFromTableEntityDetails struct {

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

	// Status of object, can set this to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Value can only contain upper case letters, underscore and numbers. It should begin with upper case letter or underscore. The value can be edited by the user.
	Identifier *string `mandatory:"false" json:"identifier"`

	// entityType
	EntityType DataEntityFromTableEntityDetailsEntityTypeEnum `mandatory:"false" json:"entityType,omitempty"`
}

func (m DataEntityFromTableEntityDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m DataEntityFromTableEntityDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDataEntityFromTableEntityDetails DataEntityFromTableEntityDetails
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeDataEntityFromTableEntityDetails
	}{
		"TABLE_ENTITY",
		(MarshalTypeDataEntityFromTableEntityDetails)(m),
	}

	return json.Marshal(&s)
}

// DataEntityFromTableEntityDetailsEntityTypeEnum Enum with underlying type: string
type DataEntityFromTableEntityDetailsEntityTypeEnum string

// Set of constants representing the allowable values for DataEntityFromTableEntityDetailsEntityTypeEnum
const (
	DataEntityFromTableEntityDetailsEntityTypeTable  DataEntityFromTableEntityDetailsEntityTypeEnum = "TABLE"
	DataEntityFromTableEntityDetailsEntityTypeView   DataEntityFromTableEntityDetailsEntityTypeEnum = "VIEW"
	DataEntityFromTableEntityDetailsEntityTypeFile   DataEntityFromTableEntityDetailsEntityTypeEnum = "FILE"
	DataEntityFromTableEntityDetailsEntityTypeQueue  DataEntityFromTableEntityDetailsEntityTypeEnum = "QUEUE"
	DataEntityFromTableEntityDetailsEntityTypeStream DataEntityFromTableEntityDetailsEntityTypeEnum = "STREAM"
	DataEntityFromTableEntityDetailsEntityTypeOther  DataEntityFromTableEntityDetailsEntityTypeEnum = "OTHER"
)

var mappingDataEntityFromTableEntityDetailsEntityType = map[string]DataEntityFromTableEntityDetailsEntityTypeEnum{
	"TABLE":  DataEntityFromTableEntityDetailsEntityTypeTable,
	"VIEW":   DataEntityFromTableEntityDetailsEntityTypeView,
	"FILE":   DataEntityFromTableEntityDetailsEntityTypeFile,
	"QUEUE":  DataEntityFromTableEntityDetailsEntityTypeQueue,
	"STREAM": DataEntityFromTableEntityDetailsEntityTypeStream,
	"OTHER":  DataEntityFromTableEntityDetailsEntityTypeOther,
}

// GetDataEntityFromTableEntityDetailsEntityTypeEnumValues Enumerates the set of values for DataEntityFromTableEntityDetailsEntityTypeEnum
func GetDataEntityFromTableEntityDetailsEntityTypeEnumValues() []DataEntityFromTableEntityDetailsEntityTypeEnum {
	values := make([]DataEntityFromTableEntityDetailsEntityTypeEnum, 0)
	for _, v := range mappingDataEntityFromTableEntityDetailsEntityType {
		values = append(values, v)
	}
	return values
}
