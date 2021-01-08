// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Fleets/Query API
//
// Java Management Service Fleets/Query API
//

package jms

// ManagedInstanceSortByEnum Enum with underlying type: string
type ManagedInstanceSortByEnum string

// Set of constants representing the allowable values for ManagedInstanceSortByEnum
const (
	ManagedInstanceSortByTimeFirstSeen ManagedInstanceSortByEnum = "timeFirstSeen"
	ManagedInstanceSortByTimeLastSeen  ManagedInstanceSortByEnum = "timeLastSeen"
)

var mappingManagedInstanceSortBy = map[string]ManagedInstanceSortByEnum{
	"timeFirstSeen": ManagedInstanceSortByTimeFirstSeen,
	"timeLastSeen":  ManagedInstanceSortByTimeLastSeen,
}

// GetManagedInstanceSortByEnumValues Enumerates the set of values for ManagedInstanceSortByEnum
func GetManagedInstanceSortByEnumValues() []ManagedInstanceSortByEnum {
	values := make([]ManagedInstanceSortByEnum, 0)
	for _, v := range mappingManagedInstanceSortBy {
		values = append(values, v)
	}
	return values
}
