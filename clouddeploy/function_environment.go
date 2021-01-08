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
	"github.com/oracle/oci-go-sdk/v31/common"
)

// FunctionEnvironment Specifies the OKE Function environment.
type FunctionEnvironment struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// Application Identifier
	ApplicationId *string `mandatory:"true" json:"applicationId"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the Function.
	FunctionId *string `mandatory:"true" json:"functionId"`

	// Environment Identifier, can be renamed
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The time the the Environment was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the Environment was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecyleDetails *string `mandatory:"false" json:"lifecyleDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The current state of the Environment.
	LifecycleState EnvironmentLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

//GetId returns Id
func (m FunctionEnvironment) GetId() *string {
	return m.Id
}

//GetDisplayName returns DisplayName
func (m FunctionEnvironment) GetDisplayName() *string {
	return m.DisplayName
}

//GetApplicationId returns ApplicationId
func (m FunctionEnvironment) GetApplicationId() *string {
	return m.ApplicationId
}

//GetCompartmentId returns CompartmentId
func (m FunctionEnvironment) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetTimeCreated returns TimeCreated
func (m FunctionEnvironment) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m FunctionEnvironment) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetLifecycleState returns LifecycleState
func (m FunctionEnvironment) GetLifecycleState() EnvironmentLifecycleStateEnum {
	return m.LifecycleState
}

//GetLifecyleDetails returns LifecyleDetails
func (m FunctionEnvironment) GetLifecyleDetails() *string {
	return m.LifecyleDetails
}

//GetFreeformTags returns FreeformTags
func (m FunctionEnvironment) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m FunctionEnvironment) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetSystemTags returns SystemTags
func (m FunctionEnvironment) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m FunctionEnvironment) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m FunctionEnvironment) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeFunctionEnvironment FunctionEnvironment
	s := struct {
		DiscriminatorParam string `json:"environmentType"`
		MarshalTypeFunctionEnvironment
	}{
		"FUNCTION",
		(MarshalTypeFunctionEnvironment)(m),
	}

	return json.Marshal(&s)
}
