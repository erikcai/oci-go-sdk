// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// MySQL as a Service API
//
// The API for the MySQL Service
//

package mysql

// WorkRequestOperationTypeEnum Enum with underlying type: string
type WorkRequestOperationTypeEnum string

// Set of constants representing the allowable values for WorkRequestOperationTypeEnum
const (
	WorkRequestOperationTypeCreateInstance           WorkRequestOperationTypeEnum = "CREATE_INSTANCE"
	WorkRequestOperationTypeUpdateInstance           WorkRequestOperationTypeEnum = "UPDATE_INSTANCE"
	WorkRequestOperationTypeDeleteInstance           WorkRequestOperationTypeEnum = "DELETE_INSTANCE"
	WorkRequestOperationTypeStartInstance            WorkRequestOperationTypeEnum = "START_INSTANCE"
	WorkRequestOperationTypeStopInstance             WorkRequestOperationTypeEnum = "STOP_INSTANCE"
	WorkRequestOperationTypeRestartInstance          WorkRequestOperationTypeEnum = "RESTART_INSTANCE"
	WorkRequestOperationTypeApplyConfigurationUpdate WorkRequestOperationTypeEnum = "APPLY_CONFIGURATION_UPDATE"
)

var mappingWorkRequestOperationType = map[string]WorkRequestOperationTypeEnum{
	"CREATE_INSTANCE":            WorkRequestOperationTypeCreateInstance,
	"UPDATE_INSTANCE":            WorkRequestOperationTypeUpdateInstance,
	"DELETE_INSTANCE":            WorkRequestOperationTypeDeleteInstance,
	"START_INSTANCE":             WorkRequestOperationTypeStartInstance,
	"STOP_INSTANCE":              WorkRequestOperationTypeStopInstance,
	"RESTART_INSTANCE":           WorkRequestOperationTypeRestartInstance,
	"APPLY_CONFIGURATION_UPDATE": WorkRequestOperationTypeApplyConfigurationUpdate,
}

// GetWorkRequestOperationTypeEnumValues Enumerates the set of values for WorkRequestOperationTypeEnum
func GetWorkRequestOperationTypeEnumValues() []WorkRequestOperationTypeEnum {
	values := make([]WorkRequestOperationTypeEnum, 0)
	for _, v := range mappingWorkRequestOperationType {
		values = append(values, v)
	}
	return values
}
