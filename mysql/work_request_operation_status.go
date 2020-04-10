// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// MySQL as a Service API
//
// The API for the MySQL Service
//

package mysql

// WorkRequestOperationStatusEnum Enum with underlying type: string
type WorkRequestOperationStatusEnum string

// Set of constants representing the allowable values for WorkRequestOperationStatusEnum
const (
	WorkRequestOperationStatusAccepted   WorkRequestOperationStatusEnum = "ACCEPTED"
	WorkRequestOperationStatusInProgress WorkRequestOperationStatusEnum = "IN_PROGRESS"
	WorkRequestOperationStatusFailed     WorkRequestOperationStatusEnum = "FAILED"
	WorkRequestOperationStatusSucceeded  WorkRequestOperationStatusEnum = "SUCCEEDED"
	WorkRequestOperationStatusCanceling  WorkRequestOperationStatusEnum = "CANCELING"
	WorkRequestOperationStatusCanceled   WorkRequestOperationStatusEnum = "CANCELED"
)

var mappingWorkRequestOperationStatus = map[string]WorkRequestOperationStatusEnum{
	"ACCEPTED":    WorkRequestOperationStatusAccepted,
	"IN_PROGRESS": WorkRequestOperationStatusInProgress,
	"FAILED":      WorkRequestOperationStatusFailed,
	"SUCCEEDED":   WorkRequestOperationStatusSucceeded,
	"CANCELING":   WorkRequestOperationStatusCanceling,
	"CANCELED":    WorkRequestOperationStatusCanceled,
}

// GetWorkRequestOperationStatusEnumValues Enumerates the set of values for WorkRequestOperationStatusEnum
func GetWorkRequestOperationStatusEnumValues() []WorkRequestOperationStatusEnum {
	values := make([]WorkRequestOperationStatusEnum, 0)
	for _, v := range mappingWorkRequestOperationStatus {
		values = append(values, v)
	}
	return values
}
