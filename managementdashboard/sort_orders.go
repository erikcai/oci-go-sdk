// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// ManagementDashboard API
//
// Management Dashboard micro-service provides APIs to support dashboard and saved search metadata preservation, as follows 1) Save and retrieve metadata to support UI activities (Create an empty dashboard, open a saved search, etc.). 2) Check authority for each CRUD operation. 3) Validate input: Are all properties present?  Any empty values?  Are all saved searches OOB when dashboard is OOB?  Etc. 4) Import and export dashboards.
//

package managementdashboard

// SortOrdersEnum Enum with underlying type: string
type SortOrdersEnum string

// Set of constants representing the allowable values for SortOrdersEnum
const (
	SortOrdersAsc  SortOrdersEnum = "ASC"
	SortOrdersDesc SortOrdersEnum = "DESC"
)

var mappingSortOrders = map[string]SortOrdersEnum{
	"ASC":  SortOrdersAsc,
	"DESC": SortOrdersDesc,
}

// GetSortOrdersEnumValues Enumerates the set of values for SortOrdersEnum
func GetSortOrdersEnumValues() []SortOrdersEnum {
	values := make([]SortOrdersEnum, 0)
	for _, v := range mappingSortOrders {
		values = append(values, v)
	}
	return values
}
