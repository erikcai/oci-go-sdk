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

// ConnectionValidationDetails Result of the validation of the connection
type ConnectionValidationDetails struct {
	DataAsset DataAssetDetails `mandatory:"false" json:"dataAsset"`

	Connection ConnectionDetails `mandatory:"false" json:"connection"`
}

func (m ConnectionValidationDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *ConnectionValidationDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DataAsset  dataassetdetails  `json:"dataAsset"`
		Connection connectiondetails `json:"connection"`
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
	return
}
