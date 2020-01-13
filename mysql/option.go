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

// Option The representation of Option
type Option struct {

	// The option name.
	Name *string `mandatory:"true" json:"name"`

	// The option actual value.
	Value *string `mandatory:"true" json:"value"`

	// Whether the option is explicitly defined by the user or configured by default.
	IsUserDefined *bool `mandatory:"false" json:"isUserDefined"`
}

func (m Option) String() string {
	return common.PointerString(m)
}
