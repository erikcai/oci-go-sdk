// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Load Balancing API
//
// API for the Load Balancing service. Use this API to manage load balancers, backend sets, and related items. For more
// information, see Overview of Load Balancing (https://docs.cloud.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm).
//

package loadbalancer

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// PathMatchCondition Path match condition for redirect path.
type PathMatchCondition struct {

	// The path string for which the redirection rule should be applied.
	AttributeValue *string `mandatory:"true" json:"attributeValue"`

	// Specifies how the load balancing service compares a [PathMatchCondition] object's `attributeValue` string
	//     against the incoming URI.
	//     *  **EXACT_MATCH** - Looks for a `attributeValue` string that exactly matches the incoming URI path.
	//     *  **FORCE_LONGEST_PREFIX_MATCH** - Looks for the `attributeValue` string with the best, longest match of the beginning
	//        portion of the incoming URI path.
	//     *  **PREFIX_MATCH** - Looks for a `attributeValue` string that matches the beginning portion of the incoming URI path.
	//     *  **SUFFIX_MATCH** - Looks for a `attributeValue` string that matches the ending portion of the incoming URI path.
	Operator PathMatchConditionOperatorEnum `mandatory:"true" json:"operator"`
}

func (m PathMatchCondition) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m PathMatchCondition) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePathMatchCondition PathMatchCondition
	s := struct {
		DiscriminatorParam string `json:"attributeName"`
		MarshalTypePathMatchCondition
	}{
		"PATH",
		(MarshalTypePathMatchCondition)(m),
	}

	return json.Marshal(&s)
}

// PathMatchConditionOperatorEnum Enum with underlying type: string
type PathMatchConditionOperatorEnum string

// Set of constants representing the allowable values for PathMatchConditionOperatorEnum
const (
	PathMatchConditionOperatorExactMatch              PathMatchConditionOperatorEnum = "EXACT_MATCH"
	PathMatchConditionOperatorForceLongestPrefixMatch PathMatchConditionOperatorEnum = "FORCE_LONGEST_PREFIX_MATCH"
	PathMatchConditionOperatorPrefixMatch             PathMatchConditionOperatorEnum = "PREFIX_MATCH"
	PathMatchConditionOperatorSuffixMatch             PathMatchConditionOperatorEnum = "SUFFIX_MATCH"
)

var mappingPathMatchConditionOperator = map[string]PathMatchConditionOperatorEnum{
	"EXACT_MATCH":                PathMatchConditionOperatorExactMatch,
	"FORCE_LONGEST_PREFIX_MATCH": PathMatchConditionOperatorForceLongestPrefixMatch,
	"PREFIX_MATCH":               PathMatchConditionOperatorPrefixMatch,
	"SUFFIX_MATCH":               PathMatchConditionOperatorSuffixMatch,
}

// GetPathMatchConditionOperatorEnumValues Enumerates the set of values for PathMatchConditionOperatorEnum
func GetPathMatchConditionOperatorEnumValues() []PathMatchConditionOperatorEnum {
	values := make([]PathMatchConditionOperatorEnum, 0)
	for _, v := range mappingPathMatchConditionOperator {
		values = append(values, v)
	}
	return values
}
