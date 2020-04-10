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

// DataEntityFromViewEntityDetails The view entity data entity details.
type DataEntityFromViewEntityDetails struct {

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

	// The external key for the object
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
	EntityType DataEntityFromViewEntityDetailsEntityTypeEnum `mandatory:"false" json:"entityType,omitempty"`
}

func (m DataEntityFromViewEntityDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m DataEntityFromViewEntityDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDataEntityFromViewEntityDetails DataEntityFromViewEntityDetails
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeDataEntityFromViewEntityDetails
	}{
		"VIEW_ENTITY",
		(MarshalTypeDataEntityFromViewEntityDetails)(m),
	}

	return json.Marshal(&s)
}

// DataEntityFromViewEntityDetailsEntityTypeEnum Enum with underlying type: string
type DataEntityFromViewEntityDetailsEntityTypeEnum string

// Set of constants representing the allowable values for DataEntityFromViewEntityDetailsEntityTypeEnum
const (
	DataEntityFromViewEntityDetailsEntityTypeTable  DataEntityFromViewEntityDetailsEntityTypeEnum = "TABLE"
	DataEntityFromViewEntityDetailsEntityTypeView   DataEntityFromViewEntityDetailsEntityTypeEnum = "VIEW"
	DataEntityFromViewEntityDetailsEntityTypeFile   DataEntityFromViewEntityDetailsEntityTypeEnum = "FILE"
	DataEntityFromViewEntityDetailsEntityTypeQueue  DataEntityFromViewEntityDetailsEntityTypeEnum = "QUEUE"
	DataEntityFromViewEntityDetailsEntityTypeStream DataEntityFromViewEntityDetailsEntityTypeEnum = "STREAM"
	DataEntityFromViewEntityDetailsEntityTypeOther  DataEntityFromViewEntityDetailsEntityTypeEnum = "OTHER"
)

var mappingDataEntityFromViewEntityDetailsEntityType = map[string]DataEntityFromViewEntityDetailsEntityTypeEnum{
	"TABLE":  DataEntityFromViewEntityDetailsEntityTypeTable,
	"VIEW":   DataEntityFromViewEntityDetailsEntityTypeView,
	"FILE":   DataEntityFromViewEntityDetailsEntityTypeFile,
	"QUEUE":  DataEntityFromViewEntityDetailsEntityTypeQueue,
	"STREAM": DataEntityFromViewEntityDetailsEntityTypeStream,
	"OTHER":  DataEntityFromViewEntityDetailsEntityTypeOther,
}

// GetDataEntityFromViewEntityDetailsEntityTypeEnumValues Enumerates the set of values for DataEntityFromViewEntityDetailsEntityTypeEnum
func GetDataEntityFromViewEntityDetailsEntityTypeEnumValues() []DataEntityFromViewEntityDetailsEntityTypeEnum {
	values := make([]DataEntityFromViewEntityDetailsEntityTypeEnum, 0)
	for _, v := range mappingDataEntityFromViewEntityDetailsEntityType {
		values = append(values, v)
	}
	return values
}
