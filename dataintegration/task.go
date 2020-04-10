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

// Task The task type contains the audit summary information and the definition of the task.
type Task struct {
	Details TaskDetails `mandatory:"false" json:"details"`

	Summary *MetadataObjectSummary `mandatory:"false" json:"summary"`
}

func (m Task) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *Task) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Details taskdetails            `json:"details"`
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
		m.Details = nn.(TaskDetails)
	} else {
		m.Details = nil
	}

	m.Summary = model.Summary
	return
}
