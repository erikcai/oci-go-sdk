// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// CreateInstanceConfigurationDetails An instance configuration that can be used to launch
type CreateInstanceConfigurationDetails struct {

	// The OCID of the compartment containing the instance configuration.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name for the instance configuration
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see
	// Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	InstanceDetails InstanceConfigurationInstanceDetails `mandatory:"false" json:"instanceDetails"`
}

func (m CreateInstanceConfigurationDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *CreateInstanceConfigurationDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DefinedTags     map[string]map[string]interface{}    `json:"definedTags"`
		DisplayName     *string                              `json:"displayName"`
		FreeformTags    map[string]string                    `json:"freeformTags"`
		InstanceDetails instanceconfigurationinstancedetails `json:"instanceDetails"`
		CompartmentId   *string                              `json:"compartmentId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	m.DefinedTags = model.DefinedTags
	m.DisplayName = model.DisplayName
	m.FreeformTags = model.FreeformTags
	nn, e := model.InstanceDetails.UnmarshalPolymorphicJSON(model.InstanceDetails.JsonData)
	if e != nil {
		return
	}
	m.InstanceDetails = nn.(InstanceConfigurationInstanceDetails)
	m.CompartmentId = model.CompartmentId
	return
}
