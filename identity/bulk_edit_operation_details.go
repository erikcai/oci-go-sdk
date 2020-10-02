// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"github.com/oracle/oci-go-sdk/v25/common"
)

// BulkEditOperationDetails The representation of BulkEditOperationDetails
type BulkEditOperationDetails struct {

	// An enum-like description of the type of operation. ADD_WHERE_ABSENT - Add defined tag only if it doesn't exist on resource SET_WHERE_PRESENT - Set defined tag value only if it is present on resource ADD_OR_SET - Combines ADD and SET (Add defined tag if doesn't exist on resource or Set defined tag value if it is already present) REMOVE - Removes the defined tag on the resource. Tag will be removed from resource irrespective of the value.
	OperationType BulkEditOperationDetailsOperationTypeEnum `mandatory:"true" json:"operationType"`

	// Array of tags to be used for operations on resources.
	DefinedTags []map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`
}

func (m BulkEditOperationDetails) String() string {
	return common.PointerString(m)
}

// BulkEditOperationDetailsOperationTypeEnum Enum with underlying type: string
type BulkEditOperationDetailsOperationTypeEnum string

// Set of constants representing the allowable values for BulkEditOperationDetailsOperationTypeEnum
const (
	BulkEditOperationDetailsOperationTypeAddWhereAbsent  BulkEditOperationDetailsOperationTypeEnum = "ADD_WHERE_ABSENT"
	BulkEditOperationDetailsOperationTypeSetWherePresent BulkEditOperationDetailsOperationTypeEnum = "SET_WHERE_PRESENT"
	BulkEditOperationDetailsOperationTypeAddOrSet        BulkEditOperationDetailsOperationTypeEnum = "ADD_OR_SET"
	BulkEditOperationDetailsOperationTypeRemove          BulkEditOperationDetailsOperationTypeEnum = "REMOVE"
)

var mappingBulkEditOperationDetailsOperationType = map[string]BulkEditOperationDetailsOperationTypeEnum{
	"ADD_WHERE_ABSENT":  BulkEditOperationDetailsOperationTypeAddWhereAbsent,
	"SET_WHERE_PRESENT": BulkEditOperationDetailsOperationTypeSetWherePresent,
	"ADD_OR_SET":        BulkEditOperationDetailsOperationTypeAddOrSet,
	"REMOVE":            BulkEditOperationDetailsOperationTypeRemove,
}

// GetBulkEditOperationDetailsOperationTypeEnumValues Enumerates the set of values for BulkEditOperationDetailsOperationTypeEnum
func GetBulkEditOperationDetailsOperationTypeEnumValues() []BulkEditOperationDetailsOperationTypeEnum {
	values := make([]BulkEditOperationDetailsOperationTypeEnum, 0)
	for _, v := range mappingBulkEditOperationDetailsOperationType {
		values = append(values, v)
	}
	return values
}
