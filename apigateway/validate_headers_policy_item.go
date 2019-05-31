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

// ValidateHeadersPolicyItem Validate whether the value or presence of an http header is as expected.
type ValidateHeadersPolicyItem struct {

	// The case-insensitive name of the header to validate.
	Header *string `mandatory:"true" json:"header"`

	// The constraint type to apply on the header.
	Constraint ValidateHeadersPolicyItemConstraintEnum `mandatory:"true" json:"constraint"`

	// The values used by the constraint type.
	Values []string `mandatory:"false" json:"values"`
}

func (m ValidateHeadersPolicyItem) String() string {
	return common.PointerString(m)
}

// ValidateHeadersPolicyItemConstraintEnum Enum with underlying type: string
type ValidateHeadersPolicyItemConstraintEnum string

// Set of constants representing the allowable values for ValidateHeadersPolicyItemConstraintEnum
const (
	ValidateHeadersPolicyItemConstraintExists    ValidateHeadersPolicyItemConstraintEnum = "EXISTS"
	ValidateHeadersPolicyItemConstraintNotExists ValidateHeadersPolicyItemConstraintEnum = "NOT_EXISTS"
	ValidateHeadersPolicyItemConstraintIn        ValidateHeadersPolicyItemConstraintEnum = "IN"
	ValidateHeadersPolicyItemConstraintNotIn     ValidateHeadersPolicyItemConstraintEnum = "NOT_IN"
)

var mappingValidateHeadersPolicyItemConstraint = map[string]ValidateHeadersPolicyItemConstraintEnum{
	"EXISTS":     ValidateHeadersPolicyItemConstraintExists,
	"NOT_EXISTS": ValidateHeadersPolicyItemConstraintNotExists,
	"IN":         ValidateHeadersPolicyItemConstraintIn,
	"NOT_IN":     ValidateHeadersPolicyItemConstraintNotIn,
}

// GetValidateHeadersPolicyItemConstraintEnumValues Enumerates the set of values for ValidateHeadersPolicyItemConstraintEnum
func GetValidateHeadersPolicyItemConstraintEnumValues() []ValidateHeadersPolicyItemConstraintEnum {
	values := make([]ValidateHeadersPolicyItemConstraintEnum, 0)
	for _, v := range mappingValidateHeadersPolicyItemConstraint {
		values = append(values, v)
	}
	return values
}
