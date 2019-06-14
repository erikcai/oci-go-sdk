// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Digital Assistant Control Plane API
//
// API to create and maintain Digital Assistant (ODA) service instances.
//

package oda

import (
	"github.com/oracle/oci-go-sdk/common"
)

// WorkRequest A description of workrequest status
type WorkRequest struct {

	// The id of the work request.
	Id *string `mandatory:"true" json:"id"`

	// The ocid of the compartment that contains the work request.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The ocid of the ODA instance to which this work request pertains.
	OdaInstanceId *string `mandatory:"true" json:"odaInstanceId"`

	// type of the operation associated with the work request
	RequestAction WorkRequestRequestActionEnum `mandatory:"true" json:"requestAction"`

	// status of current work request.
	Status WorkRequestStatusEnum `mandatory:"true" json:"status"`

	// The resources affected by this work request.
	Resources []WorkRequestResource `mandatory:"true" json:"resources"`

	// Percentage of the request completed.
	PercentComplete *float32 `mandatory:"true" json:"percentComplete"`

	// The date and time the request was created, as described in
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	TimeAccepted *common.SDKTime `mandatory:"true" json:"timeAccepted"`

	// Short message providing more detail for the current status. For example, if a work request fails
	// this may include information about the resource that failed
	StatusMessage *string `mandatory:"false" json:"statusMessage"`

	// The date and time the request was started, as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339),
	// section 14.29.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The date and time the object was finished, as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`
}

func (m WorkRequest) String() string {
	return common.PointerString(m)
}

// WorkRequestRequestActionEnum Enum with underlying type: string
type WorkRequestRequestActionEnum string

// Set of constants representing the allowable values for WorkRequestRequestActionEnum
const (
	WorkRequestRequestActionCreateOdaInstance  WorkRequestRequestActionEnum = "CREATE_ODA_INSTANCE"
	WorkRequestRequestActionDeleteOdaInstance  WorkRequestRequestActionEnum = "DELETE_ODA_INSTANCE"
	WorkRequestRequestActionPurgeOdaInstance   WorkRequestRequestActionEnum = "PURGE_ODA_INSTANCE"
	WorkRequestRequestActionRecoverOdaInstance WorkRequestRequestActionEnum = "RECOVER_ODA_INSTANCE"
)

var mappingWorkRequestRequestAction = map[string]WorkRequestRequestActionEnum{
	"CREATE_ODA_INSTANCE":  WorkRequestRequestActionCreateOdaInstance,
	"DELETE_ODA_INSTANCE":  WorkRequestRequestActionDeleteOdaInstance,
	"PURGE_ODA_INSTANCE":   WorkRequestRequestActionPurgeOdaInstance,
	"RECOVER_ODA_INSTANCE": WorkRequestRequestActionRecoverOdaInstance,
}

// GetWorkRequestRequestActionEnumValues Enumerates the set of values for WorkRequestRequestActionEnum
func GetWorkRequestRequestActionEnumValues() []WorkRequestRequestActionEnum {
	values := make([]WorkRequestRequestActionEnum, 0)
	for _, v := range mappingWorkRequestRequestAction {
		values = append(values, v)
	}
	return values
}

// WorkRequestStatusEnum Enum with underlying type: string
type WorkRequestStatusEnum string

// Set of constants representing the allowable values for WorkRequestStatusEnum
const (
	WorkRequestStatusAccepted   WorkRequestStatusEnum = "ACCEPTED"
	WorkRequestStatusInProgress WorkRequestStatusEnum = "IN_PROGRESS"
	WorkRequestStatusSucceeded  WorkRequestStatusEnum = "SUCCEEDED"
	WorkRequestStatusFailed     WorkRequestStatusEnum = "FAILED"
	WorkRequestStatusCanceling  WorkRequestStatusEnum = "CANCELING"
	WorkRequestStatusCanceled   WorkRequestStatusEnum = "CANCELED"
)

var mappingWorkRequestStatus = map[string]WorkRequestStatusEnum{
	"ACCEPTED":    WorkRequestStatusAccepted,
	"IN_PROGRESS": WorkRequestStatusInProgress,
	"SUCCEEDED":   WorkRequestStatusSucceeded,
	"FAILED":      WorkRequestStatusFailed,
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
