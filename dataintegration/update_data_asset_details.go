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

// UpdateDataAssetDetails Properties used in data asset update operations.
type UpdateDataAssetDetails struct {
	Details DataAssetDetails `mandatory:"false" json:"details"`

	RegistryInfo *RegistryInfo `mandatory:"false" json:"registryInfo"`
}

func (m UpdateDataAssetDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *UpdateDataAssetDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Details      dataassetdetails `json:"details"`
		RegistryInfo *RegistryInfo    `json:"registryInfo"`
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

	m.RegistryInfo = model.RegistryInfo
	return
}
