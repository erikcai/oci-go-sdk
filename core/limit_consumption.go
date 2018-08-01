// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"github.com/oracle/oci-go-sdk/common"
)

// LimitConsumption Regional and AD specific limit consumption values
type LimitConsumption struct {

	// Consumption name corresponding to the limit name
	Name *string `mandatory:"false" json:"name"`

	// Scope of the consumption data
	Scope LimitConsumptionScopeEnum `mandatory:"false" json:"scope,omitempty"`

	// Consumption value
	Value *string `mandatory:"false" json:"value"`
}

func (m LimitConsumption) String() string {
	return common.PointerString(m)
}

// LimitConsumptionScopeEnum Enum with underlying type: string
type LimitConsumptionScopeEnum string

// Set of constants representing the allowable values for LimitConsumptionScope
const (
	LimitConsumptionScopeRegion             LimitConsumptionScopeEnum = "REGION"
	LimitConsumptionScopeAvailabilityDomain LimitConsumptionScopeEnum = "AVAILABILITY_DOMAIN"
)

var mappingLimitConsumptionScope = map[string]LimitConsumptionScopeEnum{
	"REGION":              LimitConsumptionScopeRegion,
	"AVAILABILITY_DOMAIN": LimitConsumptionScopeAvailabilityDomain,
}

// GetLimitConsumptionScopeEnumValues Enumerates the set of values for LimitConsumptionScope
func GetLimitConsumptionScopeEnumValues() []LimitConsumptionScopeEnum {
	values := make([]LimitConsumptionScopeEnum, 0)
	for _, v := range mappingLimitConsumptionScope {
		values = append(values, v)
	}
	return values
}
