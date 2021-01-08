// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Fleets/Query API
//
// Java Management Service Fleets/Query API
//

package jms

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateFleet OperationTypeEnum = "CREATE_FLEET"
	OperationTypeUpdateFleet OperationTypeEnum = "UPDATE_FLEET"
	OperationTypeDeleteFleet OperationTypeEnum = "DELETE_FLEET"
	OperationTypeMoveFleet   OperationTypeEnum = "MOVE_FLEET"
)

var mappingOperationType = map[string]OperationTypeEnum{
	"CREATE_FLEET": OperationTypeCreateFleet,
	"UPDATE_FLEET": OperationTypeUpdateFleet,
	"DELETE_FLEET": OperationTypeDeleteFleet,
	"MOVE_FLEET":   OperationTypeMoveFleet,
}

// GetOperationTypeEnumValues Enumerates the set of values for OperationTypeEnum
func GetOperationTypeEnumValues() []OperationTypeEnum {
	values := make([]OperationTypeEnum, 0)
	for _, v := range mappingOperationType {
		values = append(values, v)
	}
	return values
}
