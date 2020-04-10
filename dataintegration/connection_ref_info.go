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

// ConnectionRefInfo auto generated description
type ConnectionRefInfo struct {
	Connection ConnectionDetails `mandatory:"false" json:"connection"`

	DataAsset DataAssetDetails `mandatory:"false" json:"dataAsset"`
}

func (m ConnectionRefInfo) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *ConnectionRefInfo) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Connection connectiondetails `json:"connection"`
		DataAsset  dataassetdetails  `json:"dataAsset"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.Connection.UnmarshalPolymorphicJSON(model.Connection.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Connection = nn.(ConnectionDetails)
	} else {
		m.Connection = nil
	}

	nn, e = model.DataAsset.UnmarshalPolymorphicJSON(model.DataAsset.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.DataAsset = nn.(DataAssetDetails)
	} else {
		m.DataAsset = nil
	}
	return
}
