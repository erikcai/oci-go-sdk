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

// ConfigurationSummary The general details of a MySQLaaS Configuration such as its id, displayName, type, and shape association.
type ConfigurationSummary struct {

	// The Oracle Cloud ID (OCID) of the MySQLaaS Configuration.
	Id *string `mandatory:"true" json:"id"`

	// OCID of the Compartment the MySQLaaS Configuration exists in.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name of the associated MySQLaaS Instance Shape.
	ShapeName *string `mandatory:"true" json:"shapeName"`

	// The MySQLaaS Configuration type, DEFAULT or CUSTOM
	Type ConfigurationSummaryTypeEnum `mandatory:"true" json:"type"`

	// The current state of the MySQLaaS Configuration.
	LifecycleState ConfigurationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// User-provided data about the MySQLaaS Instance.
	Description *string `mandatory:"false" json:"description"`

	// The display name of the MySQLaaS Configuration.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The version of MySQLaaS that this Configuration and its Options are relevant to.
	MysqlVersion *string `mandatory:"false" json:"mysqlVersion"`

	// If the Configuration is currently being applied
	// to its associated instances,  this field will be the URL for the
	// WorkRequest representing the progress of that application operation.
	ApplicationProgress *string `mandatory:"false" json:"applicationProgress"`

	// The date and time the MySQLaaS Configuration was created, as described by RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the MySQLaaS Configuration was last updated, as described by RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Simple key-value pair applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m ConfigurationSummary) String() string {
	return common.PointerString(m)
}

// ConfigurationSummaryTypeEnum Enum with underlying type: string
type ConfigurationSummaryTypeEnum string

// Set of constants representing the allowable values for ConfigurationSummaryTypeEnum
const (
	ConfigurationSummaryTypeDefault ConfigurationSummaryTypeEnum = "DEFAULT"
	ConfigurationSummaryTypeCustom  ConfigurationSummaryTypeEnum = "CUSTOM"
)

var mappingConfigurationSummaryType = map[string]ConfigurationSummaryTypeEnum{
	"DEFAULT": ConfigurationSummaryTypeDefault,
	"CUSTOM":  ConfigurationSummaryTypeCustom,
}

// GetConfigurationSummaryTypeEnumValues Enumerates the set of values for ConfigurationSummaryTypeEnum
func GetConfigurationSummaryTypeEnumValues() []ConfigurationSummaryTypeEnum {
	values := make([]ConfigurationSummaryTypeEnum, 0)
	for _, v := range mappingConfigurationSummaryType {
		values = append(values, v)
	}
	return values
}
