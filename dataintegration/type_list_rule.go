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

// TypeListRule The type list rule which defines how fields are projected.
type TypeListRule struct {

	// Object key
	Key *string `mandatory:"false" json:"key"`

	// modelVersion
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// javaRegexSyntax
	IsJavaRegexSyntax *bool `mandatory:"false" json:"isJavaRegexSyntax"`

	ConfigValues *ConfigValues `mandatory:"false" json:"configValues"`

	// Status of object, can set this to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Descriptive text for the object.
	Description *string `mandatory:"false" json:"description"`

	// skipRemainingRulesOnMatch
	IsSkipRemainingRulesOnMatch *bool `mandatory:"false" json:"isSkipRemainingRulesOnMatch"`

	// Reference to a typed object
	Scope *string `mandatory:"false" json:"scope"`

	// cascade
	IsCascade *bool `mandatory:"false" json:"isCascade"`

	// caseSensitive
	IsCaseSensitive *bool `mandatory:"false" json:"isCaseSensitive"`

	// types
	Types []BaseType `mandatory:"false" json:"types"`

	// matchingStrategy
	MatchingStrategy TypeListRuleMatchingStrategyEnum `mandatory:"false" json:"matchingStrategy,omitempty"`

	// ruleType
	RuleType TypeListRuleRuleTypeEnum `mandatory:"false" json:"ruleType,omitempty"`
}

//GetKey returns Key
func (m TypeListRule) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m TypeListRule) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m TypeListRule) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetIsJavaRegexSyntax returns IsJavaRegexSyntax
func (m TypeListRule) GetIsJavaRegexSyntax() *bool {
	return m.IsJavaRegexSyntax
}

//GetConfigValues returns ConfigValues
func (m TypeListRule) GetConfigValues() *ConfigValues {
	return m.ConfigValues
}

//GetObjectStatus returns ObjectStatus
func (m TypeListRule) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetDescription returns Description
func (m TypeListRule) GetDescription() *string {
	return m.Description
}

func (m TypeListRule) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m TypeListRule) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTypeListRule TypeListRule
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeTypeListRule
	}{
		"TYPE_LIST_RULE",
		(MarshalTypeTypeListRule)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *TypeListRule) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Key                         *string                          `json:"key"`
		ModelVersion                *string                          `json:"modelVersion"`
		ParentRef                   *ParentReference                 `json:"parentRef"`
		IsJavaRegexSyntax           *bool                            `json:"isJavaRegexSyntax"`
		ConfigValues                *ConfigValues                    `json:"configValues"`
		ObjectStatus                *int                             `json:"objectStatus"`
		Description                 *string                          `json:"description"`
		IsSkipRemainingRulesOnMatch *bool                            `json:"isSkipRemainingRulesOnMatch"`
		Scope                       *string                          `json:"scope"`
		IsCascade                   *bool                            `json:"isCascade"`
		MatchingStrategy            TypeListRuleMatchingStrategyEnum `json:"matchingStrategy"`
		IsCaseSensitive             *bool                            `json:"isCaseSensitive"`
		RuleType                    TypeListRuleRuleTypeEnum         `json:"ruleType"`
		Types                       []basetype                       `json:"types"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Key = model.Key

	m.ModelVersion = model.ModelVersion

	m.ParentRef = model.ParentRef

	m.IsJavaRegexSyntax = model.IsJavaRegexSyntax

	m.ConfigValues = model.ConfigValues

	m.ObjectStatus = model.ObjectStatus

	m.Description = model.Description

	m.IsSkipRemainingRulesOnMatch = model.IsSkipRemainingRulesOnMatch

	m.Scope = model.Scope

	m.IsCascade = model.IsCascade

	m.MatchingStrategy = model.MatchingStrategy

	m.IsCaseSensitive = model.IsCaseSensitive

	m.RuleType = model.RuleType

	m.Types = make([]BaseType, len(model.Types))
	for i, n := range model.Types {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Types[i] = nn.(BaseType)
		} else {
			m.Types[i] = nil
		}
	}
	return
}

// TypeListRuleMatchingStrategyEnum Enum with underlying type: string
type TypeListRuleMatchingStrategyEnum string

// Set of constants representing the allowable values for TypeListRuleMatchingStrategyEnum
const (
	TypeListRuleMatchingStrategyNameOrTags TypeListRuleMatchingStrategyEnum = "NAME_OR_TAGS"
	TypeListRuleMatchingStrategyTagsOnly   TypeListRuleMatchingStrategyEnum = "TAGS_ONLY"
	TypeListRuleMatchingStrategyNameOnly   TypeListRuleMatchingStrategyEnum = "NAME_ONLY"
)

var mappingTypeListRuleMatchingStrategy = map[string]TypeListRuleMatchingStrategyEnum{
	"NAME_OR_TAGS": TypeListRuleMatchingStrategyNameOrTags,
	"TAGS_ONLY":    TypeListRuleMatchingStrategyTagsOnly,
	"NAME_ONLY":    TypeListRuleMatchingStrategyNameOnly,
}

// GetTypeListRuleMatchingStrategyEnumValues Enumerates the set of values for TypeListRuleMatchingStrategyEnum
func GetTypeListRuleMatchingStrategyEnumValues() []TypeListRuleMatchingStrategyEnum {
	values := make([]TypeListRuleMatchingStrategyEnum, 0)
	for _, v := range mappingTypeListRuleMatchingStrategy {
		values = append(values, v)
	}
	return values
}

// TypeListRuleRuleTypeEnum Enum with underlying type: string
type TypeListRuleRuleTypeEnum string

// Set of constants representing the allowable values for TypeListRuleRuleTypeEnum
const (
	TypeListRuleRuleTypeInclude TypeListRuleRuleTypeEnum = "INCLUDE"
	TypeListRuleRuleTypeExclude TypeListRuleRuleTypeEnum = "EXCLUDE"
)

var mappingTypeListRuleRuleType = map[string]TypeListRuleRuleTypeEnum{
	"INCLUDE": TypeListRuleRuleTypeInclude,
	"EXCLUDE": TypeListRuleRuleTypeExclude,
}

// GetTypeListRuleRuleTypeEnumValues Enumerates the set of values for TypeListRuleRuleTypeEnum
func GetTypeListRuleRuleTypeEnumValues() []TypeListRuleRuleTypeEnum {
	values := make([]TypeListRuleRuleTypeEnum, 0)
	for _, v := range mappingTypeListRuleRuleType {
		values = append(values, v)
	}
	return values
}
