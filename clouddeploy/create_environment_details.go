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
	"github.com/oracle/oci-go-sdk/v32/common"
)

// CreateEnvironmentDetails The information about new Environment.
type CreateEnvironmentDetails interface {

	// Environment Identifier
	GetDisplayName() *string

	// Application Identifier
	GetApplicationId() *string

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}
}

type createenvironmentdetails struct {
	JsonData        []byte
	DisplayName     *string                           `mandatory:"true" json:"displayName"`
	ApplicationId   *string                           `mandatory:"true" json:"applicationId"`
	FreeformTags    map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags     map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	EnvironmentType string                            `json:"environmentType"`
}

// UnmarshalJSON unmarshals json
func (m *createenvironmentdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreateenvironmentdetails createenvironmentdetails
	s := struct {
		Model Unmarshalercreateenvironmentdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DisplayName = s.Model.DisplayName
	m.ApplicationId = s.Model.ApplicationId
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.EnvironmentType = s.Model.EnvironmentType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createenvironmentdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.EnvironmentType {
	case "FUNCTION":
		mm := CreateFunctionEnvironmentDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "LOAD_BALANCE_LISTENER":
		mm := CreateLoadBalancerListenerEnvironmentDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "COMPUTE_INSTANCE_GROUP":
		mm := CreateComputeInstanceGroupEnvironmentDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OKE_CLUSTER":
		mm := CreateOkeClusterEnvironmentDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetDisplayName returns DisplayName
func (m createenvironmentdetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetApplicationId returns ApplicationId
func (m createenvironmentdetails) GetApplicationId() *string {
	return m.ApplicationId
}

//GetFreeformTags returns FreeformTags
func (m createenvironmentdetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m createenvironmentdetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m createenvironmentdetails) String() string {
	return common.PointerString(m)
}
