// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// OraCache Public API
//
// Oracle Caching Service Public API
//

package cache

import (
	"github.com/oracle/oci-go-sdk/common"
)

// WorkRequest The details of a work request.
type WorkRequest struct {

	// The type of operation that is currently being performed.
	OperationType *string `mandatory:"true" json:"operationType"`

	// The current status of the work request
	Status WorkRequestStatusEnum `mandatory:"true" json:"status"`

	// The OCID of the work request
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment that contains the work request
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The list of resources this work request affects
	Resources []WorkRequestResource `mandatory:"true" json:"resources"`

	// The current progress of the work request
	PercentComplete *float32 `mandatory:"true" json:"percentComplete"`

	// The time this work request was created
	TimeAccepted *common.SDKTime `mandatory:"true" json:"timeAccepted"`

	// The time this work request was moved from ACCEPTED status to IN_PROGRESS status
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The time this work request reached a terminal status - SUCCEEDED, CANCELED or FAILED
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`
}

func (m WorkRequest) String() string {
	return common.PointerString(m)
}

// WorkRequestStatusEnum Enum with underlying type: string
type WorkRequestStatusEnum string

// Set of constants representing the allowable values for WorkRequestStatusEnum
const (
	WorkRequestStatusAccepted   WorkRequestStatusEnum = "ACCEPTED"
	WorkRequestStatusInProgress WorkRequestStatusEnum = "IN_PROGRESS"
	WorkRequestStatusFailed     WorkRequestStatusEnum = "FAILED"
	WorkRequestStatusSucceeded  WorkRequestStatusEnum = "SUCCEEDED"
	WorkRequestStatusCanceling  WorkRequestStatusEnum = "CANCELING"
	WorkRequestStatusCanceled   WorkRequestStatusEnum = "CANCELED"
)

var mappingWorkRequestStatus = map[string]WorkRequestStatusEnum{
	"ACCEPTED":    WorkRequestStatusAccepted,
	"IN_PROGRESS": WorkRequestStatusInProgress,
	"FAILED":      WorkRequestStatusFailed,
	"SUCCEEDED":   WorkRequestStatusSucceeded,
	"CANCELING":   WorkRequestStatusCanceling,
	"CANCELED":    WorkRequestStatusCanceled,
}

// GetWorkRequestStatusEnumValues Enumerates the set of values for WorkRequestStatusEnum
func GetWorkRequestStatusEnumValues() []WorkRequestStatusEnum {
	values := make([]WorkRequestStatusEnum, 0)
	for _, v := range mappingWorkRequestStatus {
		values = append(values, v)
	}
	return values
}
