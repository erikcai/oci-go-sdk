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

// ModelSelect auto generated description
type ModelSelect struct {

	// distinct
	IsDistinct *bool `mandatory:"false" json:"isDistinct"`

	// selectCols
	SelectCols []ShapeField `mandatory:"false" json:"selectCols"`
}

func (m ModelSelect) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ModelSelect) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeModelSelect ModelSelect
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeModelSelect
	}{
		"SELECT",
		(MarshalTypeModelSelect)(m),
	}

	return json.Marshal(&s)
}
