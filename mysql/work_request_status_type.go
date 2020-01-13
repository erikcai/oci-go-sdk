// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// MySQL as a Service API
//
// The API for the MySQL Service
//

package mysql

// WorkRequestStatusTypeEnum Enum with underlying type: string
type WorkRequestStatusTypeEnum string

// Set of constants representing the allowable values for WorkRequestStatusTypeEnum
const (
	WorkRequestStatusTypeAccepted   WorkRequestStatusTypeEnum = "ACCEPTED"
	WorkRequestStatusTypeInProgress WorkRequestStatusTypeEnum = "IN_PROGRESS"
	WorkRequestStatusTypeFailed     WorkRequestStatusTypeEnum = "FAILED"
	WorkRequestStatusTypeSucceeded  WorkRequestStatusTypeEnum = "SUCCEEDED"
	WorkRequestStatusTypeCanceling  WorkRequestStatusTypeEnum = "CANCELING"
	WorkRequestStatusTypeCanceled   WorkRequestStatusTypeEnum = "CANCELED"
)

var mappingWorkRequestStatusType = map[string]WorkRequestStatusTypeEnum{
	"ACCEPTED":    WorkRequestStatusTypeAccepted,
	"IN_PROGRESS": WorkRequestStatusTypeInProgress,
	"FAILED":      WorkRequestStatusTypeFailed,
	"SUCCEEDED":   WorkRequestStatusTypeSucceeded,
	"CANCELING":   WorkRequestStatusTypeCanceling,
	"CANCELED":    WorkRequestStatusTypeCanceled,
}

// GetWorkRequestStatusTypeEnumValues Enumerates the set of values for WorkRequestStatusTypeEnum
func GetWorkRequestStatusTypeEnumValues() []WorkRequestStatusTypeEnum {
	values := make([]WorkRequestStatusTypeEnum, 0)
	for _, v := range mappingWorkRequestStatusType {
		values = append(values, v)
	}
	return values
}
