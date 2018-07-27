// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Email Delivery Service API
//
// API for managing OCI Email Delivery services.
//

package email

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateSenderDetails The details needed for creating a sender.
type CreateSenderDetails struct {

	// The OCID of the compartment that contains the sender.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The email address of the sender.
	EmailAddress *string `mandatory:"true" json:"emailAddress"`
}

func (m CreateSenderDetails) String() string {
	return common.PointerString(m)
}
