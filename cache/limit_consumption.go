// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// OraCache Public API
//
// API for the Data Caching Service. Use this service to manage Redis replicated caches.
//

package cache

import (
	"github.com/oracle/oci-go-sdk/common"
)

// LimitConsumption Region-specific limit consumption values.
type LimitConsumption struct {

	// Consumption name corresponding to the limit name
	Name *string `mandatory:"false" json:"name"`

	// Consumption value
	Value *string `mandatory:"false" json:"value"`

	// Scope of the consumption data
	Scope LimitConsumptionScopeEnum `mandatory:"false" json:"scope,omitempty"`
}

func (m LimitConsumption) String() string {
	return common.PointerString(m)
}

// LimitConsumptionScopeEnum Enum with underlying type: string
type LimitConsumptionScopeEnum string

// Set of constants representing the allowable values for LimitConsumptionScopeEnum
const (
	LimitConsumptionScopeRegion LimitConsumptionScopeEnum = "REGION"
)

var mappingLimitConsumptionScope = map[string]LimitConsumptionScopeEnum{
	"REGION": LimitConsumptionScopeRegion,
}

// GetLimitConsumptionScopeEnumValues Enumerates the set of values for LimitConsumptionScopeEnum
func GetLimitConsumptionScopeEnumValues() []LimitConsumptionScopeEnum {
	values := make([]LimitConsumptionScopeEnum, 0)
	for _, v := range mappingLimitConsumptionScope {
		values = append(values, v)
	}
	return values
}
