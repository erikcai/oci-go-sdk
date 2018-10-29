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

// CreateLeadDetails The details for creating a new lead
type CreateLeadDetails struct {

	// The unique identifier of the application
	ApplicationId *string `mandatory:"true" json:"applicationId"`

	// First name of the lead
	FirstName *string `mandatory:"true" json:"firstName"`

	// Last name of the lead
	LastName *string `mandatory:"true" json:"lastName"`

	// Company name of the lead
	CompanyName *string `mandatory:"true" json:"companyName"`

	// Company address of the lead
	CompanyAddress *string `mandatory:"true" json:"companyAddress"`

	// Email of the lead
	Email *string `mandatory:"true" json:"email"`

	// The oracle TC ID
	OracleTcId *int64 `mandatory:"false" json:"oracleTcId"`

	// Phone number of the lead
	Phone *string `mandatory:"false" json:"phone"`

	// The additional note of the lead
	AdditionalNotes *string `mandatory:"false" json:"additionalNotes"`
}

func (m CreateLeadDetails) String() string {
	return common.PointerString(m)
}
