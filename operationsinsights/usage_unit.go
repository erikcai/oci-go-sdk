// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OperationsInsights API
//
// The OperationsInsights API for OPSI service.
//

package operationsinsights

// UsageUnitEnum Enum with underlying type: string
type UsageUnitEnum string

// Set of constants representing the allowable values for UsageUnitEnum
const (
	UsageUnitCores   UsageUnitEnum = "CORES"
	UsageUnitGb      UsageUnitEnum = "GB"
	UsageUnitMbps    UsageUnitEnum = "MBPS"
	UsageUnitPercent UsageUnitEnum = "PERCENT"
)

var mappingUsageUnit = map[string]UsageUnitEnum{
	"CORES":   UsageUnitCores,
	"GB":      UsageUnitGb,
	"MBPS":    UsageUnitMbps,
	"PERCENT": UsageUnitPercent,
}

// GetUsageUnitEnumValues Enumerates the set of values for UsageUnitEnum
func GetUsageUnitEnumValues() []UsageUnitEnum {
	values := make([]UsageUnitEnum, 0)
	for _, v := range mappingUsageUnit {
		values = append(values, v)
	}
	return values
}
