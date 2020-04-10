// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Data Integration Service API Specification
//
// Data Integration Service API Specification
//

package dataintegration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// RuleBasedFieldMap auto generated description
type RuleBasedFieldMap struct {

	// Descriptive text for the object.
	Description *string `mandatory:"false" json:"description"`

	// Object key
	Key *string `mandatory:"false" json:"key"`

	// modelVersion
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	ConfigValues *ConfigValues `mandatory:"false" json:"configValues"`

	// fromPattern
	FromPattern *string `mandatory:"false" json:"fromPattern"`

	// toPattern
	ToPattern *string `mandatory:"false" json:"toPattern"`

	// javaRegexSyntax
	IsJavaRegexSyntax *bool `mandatory:"false" json:"isJavaRegexSyntax"`

	// Status of object, can set this to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	FromRuleConfig *RuleTypeConfig `mandatory:"false" json:"fromRuleConfig"`

	ToRuleConfig *RuleTypeConfig `mandatory:"false" json:"toRuleConfig"`

	// mapType
	MapType RuleBasedFieldMapMapTypeEnum `mandatory:"false" json:"mapType,omitempty"`
}

//GetDescription returns Description
func (m RuleBasedFieldMap) GetDescription() *string {
	return m.Description
}

func (m RuleBasedFieldMap) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m RuleBasedFieldMap) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeRuleBasedFieldMap RuleBasedFieldMap
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeRuleBasedFieldMap
	}{
		"RULE_BASED_FIELD_MAP",
		(MarshalTypeRuleBasedFieldMap)(m),
	}

	return json.Marshal(&s)
}

// RuleBasedFieldMapMapTypeEnum Enum with underlying type: string
type RuleBasedFieldMapMapTypeEnum string

// Set of constants representing the allowable values for RuleBasedFieldMapMapTypeEnum
const (
	RuleBasedFieldMapMapTypeMapbyname     RuleBasedFieldMapMapTypeEnum = "MAPBYNAME"
	RuleBasedFieldMapMapTypeMapbyposition RuleBasedFieldMapMapTypeEnum = "MAPBYPOSITION"
	RuleBasedFieldMapMapTypeMapbypattern  RuleBasedFieldMapMapTypeEnum = "MAPBYPATTERN"
)

var mappingRuleBasedFieldMapMapType = map[string]RuleBasedFieldMapMapTypeEnum{
	"MAPBYNAME":     RuleBasedFieldMapMapTypeMapbyname,
	"MAPBYPOSITION": RuleBasedFieldMapMapTypeMapbyposition,
	"MAPBYPATTERN":  RuleBasedFieldMapMapTypeMapbypattern,
}

// GetRuleBasedFieldMapMapTypeEnumValues Enumerates the set of values for RuleBasedFieldMapMapTypeEnum
func GetRuleBasedFieldMapMapTypeEnumValues() []RuleBasedFieldMapMapTypeEnum {
	values := make([]RuleBasedFieldMapMapTypeEnum, 0)
	for _, v := range mappingRuleBasedFieldMapMapType {
		values = append(values, v)
	}
	return values
}
