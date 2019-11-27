// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Application Migration Service API
//
// API for the Application Migration service. Use this API to migrate applications from Oracle Cloud Infrastructure - Classic to Oracle Cloud Infrastructure.
//

package applicationmigration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateMigrationDetails Update the details and configuration of a migration.
type UpdateMigrationDetails struct {

	// Human-readable name of the migration.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description of the migration.
	Description *string `mandatory:"false" json:"description"`

	DiscoveryDetails DiscoveryDetails `mandatory:"false" json:"discoveryDetails"`

	// Automatically start the migration if the provided configuration passes validation.
	IsAutomaticMigration *bool `mandatory:"false" json:"isAutomaticMigration"`

	// Indicates whether a new service will be provisioned to host the migrated application.
	IsNewServiceRequired *bool `mandatory:"false" json:"isNewServiceRequired"`

	// Configuration required to migrate the application. In addition to the key and value, additional fields are provided to describe type type and purpose of each field. Only the value for each key is required when passing configuration to the CreateMigration operation.
	ServiceConfig map[string]ConfigurationField `mandatory:"false" json:"serviceConfig"`

	// Configuration required to migrate the application. In addition to the key and value, additional fields are provided to describe type type and purpose of each field. Only the value for each key is required when passing configuration to the CreateMigration operation.
	ApplicationConfig map[string]ConfigurationField `mandatory:"false" json:"applicationConfig"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateMigrationDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *UpdateMigrationDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName          *string                           `json:"displayName"`
		Description          *string                           `json:"description"`
		DiscoveryDetails     discoverydetails                  `json:"discoveryDetails"`
		IsAutomaticMigration *bool                             `json:"isAutomaticMigration"`
		IsNewServiceRequired *bool                             `json:"isNewServiceRequired"`
		ServiceConfig        map[string]ConfigurationField     `json:"serviceConfig"`
		ApplicationConfig    map[string]ConfigurationField     `json:"applicationConfig"`
		FreeformTags         map[string]string                 `json:"freeformTags"`
		DefinedTags          map[string]map[string]interface{} `json:"definedTags"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	nn, e = model.DiscoveryDetails.UnmarshalPolymorphicJSON(model.DiscoveryDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.DiscoveryDetails = nn.(DiscoveryDetails)
	} else {
		m.DiscoveryDetails = nil
	}

	m.IsAutomaticMigration = model.IsAutomaticMigration

	m.IsNewServiceRequired = model.IsNewServiceRequired

	m.ServiceConfig = model.ServiceConfig

	m.ApplicationConfig = model.ApplicationConfig

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags
	return
}
