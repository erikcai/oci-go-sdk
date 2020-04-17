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

// OutputLink auto generated description
type OutputLink struct {

	// Object key
	Key *string `mandatory:"false" json:"key"`

	// modelVersion
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Status of object, can set this to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Descriptive text for the object.
	Description *string `mandatory:"false" json:"description"`

	// Key of FlowPort reference
	Port *string `mandatory:"false" json:"port"`

	// The links from this output link to connect to other links in flow.
	ToLinks []string `mandatory:"false" json:"toLinks"`
}

//GetKey returns Key
func (m OutputLink) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m OutputLink) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m OutputLink) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetObjectStatus returns ObjectStatus
func (m OutputLink) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetDescription returns Description
func (m OutputLink) GetDescription() *string {
	return m.Description
}

//GetPort returns Port
func (m OutputLink) GetPort() *string {
	return m.Port
}

func (m OutputLink) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m OutputLink) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOutputLink OutputLink
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeOutputLink
	}{
		"OUTPUT_LINK",
		(MarshalTypeOutputLink)(m),
	}

	return json.Marshal(&s)
}
