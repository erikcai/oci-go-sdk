// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Data Flow API
//
// Use the Data Flow APIs to run any Apache Spark application at any scale without deploying or managing any infrastructure.
//

package dataflow

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ConsumptionSummary The regional and AD specific limit consumption values.
type ConsumptionSummary struct {

	// The consumption name corresponding to the limit name.
	Name *string `mandatory:"true" json:"name"`

	// The scope of the consumption data.
	Scope ConsumptionSummaryScopeEnum `mandatory:"true" json:"scope"`

	// The consumption value.
	Value *string `mandatory:"true" json:"value"`
}

func (m ConsumptionSummary) String() string {
	return common.PointerString(m)
}

// ConsumptionSummaryScopeEnum Enum with underlying type: string
type ConsumptionSummaryScopeEnum string

// Set of constants representing the allowable values for ConsumptionSummaryScopeEnum
const (
	ConsumptionSummaryScopeRegion             ConsumptionSummaryScopeEnum = "REGION"
	ConsumptionSummaryScopeAvailabilityDomain ConsumptionSummaryScopeEnum = "AVAILABILITY_DOMAIN"
)

var mappingConsumptionSummaryScope = map[string]ConsumptionSummaryScopeEnum{
	"REGION":              ConsumptionSummaryScopeRegion,
	"AVAILABILITY_DOMAIN": ConsumptionSummaryScopeAvailabilityDomain,
}

// GetConsumptionSummaryScopeEnumValues Enumerates the set of values for ConsumptionSummaryScopeEnum
func GetConsumptionSummaryScopeEnumValues() []ConsumptionSummaryScopeEnum {
	values := make([]ConsumptionSummaryScopeEnum, 0)
	for _, v := range mappingConsumptionSummaryScope {
		values = append(values, v)
	}
	return values
}
