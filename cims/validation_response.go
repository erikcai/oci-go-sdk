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

// ValidationResponse The representation of ValidationResponse
type ValidationResponse struct {
	IsValidUser *bool `mandatory:"false" json:"isValidUser"`
}

func (m ValidationResponse) String() string {
	return common.PointerString(m)
}
