// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// MySQL as a Service API
//
// The API for the MySQL Service
//

package mysql

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Configuration The predefined set of options when deploying a MySQLaaS
// Instance. Assumed by default per MySQLaaS Shape, or by user selection of
// an existing custom configuration.
type Configuration struct {

	// The Oracle Cloud ID (OCID) of the MySQLaaS Configuration.
	Id *string `mandatory:"true" json:"id"`

	// OCID of the Compartment the MySQLaaS Configuration exists in.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name of the associated MySQLaaS Instance Shape.
	ShapeName *string `mandatory:"true" json:"shapeName"`

	// The specific version of MySQLaaS that this Configuration and its Options are relevant to.
	MysqlVersion *string `mandatory:"true" json:"mysqlVersion"`

	// The MySQLaaS Configuration type, DEFAULT or CUSTOM.
	Type ConfigurationTypeEnum `mandatory:"true" json:"type"`

	// The date and time the MySQLaaS Configuration was created, as described by RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the MySQLaaS Configuration was last updated, as described by RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current state of the MySQLaaS Configuration.
	LifecycleState ConfigurationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The current set of option values for the Configuration.
	// These will match exactly the most recent revision of the Configuration.
	Options []Option `mandatory:"true" json:"options"`

	// User-provided data about the MySQLaaS Instance.
	Description *string `mandatory:"false" json:"description"`

	// The display name of the MySQLaaS Configuration.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// If the Configuration is currently being applied
	// to its associated instances, this field is the URL of the
	// WorkRequest representing the progress of that operation.
	ApplicationProgress *string `mandatory:"false" json:"applicationProgress"`

	// Simple key-value pair applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m Configuration) String() string {
	return common.PointerString(m)
}

// ConfigurationTypeEnum Enum with underlying type: string
type ConfigurationTypeEnum string

// Set of constants representing the allowable values for ConfigurationTypeEnum
const (
	ConfigurationTypeDefault ConfigurationTypeEnum = "DEFAULT"
	ConfigurationTypeCustom  ConfigurationTypeEnum = "CUSTOM"
)

var mappingConfigurationType = map[string]ConfigurationTypeEnum{
	"DEFAULT": ConfigurationTypeDefault,
	"CUSTOM":  ConfigurationTypeCustom,
}

// GetConfigurationTypeEnumValues Enumerates the set of values for ConfigurationTypeEnum
func GetConfigurationTypeEnumValues() []ConfigurationTypeEnum {
	values := make([]ConfigurationTypeEnum, 0)
	for _, v := range mappingConfigurationType {
		values = append(values, v)
	}
	return values
}
