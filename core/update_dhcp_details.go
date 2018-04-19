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

// UpdateDhcpDetails The representation of UpdateDhcpDetails
type UpdateDhcpDetails struct {

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "foo-value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	Options []DhcpOption `mandatory:"false" json:"options"`
}

func (m UpdateDhcpDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *UpdateDhcpDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DefinedTags  map[string]map[string]interface{} `json:"definedTags"`
		DisplayName  *string                           `json:"displayName"`
		FreeformTags map[string]string                 `json:"freeformTags"`
		Options      []dhcpoption                      `json:"options"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	m.DefinedTags = model.DefinedTags
	m.DisplayName = model.DisplayName
	m.FreeformTags = model.FreeformTags
	m.Options = make([]DhcpOption, len(model.Options))
	for i, n := range model.Options {
		nn, err := n.UnmarshalPolymorphicJSON(n.JsonData)
		if err != nil {
			return err
		}
		m.Options[i] = nn.(DhcpOption)
	}
	return
}
