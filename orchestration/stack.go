// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Orchestration
//
// Orchestration Maestro API
//

package orchestration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// Stack The representation of Stack
type Stack struct {
	Id *string `mandatory:"false" json:"id"`

	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	DisplayName *string `mandatory:"false" json:"displayName"`

	Description *string `mandatory:"false" json:"description"`

	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	LifecycleState StackLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	ConfigSource ConfigSource `mandatory:"false" json:"configSource"`

	Variables map[string]string `mandatory:"false" json:"variables"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m Stack) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *Stack) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Id             *string                           `json:"id"`
		CompartmentId  *string                           `json:"compartmentId"`
		DisplayName    *string                           `json:"displayName"`
		Description    *string                           `json:"description"`
		TimeCreated    *common.SDKTime                   `json:"timeCreated"`
		LifecycleState StackLifecycleStateEnum           `json:"lifecycleState"`
		ConfigSource   configsource                      `json:"configSource"`
		Variables      map[string]string                 `json:"variables"`
		FreeformTags   map[string]string                 `json:"freeformTags"`
		DefinedTags    map[string]map[string]interface{} `json:"definedTags"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	m.Id = model.Id
	m.CompartmentId = model.CompartmentId
	m.DisplayName = model.DisplayName
	m.Description = model.Description
	m.TimeCreated = model.TimeCreated
	m.LifecycleState = model.LifecycleState
	nn, e := model.ConfigSource.UnmarshalPolymorphicJSON(model.ConfigSource.JsonData)
	if e != nil {
		return
	}
	m.ConfigSource = nn.(ConfigSource)
	m.Variables = model.Variables
	m.FreeformTags = model.FreeformTags
	m.DefinedTags = model.DefinedTags
	return
}

// StackLifecycleStateEnum Enum with underlying type: string
type StackLifecycleStateEnum string

// Set of constants representing the allowable values for StackLifecycleState
const (
	StackLifecycleStateCreating StackLifecycleStateEnum = "CREATING"
	StackLifecycleStateActive   StackLifecycleStateEnum = "ACTIVE"
	StackLifecycleStateDeleting StackLifecycleStateEnum = "DELETING"
	StackLifecycleStateDeleted  StackLifecycleStateEnum = "DELETED"
)

var mappingStackLifecycleState = map[string]StackLifecycleStateEnum{
	"CREATING": StackLifecycleStateCreating,
	"ACTIVE":   StackLifecycleStateActive,
	"DELETING": StackLifecycleStateDeleting,
	"DELETED":  StackLifecycleStateDeleted,
}

// GetStackLifecycleStateEnumValues Enumerates the set of values for StackLifecycleState
func GetStackLifecycleStateEnumValues() []StackLifecycleStateEnum {
	values := make([]StackLifecycleStateEnum, 0)
	for _, v := range mappingStackLifecycleState {
		values = append(values, v)
	}
	return values
}
