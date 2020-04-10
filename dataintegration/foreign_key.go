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

// ForeignKey The foreign key object.
type ForeignKey struct {

	// Object key
	Key *string `mandatory:"false" json:"key"`

	// modelVersion
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value can be edited by the user and it is restricted to 1000 characters
	Name *string `mandatory:"false" json:"name"`

	// attributeRefs
	AttributeRefs []KeyAttribute `mandatory:"false" json:"attributeRefs"`

	// updateRule
	UpdateRule *int `mandatory:"false" json:"updateRule"`

	// deleteRule
	DeleteRule *int `mandatory:"false" json:"deleteRule"`

	ReferenceUniqueKey *UniqueKey `mandatory:"false" json:"referenceUniqueKey"`

	// Status of object, can set this to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`
}

func (m ForeignKey) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ForeignKey) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeForeignKey ForeignKey
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeForeignKey
	}{
		"FOREIGN_KEY",
		(MarshalTypeForeignKey)(m),
	}

	return json.Marshal(&s)
}
