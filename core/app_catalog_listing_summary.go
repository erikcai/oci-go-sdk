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

// AppCatalogListingSummary The representation of AppCatalogListingSummary
type AppCatalogListingSummary struct {

	// the region free ocid of the listing resource.
	Id *string `mandatory:"true" json:"id"`

	// The display name of the listing.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The short summary to display in list requests.
	Summary *string `mandatory:"true" json:"summary"`

	// Name of the publisher who published this listing.
	PublisherName *string `mandatory:"false" json:"publisherName"`
}

//GetId returns Id
func (m AppCatalogListingSummary) GetId() *string {
	return m.Id
}

//GetDisplayName returns DisplayName
func (m AppCatalogListingSummary) GetDisplayName() *string {
	return m.DisplayName
}

//GetSummary returns Summary
func (m AppCatalogListingSummary) GetSummary() *string {
	return m.Summary
}

//GetPublisherName returns PublisherName
func (m AppCatalogListingSummary) GetPublisherName() *string {
	return m.PublisherName
}

func (m AppCatalogListingSummary) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m AppCatalogListingSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAppCatalogListingSummary AppCatalogListingSummary
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeAppCatalogListingSummary
	}{
		"listing",
		(MarshalTypeAppCatalogListingSummary)(m),
	}

	return json.Marshal(&s)
}
