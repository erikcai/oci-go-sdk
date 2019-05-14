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

// WorkRequestResource A resource created or operated on by a work request.
type WorkRequestResource struct {

	// The action to be taken against the ODA instance
	ResourceAction WorkRequestResourceResourceActionEnum `mandatory:"true" json:"resourceAction"`

	// The resource type the work request is affects.
	ResourceType *string `mandatory:"true" json:"resourceType"`

	// The identifier of the ODA instance that is the subject of the request.
	ResourceId *string `mandatory:"true" json:"resourceId"`

	// The current state of the work request. Terminal states correspond to the action being performed.
	Status WorkRequestResourceStatusEnum `mandatory:"true" json:"status"`

	// Short message providing more detail for the current status. For example, if an operation fails
	// this may include information about the reason for the failure and possible resolution.
	StatusMessage *string `mandatory:"false" json:"statusMessage"`

	// The URI path that the user can do a GET on to access the resource metadata
	ResourceUri *string `mandatory:"false" json:"resourceUri"`
}

func (m WorkRequestResource) String() string {
	return common.PointerString(m)
}

// WorkRequestResourceResourceActionEnum Enum with underlying type: string
type WorkRequestResourceResourceActionEnum string

// Set of constants representing the allowable values for WorkRequestResourceResourceActionEnum
const (
	WorkRequestResourceResourceActionCreate WorkRequestResourceResourceActionEnum = "CREATE"
	WorkRequestResourceResourceActionUpdate WorkRequestResourceResourceActionEnum = "UPDATE"
	WorkRequestResourceResourceActionDelete WorkRequestResourceResourceActionEnum = "DELETE"
)

var mappingWorkRequestResourceResourceAction = map[string]WorkRequestResourceResourceActionEnum{
	"CREATE": WorkRequestResourceResourceActionCreate,
	"UPDATE": WorkRequestResourceResourceActionUpdate,
	"DELETE": WorkRequestResourceResourceActionDelete,
}

// GetWorkRequestResourceResourceActionEnumValues Enumerates the set of values for WorkRequestResourceResourceActionEnum
func GetWorkRequestResourceResourceActionEnumValues() []WorkRequestResourceResourceActionEnum {
	values := make([]WorkRequestResourceResourceActionEnum, 0)
	for _, v := range mappingWorkRequestResourceResourceAction {
		values = append(values, v)
	}
	return values
}

// WorkRequestResourceStatusEnum Enum with underlying type: string
type WorkRequestResourceStatusEnum string

// Set of constants representing the allowable values for WorkRequestResourceStatusEnum
const (
	WorkRequestResourceStatusAccepted   WorkRequestResourceStatusEnum = "ACCEPTED"
	WorkRequestResourceStatusInProgress WorkRequestResourceStatusEnum = "IN_PROGRESS"
	WorkRequestResourceStatusSucceeded  WorkRequestResourceStatusEnum = "SUCCEEDED"
	WorkRequestResourceStatusFailed     WorkRequestResourceStatusEnum = "FAILED"
	WorkRequestResourceStatusCanceling  WorkRequestResourceStatusEnum = "CANCELING"
	WorkRequestResourceStatusCanceled   WorkRequestResourceStatusEnum = "CANCELED"
)

var mappingWorkRequestResourceStatus = map[string]WorkRequestResourceStatusEnum{
	"ACCEPTED":    WorkRequestResourceStatusAccepted,
	"IN_PROGRESS": WorkRequestResourceStatusInProgress,
	"SUCCEEDED":   WorkRequestResourceStatusSucceeded,
	"FAILED":      WorkRequestResourceStatusFailed,
	"CANCELING":   WorkRequestResourceStatusCanceling,
	"CANCELED":    WorkRequestResourceStatusCanceled,
}

// GetWorkRequestResourceStatusEnumValues Enumerates the set of values for WorkRequestResourceStatusEnum
func GetWorkRequestResourceStatusEnumValues() []WorkRequestResourceStatusEnum {
	values := make([]WorkRequestResourceStatusEnum, 0)
	for _, v := range mappingWorkRequestResourceStatus {
		values = append(values, v)
	}
	return values
}
