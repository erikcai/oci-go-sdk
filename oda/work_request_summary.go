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

// WorkRequestSummary A description of workrequest status
type WorkRequestSummary struct {

	// The id of the work request.
	Id *string `mandatory:"true" json:"id"`

	// The ocid of the compartment that contains the work request.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The ocid of the ODA instance to which this work request pertains.
	OdaInstanceId *string `mandatory:"true" json:"odaInstanceId"`

	// type of the operation associated with the work request
	RequestAction WorkRequestSummaryRequestActionEnum `mandatory:"true" json:"requestAction"`

	// status of current work request.
	Status WorkRequestSummaryStatusEnum `mandatory:"true" json:"status"`

	// The resources affected by this work request.
	Resources []WorkRequestResource `mandatory:"true" json:"resources"`
}

func (m WorkRequestSummary) String() string {
	return common.PointerString(m)
}

// WorkRequestSummaryRequestActionEnum Enum with underlying type: string
type WorkRequestSummaryRequestActionEnum string

// Set of constants representing the allowable values for WorkRequestSummaryRequestActionEnum
const (
	WorkRequestSummaryRequestActionCreateOdaInstance  WorkRequestSummaryRequestActionEnum = "CREATE_ODA_INSTANCE"
	WorkRequestSummaryRequestActionDeleteOdaInstance  WorkRequestSummaryRequestActionEnum = "DELETE_ODA_INSTANCE"
	WorkRequestSummaryRequestActionPurgeOdaInstance   WorkRequestSummaryRequestActionEnum = "PURGE_ODA_INSTANCE"
	WorkRequestSummaryRequestActionRecoverOdaInstance WorkRequestSummaryRequestActionEnum = "RECOVER_ODA_INSTANCE"
	WorkRequestSummaryRequestActionCreateAssocation   WorkRequestSummaryRequestActionEnum = "CREATE_ASSOCATION"
	WorkRequestSummaryRequestActionDeleteAssociation  WorkRequestSummaryRequestActionEnum = "DELETE_ASSOCIATION"
)

var mappingWorkRequestSummaryRequestAction = map[string]WorkRequestSummaryRequestActionEnum{
	"CREATE_ODA_INSTANCE":  WorkRequestSummaryRequestActionCreateOdaInstance,
	"DELETE_ODA_INSTANCE":  WorkRequestSummaryRequestActionDeleteOdaInstance,
	"PURGE_ODA_INSTANCE":   WorkRequestSummaryRequestActionPurgeOdaInstance,
	"RECOVER_ODA_INSTANCE": WorkRequestSummaryRequestActionRecoverOdaInstance,
	"CREATE_ASSOCATION":    WorkRequestSummaryRequestActionCreateAssocation,
	"DELETE_ASSOCIATION":   WorkRequestSummaryRequestActionDeleteAssociation,
}

// GetWorkRequestSummaryRequestActionEnumValues Enumerates the set of values for WorkRequestSummaryRequestActionEnum
func GetWorkRequestSummaryRequestActionEnumValues() []WorkRequestSummaryRequestActionEnum {
	values := make([]WorkRequestSummaryRequestActionEnum, 0)
	for _, v := range mappingWorkRequestSummaryRequestAction {
		values = append(values, v)
	}
	return values
}

// WorkRequestSummaryStatusEnum Enum with underlying type: string
type WorkRequestSummaryStatusEnum string

// Set of constants representing the allowable values for WorkRequestSummaryStatusEnum
const (
	WorkRequestSummaryStatusAccepted   WorkRequestSummaryStatusEnum = "ACCEPTED"
	WorkRequestSummaryStatusInProgress WorkRequestSummaryStatusEnum = "IN_PROGRESS"
	WorkRequestSummaryStatusSucceeded  WorkRequestSummaryStatusEnum = "SUCCEEDED"
	WorkRequestSummaryStatusFailed     WorkRequestSummaryStatusEnum = "FAILED"
	WorkRequestSummaryStatusCanceling  WorkRequestSummaryStatusEnum = "CANCELING"
	WorkRequestSummaryStatusCanceled   WorkRequestSummaryStatusEnum = "CANCELED"
)

var mappingWorkRequestSummaryStatus = map[string]WorkRequestSummaryStatusEnum{
	"ACCEPTED":    WorkRequestSummaryStatusAccepted,
	"IN_PROGRESS": WorkRequestSummaryStatusInProgress,
	"SUCCEEDED":   WorkRequestSummaryStatusSucceeded,
	"FAILED":      WorkRequestSummaryStatusFailed,
	"CANCELING":   WorkRequestSummaryStatusCanceling,
	"CANCELED":    WorkRequestSummaryStatusCanceled,
}

// GetWorkRequestSummaryStatusEnumValues Enumerates the set of values for WorkRequestSummaryStatusEnum
func GetWorkRequestSummaryStatusEnumValues() []WorkRequestSummaryStatusEnum {
	values := make([]WorkRequestSummaryStatusEnum, 0)
	for _, v := range mappingWorkRequestSummaryStatus {
		values = append(values, v)
	}
	return values
}
