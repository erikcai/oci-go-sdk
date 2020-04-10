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

// CreateDataEntityShapeDetails The data entity summary object.
type CreateDataEntityShapeDetails struct {
	Details DataEntityDetails `mandatory:"false" json:"details"`
}

func (m CreateDataEntityShapeDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *CreateDataEntityShapeDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Details dataentitydetails `json:"details"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.Details.UnmarshalPolymorphicJSON(model.Details.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Details = nn.(DataEntityDetails)
	} else {
		m.Details = nil
	}
	return
}
