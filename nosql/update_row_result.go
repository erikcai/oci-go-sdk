// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// ndcs-control-plane API
//
// The control plane API for NoSQL Database Cloud Service HTTPS
// provides endpoints to perform NDCS operations, including creation
// and deletion of tables and indexes; population and access of data
// in tables; and access of table usage metrics.
//

package nosql

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateRowResult The result of an UpdateRow operation.
type UpdateRowResult struct {

	// An opaque version string associated with the row.
	Version *string `mandatory:"false" json:"version"`

	// The version string associated with the existing row.
	// Returned if the put fails due to options setting in the
	// request.
	ExistingVersion *string `mandatory:"false" json:"existingVersion"`

	// The map of values from a row.
	ExistingValue map[string]interface{} `mandatory:"false" json:"existingValue"`

	Usage *RequestUsage `mandatory:"false" json:"usage"`
}

func (m UpdateRowResult) String() string {
	return common.PointerString(m)
}
