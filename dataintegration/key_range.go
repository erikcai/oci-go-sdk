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

// KeyRange auto generated description
type KeyRange struct {
	Key *ShapeField `mandatory:"false" json:"key"`

	// range
	Range []string `mandatory:"false" json:"range"`
}

func (m KeyRange) String() string {
	return common.PointerString(m)
}
