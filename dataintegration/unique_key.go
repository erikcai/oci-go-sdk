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

// UniqueKey The unqique key object.
type UniqueKey struct {

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

func (m UniqueKey) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m UniqueKey) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUniqueKey UniqueKey
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeUniqueKey
	}{
		"UNIQUE_KEY",
		(MarshalTypeUniqueKey)(m),
	}

	return json.Marshal(&s)
}
