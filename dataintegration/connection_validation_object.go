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

// ConnectionValidationObject Connection validation object representing the Results of the validation of the connection, along with an key
type ConnectionValidationObject struct {
	DataAsset DataAssetDetails `mandatory:"false" json:"dataAsset"`

	Connection ConnectionDetails `mandatory:"false" json:"connection"`

	ValidationMessage *Message `mandatory:"false" json:"validationMessage"`

	// Objects will use a 36 character key as unique ID. It is system generated and cannot be edited by user
	Key *string `mandatory:"false" json:"key"`

	// The type of the object.
	ModelType *string `mandatory:"false" json:"modelType"`

	// modelVersion
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value can be edited by the user and it is restricted to 1000 characters
	Name *string `mandatory:"false" json:"name"`

	// Descriptive text for the object.
	Description *string `mandatory:"false" json:"description"`

	// The version of the object, to track changes in the object instance
	ObjectVersion *int `mandatory:"false" json:"objectVersion"`

	// Status of object, can set this to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Value can only contain upper case letters, underscore and numbers. It should begin with upper case letter or underscore. The value can be edited by the user.
	Identifier *string `mandatory:"false" json:"identifier"`
}

func (m ConnectionValidationObject) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *ConnectionValidationObject) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DataAsset         dataassetdetails  `json:"dataAsset"`
		Connection        connectiondetails `json:"connection"`
		ValidationMessage *Message          `json:"validationMessage"`
		Key               *string           `json:"key"`
		ModelType         *string           `json:"modelType"`
		ModelVersion      *string           `json:"modelVersion"`
		ParentRef         *ParentReference  `json:"parentRef"`
		Name              *string           `json:"name"`
		Description       *string           `json:"description"`
		ObjectVersion     *int              `json:"objectVersion"`
		ObjectStatus      *int              `json:"objectStatus"`
		Identifier        *string           `json:"identifier"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.DataAsset.UnmarshalPolymorphicJSON(model.DataAsset.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.DataAsset = nn.(DataAssetDetails)
	} else {
		m.DataAsset = nil
	}

	nn, e = model.Connection.UnmarshalPolymorphicJSON(model.Connection.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Connection = nn.(ConnectionDetails)
	} else {
		m.Connection = nil
	}

	m.ValidationMessage = model.ValidationMessage

	m.Key = model.Key

	m.ModelType = model.ModelType

	m.ModelVersion = model.ModelVersion

	m.ParentRef = model.ParentRef

	m.Name = model.Name

	m.Description = model.Description

	m.ObjectVersion = model.ObjectVersion

	m.ObjectStatus = model.ObjectStatus

	m.Identifier = model.Identifier
	return
}
