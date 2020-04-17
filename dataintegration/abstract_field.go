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

// AbstractField The type representing the abstract field concept.
type AbstractField struct {

	// Object key
	Key *string `mandatory:"false" json:"key"`

	// modelVersion
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	ConfigValues *ConfigValues `mandatory:"false" json:"configValues"`

	// Status of object, can set this to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value can be edited by the user and it is restricted to 1000 characters
	Name *string `mandatory:"false" json:"name"`

	// Descriptive text for the object.
	Description *string `mandatory:"false" json:"description"`
}

//GetKey returns Key
func (m AbstractField) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m AbstractField) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m AbstractField) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetConfigValues returns ConfigValues
func (m AbstractField) GetConfigValues() *ConfigValues {
	return m.ConfigValues
}

//GetObjectStatus returns ObjectStatus
func (m AbstractField) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetName returns Name
func (m AbstractField) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m AbstractField) GetDescription() *string {
	return m.Description
}

func (m AbstractField) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m AbstractField) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAbstractField AbstractField
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeAbstractField
	}{
		"FIELD",
		(MarshalTypeAbstractField)(m),
	}

	return json.Marshal(&s)
}
