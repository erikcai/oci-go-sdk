// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// AppCatalogListingResourceVersionSummary Listing Resource Version
type AppCatalogListingResourceVersionSummary struct {

	// The OCID of the listing this resource version belongs to.
	ListingId *string `mandatory:"true" json:"listingId"`

	// Date and time the listing resource version was published, in RFC3339 format.
	// Example: `2018-03-20T12:32:53.532Z`
	PublishedDate *common.SDKTime `mandatory:"true" json:"publishedDate"`

	// OCID of the listing resource.
	ResourceId *string `mandatory:"true" json:"resourceId"`

	// Resource Version.
	ResourceVersion *string `mandatory:"true" json:"resourceVersion"`
}

//GetListingId returns ListingId
func (m AppCatalogListingResourceVersionSummary) GetListingId() *string {
	return m.ListingId
}

//GetPublishedDate returns PublishedDate
func (m AppCatalogListingResourceVersionSummary) GetPublishedDate() *common.SDKTime {
	return m.PublishedDate
}

//GetResourceId returns ResourceId
func (m AppCatalogListingResourceVersionSummary) GetResourceId() *string {
	return m.ResourceId
}

//GetResourceVersion returns ResourceVersion
func (m AppCatalogListingResourceVersionSummary) GetResourceVersion() *string {
	return m.ResourceVersion
}

func (m AppCatalogListingResourceVersionSummary) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m AppCatalogListingResourceVersionSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAppCatalogListingResourceVersionSummary AppCatalogListingResourceVersionSummary
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeAppCatalogListingResourceVersionSummary
	}{
		"listingResourceVersionSummary",
		(MarshalTypeAppCatalogListingResourceVersionSummary)(m),
	}

	return json.Marshal(&s)
}
