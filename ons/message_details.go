// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// ONS Gateway API
//
// A description of the ONS Gateway API
//

package ons

import (
	"github.com/oracle/oci-go-sdk/common"
)

// MessageDetails message details
type MessageDetails struct {

	// The body of the message to be published.
	Body *string `mandatory:"true" json:"body"`

	// The title of the message to be published.
	Title *string `mandatory:"false" json:"title"`
}

func (m MessageDetails) String() string {
	return common.PointerString(m)
}
