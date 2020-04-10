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

// DataEntitySummary The data entity summary object.
type DataEntitySummary struct {
	Details DataEntityDetails `mandatory:"false" json:"details"`

	Summary *MetadataObjectSummary `mandatory:"false" json:"summary"`
}

func (m DataEntitySummary) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *DataEntitySummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Details dataentitydetails      `json:"details"`
		Summary *MetadataObjectSummary `json:"summary"`
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

	m.Summary = model.Summary
	return
}
