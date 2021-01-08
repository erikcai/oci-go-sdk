// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataintegration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v31/common"
)

// EnrichedEntity This is used to specify runtime parameters for data entities such as files that need both the data entity and the format.
type EnrichedEntity struct {
	Entity DataEntity `mandatory:"false" json:"entity"`

	DataFormat *DataFormat `mandatory:"false" json:"dataFormat"`
}

func (m EnrichedEntity) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *EnrichedEntity) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Entity     dataentity  `json:"entity"`
		DataFormat *DataFormat `json:"dataFormat"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.Entity.UnmarshalPolymorphicJSON(model.Entity.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Entity = nn.(DataEntity)
	} else {
		m.Entity = nil
	}

	m.DataFormat = model.DataFormat

	return
}
