// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DeploymentService API
//
// A description of the DeploymentService API
//

package clouddeploy

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v33/common"
)

// Environment Description of Environment.
type Environment interface {

	// Unique identifier that is immutable on creation
	GetId() *string

	// Application Identifier
	GetApplicationId() *string

	// Compartment Identifier
	GetCompartmentId() *string

	// Environment Identifier, can be renamed
	GetDisplayName() *string

	// The time the the Environment was created. An RFC3339 formatted datetime string
	GetTimeCreated() *common.SDKTime

	// The time the Environment was updated. An RFC3339 formatted datetime string
	GetTimeUpdated() *common.SDKTime

	// The current state of the Environment.
	GetLifecycleState() EnvironmentLifecycleStateEnum

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	GetLifecyleDetails() *string

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	GetSystemTags() map[string]map[string]interface{}
}

type environment struct {
	JsonData        []byte
	Id              *string                           `mandatory:"true" json:"id"`
	ApplicationId   *string                           `mandatory:"true" json:"applicationId"`
	CompartmentId   *string                           `mandatory:"true" json:"compartmentId"`
	DisplayName     *string                           `mandatory:"false" json:"displayName"`
	TimeCreated     *common.SDKTime                   `mandatory:"false" json:"timeCreated"`
	TimeUpdated     *common.SDKTime                   `mandatory:"false" json:"timeUpdated"`
	LifecycleState  EnvironmentLifecycleStateEnum     `mandatory:"false" json:"lifecycleState,omitempty"`
	LifecyleDetails *string                           `mandatory:"false" json:"lifecyleDetails"`
	FreeformTags    map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags     map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	SystemTags      map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	EnvironmentType string                            `json:"environmentType"`
}

// UnmarshalJSON unmarshals json
func (m *environment) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerenvironment environment
	s := struct {
		Model Unmarshalerenvironment
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.ApplicationId = s.Model.ApplicationId
	m.CompartmentId = s.Model.CompartmentId
	m.DisplayName = s.Model.DisplayName
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
	m.LifecycleState = s.Model.LifecycleState
	m.LifecyleDetails = s.Model.LifecyleDetails
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.EnvironmentType = s.Model.EnvironmentType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *environment) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.EnvironmentType {
	case "LOAD_BALANCE_LISTENER":
		mm := LoadBalancerListenerEnvironment{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FUNCTION":
		mm := FunctionEnvironment{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OKE_CLUSTER":
		mm := OkeClusterEnvironment{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "COMPUTE_INSTANCE_GROUP":
		mm := ComputeInstanceGroupEnvironment{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetId returns Id
func (m environment) GetId() *string {
	return m.Id
}

//GetApplicationId returns ApplicationId
func (m environment) GetApplicationId() *string {
	return m.ApplicationId
}

//GetCompartmentId returns CompartmentId
func (m environment) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetDisplayName returns DisplayName
func (m environment) GetDisplayName() *string {
	return m.DisplayName
}

//GetTimeCreated returns TimeCreated
func (m environment) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m environment) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetLifecycleState returns LifecycleState
func (m environment) GetLifecycleState() EnvironmentLifecycleStateEnum {
	return m.LifecycleState
}

//GetLifecyleDetails returns LifecyleDetails
func (m environment) GetLifecyleDetails() *string {
	return m.LifecyleDetails
}

//GetFreeformTags returns FreeformTags
func (m environment) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m environment) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetSystemTags returns SystemTags
func (m environment) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m environment) String() string {
	return common.PointerString(m)
}

// EnvironmentLifecycleStateEnum Enum with underlying type: string
type EnvironmentLifecycleStateEnum string

// Set of constants representing the allowable values for EnvironmentLifecycleStateEnum
const (
	EnvironmentLifecycleStateCreating EnvironmentLifecycleStateEnum = "CREATING"
	EnvironmentLifecycleStateUpdating EnvironmentLifecycleStateEnum = "UPDATING"
	EnvironmentLifecycleStateActive   EnvironmentLifecycleStateEnum = "ACTIVE"
	EnvironmentLifecycleStateDeleting EnvironmentLifecycleStateEnum = "DELETING"
	EnvironmentLifecycleStateDeleted  EnvironmentLifecycleStateEnum = "DELETED"
	EnvironmentLifecycleStateFailed   EnvironmentLifecycleStateEnum = "FAILED"
)

var mappingEnvironmentLifecycleState = map[string]EnvironmentLifecycleStateEnum{
	"CREATING": EnvironmentLifecycleStateCreating,
	"UPDATING": EnvironmentLifecycleStateUpdating,
	"ACTIVE":   EnvironmentLifecycleStateActive,
	"DELETING": EnvironmentLifecycleStateDeleting,
	"DELETED":  EnvironmentLifecycleStateDeleted,
	"FAILED":   EnvironmentLifecycleStateFailed,
}

// GetEnvironmentLifecycleStateEnumValues Enumerates the set of values for EnvironmentLifecycleStateEnum
func GetEnvironmentLifecycleStateEnumValues() []EnvironmentLifecycleStateEnum {
	values := make([]EnvironmentLifecycleStateEnum, 0)
	for _, v := range mappingEnvironmentLifecycleState {
		values = append(values, v)
	}
	return values
}

// EnvironmentEnvironmentTypeEnum Enum with underlying type: string
type EnvironmentEnvironmentTypeEnum string

// Set of constants representing the allowable values for EnvironmentEnvironmentTypeEnum
const (
	EnvironmentEnvironmentTypeOkeCluster           EnvironmentEnvironmentTypeEnum = "OKE_CLUSTER"
	EnvironmentEnvironmentTypeComputeInstanceGroup EnvironmentEnvironmentTypeEnum = "COMPUTE_INSTANCE_GROUP"
	EnvironmentEnvironmentTypeFunction             EnvironmentEnvironmentTypeEnum = "FUNCTION"
	EnvironmentEnvironmentTypeLoadBalanceListener  EnvironmentEnvironmentTypeEnum = "LOAD_BALANCE_LISTENER"
)

var mappingEnvironmentEnvironmentType = map[string]EnvironmentEnvironmentTypeEnum{
	"OKE_CLUSTER":            EnvironmentEnvironmentTypeOkeCluster,
	"COMPUTE_INSTANCE_GROUP": EnvironmentEnvironmentTypeComputeInstanceGroup,
	"FUNCTION":               EnvironmentEnvironmentTypeFunction,
	"LOAD_BALANCE_LISTENER":  EnvironmentEnvironmentTypeLoadBalanceListener,
}

// GetEnvironmentEnvironmentTypeEnumValues Enumerates the set of values for EnvironmentEnvironmentTypeEnum
func GetEnvironmentEnvironmentTypeEnumValues() []EnvironmentEnvironmentTypeEnum {
	values := make([]EnvironmentEnvironmentTypeEnum, 0)
	for _, v := range mappingEnvironmentEnvironmentType {
		values = append(values, v)
	}
	return values
}
