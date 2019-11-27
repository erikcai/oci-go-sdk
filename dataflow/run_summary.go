// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Data Flow API
//
// Use the Data Flow APIs to run any Apache Spark application at any scale without deploying or managing any infrastructure.
//

package dataflow

import (
	"github.com/oracle/oci-go-sdk/common"
)

// RunSummary A summary of the run.
type RunSummary struct {

	// The application ID.
	ApplicationId *string `mandatory:"true" json:"applicationId"`

	// The OCID of the compartment that contains this application.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The ID of a run.
	Id *string `mandatory:"true" json:"id"`

	// The Spark language.
	Language ApplicationLanguageEnum `mandatory:"true" json:"language"`

	// The current state of this run.
	LifecycleState RunLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time a application was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2018-04-03T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time a application was updated, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2018-04-03T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The data read by the run in bytes.
	DataReadInBytes *int64 `mandatory:"false" json:"dataReadInBytes"`

	// The data written by the run in bytes.
	DataWrittenInBytes *int64 `mandatory:"false" json:"dataWrittenInBytes"`

	// A user-friendly name. This name is not necessarily unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The detailed messages about the lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" json:"opcRequestId"`

	// The OCID of the user who created the resource.
	OwnerPrincipalId *string `mandatory:"false" json:"ownerPrincipalId"`

	// The username of the user who created the resource.  If the username of the owner does not exist,
	// `null` will be returned and the caller should refer to the ownerPrincipalId value instead.
	OwnerUserName *string `mandatory:"false" json:"ownerUserName"`

	// The duration of the run in milliseconds.
	RunDurationInMilliseconds *int64 `mandatory:"false" json:"runDurationInMilliseconds"`

	// The total number of oCPU requested by the run.
	TotalOCpu *int `mandatory:"false" json:"totalOCpu"`
}

func (m RunSummary) String() string {
	return common.PointerString(m)
}
