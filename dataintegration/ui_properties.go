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

// UiProperties auto generated description
type UiProperties struct {

	// xCoordinate
	XCoordinate *float32 `mandatory:"false" json:"xCoordinate"`

	// yCoordinate
	YCoordinate *float32 `mandatory:"false" json:"yCoordinate"`
}

func (m UiProperties) String() string {
	return common.PointerString(m)
}
