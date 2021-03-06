// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Fleets/Query API
//
// Java Management Service Fleets/Query API
//

package jms

// JreSortByEnum Enum with underlying type: string
type JreSortByEnum string

// Set of constants representing the allowable values for JreSortByEnum
const (
	JreSortByTimeFirstSeen JreSortByEnum = "timeFirstSeen"
	JreSortByTimeLastSeen  JreSortByEnum = "timeLastSeen"
	JreSortByVersion       JreSortByEnum = "version"
	JreSortByDistribution  JreSortByEnum = "distribution"
	JreSortByVendor        JreSortByEnum = "vendor"
)

var mappingJreSortBy = map[string]JreSortByEnum{
	"timeFirstSeen": JreSortByTimeFirstSeen,
	"timeLastSeen":  JreSortByTimeLastSeen,
	"version":       JreSortByVersion,
	"distribution":  JreSortByDistribution,
	"vendor":        JreSortByVendor,
}

// GetJreSortByEnumValues Enumerates the set of values for JreSortByEnum
func GetJreSortByEnumValues() []JreSortByEnum {
	values := make([]JreSortByEnum, 0)
	for _, v := range mappingJreSortBy {
		values = append(values, v)
	}
	return values
}
