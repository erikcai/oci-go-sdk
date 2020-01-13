// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// MySQL as a Service API
//
// The API for the MySQL Service
//

package mysql

import (
	"github.com/oracle/oci-go-sdk/common"
)

// WorkRequestSummary A summary of WorkRequest status.
type WorkRequestSummary struct {

	// The id of the work request.
	Id *string `mandatory:"true" json:"id"`

	// the original operation ID requested
	OperationType WorkRequestOperationTypeEnum `mandatory:"true" json:"operationType"`

	// Current status of the work request
	Status WorkRequestStatusTypeEnum `mandatory:"true" json:"status"`

	// The ocid of the compartment that contains the work request. Work requests should be scoped to the same compartment as the resource the work request affects.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// How much progress the operation has made, vs the total amount of work that must be performed.
	PercentComplete *float32 `mandatory:"true" json:"percentComplete"`

	// The time the Work Request was accepted, as described by RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeAccepted *common.SDKTime `mandatory:"true" json:"timeAccepted"`

	// The time the Work Request was started, as described by RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The time the Work Request was finished, as described by RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`
}

func (m WorkRequestSummary) String() string {
	return common.PointerString(m)
}
