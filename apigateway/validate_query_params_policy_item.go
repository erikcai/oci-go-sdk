// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// API Gateway API
//
// API for the API Gateway service. Use this API to manage gateways, deployments, and related items.
//

package apigateway

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ValidateQueryParamsPolicyItem Validate whether the value or presence of an query parameter is as expected.
type ValidateQueryParamsPolicyItem struct {

	// The case-sensitive name of the query parameter.
	Key *string `mandatory:"true" json:"key"`

	// The constraint type to apply on the query parameter.
	Constraint ValidateQueryParamsPolicyItemConstraintEnum `mandatory:"true" json:"constraint"`

	// The values used by the constraint type.
	Values []string `mandatory:"false" json:"values"`
}

func (m ValidateQueryParamsPolicyItem) String() string {
	return common.PointerString(m)
}

// ValidateQueryParamsPolicyItemConstraintEnum Enum with underlying type: string
type ValidateQueryParamsPolicyItemConstraintEnum string

// Set of constants representing the allowable values for ValidateQueryParamsPolicyItemConstraintEnum
const (
	ValidateQueryParamsPolicyItemConstraintExists    ValidateQueryParamsPolicyItemConstraintEnum = "EXISTS"
	ValidateQueryParamsPolicyItemConstraintNotExists ValidateQueryParamsPolicyItemConstraintEnum = "NOT_EXISTS"
	ValidateQueryParamsPolicyItemConstraintIn        ValidateQueryParamsPolicyItemConstraintEnum = "IN"
	ValidateQueryParamsPolicyItemConstraintNotIn     ValidateQueryParamsPolicyItemConstraintEnum = "NOT_IN"
)

var mappingValidateQueryParamsPolicyItemConstraint = map[string]ValidateQueryParamsPolicyItemConstraintEnum{
	"EXISTS":     ValidateQueryParamsPolicyItemConstraintExists,
	"NOT_EXISTS": ValidateQueryParamsPolicyItemConstraintNotExists,
	"IN":         ValidateQueryParamsPolicyItemConstraintIn,
	"NOT_IN":     ValidateQueryParamsPolicyItemConstraintNotIn,
}

// GetValidateQueryParamsPolicyItemConstraintEnumValues Enumerates the set of values for ValidateQueryParamsPolicyItemConstraintEnum
func GetValidateQueryParamsPolicyItemConstraintEnumValues() []ValidateQueryParamsPolicyItemConstraintEnum {
	values := make([]ValidateQueryParamsPolicyItemConstraintEnum, 0)
	for _, v := range mappingValidateQueryParamsPolicyItemConstraint {
		values = append(values, v)
	}
	return values
}
