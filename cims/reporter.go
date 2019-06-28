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

// Reporter The representation of Reporter
type Reporter struct {
	UserId *string `mandatory:"true" json:"userId"`

	Email *string `mandatory:"false" json:"email"`
}

func (m Reporter) String() string {
	return common.PointerString(m)
}
