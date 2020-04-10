// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Data Integration Service API Specification
//
// Data Integration Service API Specification
//

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ErrorDetails Error details.
type ErrorDetails struct {

	// A short error code that defines the error, meant for programmatic parsing. See
	// API Errors (https://docs.cloud.oracle.com/Content/API/References/apierrors.htm).
	Code *string `mandatory:"true" json:"code"`

	// A human-readable error string.
	Message *string `mandatory:"true" json:"message"`
}

func (m ErrorDetails) String() string {
	return common.PointerString(m)
}
