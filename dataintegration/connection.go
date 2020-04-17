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

// Connection The connection object.
type Connection struct {
	Details ConnectionDetails `mandatory:"false" json:"details"`

	Summary *MetadataObjectSummary `mandatory:"false" json:"summary"`

	// keyMappingFromInput
	KeyMap map[string]string `mandatory:"false" json:"keyMap"`
}

func (m Connection) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *Connection) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Details connectiondetails      `json:"details"`
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
		m.Details = nn.(ConnectionDetails)
	} else {
		m.Details = nil
	}

	m.Summary = model.Summary

	m.KeyMap = model.KeyMap
	return
}
