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

// MysqlaasInstanceLimitsCheck Describes the error if the check does not pass. Empty otherwise.
type MysqlaasInstanceLimitsCheck struct {

	// In case of error this field will have the respective error code. Otherwise, it will be null.
	ErrorCode *string `mandatory:"false" json:"errorCode"`

	// In case of error this field will have the message describing the problem. Otherwise, it will be null.
	ErrorMessage *string `mandatory:"false" json:"errorMessage"`
}

func (m MysqlaasInstanceLimitsCheck) String() string {
	return common.PointerString(m)
}
