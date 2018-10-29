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

// Publisher Model of the publisher
type Publisher struct {

	// Unique identifier of the publisher
	Id *string `mandatory:"false" json:"id"`

	// Name of the publisher
	Name *string `mandatory:"false" json:"name"`

	// Description of the publisher
	Description *string `mandatory:"false" json:"description"`

	// The year the publisher was founded
	YearFounded *int64 `mandatory:"false" json:"yearFounded"`

	// The count of employee of the publisher
	EmployeeCount *string `mandatory:"false" json:"employeeCount"`

	// The website URL of the publisher
	WebsiteUrl *string `mandatory:"false" json:"websiteUrl"`

	// The contact email of the publisher
	ContactEmail *string `mandatory:"false" json:"contactEmail"`

	// The contact phone number of the publisher
	ContactPhone *string `mandatory:"false" json:"contactPhone"`

	// The headquater address of the publisher
	HqAddress *string `mandatory:"false" json:"hqAddress"`

	// The Logo URL of the publisher
	Logo *UploadData `mandatory:"false" json:"logo"`

	// The reference links
	Links []Link `mandatory:"false" json:"links"`
}

func (m Publisher) String() string {
	return common.PointerString(m)
}
