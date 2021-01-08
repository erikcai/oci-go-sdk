// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Fleets/Query API
//
// Java Management Service Fleets/Query API
//

package jms

// ManagedInstanceTypeEnum Enum with underlying type: string
type ManagedInstanceTypeEnum string

// Set of constants representing the allowable values for ManagedInstanceTypeEnum
const (
	ManagedInstanceTypePolarisAgent ManagedInstanceTypeEnum = "POLARIS_AGENT"
)

var mappingManagedInstanceType = map[string]ManagedInstanceTypeEnum{
	"POLARIS_AGENT": ManagedInstanceTypePolarisAgent,
}

// GetManagedInstanceTypeEnumValues Enumerates the set of values for ManagedInstanceTypeEnum
func GetManagedInstanceTypeEnumValues() []ManagedInstanceTypeEnum {
	values := make([]ManagedInstanceTypeEnum, 0)
	for _, v := range mappingManagedInstanceType {
		values = append(values, v)
	}
	return values
}
