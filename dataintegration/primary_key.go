// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Data Integration Service API Specification
//
// Data Integration Service API Specification
//

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/common"
)

// PrimaryKey The primary key object.
type PrimaryKey struct {

	// The type of the key.
	ModelType PrimaryKeyModelTypeEnum `mandatory:"true" json:"modelType"`

	// Object key
	Key *string `mandatory:"false" json:"key"`

	// modelVersion
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value can be edited by the user and it is restricted to 1000 characters
	Name *string `mandatory:"false" json:"name"`

	// attributeRefs
	AttributeRefs []KeyAttribute `mandatory:"false" json:"attributeRefs"`

	// Status of object, can set this to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`
}

func (m PrimaryKey) String() string {
	return common.PointerString(m)
}

// PrimaryKeyModelTypeEnum Enum with underlying type: string
type PrimaryKeyModelTypeEnum string

// Set of constants representing the allowable values for PrimaryKeyModelTypeEnum
const (
	PrimaryKeyModelTypeForeignKey PrimaryKeyModelTypeEnum = "FOREIGN_KEY"
	PrimaryKeyModelTypePrimaryKey PrimaryKeyModelTypeEnum = "PRIMARY_KEY"
	PrimaryKeyModelTypeUniqueKey  PrimaryKeyModelTypeEnum = "UNIQUE_KEY"
)

var mappingPrimaryKeyModelType = map[string]PrimaryKeyModelTypeEnum{
	"FOREIGN_KEY": PrimaryKeyModelTypeForeignKey,
	"PRIMARY_KEY": PrimaryKeyModelTypePrimaryKey,
	"UNIQUE_KEY":  PrimaryKeyModelTypeUniqueKey,
}

// GetPrimaryKeyModelTypeEnumValues Enumerates the set of values for PrimaryKeyModelTypeEnum
func GetPrimaryKeyModelTypeEnumValues() []PrimaryKeyModelTypeEnum {
	values := make([]PrimaryKeyModelTypeEnum, 0)
	for _, v := range mappingPrimaryKeyModelType {
		values = append(values, v)
	}
	return values
}
