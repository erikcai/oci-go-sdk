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

// OptionMetadata The representation of OptionMetadata
type OptionMetadata struct {

	// The option's name.
	Name *string `mandatory:"true" json:"name"`

	// The data-type describing the range of values permissible for the option.
	Type OptionMetadataTypeEnum `mandatory:"true" json:"type"`

	// Whether the option can be set by the user.
	IsConfigurable *bool `mandatory:"true" json:"isConfigurable"`

	// If true, the option is runtime settable and takes immediate effect, and can be used with `SET PERSIST`.
	// If false, the option will only take effect on the next restart, and `SET PERSIST_ONLY` should be used.
	IsDynamic *bool `mandatory:"true" json:"isDynamic"`

	// The option's default value.
	DefaultValue *string `mandatory:"false" json:"defaultValue"`

	// Minimum value for INTEGER and NUMERIC, or length for STRING options.
	Min *float32 `mandatory:"false" json:"min"`

	// Maximum value for INTEGER and NUMERIC, or length for STRING options.
	Max *float32 `mandatory:"false" json:"max"`

	// In cases where an option has a defined/enumerated set of
	// values, a description of the possible values.
	Values []OptionMetadataValue `mandatory:"false" json:"values"`

	// Additional information on how to configure the option
	Description *string `mandatory:"false" json:"description"`
}

func (m OptionMetadata) String() string {
	return common.PointerString(m)
}

// OptionMetadataTypeEnum Enum with underlying type: string
type OptionMetadataTypeEnum string

// Set of constants representing the allowable values for OptionMetadataTypeEnum
const (
	OptionMetadataTypeBoolean   OptionMetadataTypeEnum = "BOOLEAN"
	OptionMetadataTypeString    OptionMetadataTypeEnum = "STRING"
	OptionMetadataTypeInteger   OptionMetadataTypeEnum = "INTEGER"
	OptionMetadataTypeNumeric   OptionMetadataTypeEnum = "NUMERIC"
	OptionMetadataTypeEnumvalue OptionMetadataTypeEnum = "ENUM"
	OptionMetadataTypeSet       OptionMetadataTypeEnum = "SET"
	OptionMetadataTypeFlagset   OptionMetadataTypeEnum = "FLAGSET"
)

var mappingOptionMetadataType = map[string]OptionMetadataTypeEnum{
	"BOOLEAN": OptionMetadataTypeBoolean,
	"STRING":  OptionMetadataTypeString,
	"INTEGER": OptionMetadataTypeInteger,
	"NUMERIC": OptionMetadataTypeNumeric,
	"ENUM":    OptionMetadataTypeEnumvalue,
	"SET":     OptionMetadataTypeSet,
	"FLAGSET": OptionMetadataTypeFlagset,
}

// GetOptionMetadataTypeEnumValues Enumerates the set of values for OptionMetadataTypeEnum
func GetOptionMetadataTypeEnumValues() []OptionMetadataTypeEnum {
	values := make([]OptionMetadataTypeEnum, 0)
	for _, v := range mappingOptionMetadataType {
		values = append(values, v)
	}
	return values
}
