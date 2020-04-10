// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Compliance Documents Service API
//
// API for downloading Oracle Cloud Infrastructure compliance documents from the Oracle Cloud Infrastructure Console.
//

package compdocsapi

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
