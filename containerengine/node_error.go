// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Clusters API Specification
//
// Container Engine for Kubernetes API
//

package containerengine

import (
	"github.com/oracle/oci-go-sdk/common"
)

// NodeError The properties that define an error.
type NodeError struct {

	// A short error code that defines the error, meant for programmatic parsing. See API Errors (https://docs.us-phoenix-1.oraclecloud.com/Content/API/References/apierrors.htm).
	Code *string `mandatory:"true" json:"code"`

	// A human-readable error string.
	Message *string `mandatory:"true" json:"message"`
}

func (m NodeError) String() string {
	return common.PointerString(m)
}
