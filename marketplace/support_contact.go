// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Marketplace REST API endpoint
//
// Marketplace REST API for loom plugin
//

package marketplace

import (
	"github.com/oracle/oci-go-sdk/common"
)

// SupportContact The contact info of the supporters
type SupportContact struct {

	// Name of the supporter
	Name *string `mandatory:"false" json:"name"`

	// Phone number of the supporter
	Phone *string `mandatory:"false" json:"phone"`

	// Email of the supporter
	Email *string `mandatory:"false" json:"email"`

	// Subject of the support
	Subject *string `mandatory:"false" json:"subject"`
}

func (m SupportContact) String() string {
	return common.PointerString(m)
}
