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

// SubscribedAppCatalogListingSummary The representation of SubscribedAppCatalogListingSummary
type SubscribedAppCatalogListingSummary struct {

	// the region free ocid of the listing resource.
	Id *string `mandatory:"true" json:"id"`

	// The display name of the listing.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The short summary to display in list requests.
	Summary *string `mandatory:"true" json:"summary"`

	// Date and time at which the listing was subscribed, in RFC3339 format.
	// Example: `2018-03-20T12:32:53.532Z`
	SubscribedOn *common.SDKTime `mandatory:"true" json:"subscribedOn"`

	// The compartmentID of the listing.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Listing resource version.
	ListingResourceVersion *string `mandatory:"true" json:"listingResourceVersion"`

	// Listing resource id.
	ListingResourceId *string `mandatory:"true" json:"listingResourceId"`

	// Name of the publisher who published this listing.
	PublisherName *string `mandatory:"false" json:"publisherName"`
}

//GetId returns Id
func (m SubscribedAppCatalogListingSummary) GetId() *string {
	return m.Id
}

//GetDisplayName returns DisplayName
func (m SubscribedAppCatalogListingSummary) GetDisplayName() *string {
	return m.DisplayName
}

//GetSummary returns Summary
func (m SubscribedAppCatalogListingSummary) GetSummary() *string {
	return m.Summary
}

//GetPublisherName returns PublisherName
func (m SubscribedAppCatalogListingSummary) GetPublisherName() *string {
	return m.PublisherName
}

func (m SubscribedAppCatalogListingSummary) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m SubscribedAppCatalogListingSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSubscribedAppCatalogListingSummary SubscribedAppCatalogListingSummary
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeSubscribedAppCatalogListingSummary
	}{
		"subscribedListing",
		(MarshalTypeSubscribedAppCatalogListingSummary)(m),
	}

	return json.Marshal(&s)
}
