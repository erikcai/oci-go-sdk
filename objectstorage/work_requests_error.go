// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Object Storage Service API
//
// Common set of Object and Archive Storage APIs for managing buckets and objects.
//

package objectstorage

import (
	"github.com/oracle/oci-go-sdk/common"
)

// WorkRequestsError The representation of WorkRequestsError
type WorkRequestsError struct {

	// A machine-usable code for the error that occured. Error codes are listed on
	// (https://docs.us-phoenix-1.oraclecloud.com/Content/API/References/apierrors.htm)
	Code *string `mandatory:"false" json:"code"`

	// A human readable description of the issue encountered.
	Message *string `mandatory:"false" json:"message"`

	// The time the error happened
	Timestamp *common.SDKTime `mandatory:"false" json:"timestamp"`
}

func (m WorkRequestsError) String() string {
	return common.PointerString(m)
}
