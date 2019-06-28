// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// CIMS API
//
// Cloud Incident Management Service
//

package cims

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Status Ticket Status
type Status struct {
	Code *string `mandatory:"false" json:"code"`

	Message *string `mandatory:"false" json:"message"`
}

func (m Status) String() string {
	return common.PointerString(m)
}
