// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Data Integration Service API Specification
//
// Data Integration Service API Specification
//

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/common"
)

// SortClause auto generated description
type SortClause struct {
	Field *ShapeField `mandatory:"false" json:"field"`

	// order
	Order SortClauseOrderEnum `mandatory:"false" json:"order,omitempty"`
}

func (m SortClause) String() string {
	return common.PointerString(m)
}

// SortClauseOrderEnum Enum with underlying type: string
type SortClauseOrderEnum string

// Set of constants representing the allowable values for SortClauseOrderEnum
const (
	SortClauseOrderAsc  SortClauseOrderEnum = "ASC"
	SortClauseOrderDesc SortClauseOrderEnum = "DESC"
)

var mappingSortClauseOrder = map[string]SortClauseOrderEnum{
	"ASC":  SortClauseOrderAsc,
	"DESC": SortClauseOrderDesc,
}

// GetSortClauseOrderEnumValues Enumerates the set of values for SortClauseOrderEnum
func GetSortClauseOrderEnumValues() []SortClauseOrderEnum {
	values := make([]SortClauseOrderEnum, 0)
	for _, v := range mappingSortClauseOrder {
		values = append(values, v)
	}
	return values
}
