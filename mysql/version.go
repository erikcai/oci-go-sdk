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

// Version A supported MySQL Version.
type Version struct {

	// The specific version identifier
	Version *string `mandatory:"false" json:"version"`

	// A link to a page describing the version.
	Description *string `mandatory:"false" json:"description"`
}

func (m Version) String() string {
	return common.PointerString(m)
}
