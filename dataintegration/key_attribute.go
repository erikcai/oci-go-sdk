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

// KeyAttribute An attribute within a key.
type KeyAttribute struct {

	// position
	Position *int `mandatory:"false" json:"position"`

	Attribute *ShapeField `mandatory:"false" json:"attribute"`
}

func (m KeyAttribute) String() string {
	return common.PointerString(m)
}
