// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// KAM API
//
// description: |
//   Kubernetes Add-on Manager API for installing, uninstalling and upgrading
//   OKE add-ons or Marketplace applications on an OKE cluster
//

package kam

// SortByEnum Enum with underlying type: string
type SortByEnum string

// Set of constants representing the allowable values for SortByEnum
const (
	SortByTimeaccepted SortByEnum = "TIMEACCEPTED"
	SortByTimeupdated  SortByEnum = "TIMEUPDATED"
)

var mappingSortBy = map[string]SortByEnum{
	"TIMEACCEPTED": SortByTimeaccepted,
	"TIMEUPDATED":  SortByTimeupdated,
}

// GetSortByEnumValues Enumerates the set of values for SortByEnum
func GetSortByEnumValues() []SortByEnum {
	values := make([]SortByEnum, 0)
	for _, v := range mappingSortBy {
		values = append(values, v)
	}
	return values
}
