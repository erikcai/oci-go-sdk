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

// WorkRequestError Errors related to a specific work request.
type WorkRequestError struct {

	// A machine-usable code for the error that occured. Error codes are listed on
	// (https://docs.cloud.oracle.com/Content/API/References/apierrors.htm)
	Code *string `mandatory:"true" json:"code"`

	// A human-readable error string.
	Message *string `mandatory:"true" json:"message"`

	// The date and time the error occurred, as described by RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	Timestamp *common.SDKTime `mandatory:"true" json:"timestamp"`
}

func (m WorkRequestError) String() string {
	return common.PointerString(m)
}
