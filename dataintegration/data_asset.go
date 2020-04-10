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

// DataAsset The data asset type.
type DataAsset struct {
	Details DataAssetDetails `mandatory:"false" json:"details"`

	Summary *MetadataObjectSummary `mandatory:"false" json:"summary"`

	// keyMappingFromInput
	KeyMap map[string]string `mandatory:"false" json:"keyMap"`
}

func (m DataAsset) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *DataAsset) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Details dataassetdetails       `json:"details"`
		Summary *MetadataObjectSummary `json:"summary"`
		KeyMap  map[string]string      `json:"keyMap"`
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
		m.Details = nn.(DataAssetDetails)
	} else {
		m.Details = nil
	}

	m.Summary = model.Summary

	m.KeyMap = model.KeyMap
	return
}
