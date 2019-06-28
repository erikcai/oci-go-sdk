// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// CIMS API
//
// Cloud Incident Management Service
//

package cims

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// Item The representation of Item
type Item interface {
	GetId() *string

	GetName() *string
}

type item struct {
	JsonData []byte
	Id       *string `mandatory:"false" json:"id"`
	Name     *string `mandatory:"false" json:"name"`
	Type     string  `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *item) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshaleritem item
	s := struct {
		Model Unmarshaleritem
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.Name = s.Model.Name
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *item) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "limit":
		mm := LimitItem{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetId returns Id
func (m item) GetId() *string {
	return m.Id
}

//GetName returns Name
func (m item) GetName() *string {
	return m.Name
}

func (m item) String() string {
	return common.PointerString(m)
}
