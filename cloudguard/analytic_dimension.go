// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

// AnalyticDimensionEnum Enum with underlying type: string
type AnalyticDimensionEnum string

// Set of constants representing the allowable values for AnalyticDimensionEnum
const (
	AnalyticDimensionResourceType  AnalyticDimensionEnum = "RESOURCE_TYPE"
	AnalyticDimensionRegion        AnalyticDimensionEnum = "REGION"
	AnalyticDimensionCompartmentId AnalyticDimensionEnum = "COMPARTMENT_ID"
	AnalyticDimensionRiskLevel     AnalyticDimensionEnum = "RISK_LEVEL"
)

var mappingAnalyticDimension = map[string]AnalyticDimensionEnum{
	"RESOURCE_TYPE":  AnalyticDimensionResourceType,
	"REGION":         AnalyticDimensionRegion,
	"COMPARTMENT_ID": AnalyticDimensionCompartmentId,
	"RISK_LEVEL":     AnalyticDimensionRiskLevel,
}

// GetAnalyticDimensionEnumValues Enumerates the set of values for AnalyticDimensionEnum
func GetAnalyticDimensionEnumValues() []AnalyticDimensionEnum {
	values := make([]AnalyticDimensionEnum, 0)
	for _, v := range mappingAnalyticDimension {
		values = append(values, v)
	}
	return values
}
