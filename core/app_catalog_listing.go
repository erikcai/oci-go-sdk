// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"github.com/oracle/oci-go-sdk/common"
)

// AppCatalogListing Listing details.
type AppCatalogListing struct {

	// The OCID of the listing.
	Id *string `mandatory:"true" json:"id"`

	// Date and time the listing was last updated, in RFC3339 format.
	// Example: `2018-03-20T12:32:53.532Z`
	LastUpdated *common.SDKTime `mandatory:"true" json:"lastUpdated"`

	// Name of the listing.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Date and time the listing was published, in RFC3339 format.
	// Example: `2018-03-20T12:32:53.532Z`
	PublishedDate *common.SDKTime `mandatory:"true" json:"publishedDate"`

	// Name of the publisher who published this listing.
	PublisherName *string `mandatory:"true" json:"publisherName"`

	// Summary of the listing.
	Summary *string `mandatory:"true" json:"summary"`

	// Listing's contact URL.
	ContactUrl *string `mandatory:"false" json:"contactUrl"`

	// Description of the listing.
	Description *string `mandatory:"false" json:"description"`

	// Publisher's logo URL.
	PublisherLogoUrl *string `mandatory:"false" json:"publisherLogoUrl"`
}

func (m AppCatalogListing) String() string {
	return common.PointerString(m)
}
