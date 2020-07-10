// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Connectors API
//
// A description of the Connectors API
//

package sch

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// ServiceConnector Description of ServiceConnector.
type ServiceConnector struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// The ServiceConnector name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The time the the ServiceConnector was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the ServiceConnector was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current state of the ServiceConnector.
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The description of the resource.
	Description *string `mandatory:"false" json:"description"`

	// A message describing the current state in more detail. For example, can be used to provide actionable
	// information for a resource in failed state.
	LifecyleDetails *string `mandatory:"false" json:"lifecyleDetails"`

	Source SourceDetails `mandatory:"false" json:"source"`

	// The list of the tasks.
	Tasks []TaskDetails `mandatory:"false" json:"tasks"`

	Target TargetDetails `mandatory:"false" json:"target"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m ServiceConnector) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *ServiceConnector) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description     *string                           `json:"description"`
		LifecyleDetails *string                           `json:"lifecyleDetails"`
		Source          sourcedetails                     `json:"source"`
		Tasks           []taskdetails                     `json:"tasks"`
		Target          targetdetails                     `json:"target"`
		FreeformTags    map[string]string                 `json:"freeformTags"`
		DefinedTags     map[string]map[string]interface{} `json:"definedTags"`
		Id              *string                           `json:"id"`
		DisplayName     *string                           `json:"displayName"`
		CompartmentId   *string                           `json:"compartmentId"`
		TimeCreated     *common.SDKTime                   `json:"timeCreated"`
		TimeUpdated     *common.SDKTime                   `json:"timeUpdated"`
		LifecycleState  LifecycleStateEnum                `json:"lifecycleState"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.LifecyleDetails = model.LifecyleDetails

	nn, e = model.Source.UnmarshalPolymorphicJSON(model.Source.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Source = nn.(SourceDetails)
	} else {
		m.Source = nil
	}

	m.Tasks = make([]TaskDetails, len(model.Tasks))
	for i, n := range model.Tasks {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Tasks[i] = nn.(TaskDetails)
		} else {
			m.Tasks[i] = nil
		}
	}

	nn, e = model.Target.UnmarshalPolymorphicJSON(model.Target.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Target = nn.(TargetDetails)
	} else {
		m.Target = nil
	}

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.Id = model.Id

	m.DisplayName = model.DisplayName

	m.CompartmentId = model.CompartmentId

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleState = model.LifecycleState

	return
}
