// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Oracle Resource Manager
//
// Oracle Resource Manager API
//

package resourcemanager

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// CreateStackDetails The representation of CreateStackDetails
type CreateStackDetails struct {
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	ConfigSource CreateConfigSourceDetails `mandatory:"true" json:"configSource"`

	DisplayName *string `mandatory:"false" json:"displayName"`

	Description *string `mandatory:"false" json:"description"`

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

func (m CreateStackDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *CreateStackDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName   *string                           `json:"displayName"`
		Description   *string                           `json:"description"`
		Variables     map[string]string                 `json:"variables"`
		FreeformTags  map[string]string                 `json:"freeformTags"`
		DefinedTags   map[string]map[string]interface{} `json:"definedTags"`
		CompartmentId *string                           `json:"compartmentId"`
		ConfigSource  createconfigsourcedetails         `json:"configSource"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	m.DisplayName = model.DisplayName
	m.Description = model.Description
	m.Variables = model.Variables
	m.FreeformTags = model.FreeformTags
	m.DefinedTags = model.DefinedTags
	m.CompartmentId = model.CompartmentId
	nn, e := model.ConfigSource.UnmarshalPolymorphicJSON(model.ConfigSource.JsonData)
	if e != nil {
		return
	}
	m.ConfigSource = nn.(CreateConfigSourceDetails)
	return
}
